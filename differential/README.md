# Differential codec harness

This module compares protocol 2168 payloads between the generated Go codec and the pinned gophertunnel fork at `be6713da4dc051a4197f897d04835e89e9c54321`.

Pool parity accepts five deprecated oracle-only IDs that are still registered by gophertunnel but absent from the manifest: 55 `AdventureSettings`, 117 `ScriptCustomEvent`, 163 `FilterText`, 173 `PhotoInfoRequest`, and 197 `ClientCheatAbility`. The canonical manifest also contains server-only ID 16 `ServerPlayerPostMovePosition`, which the pinned oracle reserves and does not register.

Run `go test ./...` for the committed corpus. Run `go test -fuzz=FuzzDifferential -fuzztime=10m -promote` to fuzz and promote successful cases.

## Comparison rules

The contract is one-directional: every input the pinned oracle accepts must
decode here and re-encode byte-identically. Three structural allowances,
each verified against oracle source:

- **Bool canonicalization**: fields the manifest types as bool but the oracle
  types as passthrough uint8 re-encode as canonical 0/1 here. Only reachable
  with non-canonical bool bytes, which vanilla peers never send.
- **Oracle truncation leniency**: the oracle zero-pads truncated fixed-width
  reads (unchecked Read counts), has no varint byte cap, and skips NBT
  validation. The generated runtime rejects these malformed inputs.
- **Accepted divergences** (`accepted-divergences.json`): per-packet classes
  where the oracle rewrites decoded state on encode — derived enum names and
  ordinals, presence flags derived from values, zeroed reserved or
  inactive-union fields, empty-NBT compound insertion, lossy sound-position
  conversion. The generated codec preserves wire bytes in every adjudicated
  case.
