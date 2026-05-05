# ProvisionInstanceState

This property indicates the current state of this step.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ProvisionInstanceStateProvisionInstanceStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ProvisionInstanceState("custom_value")
```


## Values

| Name                                                                          | Value                                                                         |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ProvisionInstanceStateProvisionInstanceStateUnspecified`                     | PROVISION_INSTANCE_STATE_UNSPECIFIED                                          |
| `ProvisionInstanceStateProvisionInstanceStateInit`                            | PROVISION_INSTANCE_STATE_INIT                                                 |
| `ProvisionInstanceStateProvisionInstanceStateCreateConnectorActionsForTarget` | PROVISION_INSTANCE_STATE_CREATE_CONNECTOR_ACTIONS_FOR_TARGET                  |
| `ProvisionInstanceStateProvisionInstanceStateSendingNotifications`            | PROVISION_INSTANCE_STATE_SENDING_NOTIFICATIONS                                |
| `ProvisionInstanceStateProvisionInstanceStateWaiting`                         | PROVISION_INSTANCE_STATE_WAITING                                              |
| `ProvisionInstanceStateProvisionInstanceStateWebhook`                         | PROVISION_INSTANCE_STATE_WEBHOOK                                              |
| `ProvisionInstanceStateProvisionInstanceStateWebhookWaiting`                  | PROVISION_INSTANCE_STATE_WEBHOOK_WAITING                                      |
| `ProvisionInstanceStateProvisionInstanceStateExternalTicket`                  | PROVISION_INSTANCE_STATE_EXTERNAL_TICKET                                      |
| `ProvisionInstanceStateProvisionInstanceStateExternalTicketWaiting`           | PROVISION_INSTANCE_STATE_EXTERNAL_TICKET_WAITING                              |
| `ProvisionInstanceStateProvisionInstanceStateAccountLifecycleActions`         | PROVISION_INSTANCE_STATE_ACCOUNT_LIFECYCLE_ACTIONS                            |
| `ProvisionInstanceStateProvisionInstanceStateAccountLifecycleActionsWaiting`  | PROVISION_INSTANCE_STATE_ACCOUNT_LIFECYCLE_ACTIONS_WAITING                    |
| `ProvisionInstanceStateProvisionInstanceStateDone`                            | PROVISION_INSTANCE_STATE_DONE                                                 |