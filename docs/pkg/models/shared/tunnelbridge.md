# TunnelBridge

TunnelBridge is the API view of a bridge — the customer-facing entity
 for managing a wormhole tunnel appliance.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Appliance`                                                              | [*shared.TunnelAppliance](../../../pkg/models/shared/tunnelappliance.md) | :heavy_minus_sign:                                                       | N/A                                                                      |
| `CreatedAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | N/A                                                                      |
| `DeletedAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Description`                                                            | `*string`                                                                | :heavy_minus_sign:                                                       | The description field.                                                   |
| `DisplayName`                                                            | `*string`                                                                | :heavy_minus_sign:                                                       | The displayName field.                                                   |
| `ID`                                                                     | `*string`                                                                | :heavy_minus_sign:                                                       | The id field.                                                            |
| `UpdatedAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | N/A                                                                      |