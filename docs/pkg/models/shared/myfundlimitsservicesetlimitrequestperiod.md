# MyFundLimitsServiceSetLimitRequestPeriod

Optional period override. Only valid together with the limit it denominates.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MyFundLimitsServiceSetLimitRequestPeriodPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MyFundLimitsServiceSetLimitRequestPeriod("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `MyFundLimitsServiceSetLimitRequestPeriodPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED                                         |
| `MyFundLimitsServiceSetLimitRequestPeriodPeriodKindDaily`       | PERIOD_KIND_DAILY                                               |
| `MyFundLimitsServiceSetLimitRequestPeriodPeriodKindWeekly`      | PERIOD_KIND_WEEKLY                                              |
| `MyFundLimitsServiceSetLimitRequestPeriodPeriodKindMonthly`     | PERIOD_KIND_MONTHLY                                             |
| `MyFundLimitsServiceSetLimitRequestPeriodPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY                                           |
| `MyFundLimitsServiceSetLimitRequestPeriodPeriodKindYearly`      | PERIOD_KIND_YEARLY                                              |