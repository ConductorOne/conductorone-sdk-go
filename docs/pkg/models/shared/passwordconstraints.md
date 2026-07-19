# PasswordConstraints

PasswordConstraints sets the complexity rules a user's password must satisfy.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `CheckBreached`                                                 | `*bool`                                                         | :heavy_minus_sign:                                              | Reject passwords found in known-breach corpora.                 |
| `HistoryDepth`                                                  | `*int`                                                          | :heavy_minus_sign:                                              | Number of previous passwords to remember and disallow reuse of. |
| `MinLength`                                                     | `*int`                                                          | :heavy_minus_sign:                                              | Minimum length, in characters.                                  |
| `RequireMixedCase`                                              | `*bool`                                                         | :heavy_minus_sign:                                              | Require both uppercase and lowercase letters.                   |
| `RequireNumber`                                                 | `*bool`                                                         | :heavy_minus_sign:                                              | Require at least one digit.                                     |
| `RequireSymbol`                                                 | `*bool`                                                         | :heavy_minus_sign:                                              | Require at least one symbol.                                    |