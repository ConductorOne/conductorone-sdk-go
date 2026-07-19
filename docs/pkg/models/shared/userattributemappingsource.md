# UserAttributeMappingSource

The UserAttributeMappingSource message.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AppID`                                                             | `*string`                                                           | :heavy_minus_sign:                                                  | The appId field.                                                    |
| `AppUserID`                                                         | `*string`                                                           | :heavy_minus_sign:                                                  | The appUserId field.                                                |
| `AppUserProfileAttributeKey`                                        | `*string`                                                           | :heavy_minus_sign:                                                  | The appUserProfileAttributeKey field.                               |
| `Priority`                                                          | `*int64`                                                            | :heavy_minus_sign:                                                  | Lower number = higher precedence; sources[0] is the winning source. |
| `UserAttributeMappingID`                                            | `*string`                                                           | :heavy_minus_sign:                                                  | The userAttributeMappingId field.                                   |
| `Value`                                                             | `*string`                                                           | :heavy_minus_sign:                                                  | The value field.                                                    |