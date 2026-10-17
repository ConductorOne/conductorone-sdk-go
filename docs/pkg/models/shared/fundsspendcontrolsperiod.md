# FundsSpendControlsPeriod

The period field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FundsSpendControlsPeriodPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FundsSpendControlsPeriod("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `FundsSpendControlsPeriodPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED                         |
| `FundsSpendControlsPeriodPeriodKindDaily`       | PERIOD_KIND_DAILY                               |
| `FundsSpendControlsPeriodPeriodKindWeekly`      | PERIOD_KIND_WEEKLY                              |
| `FundsSpendControlsPeriodPeriodKindMonthly`     | PERIOD_KIND_MONTHLY                             |
| `FundsSpendControlsPeriodPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY                           |
| `FundsSpendControlsPeriodPeriodKindYearly`      | PERIOD_KIND_YEARLY                              |