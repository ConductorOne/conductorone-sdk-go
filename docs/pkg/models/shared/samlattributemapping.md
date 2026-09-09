# SAMLAttributeMapping

SAMLAttributeMapping releases one user attribute to the service provider
 as one Attribute in the assertion's AttributeStatement.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `FriendlyName`                                                                     | `*string`                                                                          | :heavy_minus_sign:                                                                 | Optional FriendlyName, for service providers that display it.                      |
| `Name`                                                                             | `string`                                                                           | :heavy_check_mark:                                                                 | The Name attribute, dictated by the service provider.                              |
| `NameFormat`                                                                       | [*shared.NameFormat](../../../pkg/models/shared/nameformat.md)                     | :heavy_minus_sign:                                                                 | The NameFormat attribute.                                                          |
| `UserAttributeMappingID`                                                           | `string`                                                                           | :heavy_check_mark:                                                                 | The user attribute mapping that resolves the value, including its fallback<br/> chain. |