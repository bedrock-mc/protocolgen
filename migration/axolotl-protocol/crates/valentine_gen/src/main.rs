#![allow(clippy::collapsible_if)]
#![allow(clippy::manual_pattern_char_comparison)]
#![allow(clippy::collapsible_match)]
#![allow(clippy::unnecessary_map_or)]
#![allow(clippy::double_ended_iterator_last)]
#![allow(clippy::manual_strip)]
#![allow(clippy::manual_find)]
#![allow(clippy::too_many_arguments)]
#![allow(clippy::needless_late_init)]
#![allow(clippy::useless_format)]
#![allow(dead_code)]

use std::collections::HashSet;
use std::fs::{self, File};
use std::io::{Read, Write};
use std::path::{Path, PathBuf};

use generator::context::GlobalRegistry;
use proc_macro2::Span;
use quote::quote;
use serde::Deserialize;
use syn::{LitStr, parse2};
use toml_edit::{Array, DocumentMut};
use tracing::{info, warn};
use tracing_subscriber::{EnvFilter, fmt};

mod generator;
mod ir;
mod overrides;
mod parser;
mod wire_manifest;

#[derive(Debug, Clone, Deserialize)]
struct BedrockVersionJson {
    #[serde(rename = "version")]
    protocol_version: i32,
    #[serde(rename = "minecraftVersion")]
    minecraft_version: String,
    #[serde(rename = "majorVersion")]
    major_version: String,
    #[serde(rename = "releaseType")]
    release_type: String,
}

#[derive(Debug, Clone)]
struct VersionDecl {
    module_name: String,
    feature: String,
    crate_name: String,
    meta: BedrockVersionJson,
}

