# ConnectorCatalogServiceConfigurationSchemaRequest

ConnectorCatalogServiceConfigurationSchemaRequest is the request for retrieving a connector's configuration schema.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `AppID`                                                                                 | `*string`                                                                               | :heavy_minus_sign:                                                                      | The ID of the app associated with the connector. Optional.                              |
| `CatalogID`                                                                             | `*string`                                                                               | :heavy_minus_sign:                                                                      | The catalog entry ID identifying the connector type.                                    |
| `ConnectorID`                                                                           | `*string`                                                                               | :heavy_minus_sign:                                                                      | The ID of an existing connector to retrieve its current configuration schema. Optional. |