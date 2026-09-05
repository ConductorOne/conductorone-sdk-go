# AccessReviewReportState

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccessReviewReportStateReportStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccessReviewReportState("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `AccessReviewReportStateReportStateUnspecified` | REPORT_STATE_UNSPECIFIED                        |
| `AccessReviewReportStateReportStatePending`     | REPORT_STATE_PENDING                            |
| `AccessReviewReportStateReportStateOk`          | REPORT_STATE_OK                                 |
| `AccessReviewReportStateReportStateError`       | REPORT_STATE_ERROR                              |