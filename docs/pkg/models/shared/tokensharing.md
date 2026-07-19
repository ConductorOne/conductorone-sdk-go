# TokenSharing

Token sharing model. SHARED = admin authorizes once; PER_USER = each user
 authenticates independently. PER_USER is supported for OAuth2
 authorization_code, bearer_token, custom_header, and basic_auth. Defaults
 to SHARED at runtime.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TokenSharingMcpServerTokenSharingUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TokenSharing("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `TokenSharingMcpServerTokenSharingUnspecified` | MCP_SERVER_TOKEN_SHARING_UNSPECIFIED           |
| `TokenSharingMcpServerTokenSharingShared`      | MCP_SERVER_TOKEN_SHARING_SHARED                |
| `TokenSharingMcpServerTokenSharingPerUser`     | MCP_SERVER_TOKEN_SHARING_PER_USER              |