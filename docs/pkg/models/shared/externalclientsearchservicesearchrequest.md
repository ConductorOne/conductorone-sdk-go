# ExternalClientSearchServiceSearchRequest

The ExternalClientSearchServiceSearchRequest message.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `PageSize`                                                                  | **int*                                                                      | :heavy_minus_sign:                                                          | The pageSize field.                                                         |
| `PageToken`                                                                 | **string*                                                                   | :heavy_minus_sign:                                                          | The pageToken field.                                                        |
| `Query`                                                                     | **string*                                                                   | :heavy_minus_sign:                                                          | Free-text search on client_name and user display name                       |
| `Users`                                                                     | [][shared.UserRef](../../../pkg/models/shared/userref.md)                   | :heavy_minus_sign:                                                          | Filter by specific user IDs                                                 |
| `WellKnownClients`                                                          | [][shared.WellKnownClients](../../../pkg/models/shared/wellknownclients.md) | :heavy_minus_sign:                                                          | Filter by well-known client type (e.g., CLAUDE_CODE, CURSOR, etc.)          |