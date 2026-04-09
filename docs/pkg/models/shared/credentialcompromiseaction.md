# CredentialCompromiseAction

The credentialCompromiseAction field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CredentialCompromiseActionSsfRevocationActionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CredentialCompromiseAction("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `CredentialCompromiseActionSsfRevocationActionUnspecified` | SSF_REVOCATION_ACTION_UNSPECIFIED                          |
| `CredentialCompromiseActionSsfRevocationActionRevokeAll`   | SSF_REVOCATION_ACTION_REVOKE_ALL                           |
| `CredentialCompromiseActionSsfRevocationActionLogOnly`     | SSF_REVOCATION_ACTION_LOG_ONLY                             |