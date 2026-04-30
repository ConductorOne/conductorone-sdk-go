# HooksServiceCreateRequestEvent

The event field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.HooksServiceCreateRequestEventHookEventTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.HooksServiceCreateRequestEvent("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `HooksServiceCreateRequestEventHookEventTypeUnspecified` | HOOK_EVENT_TYPE_UNSPECIFIED                              |
| `HooksServiceCreateRequestEventHookEventTypePreToolUse`  | HOOK_EVENT_TYPE_PRE_TOOL_USE                             |
| `HooksServiceCreateRequestEventHookEventTypePostToolUse` | HOOK_EVENT_TYPE_POST_TOOL_USE                            |