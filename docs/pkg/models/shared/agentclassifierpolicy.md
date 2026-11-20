# AgentClassifierPolicy

AgentClassifierPolicy is the tenant's full classifier policy: the ordered
 rule cascade plus the no-match default.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `AgentID`                                                                         | `*string`                                                                         | :heavy_minus_sign:                                                                | Optional agent scoping. Empty is the tenant-wide default policy.                  |
| `CreatedAt`                                                                       | [*time.Time](https://pkg.go.dev/time#Time)                                        | :heavy_minus_sign:                                                                | N/A                                                                               |
| `DefaultDenyReason`                                                               | `*string`                                                                         | :heavy_minus_sign:                                                                | The defaultDenyReason field.                                                      |
| `DefaultOutcome`                                                                  | [*shared.DefaultOutcome](../../../pkg/models/shared/defaultoutcome.md)            | :heavy_minus_sign:                                                                | The defaultOutcome field.                                                         |
| `Rules`                                                                           | [][shared.AgentClassifierRule](../../../pkg/models/shared/agentclassifierrule.md) | :heavy_minus_sign:                                                                | The rules field.                                                                  |
| `TenantID`                                                                        | `*string`                                                                         | :heavy_minus_sign:                                                                | The tenantId field.                                                               |
| `UpdatedAt`                                                                       | [*time.Time](https://pkg.go.dev/time#Time)                                        | :heavy_minus_sign:                                                                | N/A                                                                               |