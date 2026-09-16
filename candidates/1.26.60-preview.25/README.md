# 1.26.60 preview 25 candidate inputs

These inputs target protocol 2211. They contain pinned sources and reviewed
schema corrections, not a release manifest or generated packet codecs.

The corrections restore missing CompoundTag annotations for 16 packet fields,
describe item-stack IDs as signed 32-bit values with compression selected at
each use, and describe colour attributes as four fixed 32-bit RGBA components.
They also align DeathInfo's nested message fields with Endstone's flat layout
and describe CraftingData recipe ingredients as a string map followed by two
compressed signed integers. Item-stack requests keep their separate ingredient
representation: a descriptor union followed by a 16-bit count.
Each operation checks the original schema node's fingerprint and records
evidence from the exact target documentation.

Agreement after a correction is reviewed agreement, not independent confirmation
from the unmodified Mojang schema. NBT type annotations do not choose an encoding
or prove root framing. LevelEventGeneric's special compound payload remains
unresolved. Reconciliation, packet directions, NBT encoding overlays, consumer
metadata and independent target verification must be completed before emission.

To ingest the pinned source checkouts:

```sh
make ingest-1.26.60 \
  MOJANG_DIR=/path/to/bedrock-protocol-docs/json \
  ENDSTONE_DIR=/path/to/endstone-protocol-docs
```

The command writes source claims to `/tmp/protocolgen-1.26.60-preview.25-*-claims.json`.
CI runs this step to check source hashes and correction fingerprints. It does
not certify this candidate as a complete protocol implementation.
