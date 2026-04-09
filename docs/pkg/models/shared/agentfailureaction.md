# AgentFailureAction

The action to take if the agent fails to approve, deny, or reassign the task.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AgentFailureActionApprovalAgentFailureActionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AgentFailureAction("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AgentFailureActionApprovalAgentFailureActionUnspecified`           | APPROVAL_AGENT_FAILURE_ACTION_UNSPECIFIED                           |
| `AgentFailureActionApprovalAgentFailureActionReassignToUsers`       | APPROVAL_AGENT_FAILURE_ACTION_REASSIGN_TO_USERS                     |
| `AgentFailureActionApprovalAgentFailureActionReassignToSuperAdmins` | APPROVAL_AGENT_FAILURE_ACTION_REASSIGN_TO_SUPER_ADMINS              |
| `AgentFailureActionApprovalAgentFailureActionSkipPolicyStep`        | APPROVAL_AGENT_FAILURE_ACTION_SKIP_POLICY_STEP                      |