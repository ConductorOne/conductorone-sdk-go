# DelegateStatus

Filter for users based on their delegate status.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DelegateStatusDelegateStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DelegateStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `DelegateStatusDelegateStatusUnspecified` | DELEGATE_STATUS_UNSPECIFIED               |
| `DelegateStatusDelegateStatusHasDelegate` | DELEGATE_STATUS_HAS_DELEGATE              |
| `DelegateStatusDelegateStatusNoDelegate`  | DELEGATE_STATUS_NO_DELEGATE               |