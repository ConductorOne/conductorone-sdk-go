# MCPServerHostedConfigTokenSharing

Token sharing model for the configured auth method. SHARED means the
 admin authorizes once and the credential applies to every tool call;
 PER_USER means each user authenticates independently. PER_USER is
 supported for OAuth2 authorization-code, bearer_token, custom_header,
 and basic_auth methods; sending PER_USER alongside any other auth
 method is rejected with InvalidArgument.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPServerHostedConfigTokenSharingMcpServerTokenSharingUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPServerHostedConfigTokenSharing("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `MCPServerHostedConfigTokenSharingMcpServerTokenSharingUnspecified` | MCP_SERVER_TOKEN_SHARING_UNSPECIFIED                                |
| `MCPServerHostedConfigTokenSharingMcpServerTokenSharingShared`      | MCP_SERVER_TOKEN_SHARING_SHARED                                     |
| `MCPServerHostedConfigTokenSharingMcpServerTokenSharingPerUser`     | MCP_SERVER_TOKEN_SHARING_PER_USER                                   |