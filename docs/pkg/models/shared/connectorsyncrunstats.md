# ConnectorSyncRunStats

The ConnectorSyncRunStats message.


## Fields

| Field                                 | Type                                  | Required                              | Description                           |
| ------------------------------------- | ------------------------------------- | ------------------------------------- | ------------------------------------- |
| `C1zFileSizeBytes`                    | `*int64`                              | :heavy_minus_sign:                    | The c1zFileSizeBytes field.           |
| `Entitlements`                        | `*int64`                              | :heavy_minus_sign:                    | The entitlements field.               |
| `EntitlementsByResourceType`          | map[string]`int64`                    | :heavy_minus_sign:                    | The entitlementsByResourceType field. |
| `Grants`                              | `*int64`                              | :heavy_minus_sign:                    | The grants field.                     |
| `GrantsByResourceType`                | map[string]`int64`                    | :heavy_minus_sign:                    | The grantsByResourceType field.       |
| `ResourceTypes`                       | `*int64`                              | :heavy_minus_sign:                    | The resourceTypes field.              |
| `Resources`                           | `*int64`                              | :heavy_minus_sign:                    | The resources field.                  |
| `ResourcesByResourceType`             | map[string]`int64`                    | :heavy_minus_sign:                    | The resourcesByResourceType field.    |