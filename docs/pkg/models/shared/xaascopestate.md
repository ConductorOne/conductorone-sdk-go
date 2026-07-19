# XAAScopeState

Approval/lifecycle state.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.XAAScopeStateXaaScopeStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.XAAScopeState("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `XAAScopeStateXaaScopeStateUnspecified`   | XAA_SCOPE_STATE_UNSPECIFIED               |
| `XAAScopeStateXaaScopeStatePendingReview` | XAA_SCOPE_STATE_PENDING_REVIEW            |
| `XAAScopeStateXaaScopeStateEnabled`       | XAA_SCOPE_STATE_ENABLED                   |
| `XAAScopeStateXaaScopeStateDisabled`      | XAA_SCOPE_STATE_DISABLED                  |
| `XAAScopeStateXaaScopeStateRemoved`       | XAA_SCOPE_STATE_REMOVED                   |