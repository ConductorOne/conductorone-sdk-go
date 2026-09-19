# CredentialRevocationState

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CredentialRevocationStateCredentialRevocationStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CredentialRevocationState("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `CredentialRevocationStateCredentialRevocationStateUnspecified`  | CREDENTIAL_REVOCATION_STATE_UNSPECIFIED                          |
| `CredentialRevocationStateCredentialRevocationStateReady`        | CREDENTIAL_REVOCATION_STATE_READY                                |
| `CredentialRevocationStateCredentialRevocationStateRevoking`     | CREDENTIAL_REVOCATION_STATE_REVOKING                             |
| `CredentialRevocationStateCredentialRevocationStateRevokeFailed` | CREDENTIAL_REVOCATION_STATE_REVOKE_FAILED                        |
| `CredentialRevocationStateCredentialRevocationStateUnavailable`  | CREDENTIAL_REVOCATION_STATE_UNAVAILABLE                          |