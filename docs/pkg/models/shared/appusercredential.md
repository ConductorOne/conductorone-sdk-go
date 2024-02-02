# AppUserCredential

Application User that represents an account in the application.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `EncryptedData`                                                        | [*shared.EncryptedData](../../../pkg/models/shared/encrypteddata.md)   | :heavy_minus_sign:                                                     | EncryptedData is a message that contains encrypted bytes and metadata. |
| `AppID`                                                                | **string*                                                              | :heavy_minus_sign:                                                     | The ID of the application.                                             |
| `AppUserID`                                                            | **string*                                                              | :heavy_minus_sign:                                                     | A unique idenditfier of the application user.                          |
| `ConnectorID`                                                          | **string*                                                              | :heavy_minus_sign:                                                     | The ID of the connector that generated this credential.                |
| `CreatedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `DeletedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `ID`                                                                   | **string*                                                              | :heavy_minus_sign:                                                     | A unique idenditfier of the credential.                                |
| `TicketID`                                                             | **string*                                                              | :heavy_minus_sign:                                                     | The ticket ID of the request that generated this credential.           |
| `UpdatedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |