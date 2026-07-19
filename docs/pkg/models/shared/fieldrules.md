# FieldRules

FieldRules encapsulates the rules for each type of field. Depending on the
 field, the correct set should be used to ensure proper validations.

This message contains a oneof named type. Only a single field of the following list may be set at a time:
  - float
  - double
  - int32
  - int64
  - uint32
  - uint64
  - sint32
  - sint64
  - fixed32
  - fixed64
  - sfixed32
  - sfixed64
  - bool
  - string
  - bytes
  - enum
  - repeated
  - map
  - any
  - duration
  - timestamp



## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Any`                                                                  | [*shared.AnyRules](../../../pkg/models/shared/anyrules.md)             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Bool`                                                                 | [*shared.BoolRules](../../../pkg/models/shared/boolrules.md)           | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Bytes`                                                                | [*shared.BytesRules](../../../pkg/models/shared/bytesrules.md)         | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Double`                                                               | [*shared.DoubleRules](../../../pkg/models/shared/doublerules.md)       | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Duration`                                                             | [*shared.DurationRules](../../../pkg/models/shared/durationrules.md)   | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Enum`                                                                 | [*shared.EnumRules](../../../pkg/models/shared/enumrules.md)           | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Fixed32`                                                              | [*shared.Fixed32Rules](../../../pkg/models/shared/fixed32rules.md)     | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Fixed64`                                                              | [*shared.Fixed64Rules](../../../pkg/models/shared/fixed64rules.md)     | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Float`                                                                | [*shared.FloatRules](../../../pkg/models/shared/floatrules.md)         | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Int32`                                                                | [*shared.Int32Rules](../../../pkg/models/shared/int32rules.md)         | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Int64`                                                                | [*shared.Int64Rules](../../../pkg/models/shared/int64rules.md)         | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Map`                                                                  | [*shared.MapRules](../../../pkg/models/shared/maprules.md)             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Message`                                                              | [*shared.MessageRules](../../../pkg/models/shared/messagerules.md)     | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Repeated`                                                             | [*shared.RepeatedRules](../../../pkg/models/shared/repeatedrules.md)   | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Sfixed32`                                                             | [*shared.SFixed32Rules](../../../pkg/models/shared/sfixed32rules.md)   | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Sfixed64`                                                             | [*shared.SFixed64Rules](../../../pkg/models/shared/sfixed64rules.md)   | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Sint32`                                                               | [*shared.SInt32Rules](../../../pkg/models/shared/sint32rules.md)       | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Sint64`                                                               | [*shared.SInt64Rules](../../../pkg/models/shared/sint64rules.md)       | :heavy_minus_sign:                                                     | N/A                                                                    |
| `String`                                                               | [*shared.StringRules](../../../pkg/models/shared/stringrules.md)       | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Timestamp`                                                            | [*shared.TimestampRules](../../../pkg/models/shared/timestamprules.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Uint32`                                                               | [*shared.UInt32Rules](../../../pkg/models/shared/uint32rules.md)       | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Uint64`                                                               | [*shared.UInt64Rules](../../../pkg/models/shared/uint64rules.md)       | :heavy_minus_sign:                                                     | N/A                                                                    |