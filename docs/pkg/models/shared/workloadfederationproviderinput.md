# WorkloadFederationProviderInput

WorkloadFederationProvider represents a tenant-level OIDC issuer registration.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Description`                                                                    | **string*                                                                        | :heavy_minus_sign:                                                               | A description of what this provider is for.                                      |
| `Disabled`                                                                       | **bool*                                                                          | :heavy_minus_sign:                                                               | Whether the provider is disabled. Disabled providers reject all token exchanges. |
| `DisplayName`                                                                    | **string*                                                                        | :heavy_minus_sign:                                                               | The display name of the provider.                                                |