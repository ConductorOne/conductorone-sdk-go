# AppUserMergeCelRule

A rule that matches when both CEL expressions produce the same key. Each expression
 reads the variable `account` and returns a string or a list of strings; an empty key
 never matches.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ExistingAppUserKeyCel`                                           | `*string`                                                         | :heavy_minus_sign:                                                | The expression evaluated against each existing account. Required. |
| `IncomingKeyCel`                                                  | `*string`                                                         | :heavy_minus_sign:                                                | The expression evaluated against the incoming account record.     |