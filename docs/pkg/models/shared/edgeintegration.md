# EdgeIntegration

EdgeIntegration is reusable application-owned Edge capability configuration.
 Its typed configuration and use site determine its role; Socket package
 authorization is the first supported type. It stores only credential
 references and never contains or returns a secret value.
 Creation requires display_name and socket_packages. Updates may omit fields
 not listed in update_mask.

This message contains a oneof named configuration. Only a single field of the following list may be set at a time:
  - socketPackages



## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `AppID`                                                                                              | `*string`                                                                                            | :heavy_minus_sign:                                                                                   | Application owning this integration.                                                                 |
| `ConfigurationGeneration`                                                                            | `*string`                                                                                            | :heavy_minus_sign:                                                                                   | Changes only when typed configuration or its credential reference changes.                           |
| `CreatedAt`                                                                                          | [*time.Time](https://pkg.go.dev/time#Time)                                                           | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `DisplayName`                                                                                        | `*string`                                                                                            | :heavy_minus_sign:                                                                                   | Human-readable integration name. Required on creation; may be omitted<br/> when updating other fields. |
| `ID`                                                                                                 | `*string`                                                                                            | :heavy_minus_sign:                                                                                   | Integration identifier. Assigned on creation.                                                        |
| `ReferenceCount`                                                                                     | `*string`                                                                                            | :heavy_minus_sign:                                                                                   | Number of saved checks referencing this integration, including disabled checks.                      |
| `SocketPackages`                                                                                     | [*shared.EdgeSocketPackagesIntegration](../../../pkg/models/shared/edgesocketpackagesintegration.md) | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `UpdatedAt`                                                                                          | [*time.Time](https://pkg.go.dev/time#Time)                                                           | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `VaultSecretRef`                                                                                     | [*shared.EdgeIntegrationVaultSecretRef](../../../pkg/models/shared/edgeintegrationvaultsecretref.md) | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |