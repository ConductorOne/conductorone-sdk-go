# FundPolicyPeriod

The root period every amount in the tenant is denominated in.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FundPolicyPeriodPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FundPolicyPeriod("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `FundPolicyPeriodPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED                 |
| `FundPolicyPeriodPeriodKindDaily`       | PERIOD_KIND_DAILY                       |
| `FundPolicyPeriodPeriodKindWeekly`      | PERIOD_KIND_WEEKLY                      |
| `FundPolicyPeriodPeriodKindMonthly`     | PERIOD_KIND_MONTHLY                     |
| `FundPolicyPeriodPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY                   |
| `FundPolicyPeriodPeriodKindYearly`      | PERIOD_KIND_YEARLY                      |