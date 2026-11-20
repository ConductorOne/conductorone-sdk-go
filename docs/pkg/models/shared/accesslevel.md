# ~~AccessLevel~~

Deprecated: set the automatic and request access levels and tool_enablement instead.
 Ignored on write; returned for registrations saved before it was deprecated.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
| `AccessLevelMcpSystemToolsetCustom`      | MCP_SYSTEM_TOOLSET_CUSTOM                |
| `AccessLevelMcpSystemToolsetWrite`       | MCP_SYSTEM_TOOLSET_WRITE                 |