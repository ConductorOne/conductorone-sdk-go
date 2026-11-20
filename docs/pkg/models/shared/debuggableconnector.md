# DebuggableConnector

DebuggableConnector contains the minimal metadata needed to select a diagnostic target.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `AppID`                                                        | `*string`                                                      | :heavy_minus_sign:                                             | Application containing the connector.                          |
| `Authored`                                                     | `*bool`                                                        | :heavy_minus_sign:                                             | Whether this is a live tenant-authored installation.           |
| `CatalogID`                                                    | `*string`                                                      | :heavy_minus_sign:                                             | Catalog entry associated with the connector.                   |
| `ConnectorID`                                                  | `*string`                                                      | :heavy_minus_sign:                                             | Connector to pass to diagnostic reads.                         |
| `DisplayName`                                                  | `*string`                                                      | :heavy_minus_sign:                                             | Connector display name.                                        |
| `LastCompletedAt`                                              | [*time.Time](https://pkg.go.dev/time#Time)                     | :heavy_minus_sign:                                             | N/A                                                            |
| `SyncStatus`                                                   | [*shared.SyncStatus](../../../pkg/models/shared/syncstatus.md) | :heavy_minus_sign:                                             | Current connector sync status.                                 |