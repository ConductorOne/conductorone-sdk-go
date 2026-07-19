# EmailOTPConstraints

EmailOTPConstraints configures one-time codes delivered by email.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `CodeLength`                                                         | `*int`                                                               | :heavy_minus_sign:                                                   | Number of digits in each code.                                       |
| `MaxAttempts`                                                        | `*int`                                                               | :heavy_minus_sign:                                                   | Number of incorrect attempts allowed before the code is invalidated. |
| `TTLSeconds`                                                         | `*int`                                                               | :heavy_minus_sign:                                                   | How long a code remains valid, in seconds.                           |