# TraceAllocation

TraceAllocation records one scope's historical limit and its source.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `BudgetKey`                                                                          | `*string`                                                                            | :heavy_minus_sign:                                                                   | Stable identifier for the scope's budget account.                                    |
| `Limit`                                                                              | [*shared.ResolvedLimit](../../../pkg/models/shared/resolvedlimit.md)                 | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Period`                                                                             | [*shared.TraceAllocationPeriod](../../../pkg/models/shared/traceallocationperiod.md) | :heavy_minus_sign:                                                                   | Length of the budget period.                                                         |
| `Source`                                                                             | [*shared.AuthorityRef](../../../pkg/models/shared/authorityref.md)                   | :heavy_minus_sign:                                                                   | N/A                                                                                  |