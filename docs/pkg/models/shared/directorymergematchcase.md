# DirectoryMergeMatchCase

DirectoryMergeMatchCase defines a pair of CEL key extractors for matching.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `AppUserKeyCel`                                                      | `*string`                                                            | :heavy_minus_sign:                                                   | CEL expression evaluated against an AppUser to produce match key(s). |
| `UserKeyCel`                                                         | `*string`                                                            | :heavy_minus_sign:                                                   | CEL expression evaluated against a User to produce match key(s).     |