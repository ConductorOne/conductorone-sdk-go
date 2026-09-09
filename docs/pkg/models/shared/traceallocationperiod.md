# TraceAllocationPeriod

The period field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TraceAllocationPeriodPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TraceAllocationPeriod("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `TraceAllocationPeriodPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED                      |
| `TraceAllocationPeriodPeriodKindDaily`       | PERIOD_KIND_DAILY                            |
| `TraceAllocationPeriodPeriodKindWeekly`      | PERIOD_KIND_WEEKLY                           |
| `TraceAllocationPeriodPeriodKindMonthly`     | PERIOD_KIND_MONTHLY                          |
| `TraceAllocationPeriodPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY                        |
| `TraceAllocationPeriodPeriodKindYearly`      | PERIOD_KIND_YEARLY                           |