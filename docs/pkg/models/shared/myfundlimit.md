# MyFundLimit

MyFundLimit is one of the caller's own per-app limits. It carries no user id:
 it is always the caller's.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `AppID`                                                              | `*string`                                                            | :heavy_minus_sign:                                                   | The C1 App this limit applies to.                                    |
| `Controls`                                                           | [*shared.SpendControls](../../../pkg/models/shared/spendcontrols.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `CreatedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `DeletedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `UpdatedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |