# AppCap

AppCap is one app's tenant-wide ceiling as the API renders it.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `AppID`                                                              | `*string`                                                            | :heavy_minus_sign:                                                   | The C1 App the spend is attributed to.                               |
| `Controls`                                                           | [*shared.SpendControls](../../../pkg/models/shared/spendcontrols.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `CreatedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |
| `TenantID`                                                           | `*string`                                                            | :heavy_minus_sign:                                                   | The tenantId field.                                                  |
| `UpdatedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |