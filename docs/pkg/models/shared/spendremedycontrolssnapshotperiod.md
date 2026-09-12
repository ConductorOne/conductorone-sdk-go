# SpendRemedyControlsSnapshotPeriod

The period field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SpendRemedyControlsSnapshotPeriodPeriodKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SpendRemedyControlsSnapshotPeriod("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `SpendRemedyControlsSnapshotPeriodPeriodKindUnspecified` | PERIOD_KIND_UNSPECIFIED                                  |
| `SpendRemedyControlsSnapshotPeriodPeriodKindDaily`       | PERIOD_KIND_DAILY                                        |
| `SpendRemedyControlsSnapshotPeriodPeriodKindWeekly`      | PERIOD_KIND_WEEKLY                                       |
| `SpendRemedyControlsSnapshotPeriodPeriodKindMonthly`     | PERIOD_KIND_MONTHLY                                      |
| `SpendRemedyControlsSnapshotPeriodPeriodKindQuarterly`   | PERIOD_KIND_QUARTERLY                                    |
| `SpendRemedyControlsSnapshotPeriodPeriodKindYearly`      | PERIOD_KIND_YEARLY                                       |