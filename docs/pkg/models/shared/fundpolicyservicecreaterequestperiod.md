# FundPolicyServiceCreateRequestPeriod

The period field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FundPolicyServiceCreateRequestPeriodPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FundPolicyServiceCreateRequestPeriod("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `FundPolicyServiceCreateRequestPeriodPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED                                     |
| `FundPolicyServiceCreateRequestPeriodPeriodKindDaily`       | PERIOD_KIND_DAILY                                           |
| `FundPolicyServiceCreateRequestPeriodPeriodKindWeekly`      | PERIOD_KIND_WEEKLY                                          |
| `FundPolicyServiceCreateRequestPeriodPeriodKindMonthly`     | PERIOD_KIND_MONTHLY                                         |
| `FundPolicyServiceCreateRequestPeriodPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY                                       |
| `FundPolicyServiceCreateRequestPeriodPeriodKindYearly`      | PERIOD_KIND_YEARLY                                          |