#[derive(Debug, Clone)]
struct CliArgs {
    source: ProtocolSource,
    versions: Vec<String>,
    all: bool,
    latest: bool,
    list_versions: bool,
    log_filter: String,
    mojang_docs: Option<PathBuf>,
    endstone_docs: Option<PathBuf>,
    overrides: Option<PathBuf>,
    output_dir: Option<PathBuf>,
    emit_wire_manifest: Option<PathBuf>,
    gen_proto: bool,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
enum ProtocolSource {
    Mojang,
    Endstone,
}

fn print_usage() {
    println!(
        r#"valentine-gen - generate Bedrock protocol crates for Valentine

USAGE:
  cargo run -p valentine_gen -- [OPTIONS]

VERSION SELECTION:
  --latest                Generate only the latest schema version (default)
  --all                   Generate all supported schema versions
  --versions <LIST>       Generate a comma-separated list, e.g. "1.21.130,1.20.80"
  --list-versions         Print available Bedrock versions and exit

GENERATION TARGETS:
  --proto                 Generate protocol code (default)

OTHER OPTIONS:
  --source <NAME>         Protocol source: endstone (default) or mojang
  --mojang-docs <DIR>     Path to a bedrock-protocol-docs checkout (defaults to ./bedrock-protocol-docs)
  --endstone-docs <DIR>   Path to an endstone protocol-docs checkout (defaults to ./endstone-docs)
  --overrides <DIR>       Mojang correction JSON directory (defaults to ./overrides)
  --output-dir <DIR>      Valentine output root (required for schema sources)
  --emit-wire-manifest <FILE>
                         Emit the selected version's fully resolved wire manifest as JSON
  --log <FILTER>          tracing filter (default: "info"), e.g. "debug" or "valentine_gen=debug"
  -h, --help              Print help and exit
"#
    );
}

fn parse_args() -> Result<CliArgs, String> {
    let mut versions: Vec<String> = Vec::new();
    let mut all = false;
    let mut latest = false;
    let mut list_versions = false;
    let mut log_filter = "info".to_string();
    let mut source = ProtocolSource::Endstone;
    let mut mojang_docs: Option<PathBuf> = None;
    let mut endstone_docs: Option<PathBuf> = None;
    let mut overrides: Option<PathBuf> = None;
    let mut output_dir: Option<PathBuf> = None;
    let mut emit_wire_manifest: Option<PathBuf> = None;

    let mut gen_proto = true;

    let mut it = std::env::args().skip(1).peekable();
    while let Some(arg) = it.next() {
        match arg.as_str() {
            "-h" | "--help" => {
                print_usage();
                std::process::exit(0);
            }
            "--all" => all = true,
            "--latest" => latest = true,
            "--list-versions" | "--list" => list_versions = true,
            "--source" => {
                let raw = it
                    .next()
                    .ok_or_else(|| "--source expects endstone or mojang".to_string())?;
                source = parse_source(&raw)?;
            }
            "--proto" | "--protocol" => gen_proto = true,
            "--versions" | "--version" | "-v" => {
                let raw = it
                    .next()
                    .ok_or_else(|| "--versions expects a value".to_string())?;
                for v in raw.split(',').map(|s| s.trim()).filter(|s| !s.is_empty()) {
                    versions.push(v.to_string());
                }
            }
            "--log" | "--log-level" => {
                log_filter = it
                    .next()
                    .ok_or_else(|| "--log expects a value".to_string())?;
            }
            "--endstone-docs" => {
                let raw = it
                    .next()
                    .ok_or_else(|| "--endstone-docs expects a path".to_string())?;
                endstone_docs = Some(PathBuf::from(raw));
            }
            "--mojang-docs" => {
                let raw = it
                    .next()
                    .ok_or_else(|| "--mojang-docs expects a path".to_string())?;
                mojang_docs = Some(PathBuf::from(raw));
            }
            "--overrides" => {
                let raw = it
                    .next()
                    .ok_or_else(|| "--overrides expects a path".to_string())?;
                overrides = Some(PathBuf::from(raw));
            }
            "--output-dir" => {
                let raw = it
                    .next()
                    .ok_or_else(|| "--output-dir expects a path".to_string())?;
                output_dir = Some(PathBuf::from(raw));
            }
            "--emit-wire-manifest" => {
                let raw = it
                    .next()
                    .ok_or_else(|| "--emit-wire-manifest expects a path".to_string())?;
                emit_wire_manifest = Some(PathBuf::from(raw));
            }
            _ if arg.starts_with("--versions=") => {
                let raw = arg.trim_start_matches("--versions=");
                for v in raw.split(',').map(|s| s.trim()).filter(|s| !s.is_empty()) {
                    versions.push(v.to_string());
                }
            }
            _ if arg.starts_with("--log=") => {
                log_filter = arg.trim_start_matches("--log=").to_string();
            }
            _ if arg.starts_with("--source=") => {
                source = parse_source(arg.trim_start_matches("--source="))?;
            }
            _ if arg.starts_with("--endstone-docs=") => {
                endstone_docs = Some(PathBuf::from(arg.trim_start_matches("--endstone-docs=")));
            }
            _ if arg.starts_with("--mojang-docs=") => {
                mojang_docs = Some(PathBuf::from(arg.trim_start_matches("--mojang-docs=")));
            }
            _ if arg.starts_with("--overrides=") => {
                overrides = Some(PathBuf::from(arg.trim_start_matches("--overrides=")));
            }
            _ if arg.starts_with("--output-dir=") => {
                output_dir = Some(PathBuf::from(arg.trim_start_matches("--output-dir=")));
            }
            _ if arg.starts_with("--emit-wire-manifest=") => {
                emit_wire_manifest = Some(PathBuf::from(
                    arg.trim_start_matches("--emit-wire-manifest="),
                ));
            }
            _ => return Err(format!("Unknown argument: {arg}")),
        }
    }

    if all && (latest || !versions.is_empty()) {
        return Err("Use either --all, --latest, or --versions (not multiple)".to_string());
    }
    if latest && !versions.is_empty() {
        return Err("Use either --latest or --versions (not both)".to_string());
    }

    Ok(CliArgs {
        source,
        versions,
        all,
        latest,
        list_versions,
        log_filter,
        mojang_docs,
        endstone_docs,
        overrides,
        output_dir,
        emit_wire_manifest,
        gen_proto,
    })
}

fn parse_source(value: &str) -> Result<ProtocolSource, String> {
    match value {
        "mojang" => Ok(ProtocolSource::Mojang),
        "endstone" => Ok(ProtocolSource::Endstone),
        other => Err(format!(
            "unknown protocol source {other:?}; expected endstone or mojang"
        )),
    }
}

fn init_tracing(filter: &str) {
    let env_filter = EnvFilter::try_new(filter)
        .or_else(|_| EnvFilter::try_new(format!("valentine_gen={filter}")))
        .unwrap_or_else(|_| EnvFilter::new("info"));
    fmt().with_env_filter(env_filter).init();
}

fn version_to_feature(version: &str) -> String {
    format!("bedrock_{}", version.replace('.', "_"))
}

fn version_to_module(version: &str) -> String {
    format!("v{}", version.replace('.', "_"))
}

fn version_to_crate(version: &str) -> String {
    format!("valentine_bedrock_{}", version.replace('.', "_"))
}

fn parse_version(version: &str) -> Vec<u64> {
    version
        .split('.')
        .map(|x| x.parse::<u64>().unwrap_or(0))
        .collect()
}

fn latest_version(versions: &[String]) -> Option<String> {
    versions
        .iter()
        .max_by(|a, b| parse_version(a).cmp(&parse_version(b)))
        .cloned()
}

fn generate_schema_source(
    args: &CliArgs,
    root: &Path,
    valentine_root: &Path,
) -> Result<(), Box<dyn std::error::Error>> {
    let (docs_arg, default_dir) = if args.source == ProtocolSource::Endstone {
        (args.endstone_docs.clone(), "endstone-docs")
    } else {
        (args.mojang_docs.clone(), "bedrock-protocol-docs")
    };
    let docs_root = docs_arg
        .clone()
        .map(|path| {
            if path.is_relative() {
                root.join(path)
            } else {
                path
            }
        })
        .unwrap_or_else(|| root.join(default_dir));
    let override_root = args
        .overrides
        .clone()
        .map(|path| {
            if path.is_relative() {
                root.join(path)
            } else {
                path
            }
        })
        .unwrap_or_else(|| {
            // Each schema source gets its own correction set. Endstone's is
            // expected to stay empty: the corpus is dumped from the server, so
            // a needed correction means the dumper is wrong and belongs
            // upstream rather than in a local patch.
            root.join(if args.source == ProtocolSource::Endstone {
                "overrides-endstone"
            } else {
                "overrides"
            })
        });

    let available = if args.source == ProtocolSource::Endstone {
        parser::endstone::discover_versions(&docs_root)?
    } else {
        parser::mojang::discover_versions(&docs_root)?
    };
    if available.is_empty() {
        return Err(format!(
            "no protocol schema version metadata found below {}",
            docs_root.display()
        )
        .into());
    }
    if args.list_versions {
        for version in &available {
            println!(
                "{} (protocol {})",
                version.minecraft_version, version.protocol_version
            );
        }
        return Ok(());
    }

    let selected = if args.all {
        available.clone()
    } else if !args.versions.is_empty() {
        let mut selected = Vec::new();
        for requested in &args.versions {
            if let Some(version) = available
                .iter()
                .find(|version| version.minecraft_version == *requested)
            {
                selected.push(version.clone());
            } else {
                warn!(version = %requested, "Requested version is not present in the Mojang schema snapshot");
            }
        }
        selected
    } else if args.latest {
        vec![
            available
                .last()
                .cloned()
                .ok_or("No schema versions available")?,
        ]
    } else {
        vec![
            available
                .last()
                .cloned()
                .ok_or("No Mojang schema versions available")?,
        ]
    };
    if selected.is_empty() {
        return Err("No Mojang schema versions selected for generation".into());
    }
    if args.emit_wire_manifest.is_some() && selected.len() != 1 {
        return Err("--emit-wire-manifest requires exactly one selected version".into());
    }

    let parse_result = if args.gen_proto {
        Some(if args.source == ProtocolSource::Endstone {
            parser::endstone::parse(&docs_root, &override_root)?
        } else {
            parser::mojang::parse(&docs_root, &override_root)?
        })
    } else {
        None
    };
    if let (Some(path), Some(parsed)) = (&args.emit_wire_manifest, &parse_result) {
        let version = &selected[0];
        let path = if path.is_relative() {
            root.join(path)
        } else {
            path.clone()
        };
        let manifest = wire_manifest::build(
            parsed,
            "mojang",
            Some(version.minecraft_version.clone()),
            Some(version.protocol_version),
        )?;
        wire_manifest::write(&path, &manifest)?;
        info!(path = %path.display(), "Emitted resolved wire manifest");
    }
    let mut version_decls = Vec::new();
    let mut global_registry = GlobalRegistry::new();
    let bedrock_versions_dir = valentine_root.join("bedrock_versions");
    fs::create_dir_all(&bedrock_versions_dir)?;

    for version in selected {
        let module_name = version_to_module(&version.minecraft_version);
        let feature = version_to_feature(&version.minecraft_version);
        let crate_name = version_to_crate(&version.minecraft_version);
        let crate_dir = bedrock_versions_dir.join(&module_name);
        let crate_src_dir = crate_dir.join("src");
        fs::create_dir_all(&crate_src_dir)?;

        let mut crate_dependencies = HashSet::new();
        if let Some(parse_result) = &parse_result {
            info!(
                minecraft_version = %version.minecraft_version,
                protocol_version = version.protocol_version,
                module = %module_name,
                "Generating protocol from schema source"
            );
            let outcome = generator::generate_protocol_module(
                &crate_name,
                "",
                parse_result,
                &crate_src_dir,
                &mut global_registry,
            )?;
            crate_dependencies.extend(outcome.crate_dependencies);
        }
        write_version_crate(&crate_dir, &crate_src_dir, &crate_name, &crate_dependencies)?;

        let major_version = version
            .minecraft_version
            .split('.')
            .take(2)
            .collect::<Vec<_>>()
            .join(".");
        version_decls.push(VersionDecl {
            module_name,
            feature,
            crate_name,
            meta: BedrockVersionJson {
                protocol_version: version.protocol_version,
                minecraft_version: version.minecraft_version,
                major_version,
                release_type: "release".to_string(),
            },
        });
    }

    // A Mojang checkout contains one release, while the checked-in Valentine
    // surface may intentionally retain older generated crates. Preserve those
    // declarations when adding a new Mojang-generated version in place.
    for existing in read_existing_version_decls(valentine_root)? {
        if !version_decls
            .iter()
            .any(|decl| decl.module_name == existing.module_name)
        {
            version_decls.push(existing);
        }
    }

    let default_version = latest_version(
        &version_decls
            .iter()
            .map(|decl| decl.meta.minecraft_version.clone())
            .collect::<Vec<_>>(),
    )
    .ok_or("No Mojang versions for default")?;
    let default_feature = version_to_feature(&default_version);
    version_decls.sort_by(|a, b| a.module_name.cmp(&b.module_name));
    write_generated_surface(valentine_root, &version_decls, &default_feature)?;
    Ok(())
}

fn read_existing_version_decls(
    valentine_root: &Path,
) -> Result<Vec<VersionDecl>, Box<dyn std::error::Error>> {
    let path = valentine_root
        .join("src")
        .join("bedrock")
        .join("version.rs");
    if !path.exists() {
        return Ok(Vec::new());
    }
    let source = fs::read_to_string(path)?;
    let mut declarations = Vec::new();
    for block in source.split("#[cfg(feature = \"").skip(1) {
        let Some((feature, block)) = block.split_once("\")]") else {
            continue;
        };
        let Some(module_name) = quoted_after(block, "pub mod ", " {") else {
            continue;
        };
        let Some(minecraft_version) = quoted_after(block, "GAME_VERSION: &str = \"", "\"") else {
            continue;
        };
        let Some(protocol) = quoted_after(block, "PROTOCOL_VERSION: i32 = ", "i32") else {
            continue;
        };
        let Some(major_version) = quoted_after(block, "MAJOR_VERSION: &str = \"", "\"") else {
            continue;
        };
        let Some(release_type) = quoted_after(block, "RELEASE_TYPE: &str = \"", "\"") else {
            continue;
        };
        let crate_name = module_name.replacen("v", "valentine_bedrock_", 1);
        if !valentine_root
            .join("bedrock_versions")
            .join(&module_name)
            .exists()
        {
            continue;
        }
        declarations.push(VersionDecl {
            module_name,
            feature: feature.to_string(),
            crate_name,
            meta: BedrockVersionJson {
                protocol_version: protocol.parse()?,
                minecraft_version,
                major_version,
                release_type,
            },
        });
    }
    Ok(declarations)
}

fn quoted_after(source: &str, prefix: &str, suffix: &str) -> Option<String> {
    source
        .split_once(prefix)
        .and_then(|(_, rest)| rest.split_once(suffix))
        .map(|(value, _)| value.to_string())
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args = match parse_args() {
        Ok(args) => args,
        Err(e) => {
            eprintln!("{e}");
            print_usage();
            return Err(e.into());
        }
    };
    init_tracing(&args.log_filter);
    let manifest_dir = std::env::var("CARGO_MANIFEST_DIR")?;
    let root = Path::new(&manifest_dir);
    let valentine_root = args
        .output_dir
        .clone()
        .map(|path| {
            if path.is_relative() {
                root.join(path)
            } else {
                path
            }
        })
        .ok_or("schema-source generation requires --output-dir")?;
    generate_schema_source(&args, root, &valentine_root)
}

fn write_generated_surface(
    valentine_root: &Path,
    version_decls: &[VersionDecl],
    default_feature: &str,
) -> Result<(), Box<dyn std::error::Error>> {
    let bedrock_src_dir = valentine_root.join("src").join("bedrock");
    let protocol_mod_dir = bedrock_src_dir.join("protocol");
    fs::create_dir_all(&protocol_mod_dir)?;

    let protocol_items: Vec<_> = version_decls
        .iter()
        .map(|version| {
            let module_ident = syn::Ident::new(&version.module_name, Span::call_site());
            let crate_ident = syn::Ident::new(&version.crate_name, Span::call_site());
            let feature = LitStr::new(&version.feature, Span::call_site());
            quote! {
                #[cfg(feature = #feature)]
                pub use #crate_ident as #module_ident;
            }
        })
        .collect();
    let version_items: Vec<_> = version_decls
        .iter()
        .map(|version| {
            let module_ident = syn::Ident::new(&version.module_name, Span::call_site());
            let feature = LitStr::new(&version.feature, Span::call_site());
            let game_version = LitStr::new(&version.meta.minecraft_version, Span::call_site());
            let major_version = LitStr::new(&version.meta.major_version, Span::call_site());
            let release_type = LitStr::new(&version.meta.release_type, Span::call_site());
            let protocol_version = version.meta.protocol_version;
            quote! {
                #[cfg(feature = #feature)]
                pub mod #module_ident {
                    pub use super::super::protocol::#module_ident::*;

                    pub const GAME_VERSION: &str = #game_version;
                    pub const PROTOCOL_VERSION: i32 = #protocol_version;
                    pub const MAJOR_VERSION: &str = #major_version;
                    pub const RELEASE_TYPE: &str = #release_type;

                    pub const INFO: super::BedrockVersionInfo = super::BedrockVersionInfo {
                        minecraft_version: GAME_VERSION,
                        protocol_version: PROTOCOL_VERSION,
                        major_version: MAJOR_VERSION,
                        release_type: RELEASE_TYPE,
                    };
                }
            }
        })
        .collect();
    let reexport_items: Vec<_> = version_decls
        .iter()
        .map(|version| {
            let module_ident = syn::Ident::new(&version.module_name, Span::call_site());
            let feature = LitStr::new(&version.feature, Span::call_site());
            quote! {
                #[cfg(feature = #feature)]
                pub use self::version::#module_ident;
            }
        })
        .collect();

