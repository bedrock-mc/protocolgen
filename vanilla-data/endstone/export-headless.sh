#!/usr/bin/env bash
set -euo pipefail
export LC_ALL=C

# Build the exact Endstone revision pinned by vanilla-source.json, run its
# in-process BDS exporter without a GUI, and authenticate every output before
# handing it to the Go packet-capture bot.

usage() {
  echo "usage: $0 -source SOURCE_JSON -bds-dir BDS_DIR -out OUTPUT_DIR -build-dir BUILD_DIR" >&2
  exit 2
}

source_file=
bds_dir=
output_dir=
build_dir=
while (($#)); do
  case "$1" in
    -source) source_file=${2-}; shift 2 ;;
    -bds-dir) bds_dir=${2-}; shift 2 ;;
    -out) output_dir=${2-}; shift 2 ;;
    -build-dir) build_dir=${2-}; shift 2 ;;
    *) usage ;;
  esac
done
test -n "$source_file" && test -n "$bds_dir" && test -n "$output_dir" && test -n "$build_dir" || usage
workspace=${GITHUB_WORKSPACE:-$(cd "$(dirname "$0")/../.." && pwd)}

endstone_repo=$(jq -er '.endstone.repository' "$source_file")
endstone_revision=$(jq -er '.endstone.revision' "$source_file")
endstone_bds=$(jq -er '.endstone.bds_version' "$source_file")
endstone_patch_sha=$(jq -er '.endstone.headless_patch_sha256' "$source_file")
bds_version=$(jq -er '.bds.version' "$source_file")
target_version=$(jq -er '.minecraft_version' "$source_file")
protocol_version=$(jq -er '.protocol_version' "$source_file")
test "$endstone_bds" = "$bds_version"
test "$endstone_revision" != "" && test "${#endstone_revision}" -eq 40
test "$endstone_patch_sha" = "sha256:cba32423594101ee2cf19aec4be52c253998ccac3f1afa15a6c42fb88051b28f"

repo_dir="$build_dir/endstone"
rm -rf "$repo_dir"
git clone --no-checkout "$endstone_repo" "$repo_dir"
git -C "$repo_dir" checkout --detach "$endstone_revision"
test "$(git -C "$repo_dir" rev-parse HEAD)" = "$endstone_revision"

actual_bds_version=$(sed -n 's/^version = "\([^"]*\)"/\1/p' "$repo_dir/scripts/configs/linux.toml" | head -n 1)
test "$actual_bds_version" = "$bds_version"

patch_file="$workspace/vanilla-data/endstone/headless-export.patch"
test "sha256:$(sha256sum "$patch_file" | cut -d' ' -f1)" = "$endstone_patch_sha"
git -C "$repo_dir" apply --check "$patch_file"
git -C "$repo_dir" apply "$patch_file"
cp "$workspace/vanilla-data/endstone/headless_vanilla_data.h" "$repo_dir/src/endstone/core/devtools/headless_vanilla_data.h"
cp "$workspace/vanilla-data/endstone/headless_vanilla_data.cpp" "$repo_dir/src/endstone/core/devtools/headless_vanilla_data.cpp"

python3 -m pip install --disable-pip-version-check --user conan
conan install "$repo_dir" --build=missing -s build_type=Release
(cd "$repo_dir" && cmake --preset conan-release)
(cd "$repo_dir" && cmake --build --preset conan-release --parallel)
python3 -m pip install --disable-pip-version-check --user "$repo_dir" -C "build-dir=$repo_dir/build"

rm -rf "$output_dir"
mkdir -p "$output_dir"
export ENDSTONE_VANILLA_DATA_OUT="$output_dir"
export ENDSTONE_VANILLA_DATA_TIMEOUT_MS=300000
export LD_LIBRARY_PATH="$repo_dir/build/Release:$repo_dir/build/Release/generators:${LD_LIBRARY_PATH-}"

(
  cd "$bds_dir"
  python3 -m endstone --server-folder "$bds_dir" --yes
) >"$build_dir/endstone.log" 2>&1 &
server_pid=$!
cleanup() {
  kill "$server_pid" 2>/dev/null || true
  wait "$server_pid" 2>/dev/null || true
}
trap cleanup EXIT

for attempt in $(seq 1 300); do
  test -f "$output_dir/.ready" && break
  if ! kill -0 "$server_pid" 2>/dev/null; then
    cat "$build_dir/endstone.log" >&2
    exit 1
  fi
  sleep 1
done
test -f "$output_dir/.ready"
rm "$output_dir/.ready"
test -s "$output_dir/block_palette.nbt"

files=()
for file in "$output_dir"/*; do
  test -f "$file" || continue
  files+=("$file")
done
test "${#files[@]}" -gt 0

file_json='[]'
for file in "${files[@]}"; do
  name=$(basename "$file")
  bytes=$(wc -c <"$file" | tr -d ' ')
  sha=$(sha256sum "$file" | cut -d' ' -f1)
  file_json=$(jq -c --arg file "$name" --arg sha "sha256:$sha" --argjson bytes "$bytes" '. + [{file: $file, bytes: $bytes, sha256: $sha}]' <<<"$file_json")
done
jq -n \
  --argjson schema_version 1 \
  --arg minecraft_version "$target_version" \
  --argjson protocol_version "$protocol_version" \
  --arg bds_version "$bds_version" \
  --arg repository "$endstone_repo" \
  --arg revision "$endstone_revision" \
  --arg patch "$endstone_patch_sha" \
  --argjson files "$file_json" \
  '{schema_version: $schema_version, target: {minecraft_version: $minecraft_version, protocol_version: $protocol_version}, bds_version: $bds_version, endstone: {repository: $repository, revision: $revision, bds_version: $bds_version, headless_patch_sha256: $patch}, files: $files}' \
  >"$output_dir/endstone-export.json"

echo "headless Endstone vanilla data exported for $target_version ($bds_version)"
