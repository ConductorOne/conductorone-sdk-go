# SubjectAppLimitServiceSetLimitRequestPeriod

Optional period override. Only valid together with the limit it denominates.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SubjectAppLimitServiceSetLimitRequestPeriodPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SubjectAppLimitServiceSetLimitRequestPeriod("custom_value")
```


## Values

| Name                                                               | Value                                                              |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `SubjectAppLimitServiceSetLimitRequestPeriodPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED                                            |
| `SubjectAppLimitServiceSetLimitRequestPeriodPeriodKindDaily`       | PERIOD_KIND_DAILY                                                  |
| `SubjectAppLimitServiceSetLimitRequestPeriodPeriodKindWeekly`      | PERIOD_KIND_WEEKLY                                                 |
| `SubjectAppLimitServiceSetLimitRequestPeriodPeriodKindMonthly`     | PERIOD_KIND_MONTHLY                                                |
| `SubjectAppLimitServiceSetLimitRequestPeriodPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY                                              |
| `SubjectAppLimitServiceSetLimitRequestPeriodPeriodKindYearly`      | PERIOD_KIND_YEARLY                                                 |