    let bedrock_tokens = quote! {
        #![allow(non_camel_case_types)]
        #![allow(non_snake_case)]
        #![allow(dead_code)]
        #![allow(unused_imports)]
        #![allow(clippy::redundant_field_names)]
        #![allow(clippy::manual_flatten)]

        pub mod codec;
        pub mod error;
        pub mod protocol;
        pub mod version;
        pub mod context;
        pub mod borrowed;

        #(#reexport_items)*
    };
    let bedrock_syntax = parse2(bedrock_tokens)?;
    let bedrock_formatted = prettyplease::unparse(&bedrock_syntax);
    fs::write(
        bedrock_src_dir.join("mod.rs"),
        format!(
            "// Generated by valentine_gen\n// Do not edit: see crates/valentine_gen for generator.\n\n{}",
            bedrock_formatted
        ),
    )?;

    let protocol_tokens = quote! {
        #![allow(non_camel_case_types)]
        #![allow(non_snake_case)]
        #![allow(dead_code)]
        #![allow(unused_imports)]
        #![allow(clippy::redundant_field_names)]
        #![allow(clippy::manual_flatten)]

        #(#protocol_items)*
    };
    let protocol_syntax = parse2(protocol_tokens)?;
    fs::write(
        protocol_mod_dir.join("mod.rs"),
        format!(
            "// Generated by valentine_gen\n// Do not edit: see crates/valentine_gen for generator.\n\n{}",
            prettyplease::unparse(&protocol_syntax)
        ),
    )?;

