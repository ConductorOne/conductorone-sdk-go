# MCPAutomaticAccessProvisioningAccessLevel

The system toolset granted through automatic access, independently of requestable access.
 Unspecified inherits the legacy shared access_level, then defaults to Read tools.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPAutomaticAccessProvisioningAccessLevelMcpSystemToolsetUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPAutomaticAccessProvisioningAccessLevel("custom_value")
```


## Values

| Name                                                                   | Value                                                                  |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `MCPAutomaticAccessProvisioningAccessLevelMcpSystemToolsetUnspecified` | MCP_SYSTEM_TOOLSET_UNSPECIFIED                                         |
| `MCPAutomaticAccessProvisioningAccessLevelMcpSystemToolsetRead`        | MCP_SYSTEM_TOOLSET_READ                                                |
| `MCPAutomaticAccessProvisioningAccessLevelMcpSystemToolsetAllApproved` | MCP_SYSTEM_TOOLSET_ALL_APPROVED                                        |
| `MCPAutomaticAccessProvisioningAccessLevelMcpSystemToolsetCustom`      | MCP_SYSTEM_TOOLSET_CUSTOM                                              |
| `MCPAutomaticAccessProvisioningAccessLevelMcpSystemToolsetWrite`       | MCP_SYSTEM_TOOLSET_WRITE                                               |