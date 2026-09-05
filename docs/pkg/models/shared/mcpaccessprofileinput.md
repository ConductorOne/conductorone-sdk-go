# MCPAccessProfileInput

MCPAccessProfile represents an admin-curated grouping of MCP tools.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `AppEntitlementID`                                     | `*string`                                              | :heavy_minus_sign:                                     | The ID of the AppEntitlement created for this profile. |
| `AppID`                                                | `*string`                                              | :heavy_minus_sign:                                     | App identifier (app that owns the connector).          |
| `ConnectorID`                                          | `*string`                                              | :heavy_minus_sign:                                     | Connector identifier.                                  |
| `CreatedAt`                                            | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | N/A                                                    |
| `DeletedAt`                                            | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | N/A                                                    |
| `Description`                                          | `*string`                                              | :heavy_minus_sign:                                     | Description of what access this profile grants.        |
| `DisplayName`                                          | `*string`                                              | :heavy_minus_sign:                                     | Display name for the profile.                          |
| `ID`                                                   | `*string`                                              | :heavy_minus_sign:                                     | Unique identifier for this access profile.             |
| `ToolCount`                                            | `*int`                                                 | :heavy_minus_sign:                                     | The number of tools currently bound to this profile.   |
| `UpdatedAt`                                            | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | N/A                                                    |