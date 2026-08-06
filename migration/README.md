# Axolotl protocol lift

This directory is an immutable, history-preserving migration record. It is not
an input corpus and it is not compiled by protocolgen.

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

The exact filter command and source pin are part of the implementation handoff;
do not rewrite this imported history when later slices replace individual
pieces with Go implementations.
