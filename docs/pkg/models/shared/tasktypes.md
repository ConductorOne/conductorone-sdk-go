# TaskTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskTypesTaskTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskTypes("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `TaskTypesTaskTypeUnspecified` | TASK_TYPE_UNSPECIFIED          |
| `TaskTypesTaskTypeRequest`     | TASK_TYPE_REQUEST              |
| `TaskTypesTaskTypeRevoke`      | TASK_TYPE_REVOKE               |
| `TaskTypesTaskTypeReview`      | TASK_TYPE_REVIEW               |