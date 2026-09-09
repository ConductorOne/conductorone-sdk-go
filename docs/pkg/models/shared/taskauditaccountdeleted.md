# TaskAuditAccountDeleted

TaskAuditAccountDeleted records an account deletion reported by a connector
 while completing a revoke action.


## Fields

| Field                          | Type                           | Required                       | Description                    |
| ------------------------------ | ------------------------------ | ------------------------------ | ------------------------------ |
| `AppID`                        | `*string`                      | :heavy_minus_sign:             | The appId field.               |
| `AppUserID`                    | `*string`                      | :heavy_minus_sign:             | The appUserId field.           |
| `ConnectorResourceID`          | `*string`                      | :heavy_minus_sign:             | The connectorResourceId field. |
| `DisplayName`                  | `*string`                      | :heavy_minus_sign:             | The displayName field.         |
| `Email`                        | `*string`                      | :heavy_minus_sign:             | The email field.               |
| `Username`                     | `*string`                      | :heavy_minus_sign:             | The username field.            |