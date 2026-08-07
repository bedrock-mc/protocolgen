# Differential codec harness

This module compares protocol 2168 payloads between the generated Go codec and the pinned gophertunnel fork at `be6713da4dc051a4197f897d04835e89e9c54321`.

Pool parity accepts five deprecated oracle-only IDs that are still registered by gophertunnel but absent from the manifest: 55 `AdventureSettings`, 117 `ScriptCustomEvent`, 163 `FilterText`, 173 `PhotoInfoRequest`, and 197 `ClientCheatAbility`. The canonical manifest also contains server-only ID 16 `ServerPlayerPostMovePosition`, which the pinned oracle reserves and does not register.

Run `go test ./...` for the committed corpus. Run `go test -fuzz=FuzzDifferential -fuzztime=10m -promote` to fuzz and promote successful cases.
