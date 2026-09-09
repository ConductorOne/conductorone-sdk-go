# AppManagedStateBinding

AppManagedStateBinding records whether a connector-discovered application is managed in ConductorOne.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `AppID`                                                                   | `*string`                                                                 | :heavy_minus_sign:                                                        | Application that owns the connector which discovered this application.    |
| `CreatedAt`                                                               | [*time.Time](https://pkg.go.dev/time#Time)                                | :heavy_minus_sign:                                                        | N/A                                                                       |
| `DeletedAt`                                                               | [*time.Time](https://pkg.go.dev/time#Time)                                | :heavy_minus_sign:                                                        | N/A                                                                       |
| `DisplayName`                                                             | `*string`                                                                 | :heavy_minus_sign:                                                        | Display name of the discovered application.                               |
| `ResourceID`                                                              | `*string`                                                                 | :heavy_minus_sign:                                                        | Resource ID of the discovered application.                                |
| `ResourceTypeID`                                                          | `*string`                                                                 | :heavy_minus_sign:                                                        | Resource type used by the connector to represent discovered applications. |
| `State`                                                                   | [*shared.AppManagedState](../../../pkg/models/shared/appmanagedstate.md)  | :heavy_minus_sign:                                                        | N/A                                                                       |
| `UpdatedAt`                                                               | [*time.Time](https://pkg.go.dev/time#Time)                                | :heavy_minus_sign:                                                        | N/A                                                                       |