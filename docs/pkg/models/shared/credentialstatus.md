# CredentialStatus

Lifecycle status of this credential record.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CredentialStatusTunnelCredentialStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CredentialStatus("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `CredentialStatusTunnelCredentialStatusUnspecified` | TUNNEL_CREDENTIAL_STATUS_UNSPECIFIED                |
| `CredentialStatusTunnelCredentialStatusActive`      | TUNNEL_CREDENTIAL_STATUS_ACTIVE                     |
| `CredentialStatusTunnelCredentialStatusRevoked`     | TUNNEL_CREDENTIAL_STATUS_REVOKED                    |
| `CredentialStatusTunnelCredentialStatusExpired`     | TUNNEL_CREDENTIAL_STATUS_EXPIRED                    |