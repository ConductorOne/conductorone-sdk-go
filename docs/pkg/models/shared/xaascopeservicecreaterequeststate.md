# XAAScopeServiceCreateRequestState

Initial state. UNSPECIFIED defaults to PENDING_REVIEW.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.XAAScopeServiceCreateRequestStateXaaScopeStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.XAAScopeServiceCreateRequestState("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `XAAScopeServiceCreateRequestStateXaaScopeStateUnspecified`   | XAA_SCOPE_STATE_UNSPECIFIED                                   |
| `XAAScopeServiceCreateRequestStateXaaScopeStatePendingReview` | XAA_SCOPE_STATE_PENDING_REVIEW                                |
| `XAAScopeServiceCreateRequestStateXaaScopeStateEnabled`       | XAA_SCOPE_STATE_ENABLED                                       |
| `XAAScopeServiceCreateRequestStateXaaScopeStateDisabled`      | XAA_SCOPE_STATE_DISABLED                                      |
| `XAAScopeServiceCreateRequestStateXaaScopeStateRemoved`       | XAA_SCOPE_STATE_REMOVED                                       |