# BatonResourceActionRef

BatonResourceActionRef describes dispatch to a connector resource-create
 action (for example, a group template that creates a group in the connected
 application).


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `AppID`                                                           | `*string`                                                         | :heavy_minus_sign:                                                | The app the resource is created in.                               |
| `BatonActionDisplayName`                                          | `*string`                                                         | :heavy_minus_sign:                                                | The connector-defined display name of the resource-create action. |
| `BatonActionName`                                                 | `*string`                                                         | :heavy_minus_sign:                                                | The connector-defined name of the resource-create action.         |
| `ConnectorID`                                                     | `*string`                                                         | :heavy_minus_sign:                                                | The connector that executes the resource-create action.           |
| `ResourceTypeID`                                                  | `*string`                                                         | :heavy_minus_sign:                                                | The type of resource the action creates (for example, "group").   |