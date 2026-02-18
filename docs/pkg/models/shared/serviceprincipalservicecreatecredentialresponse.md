# ServicePrincipalServiceCreateCredentialResponse

The ServicePrincipalServiceCreateCredentialResponse message.


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ServicePrincipalCredential`                                                                   | [*shared.ServicePrincipalCredential](../../../pkg/models/shared/serviceprincipalcredential.md) | :heavy_minus_sign:                                                                             | ServicePrincipalCredential represents a client credential for a service principal.             |
| `ClientSecret`                                                                                 | **string*                                                                                      | :heavy_minus_sign:                                                                             | The client secret. Shown exactly once at creation -- cannot be retrieved again.                |