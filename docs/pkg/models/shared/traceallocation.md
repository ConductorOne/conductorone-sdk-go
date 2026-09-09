# TraceAllocation

TraceAllocation records one admitted scope, the row that supplied its
 resolved limit, and the historical limit and period.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `BudgetKey`                                                                          | `*string`                                                                            | :heavy_minus_sign:                                                                   | The budgetKey field.                                                                 |
| `Limit`                                                                              | [*shared.ResolvedLimit](../../../pkg/models/shared/resolvedlimit.md)                 | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Period`                                                                             | [*shared.TraceAllocationPeriod](../../../pkg/models/shared/traceallocationperiod.md) | :heavy_minus_sign:                                                                   | The period field.                                                                    |
| `Source`                                                                             | [*shared.AuthorityRef](../../../pkg/models/shared/authorityref.md)                   | :heavy_minus_sign:                                                                   | N/A                                                                                  |