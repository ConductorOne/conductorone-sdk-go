# InitialState

Initial lifecycle state. Omitted or ACTIVE creates an active server.
 DISABLED creates a server that denies connections until it is enabled.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.InitialStateVirtualMcpServerStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.InitialState("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `InitialStateVirtualMcpServerStateUnspecified` | VIRTUAL_MCP_SERVER_STATE_UNSPECIFIED           |
| `InitialStateVirtualMcpServerStateActive`      | VIRTUAL_MCP_SERVER_STATE_ACTIVE                |
| `InitialStateVirtualMcpServerStateDisabled`    | VIRTUAL_MCP_SERVER_STATE_DISABLED              |