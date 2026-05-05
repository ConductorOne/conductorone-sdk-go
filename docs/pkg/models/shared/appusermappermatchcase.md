# AppUserMapperMatchCase

AppUserMapperMatchCase defines a single matching rule for uplift account mapping.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `AppUserKeyCel`                                                      | `*string`                                                            | :heavy_minus_sign:                                                   | CEL expression evaluated against an AppUser to produce match key(s). |
| `UserKeyCel`                                                         | `*string`                                                            | :heavy_minus_sign:                                                   | CEL expression evaluated against a User to produce match key(s).     |