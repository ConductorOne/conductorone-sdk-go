# TaskTypeActionType

Flavor of action the ticket represents — mirrors the snapshot's
 target_ref variant.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskTypeActionTypeTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskTypeActionType("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `TaskTypeActionTypeTypeUnspecified`    | TYPE_UNSPECIFIED                       |
| `TaskTypeActionTypeTypeGrant`          | TYPE_GRANT                             |
| `TaskTypeActionTypeTypeWorkflow`       | TYPE_WORKFLOW                          |
| `TaskTypeActionTypeTypeResourceAction` | TYPE_RESOURCE_ACTION                   |
| `TaskTypeActionTypeTypeToolCall`       | TYPE_TOOL_CALL                         |
| `TaskTypeActionTypeTypeManual`         | TYPE_MANUAL                            |