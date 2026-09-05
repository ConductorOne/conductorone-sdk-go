# FundPolicy

FundPolicy is the tenant's fund policy as the API renders it. Every field is
 server-owned on the way out; requests name the fields they change rather than
 sending this message back.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `CreatedAt`                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                 | :heavy_minus_sign:                                                         | N/A                                                                        |
| `CurrencyCode`                                                             | `*string`                                                                  | :heavy_minus_sign:                                                         | ISO 4217. Set at Create and immutable thereafter.                          |
| `DefaultLimit`                                                             | [*shared.SpendLimit](../../../pkg/models/shared/spendlimit.md)             | :heavy_minus_sign:                                                         | N/A                                                                        |
| `OrgCeiling`                                                               | [*shared.SpendControls](../../../pkg/models/shared/spendcontrols.md)       | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Period`                                                                   | [*shared.FundPolicyPeriod](../../../pkg/models/shared/fundpolicyperiod.md) | :heavy_minus_sign:                                                         | The root period every amount in the tenant is denominated in.              |
| `TenantID`                                                                 | `*string`                                                                  | :heavy_minus_sign:                                                         | The tenantId field.                                                        |
| `UpdatedAt`                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                 | :heavy_minus_sign:                                                         | N/A                                                                        |