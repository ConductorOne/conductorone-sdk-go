# TOTPConstraints

TOTPConstraints configures authenticator-app one-time codes (RFC 6238).


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `CodeLength`                                                        | `*int`                                                              | :heavy_minus_sign:                                                  | Number of digits in each code.                                      |
| `PeriodSeconds`                                                     | `*int`                                                              | :heavy_minus_sign:                                                  | How often a new code is generated, in seconds (typically 30 or 60). |
| `SkewTolerance`                                                     | `*int`                                                              | :heavy_minus_sign:                                                  | How many adjacent time windows to accept, to tolerate clock drift.  |