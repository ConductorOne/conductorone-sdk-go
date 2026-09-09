# ReportRunStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ReportRunStatusReportRunStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ReportRunStatus("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `ReportRunStatusReportRunStatusUnspecified`  | REPORT_RUN_STATUS_UNSPECIFIED                |
| `ReportRunStatusReportRunStatusPending`      | REPORT_RUN_STATUS_PENDING                    |
| `ReportRunStatusReportRunStatusSucceeded`    | REPORT_RUN_STATUS_SUCCEEDED                  |
| `ReportRunStatusReportRunStatusFailed`       | REPORT_RUN_STATUS_FAILED                     |
| `ReportRunStatusReportRunStatusStaleProgram` | REPORT_RUN_STATUS_STALE_PROGRAM              |