    let version_tokens = quote! {
        #![allow(non_camel_case_types)]
        #![allow(non_snake_case)]
        #![allow(dead_code)]
        #![allow(unused_imports)]

        pub use valentine_bedrock_core::bedrock::version::BedrockVersionInfo;

        #(#version_items)*
    };
    let version_syntax = parse2(version_tokens)?;
    fs::write(
        bedrock_src_dir.join("version.rs"),
        format!(
            "// Generated by valentine_gen\n// Do not edit: see crates/valentine_gen for generator.\n\n{}",
            prettyplease::unparse(&version_syntax)
        ),
    )?;

    update_valentine_manifest_at(valentine_root, default_feature, version_decls)
}

fn update_valentine_manifest_at(
    valentine_root: &Path,
    default_feature: &str,
    versions: &[VersionDecl],
) -> Result<(), Box<dyn std::error::Error>> {
    let valentine_cargo = valentine_root.join("Cargo.toml");
    if !valentine_cargo.exists() {
        // `--output-dir` is also used for scratch generation/parity runs. Such
        // directories intentionally contain generated sources only.
        return Ok(());
    }
    let mut contents = String::new();
    {
        let mut f = File::open(&valentine_cargo)?;
        f.read_to_string(&mut contents)?;
    }
    let mut doc: DocumentMut = contents.parse()?;

    // Ensure tables exist (do this before borrowing them mutably).
    if !doc.as_table().contains_key("dependencies") {
        doc["dependencies"] = toml_edit::table();
    }
    if !doc.as_table().contains_key("features") {
        doc["features"] = toml_edit::table();
    }

    // Dependencies: remove stale generated crates then insert current set.
    {
        let deps_tbl = doc["dependencies"]
            .as_table_mut()
            .ok_or("Cargo.toml missing [dependencies] table")?;

        let existing_deps: Vec<String> = deps_tbl.iter().map(|(k, _)| k.to_string()).collect();
        for key in existing_deps {
            if key.starts_with("valentine_bedrock_") && key != "valentine_bedrock_core" {
                deps_tbl.remove(&key);
            }
        }

        for vd in versions {
            let mut dep = toml_edit::InlineTable::new();
            dep.insert(
                "path",
                toml_edit::Value::from(format!("bedrock_versions/{}", vd.module_name)),
            );
            dep.insert("optional", toml_edit::Value::from(true));
            deps_tbl.insert(
                &vd.crate_name,
                toml_edit::Item::Value(toml_edit::Value::InlineTable(dep)),
            );
        }
    }

    // Features: remove stale generated entries then insert current set.
    {
        let features_tbl = doc["features"]
            .as_table_mut()
            .ok_or("Cargo.toml missing [features] table")?;

        let existing_keys: Vec<String> = features_tbl.iter().map(|(k, _)| k.to_string()).collect();
        let had_default = features_tbl.contains_key("default");
        for key in existing_keys {
            if key.starts_with("bedrock_") {
                features_tbl.remove(&key);
            }
        }

        // Only seed a default when the manifest has none. Which protocol the
        // workspace speaks by default is a deliberate decision - regenerating
        // must never silently move it to whichever version was generated last,
        // because that changes runtime behaviour and breaks every consumer
        // pinned to the previous one.
        if !had_default {
            let mut default_arr = Array::new();
            default_arr.push(default_feature);
            features_tbl.insert(
                "default",
                toml_edit::Item::Value(toml_edit::Value::Array(default_arr)),
            );
        }

        for vd in versions {
            let mut arr = Array::new();
            arr.push(format!("dep:{}", vd.crate_name));
            features_tbl.insert(
                &vd.feature,
                toml_edit::Item::Value(toml_edit::Value::Array(arr)),
            );
        }
    }

    let new_contents = doc.to_string();
    if new_contents != contents {
        let mut f = File::create(&valentine_cargo)?;
        f.write_all(new_contents.as_bytes())?;
    }

    Ok(())
}

