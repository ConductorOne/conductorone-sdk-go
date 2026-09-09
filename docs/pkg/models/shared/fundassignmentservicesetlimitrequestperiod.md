# FundAssignmentServiceSetLimitRequestPeriod

Optional period override. Only valid together with the limit it denominates.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FundAssignmentServiceSetLimitRequestPeriodPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FundAssignmentServiceSetLimitRequestPeriod("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `FundAssignmentServiceSetLimitRequestPeriodPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED                                           |
| `FundAssignmentServiceSetLimitRequestPeriodPeriodKindDaily`       | PERIOD_KIND_DAILY                                                 |
| `FundAssignmentServiceSetLimitRequestPeriodPeriodKindWeekly`      | PERIOD_KIND_WEEKLY                                                |
| `FundAssignmentServiceSetLimitRequestPeriodPeriodKindMonthly`     | PERIOD_KIND_MONTHLY                                               |
| `FundAssignmentServiceSetLimitRequestPeriodPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY                                             |
| `FundAssignmentServiceSetLimitRequestPeriodPeriodKindYearly`      | PERIOD_KIND_YEARLY                                                |