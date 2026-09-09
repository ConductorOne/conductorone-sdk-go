# PeriodKind

Length of the denied budget period.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PeriodKindPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PeriodKind("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PeriodKindPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED           |
| `PeriodKindPeriodKindDaily`       | PERIOD_KIND_DAILY                 |
| `PeriodKindPeriodKindWeekly`      | PERIOD_KIND_WEEKLY                |
| `PeriodKindPeriodKindMonthly`     | PERIOD_KIND_MONTHLY               |
| `PeriodKindPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY             |
| `PeriodKindPeriodKindYearly`      | PERIOD_KIND_YEARLY                |