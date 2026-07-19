# TaskAuditViewSource

The source field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskAuditViewSourceSourceUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskAuditViewSource("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `TaskAuditViewSourceSourceUnspecified`   | SOURCE_UNSPECIFIED                       |
| `TaskAuditViewSourceSourceC1`            | SOURCE_C1                                |
| `TaskAuditViewSourceSourceJira`          | SOURCE_JIRA                              |
| `TaskAuditViewSourceSourceSlack`         | SOURCE_SLACK                             |
| `TaskAuditViewSourceSourceCopilotAgents` | SOURCE_COPILOT_AGENTS                    |