# AppPopulationReportState

The state field tracks the state of the AppPopulationReport. This state field can be one of REPORT_STATE_PENDING, REPORT_STATE_UNSPECIFIED, REPORT_STATE_OK, REPORT_STATE_ERROR.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AppPopulationReportStateReportStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AppPopulationReportState("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `AppPopulationReportStateReportStateUnspecified` | REPORT_STATE_UNSPECIFIED                         |
| `AppPopulationReportStateReportStatePending`     | REPORT_STATE_PENDING                             |
| `AppPopulationReportStateReportStateOk`          | REPORT_STATE_OK                                  |
| `AppPopulationReportStateReportStateError`       | REPORT_STATE_ERROR                               |