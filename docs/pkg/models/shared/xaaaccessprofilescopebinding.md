# XAAAccessProfileScopeBinding

XAAAccessProfileScopeBinding is a binding between an access profile and a
 scope. Both ends belong to one resource server.


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `AccessProfileID`                              | `*string`                                      | :heavy_minus_sign:                             | The access profile end of the binding.         |
| `AppID`                                        | `*string`                                      | :heavy_minus_sign:                             | The application that owns the resource server. |
| `CreatedAt`                                    | [*time.Time](https://pkg.go.dev/time#Time)     | :heavy_minus_sign:                             | N/A                                            |
| `DeletedAt`                                    | [*time.Time](https://pkg.go.dev/time#Time)     | :heavy_minus_sign:                             | N/A                                            |
| `UpdatedAt`                                    | [*time.Time](https://pkg.go.dev/time#Time)     | :heavy_minus_sign:                             | N/A                                            |
| `XaaResourceServerID`                          | `*string`                                      | :heavy_minus_sign:                             | The resource server both ends belong to.       |
| `XaaScopeID`                                   | `*string`                                      | :heavy_minus_sign:                             | The scope end of the binding.                  |