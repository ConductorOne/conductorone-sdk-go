# VirtualMCPServerState

Current lifecycle state.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.VirtualMCPServerStateVirtualMcpServerStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.VirtualMCPServerState("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `VirtualMCPServerStateVirtualMcpServerStateUnspecified` | VIRTUAL_MCP_SERVER_STATE_UNSPECIFIED                    |
| `VirtualMCPServerStateVirtualMcpServerStateActive`      | VIRTUAL_MCP_SERVER_STATE_ACTIVE                         |
| `VirtualMCPServerStateVirtualMcpServerStateDisabled`    | VIRTUAL_MCP_SERVER_STATE_DISABLED                       |