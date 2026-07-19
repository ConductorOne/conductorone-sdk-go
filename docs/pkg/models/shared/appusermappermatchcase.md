# AppUserMapperMatchCase

AppUserMapperMatchCase defines a single matching rule for uplift account mapping.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `AppID`                                                              | `*string`                                                            | :heavy_minus_sign:                                                   | The app this match case belongs to.                                  |
| `AppUserKeyCel`                                                      | `*string`                                                            | :heavy_minus_sign:                                                   | CEL expression evaluated against an AppUser to produce match key(s). |
| `CaseIndex`                                                          | `*int64`                                                             | :heavy_minus_sign:                                                   | The ordered index of this match case within the mapper.              |
| `UserKeyCel`                                                         | `*string`                                                            | :heavy_minus_sign:                                                   | CEL expression evaluated against a User to produce match key(s).     |