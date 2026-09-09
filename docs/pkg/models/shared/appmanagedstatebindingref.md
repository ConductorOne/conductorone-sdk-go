# AppManagedStateBindingRef

AppManagedStateBindingRef identifies an application discovered by a connector.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `AppID`                                                   | `*string`                                                 | :heavy_minus_sign:                                        | ID of the application that owns the connector.            |
| `ResourceID`                                              | `*string`                                                 | :heavy_minus_sign:                                        | Resource ID of the discovered application.                |
| `ResourceTypeID`                                          | `*string`                                                 | :heavy_minus_sign:                                        | ID of the resource type used for discovered applications. |