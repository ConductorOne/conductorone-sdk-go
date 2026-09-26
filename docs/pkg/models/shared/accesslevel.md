# AccessLevel

Unspecified defaults to read-only access.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccessLevelMcpSystemToolsetUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccessLevel("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `AccessLevelMcpSystemToolsetUnspecified` | MCP_SYSTEM_TOOLSET_UNSPECIFIED           |
| `AccessLevelMcpSystemToolsetRead`        | MCP_SYSTEM_TOOLSET_READ                  |
| `AccessLevelMcpSystemToolsetAllApproved` | MCP_SYSTEM_TOOLSET_ALL_APPROVED          |