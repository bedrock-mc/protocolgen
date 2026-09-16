# Mojang U6 metadata

The importer understands `x-enum-binary-value`, introduced in Mojang's
`v1.26.60-preview.21` schemas. This is an array parallel to `enum`, not an
array of positions: values may be negative, nonconsecutive, or shared by
several names. For example, `GameType` contains both `Survival = 0` and
`WorldDefault = 0`.

Numeric uses select `Enum-as-Value`. The importer combines serialization
flags from the referenced definition and the field using it, including
`Compression`, and keeps the field's underlying type override. With U6
metadata, a named enum without `Enum-as-Value` is a length-prefixed string.
Malformed numeric mappings remain unresolved; the importer never substitutes
array positions. Older unmapped enums retain their unresolved claims and
version-pinned adjudications.

The manifest preserves aliases. Go emits a constant for every name. Rust
emits every named variant, encodes aliases to the same number, and decodes
each number to the first alias in source order. Normalization preserves that
order among equal values. Unknown numeric values still round-trip through
the existing open-enum codecs.

The changelog and update guide also consume the official mapping. A newly
published mapping is reported as added metadata, rather than compared with
previously guessed positions and presented as a wire change.

Both tools treat fields with a default as always present, even when that
default is `false`, `0`, or `null`. The update guide also carries Mojang
`x-runtime-constraint-description` text into field comments without inventing
validation code.

## Variant selectors

For a wire `oneOf` marked `x-control-value-type: uint32`, the selector is
always an unsigned VarUInt32. Alternative positions are zero-based: a list
of three alternatives uses values 0, 1 and 2. Write that index first, then
write all fields in the chosen alternative. The selector does not need a
`Compression` flag or per-alternative `x-ordinal-index` annotations. Enum
binary values and fields inside an alternative do not replace this index.

For example, [PlayerListPacketPayload](https://github.com/Mojang/bedrock-protocol-docs/blob/19e25de129227373e97dfa340af93ccf12c0feb9/json/PlayerListPacketPayload.json)
puts Remove first and Add second. Their outer indices are 0 and 1. The
branches still carry their separate Action byte, whose values are 1 and 0
respectively. Reading the outer selector chooses the payload layout inside
this packet; it does not select the next packet.

Plain JSON `oneOf` is not enough to establish a binary selector. For example,
[Color255RGB](https://github.com/Mojang/bedrock-protocol-docs/blob/19e25de129227373e97dfa340af93ccf12c0feb9/json/Color255RGB.json)
offers hex-string and numeric-array representations without a wire-selector
marker. Such alternatives remain unresolved. Unsupported selector types and
empty or malformed alternatives also remain unresolved.

The newly reached schemas include arrays with `maxProperties`, an object-only
JSON Schema keyword. The importer ignores that misplaced keyword rather than
treating it as an array limit. Valid `maxItems` metadata still applies.

## Verification against published schemas

The implementation was checked against Mojang commit
[`19e25de129227373e97dfa340af93ccf12c0feb9`](https://github.com/Mojang/bedrock-protocol-docs/tree/19e25de129227373e97dfa340af93ccf12c0feb9/json),
whose schemas identify Minecraft `1.26.60-beta.23`, protocol `2208`.
The source tree digest for its `json` directory is
`sha256:c5a06f6dfbffa071f39c37de9af28b70d37a3283316c92f3cc2e2ca1a6a111d0`.

Raw ingestion produces 745 top-level field claims. Enum metadata resolved
146 numeric enum occurrences and 19 string enum occurrences, reducing
unresolved occurrences from 238 to 73. Applying the variant-selector rule
resolves all 44 previous selector gaps and exposes 50 previously hidden
unresolved nodes inside their alternatives. The current total is **79**:
50 ItemStackNetIdVariant occurrences, 18 untyped payloads, 10 colour
alternatives, and one WebToken self-reference. These are nested occurrences,
not 79 distinct packet definitions. This remains an ingestion audit, not a
fully reconciled or shipping 1.26.60 snapshot.

Regenerated 1.26.40 and 1.26.50 snapshots need nine fewer adjudications each:
CommandBlockUpdate.Target, BookEdit.Operation, and seven sound-update fields
now have two agreeing source claims. Remaining changed adjudication contexts
retain identical selected-source fingerprints. The released packet wire
layouts and generated public type names are unchanged. The 1.26.44 hotfix
inherits the same metadata updates from 1.26.40.

The released snapshots retain their existing source pins. New optional-field
metadata must be reconciled with evidence for the same protocol; it must not
be copied into older snapshots as an assumed correction. Mojang still omits
defaulted fields from `required`, so a missing name alone does not prove an
optional presence byte.
