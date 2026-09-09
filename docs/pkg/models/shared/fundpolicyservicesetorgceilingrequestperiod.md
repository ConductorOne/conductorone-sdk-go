# FundPolicyServiceSetOrgCeilingRequestPeriod

Optional period override for the ceiling. Only valid together with limit.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FundPolicyServiceSetOrgCeilingRequestPeriodPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FundPolicyServiceSetOrgCeilingRequestPeriod("custom_value")
```


## Values

| Name                                                               | Value                                                              |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `FundPolicyServiceSetOrgCeilingRequestPeriodPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED                                            |
| `FundPolicyServiceSetOrgCeilingRequestPeriodPeriodKindDaily`       | PERIOD_KIND_DAILY                                                  |
| `FundPolicyServiceSetOrgCeilingRequestPeriodPeriodKindWeekly`      | PERIOD_KIND_WEEKLY                                                 |
| `FundPolicyServiceSetOrgCeilingRequestPeriodPeriodKindMonthly`     | PERIOD_KIND_MONTHLY                                                |
| `FundPolicyServiceSetOrgCeilingRequestPeriodPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY                                              |
| `FundPolicyServiceSetOrgCeilingRequestPeriodPeriodKindYearly`      | PERIOD_KIND_YEARLY                                                 |