// Note: No cleanup logic here; assume old generated files are managed manually.

fn write_version_crate(
    crate_dir: &Path,
    crate_src_dir: &Path,
    crate_name: &str,
    crate_dependencies: &HashSet<String>,
) -> Result<(), Box<dyn std::error::Error>> {
    fs::create_dir_all(crate_dir)?;
    fs::create_dir_all(crate_src_dir)?;

    let cargo_toml = format!(
        r#"[package]
name = "{crate_name}"
version = "0.1.0"
edition = "2024"

[dependencies]
bitflags = "2"
bytes = "1"
uuid = "1.8.0"
valentine_bedrock_core = {{ path = "../../bedrock_core" }}
"#
    );
    let mut extra_deps: Vec<_> = crate_dependencies.iter().collect();
    extra_deps.sort();

    let mut cargo_toml = cargo_toml;
    for dep in extra_deps {
        let dep_path = dep.replacen("valentine_bedrock_", "v", 1);
        cargo_toml.push_str(&format!(r#"{dep} = {{ path = "../{dep_path}" }}"#));
        cargo_toml.push('\n');
    }
    let mut cargo_file = File::create(crate_dir.join("Cargo.toml"))?;
    cargo_file.write_all(cargo_toml.as_bytes())?;

    Ok(())
}
