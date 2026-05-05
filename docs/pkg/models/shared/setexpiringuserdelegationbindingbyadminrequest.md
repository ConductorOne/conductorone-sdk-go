# SetExpiringUserDelegationBindingByAdminRequest

SetExpiringUserDelegationBindingByAdminRequest is the request for an admin to set a temporary delegation binding for a user.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `DelegatedUserID`                                                                 | `*string`                                                                         | :heavy_minus_sign:                                                                | The ID of the user who will act as delegate. Empty string removes the delegation. |
| `DelegationExpireAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                                        | :heavy_minus_sign:                                                                | N/A                                                                               |
| `DelegationStartAt`                                                               | [*time.Time](https://pkg.go.dev/time#Time)                                        | :heavy_minus_sign:                                                                | N/A                                                                               |