# ApprovalInstanceState

The state of the approval instance

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ApprovalInstanceStateApprovalInstanceStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ApprovalInstanceState("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ApprovalInstanceStateApprovalInstanceStateUnspecified`          | APPROVAL_INSTANCE_STATE_UNSPECIFIED                              |
| `ApprovalInstanceStateApprovalInstanceStateInit`                 | APPROVAL_INSTANCE_STATE_INIT                                     |
| `ApprovalInstanceStateApprovalInstanceStateSendingNotifications` | APPROVAL_INSTANCE_STATE_SENDING_NOTIFICATIONS                    |
| `ApprovalInstanceStateApprovalInstanceStateWaiting`              | APPROVAL_INSTANCE_STATE_WAITING                                  |
| `ApprovalInstanceStateApprovalInstanceStateDone`                 | APPROVAL_INSTANCE_STATE_DONE                                     |