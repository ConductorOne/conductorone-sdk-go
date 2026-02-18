# PaperSecretAdminServiceSearchAuditEventsRequest

The PaperSecretAdminServiceSearchAuditEventsRequest message.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ActorEmail`                                                  | **string*                                                     | :heavy_minus_sign:                                            | Filter by external email (partial match via full-text search) |
| `ActorUserID`                                                 | **string*                                                     | :heavy_minus_sign:                                            | Filter by C1 user ID (internal users)                         |
| `ClientIP`                                                    | **string*                                                     | :heavy_minus_sign:                                            | Filter by client IP (exact match)                             |
| `PageSize`                                                    | **int*                                                        | :heavy_minus_sign:                                            | The pageSize field.                                           |
| `PageToken`                                                   | **string*                                                     | :heavy_minus_sign:                                            | The pageToken field.                                          |
| `VaultID`                                                     | **string*                                                     | :heavy_minus_sign:                                            | Filter by specific vault                                      |