# GetCustomAnalysisResultResponseStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GetCustomAnalysisResultResponseStatusRunStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GetCustomAnalysisResultResponseStatus("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `GetCustomAnalysisResultResponseStatusRunStatusUnspecified` | RUN_STATUS_UNSPECIFIED                                      |
| `GetCustomAnalysisResultResponseStatusRunStatusRunning`     | RUN_STATUS_RUNNING                                          |
| `GetCustomAnalysisResultResponseStatusRunStatusCompleted`   | RUN_STATUS_COMPLETED                                        |
| `GetCustomAnalysisResultResponseStatusRunStatusFailed`      | RUN_STATUS_FAILED                                           |