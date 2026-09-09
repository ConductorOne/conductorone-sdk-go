# SubjectAppLimit

SubjectAppLimit is one subject's per-app limit row as the admin plane
 renders it. The subject is named explicitly, unlike MyFundLimit, which is
 always the caller's.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `AppID`                                                              | `*string`                                                            | :heavy_minus_sign:                                                   | The C1 App this limit applies to.                                    |
| `Controls`                                                           | [*shared.SpendControls](../../../pkg/models/shared/spendcontrols.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `CreatedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `DeletedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `TenantID`                                                           | `*string`                                                            | :heavy_minus_sign:                                                   | The tenantId field.                                                  |
| `UpdatedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `UserID`                                                             | `*string`                                                            | :heavy_minus_sign:                                                   | Canonical c1.models.user.v2.User id.                                 |