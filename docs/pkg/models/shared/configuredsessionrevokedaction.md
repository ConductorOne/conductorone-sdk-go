# ConfiguredSessionRevokedAction

Step 3: Action preview.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ConfiguredSessionRevokedActionSsfRevocationActionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ConfiguredSessionRevokedAction("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ConfiguredSessionRevokedActionSsfRevocationActionUnspecified` | SSF_REVOCATION_ACTION_UNSPECIFIED                              |
| `ConfiguredSessionRevokedActionSsfRevocationActionRevokeAll`   | SSF_REVOCATION_ACTION_REVOKE_ALL                               |
| `ConfiguredSessionRevokedActionSsfRevocationActionLogOnly`     | SSF_REVOCATION_ACTION_LOG_ONLY                                 |