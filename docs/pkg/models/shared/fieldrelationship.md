# FieldRelationship

FieldRelationships can be used during form validation, or they can represent
 information that is necessary to when it comes to visually rendering the form

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - requiredTogether
  - atLeastOne
  - mutuallyExclusive



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AtLeastOne`                                                                 | [*shared.AtLeastOne](../../../pkg/models/shared/atleastone.md)               | :heavy_minus_sign:                                                           | The AtLeastOne message.                                                      |
| `MutuallyExclusive`                                                          | [*shared.MutuallyExclusive](../../../pkg/models/shared/mutuallyexclusive.md) | :heavy_minus_sign:                                                           | The MutuallyExclusive message.                                               |
| `RequiredTogether`                                                           | [*shared.RequiredTogether](../../../pkg/models/shared/requiredtogether.md)   | :heavy_minus_sign:                                                           | The RequiredTogether message.                                                |
| `FieldNames`                                                                 | []*string*                                                                   | :heavy_minus_sign:                                                           | The names of the fields that share this relationship                         |