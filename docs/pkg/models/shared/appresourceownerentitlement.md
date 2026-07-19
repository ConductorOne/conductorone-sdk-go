# AppResourceOwnerEntitlement

AppResourceOwnerEntitlement represents an entitlement ownership source for an app resource.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AppEntitlement`                                                       | [*shared.AppEntitlement](../../../pkg/models/shared/appentitlement.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `AppID`                                                                | `*string`                                                              | :heavy_minus_sign:                                                     | The appId field.                                                       |
| `CreatedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `ResourceID`                                                           | `*string`                                                              | :heavy_minus_sign:                                                     | The resourceId field.                                                  |
| `ResourceTypeID`                                                       | `*string`                                                              | :heavy_minus_sign:                                                     | The resourceTypeId field.                                              |
| `RoleSlug`                                                             | `*string`                                                              | :heavy_minus_sign:                                                     | The roleSlug field.                                                    |