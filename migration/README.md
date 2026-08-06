# Axolotl protocol lift

This directory is an immutable, history-preserving migration record. It is not
an input corpus or a second canonical protocol source. Its small destination
workspace is compiled only for migration-scope tests and lint.

The source was `axolotl-stack` `origin/main` at `781dfcb0ab443476b62df3c983750c0c1527a95a`.
Because the source and destination repositories have unrelated histories, the
protocol-specific source was filtered in a temporary clone with `git filter-repo`
and merged with `--allow-unrelated-histories --no-ff`. The filtered source tip is
recorded by the merge parent; the filter retained the source commit ancestry for
the selected paths rather than copying a snapshot.

Retained paths are the protocol parser/lowering and wire-manifest code, the
required `bedrock_core` codec/wire vocabulary, the source correction metadata,
and `tools/wire-conformance`. The filter intentionally excluded
`src/data_generator/**`, gameplay data, the Prismarine/minecraft-data and
bedrock-data submodules, and all source corpora. The retained Rust code is a
reference for extraction and parity only; v2 emitters consume the canonical
manifest, never generated Rust.

The exact history-lift method was:

```sh
git clone /Users/hashim/Coding/Go/Lunar/worktrees/axolotl-stack/update-1.26.30-clean /tmp/protocolgen-axolotl-history
git -C /tmp/protocolgen-axolotl-history fetch origin main
git -C /tmp/protocolgen-axolotl-history filter-repo --force \
  --refs origin/main --to-subdirectory-filter migration/axolotl-protocol \
  --path crates/valentine_gen/src/generator \
  --path crates/valentine_gen/src/ir.rs \
  --path crates/valentine_gen/src/main.rs \
  --path crates/valentine_gen/src/overrides.rs \
  --path crates/valentine_gen/src/parser \
  --path crates/valentine_gen/src/wire_manifest.rs \
  --path crates/valentine_gen/overrides \
  --path crates/valentine_gen/overrides-endstone \
  --path crates/valentine_gen/Cargo.toml \
  --path crates/valentine_gen/README.md \
  --path crates/valentine_gen/MOJANG_PARITY.md \
  --path crates/valentine/bedrock_core \
  --path tools/wire-conformance
git fetch /tmp/protocolgen-axolotl-history refs/heads/origin/main:refs/remotes/axolotl-filtered/main
git merge --no-ff --allow-unrelated-histories --no-commit axolotl-filtered/main
```

The filtered source tip is `7db3f1ab669968d167f20ad001e9bab163ea2936`, derived
from source `origin/main` `781dfcb0ab443476b62df3c983750c0c1527a95a`. The
resulting lift was committed as `19c865de0a9ba6dfb716049947a8654c87eeb06f`.
Do not rewrite this imported history when later slices replace individual
pieces with Go implementations.

The small workspace manifest at `migration/axolotl-protocol/Cargo.toml` is
destination scaffolding for formatting and testing the retained core/generator
scope. It does not vendor a lockfile or any source corpus.
