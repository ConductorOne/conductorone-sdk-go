# SpendControlsPeriod

Only valid together with limit: a period without its amount would
 reinterpret some other layer's number in a cadence that layer never
 agreed to.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SpendControlsPeriodPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SpendControlsPeriod("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `SpendControlsPeriodPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED                    |
| `SpendControlsPeriodPeriodKindDaily`       | PERIOD_KIND_DAILY                          |
| `SpendControlsPeriodPeriodKindWeekly`      | PERIOD_KIND_WEEKLY                         |
| `SpendControlsPeriodPeriodKindMonthly`     | PERIOD_KIND_MONTHLY                        |
| `SpendControlsPeriodPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY                      |
| `SpendControlsPeriodPeriodKindYearly`      | PERIOD_KIND_YEARLY                         |