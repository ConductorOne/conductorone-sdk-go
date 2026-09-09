# AppManagedStateBindingView

AppManagedStateBindingView contains a managed-state binding and paths to its related objects.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `AppManagementStateBinding`                                                            | [*shared.AppManagedStateBinding](../../../pkg/models/shared/appmanagedstatebinding.md) | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `AppPath`                                                                              | `*string`                                                                              | :heavy_minus_sign:                                                                     | Path of the application that owns the connector.                                       |
| `ResourcePath`                                                                         | `*string`                                                                              | :heavy_minus_sign:                                                                     | Path of the connector resource representing the discovered application.                |