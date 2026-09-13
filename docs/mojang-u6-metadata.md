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

## Verification against published schemas

The implementation was checked against Mojang commit
[`19e25de129227373e97dfa340af93ccf12c0feb9`](https://github.com/Mojang/bedrock-protocol-docs/tree/19e25de129227373e97dfa340af93ccf12c0feb9/json),
whose schemas identify Minecraft `1.26.60-beta.23`, protocol `2208`.
The source tree digest for its `json` directory is
`sha256:c5a06f6dfbffa071f39c37de9af28b70d37a3283316c92f3cc2e2ca1a6a111d0`.

Raw ingestion produces 745 top-level field claims. Compared with the previous
importer, 146 numeric enum occurrences and 19 string enum occurrences become
resolved; unresolved node occurrences fall from 238 to 73. These counts
include nested schemas reached from packet fields. This is an ingestion
audit, not a fully reconciled or shipping 1.26.60 snapshot.

The released snapshots retain their existing source pins. New optional-field
metadata must be reconciled with evidence for the same protocol; it must not
be copied into older snapshots as an assumed correction. Mojang still omits
defaulted fields from `required`, so a missing name alone does not prove an
optional presence byte.
