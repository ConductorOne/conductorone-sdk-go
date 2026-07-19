# ServerType

Server type (hosted or external).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ServerTypeMcpServerTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ServerType("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `ServerTypeMcpServerTypeUnspecified` | MCP_SERVER_TYPE_UNSPECIFIED          |
| `ServerTypeMcpServerTypeHosted`      | MCP_SERVER_TYPE_HOSTED               |
| `ServerTypeMcpServerTypeExternal`    | MCP_SERVER_TYPE_EXTERNAL             |