# Source

The source field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SourceSourceUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Source("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `SourceSourceUnspecified`   | SOURCE_UNSPECIFIED          |
| `SourceSourceC1`            | SOURCE_C1                   |
| `SourceSourceJira`          | SOURCE_JIRA                 |
| `SourceSourceSlack`         | SOURCE_SLACK                |
| `SourceSourceCopilotAgents` | SOURCE_COPILOT_AGENTS       |