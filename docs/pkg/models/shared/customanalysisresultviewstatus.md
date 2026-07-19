# CustomAnalysisResultViewStatus

Execution status of this analysis (e.g., running, completed, failed).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CustomAnalysisResultViewStatusRunStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CustomAnalysisResultViewStatus("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `CustomAnalysisResultViewStatusRunStatusUnspecified` | RUN_STATUS_UNSPECIFIED                               |
| `CustomAnalysisResultViewStatusRunStatusRunning`     | RUN_STATUS_RUNNING                                   |
| `CustomAnalysisResultViewStatusRunStatusCompleted`   | RUN_STATUS_COMPLETED                                 |
| `CustomAnalysisResultViewStatusRunStatusFailed`      | RUN_STATUS_FAILED                                    |