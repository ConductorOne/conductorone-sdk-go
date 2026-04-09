# CredentialChangeAction

The credentialChangeAction field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CredentialChangeActionSsfRevocationActionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CredentialChangeAction("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `CredentialChangeActionSsfRevocationActionUnspecified` | SSF_REVOCATION_ACTION_UNSPECIFIED                      |
| `CredentialChangeActionSsfRevocationActionRevokeAll`   | SSF_REVOCATION_ACTION_REVOKE_ALL                       |
| `CredentialChangeActionSsfRevocationActionLogOnly`     | SSF_REVOCATION_ACTION_LOG_ONLY                         |