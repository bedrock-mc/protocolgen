#!/usr/bin/env bash
# Checks out the Cinnabar commit a retail corpus pins: only its fixtures, the
# vendored Valentine codec and fixturegen's go.mod.
set -euo pipefail

if [[ $# -ne 2 ]]; then
  echo "usage: $0 CORPUS_JSON CHECKOUT_DIR" >&2
  exit 2
fi

corpus=$1
dir=$2
repository=$(jq -er .source.repository "$corpus")
commit=$(jq -er .source.commit "$corpus")
fixtures=$(jq -er .source.directory "$corpus")

if [[ "$(git -C "$dir" rev-parse --verify --quiet HEAD 2>/dev/null || true)" == "$commit" ]]; then
  exit 0
fi
if [[ ! -d "$dir/.git" ]]; then
  git init --quiet "$dir"
  git -C "$dir" remote add origin "$repository"
  git -C "$dir" config remote.origin.promisor true
  git -C "$dir" config remote.origin.partialclonefilter blob:none
fi
git -C "$dir" sparse-checkout set --no-cone \
  "/$fixtures/" /crates/protocol/vendor/valentine/ /tools/fixturegen/go.mod
git -C "$dir" fetch --quiet --depth 1 --filter=blob:none origin "$commit"
git -C "$dir" checkout --quiet --detach "$commit"
test "$(git -C "$dir" rev-parse HEAD)" == "$commit"
