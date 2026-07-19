# Source

Whether the assignment is direct or conferred through a group.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SourceAssignmentSourceUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Source("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `SourceAssignmentSourceUnspecified` | ASSIGNMENT_SOURCE_UNSPECIFIED       |
| `SourceAssignmentSourceDirect`      | ASSIGNMENT_SOURCE_DIRECT            |
| `SourceAssignmentSourceGroup`       | ASSIGNMENT_SOURCE_GROUP             |