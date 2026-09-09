# FundAssignment

FundAssignment is one principal's fund exception as the API renders it.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Controls`                                                           | [*shared.SpendControls](../../../pkg/models/shared/spendcontrols.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `CreatedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `TenantID`                                                           | `*string`                                                            | :heavy_minus_sign:                                                   | The tenantId field.                                                  |
| `UpdatedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `UserID`                                                             | `*string`                                                            | :heavy_minus_sign:                                                   | Canonical c1.models.user.v2.User id, every UserType.                 |