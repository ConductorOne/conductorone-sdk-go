# UserOwnerEntitlement

UserOwnerEntitlement represents an entitlement ownership source for a canonical User (a service account).


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AppEntitlement`                                                       | [*shared.AppEntitlement](../../../pkg/models/shared/appentitlement.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `CreatedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `RoleSlug`                                                             | `*string`                                                              | :heavy_minus_sign:                                                     | The roleSlug field.                                                    |
| `UserID`                                                               | `*string`                                                              | :heavy_minus_sign:                                                     | The userId field.                                                      |