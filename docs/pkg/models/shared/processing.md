# Processing

The processing state of a task as defined by the `processing_enum`

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ProcessingTaskProcessingTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Processing("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ProcessingTaskProcessingTypeUnspecified` | TASK_PROCESSING_TYPE_UNSPECIFIED          |
| `ProcessingTaskProcessingTypeProcessing`  | TASK_PROCESSING_TYPE_PROCESSING           |
| `ProcessingTaskProcessingTypeWaiting`     | TASK_PROCESSING_TYPE_WAITING              |
| `ProcessingTaskProcessingTypeDone`        | TASK_PROCESSING_TYPE_DONE                 |