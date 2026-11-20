# SpendRemedyReviewPeriod

The period field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SpendRemedyReviewPeriodPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SpendRemedyReviewPeriod("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `SpendRemedyReviewPeriodPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED                        |
| `SpendRemedyReviewPeriodPeriodKindDaily`       | PERIOD_KIND_DAILY                              |
| `SpendRemedyReviewPeriodPeriodKindWeekly`      | PERIOD_KIND_WEEKLY                             |
| `SpendRemedyReviewPeriodPeriodKindMonthly`     | PERIOD_KIND_MONTHLY                            |
| `SpendRemedyReviewPeriodPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY                          |
| `SpendRemedyReviewPeriodPeriodKindYearly`      | PERIOD_KIND_YEARLY                             |