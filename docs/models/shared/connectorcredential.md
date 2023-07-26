# ConnectorCredential

 ConnectorCredential is used by a connector to authenticate with conductor one.



## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `AppID`                                                               | **string*                                                             | :heavy_minus_sign:                                                    |  The appId of the app the connector is attached to.<br/>              |
| `ClientID`                                                            | **string*                                                             | :heavy_minus_sign:                                                    |  The client id of the ConnectorCredential.<br/>                       |
| `ConnectorID`                                                         | **string*                                                             | :heavy_minus_sign:                                                    |  The connectorId of the connector the credential is associated with.<br/> |
| `CreatedAt`                                                           | [*time.Time](https://pkg.go.dev/time#Time)                            | :heavy_minus_sign:                                                    | N/A                                                                   |
| `DeletedAt`                                                           | [*time.Time](https://pkg.go.dev/time#Time)                            | :heavy_minus_sign:                                                    | N/A                                                                   |
| `DisplayName`                                                         | **string*                                                             | :heavy_minus_sign:                                                    |  The display name of the ConnectorCredential.<br/>                    |
| `ExpiresTime`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                            | :heavy_minus_sign:                                                    | N/A                                                                   |
| `ID`                                                                  | **string*                                                             | :heavy_minus_sign:                                                    |  The id of the ConnectorCredential.<br/>                              |
| `LastUsedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                            | :heavy_minus_sign:                                                    | N/A                                                                   |
| `UpdatedAt`                                                           | [*time.Time](https://pkg.go.dev/time#Time)                            | :heavy_minus_sign:                                                    | N/A                                                                   |