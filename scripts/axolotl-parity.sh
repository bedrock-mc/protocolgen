#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 2 ]]; then
  echo "usage: $0 AXOLOTL_ROOT CANONICAL_MANIFEST" >&2
  exit 2
fi

axolotl_root=$1
canonical_manifest=$2
work_dir=$(mktemp -d "${TMPDIR:-/tmp}/protocolgen-axolotl-parity.XXXXXX")
trap 'rm -rf "$work_dir"' EXIT
generator_args=(
  --source mojang
  --versions 1.26.40
  --output-dir "$work_dir/valentine"
  --emit-wire-manifest "$work_dir/axolotl-v1.json"
)
if [[ -n "${AXOLOTL_MOJANG_DOCS:-}" ]]; then
  generator_args+=(--mojang-docs "$AXOLOTL_MOJANG_DOCS")
fi

cargo run --manifest-path "$axolotl_root/Cargo.toml" -p valentine_gen -- \
  "${generator_args[@]}"

go run ./cmd/protocolgen parity \
  -manifest "$canonical_manifest" \
  -axolotl "$work_dir/axolotl-v1.json"
