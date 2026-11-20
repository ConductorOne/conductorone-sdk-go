# DirectoryMergeCustomRule

DirectoryMergeCustomRule is a custom rule defined by a pair of CEL expressions whose keys are compared
 exactly as returned.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `AppUserKeyCel`                                                                  | `*string`                                                                        | :heavy_minus_sign:                                                               | A CEL expression over `account` returning a string or list of strings. Required. |
| `Description`                                                                    | `*string`                                                                        | :heavy_minus_sign:                                                               | Optional.                                                                        |
| `DisplayName`                                                                    | `*string`                                                                        | :heavy_minus_sign:                                                               | Required. Unique among the directory's rules, ignoring case.                     |
| `UserKeyCel`                                                                     | `*string`                                                                        | :heavy_minus_sign:                                                               | A CEL expression over `subject` returning a string or list of strings. Required. |