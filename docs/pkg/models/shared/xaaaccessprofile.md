# XAAAccessProfile

XAAAccessProfile is a requestable bundle of scopes for one resource server.


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `AppEntitlementID`                                    | `*string`                                             | :heavy_minus_sign:                                    | The AppEntitlement created for this profile.          |
| `AppID`                                               | `*string`                                             | :heavy_minus_sign:                                    | The application that owns the resource server.        |
| `CreatedAt`                                           | [*time.Time](https://pkg.go.dev/time#Time)            | :heavy_minus_sign:                                    | N/A                                                   |
| `DeletedAt`                                           | [*time.Time](https://pkg.go.dev/time#Time)            | :heavy_minus_sign:                                    | N/A                                                   |
| `Description`                                         | `*string`                                             | :heavy_minus_sign:                                    | Description of what access this profile grants.       |
| `DisplayName`                                         | `*string`                                             | :heavy_minus_sign:                                    | Display name for the profile.                         |
| `ID`                                                  | `*string`                                             | :heavy_minus_sign:                                    | Unique identifier for this access profile.            |
| `ScopeCount`                                          | `*int`                                                | :heavy_minus_sign:                                    | The number of scopes currently bound to this profile. |
| `UpdatedAt`                                           | [*time.Time](https://pkg.go.dev/time#Time)            | :heavy_minus_sign:                                    | N/A                                                   |
| `XaaResourceServerID`                                 | `*string`                                             | :heavy_minus_sign:                                    | The resource server this profile grants access to.    |