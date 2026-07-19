# MCPServerViewTokenSharing

Token sharing model in use. Read-only; derived from stored config.
 For rows stored under the legacy SERVICE/PASSTHROUGH OAuth2 modes,
 this is synthesized as SHARED/PER_USER respectively.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerViewTokenSharingMcpServerTokenSharingUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerViewTokenSharing("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `MCPServerViewTokenSharingMcpServerTokenSharingUnspecified` | MCP_SERVER_TOKEN_SHARING_UNSPECIFIED                        |
| `MCPServerViewTokenSharingMcpServerTokenSharingShared`      | MCP_SERVER_TOKEN_SHARING_SHARED                             |
| `MCPServerViewTokenSharingMcpServerTokenSharingPerUser`     | MCP_SERVER_TOKEN_SHARING_PER_USER                           |