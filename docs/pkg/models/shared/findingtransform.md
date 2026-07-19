# FindingTransform

FindingTransform is a single mutation applied to a finding by a matched
 transformation rule. Structured ops are the v1 authoring surface; a future
 raw-CEL arm continues numbering at 103.

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - setSeverity
  - setTags
  - removeTags



## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `RemoveTags`                                                     | [*shared.RemoveTags](../../../pkg/models/shared/removetags.md)   | :heavy_minus_sign:                                               | N/A                                                              |
| `SetSeverity`                                                    | [*shared.SetSeverity](../../../pkg/models/shared/setseverity.md) | :heavy_minus_sign:                                               | N/A                                                              |
| `SetTags`                                                        | [*shared.SetTags](../../../pkg/models/shared/settags.md)         | :heavy_minus_sign:                                               | N/A                                                              |