# ServicePrincipalServiceCreateCredentialResponse

The ServicePrincipalServiceCreateCredentialResponse message.


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ClientSecret`                                                                                 | `*string`                                                                                      | :heavy_minus_sign:                                                                             | The client secret. Shown exactly once at creation -- cannot be retrieved again.                |
| `Credential`                                                                                   | [*shared.ServicePrincipalCredential](../../../pkg/models/shared/serviceprincipalcredential.md) | :heavy_minus_sign:                                                                             | N/A                                                                                            |