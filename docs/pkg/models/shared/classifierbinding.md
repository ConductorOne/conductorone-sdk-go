# ClassifierBinding

ClassifierBinding attaches one Classifier to one enforcement target.


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ClassifierID`                                          | `string`                                                | :heavy_check_mark:                                      | The classifierId field.                                 |
| `CreatedAt`                                             | [*time.Time](https://pkg.go.dev/time#Time)              | :heavy_minus_sign:                                      | N/A                                                     |
| `DeletedAt`                                             | [*time.Time](https://pkg.go.dev/time#Time)              | :heavy_minus_sign:                                      | N/A                                                     |
| `ID`                                                    | `*string`                                               | :heavy_minus_sign:                                      | The id field.                                           |
| `Surface`                                               | [shared.Surface](../../../pkg/models/shared/surface.md) | :heavy_check_mark:                                      | The surface field.                                      |
| `TargetID`                                              | `*string`                                               | :heavy_minus_sign:                                      | The targetId field.                                     |
| `TenantID`                                              | `*string`                                               | :heavy_minus_sign:                                      | The tenantId field.                                     |
| `UpdatedAt`                                             | [*time.Time](https://pkg.go.dev/time#Time)              | :heavy_minus_sign:                                      | N/A                                                     |