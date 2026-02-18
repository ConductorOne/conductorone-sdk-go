# ServicePrincipalInput

ServicePrincipal represents a tenant-managed non-human identity.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `User`                                                                                  | [*shared.UserInput](../../../pkg/models/shared/userinput.md)                            | :heavy_minus_sign:                                                                      | The User object provides all of the details for an user, as well as some configuration. |
| `DisplayName`                                                                           | **string*                                                                               | :heavy_minus_sign:                                                                      | The display name of the service principal.                                              |