# MCPServerViewDataSensitivity

Data sensitivity classification.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerViewDataSensitivityMcpServerDataSensitivityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerViewDataSensitivity("custom_value")
```


## Values

| Name                                                               | Value                                                              |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `MCPServerViewDataSensitivityMcpServerDataSensitivityUnspecified`  | MCP_SERVER_DATA_SENSITIVITY_UNSPECIFIED                            |
| `MCPServerViewDataSensitivityMcpServerDataSensitivityPublic`       | MCP_SERVER_DATA_SENSITIVITY_PUBLIC                                 |
| `MCPServerViewDataSensitivityMcpServerDataSensitivityInternal`     | MCP_SERVER_DATA_SENSITIVITY_INTERNAL                               |
| `MCPServerViewDataSensitivityMcpServerDataSensitivityConfidential` | MCP_SERVER_DATA_SENSITIVITY_CONFIDENTIAL                           |
| `MCPServerViewDataSensitivityMcpServerDataSensitivityRestricted`   | MCP_SERVER_DATA_SENSITIVITY_RESTRICTED                             |