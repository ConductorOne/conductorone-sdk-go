# FieldRelationship

FieldRelationships can be used during form validation, or they can represent
 information that is necessary to when it comes to visually rendering the form

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - requiredTogether
  - atLeastOne
  - mutuallyExclusive
  - dependentOn



## Fields

| Field                                                                                                                | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `AtLeastOne`                                                                                                         | [*shared.AtLeastOne](../../../pkg/models/shared/atleastone.md)                                                       | :heavy_minus_sign:                                                                                                   | The AtLeastOne message.                                                                                              |
| `DependentOn`                                                                                                        | [*shared.DependentOn](../../../pkg/models/shared/dependenton.md)                                                     | :heavy_minus_sign:                                                                                                   | DependentOn means the fields in field_names are only valid if all fields<br/> in dependency_field_names are also present |
| `MutuallyExclusive`                                                                                                  | [*shared.MutuallyExclusive](../../../pkg/models/shared/mutuallyexclusive.md)                                         | :heavy_minus_sign:                                                                                                   | The MutuallyExclusive message.                                                                                       |
| `RequiredTogether`                                                                                                   | [*shared.RequiredTogether](../../../pkg/models/shared/requiredtogether.md)                                           | :heavy_minus_sign:                                                                                                   | The RequiredTogether message.                                                                                        |
| `FieldNames`                                                                                                         | []*string*                                                                                                           | :heavy_minus_sign:                                                                                                   | The names of the fields that share this relationship                                                                 |