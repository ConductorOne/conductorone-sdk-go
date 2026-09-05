# Event

The event field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EventHookEventTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Event("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `EventHookEventTypeUnspecified` | HOOK_EVENT_TYPE_UNSPECIFIED     |
| `EventHookEventTypePreToolUse`  | HOOK_EVENT_TYPE_PRE_TOOL_USE    |
| `EventHookEventTypePostToolUse` | HOOK_EVENT_TYPE_POST_TOOL_USE   |
| `EventHookEventTypePreOutput`   | HOOK_EVENT_TYPE_PRE_OUTPUT      |