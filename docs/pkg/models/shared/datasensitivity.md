# DataSensitivity

Data sensitivity classification.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DataSensitivityMcpServerDataSensitivityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DataSensitivity("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `DataSensitivityMcpServerDataSensitivityUnspecified`  | MCP_SERVER_DATA_SENSITIVITY_UNSPECIFIED               |
| `DataSensitivityMcpServerDataSensitivityPublic`       | MCP_SERVER_DATA_SENSITIVITY_PUBLIC                    |
| `DataSensitivityMcpServerDataSensitivityInternal`     | MCP_SERVER_DATA_SENSITIVITY_INTERNAL                  |
| `DataSensitivityMcpServerDataSensitivityConfidential` | MCP_SERVER_DATA_SENSITIVITY_CONFIDENTIAL              |
| `DataSensitivityMcpServerDataSensitivityRestricted`   | MCP_SERVER_DATA_SENSITIVITY_RESTRICTED                |