# PerCredentialDuration

PerCredentialDuration overrides session lifetimes for sessions established
 with a particular credential type — stronger credentials can earn longer
 sessions.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AccessTokenTTLSeconds`                                                | `*int`                                                                 | :heavy_minus_sign:                                                     | Access-token lifetime for this credential type, in seconds.            |
| `CredentialType`                                                       | [*shared.CredentialType](../../../pkg/models/shared/credentialtype.md) | :heavy_minus_sign:                                                     | The credentialType field.                                              |
| `MaxSessionDurationSeconds`                                            | `*int`                                                                 | :heavy_minus_sign:                                                     | Maximum total session duration for this credential type, in seconds.   |