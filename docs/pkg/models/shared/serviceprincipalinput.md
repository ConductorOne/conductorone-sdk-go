# ServicePrincipalInput

ServicePrincipal represents a tenant-managed non-human identity.


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `DisplayName`                                                                                  | `*string`                                                                                      | :heavy_minus_sign:                                                                             | The display name of the service principal.                                                     |
| `ObjectPermissions`                                                                            | [*shared.UserActorObjectPermissions](../../../pkg/models/shared/useractorobjectpermissions.md) | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `User`                                                                                         | [*shared.UserInput](../../../pkg/models/shared/userinput.md)                                   | :heavy_minus_sign:                                                                             | N/A                                                                                            |