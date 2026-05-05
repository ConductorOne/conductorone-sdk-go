# WebhookInstanceState

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.WebhookInstanceStateWebhookStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.WebhookInstanceState("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `WebhookInstanceStateWebhookStateUnspecified`     | WEBHOOK_STATE_UNSPECIFIED                         |
| `WebhookInstanceStateWebhookStatePending`         | WEBHOOK_STATE_PENDING                             |
| `WebhookInstanceStateWebhookStateRunning`         | WEBHOOK_STATE_RUNNING                             |
| `WebhookInstanceStateWebhookStateError`           | WEBHOOK_STATE_ERROR                               |
| `WebhookInstanceStateWebhookStateWaitingCallback` | WEBHOOK_STATE_WAITING_CALLBACK                    |
| `WebhookInstanceStateWebhookStateProcessResponse` | WEBHOOK_STATE_PROCESS_RESPONSE                    |
| `WebhookInstanceStateWebhookStateSuccess`         | WEBHOOK_STATE_SUCCESS                             |
| `WebhookInstanceStateWebhookStateFatalError`      | WEBHOOK_STATE_FATAL_ERROR                         |