# AccountDisabledAction

The accountDisabledAction field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccountDisabledActionSsfRevocationActionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountDisabledAction("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `AccountDisabledActionSsfRevocationActionUnspecified` | SSF_REVOCATION_ACTION_UNSPECIFIED                     |
| `AccountDisabledActionSsfRevocationActionRevokeAll`   | SSF_REVOCATION_ACTION_REVOKE_ALL                      |
| `AccountDisabledActionSsfRevocationActionLogOnly`     | SSF_REVOCATION_ACTION_LOG_ONLY                        |