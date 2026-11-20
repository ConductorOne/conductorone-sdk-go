# MCPRequestAccessProvisioningAccessLevel

The system toolset selected for requestable access, independently of automatic access.
 Unspecified inherits the legacy shared access_level, then defaults to Read tools.
 Access-profile associations remain stored on each toolset.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPRequestAccessProvisioningAccessLevelMcpSystemToolsetUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPRequestAccessProvisioningAccessLevel("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `MCPRequestAccessProvisioningAccessLevelMcpSystemToolsetUnspecified` | MCP_SYSTEM_TOOLSET_UNSPECIFIED                                       |
| `MCPRequestAccessProvisioningAccessLevelMcpSystemToolsetRead`        | MCP_SYSTEM_TOOLSET_READ                                              |
| `MCPRequestAccessProvisioningAccessLevelMcpSystemToolsetAllApproved` | MCP_SYSTEM_TOOLSET_ALL_APPROVED                                      |
| `MCPRequestAccessProvisioningAccessLevelMcpSystemToolsetCustom`      | MCP_SYSTEM_TOOLSET_CUSTOM                                            |
| `MCPRequestAccessProvisioningAccessLevelMcpSystemToolsetWrite`       | MCP_SYSTEM_TOOLSET_WRITE                                             |