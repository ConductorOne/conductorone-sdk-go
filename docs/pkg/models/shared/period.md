# Period

Optional period override. Only valid together with the limit it denominates.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PeriodPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Period("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `PeriodPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED       |
| `PeriodPeriodKindDaily`       | PERIOD_KIND_DAILY             |
| `PeriodPeriodKindWeekly`      | PERIOD_KIND_WEEKLY            |
| `PeriodPeriodKindMonthly`     | PERIOD_KIND_MONTHLY           |
| `PeriodPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY         |
| `PeriodPeriodKindYearly`      | PERIOD_KIND_YEARLY            |