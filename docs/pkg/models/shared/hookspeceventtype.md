# HookSpecEventType

Event the stamped hook fires on. When UNSPECIFIED, instantiation falls
 back to post_hook. Required to mint a pre-output hook, which post_hook
 cannot express.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.HookSpecEventTypeHookSpecEventTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.HookSpecEventType("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `HookSpecEventTypeHookSpecEventTypeUnspecified` | HOOK_SPEC_EVENT_TYPE_UNSPECIFIED                |
| `HookSpecEventTypeHookSpecEventTypePreToolUse`  | HOOK_SPEC_EVENT_TYPE_PRE_TOOL_USE               |
| `HookSpecEventTypeHookSpecEventTypePostToolUse` | HOOK_SPEC_EVENT_TYPE_POST_TOOL_USE              |
| `HookSpecEventTypeHookSpecEventTypePreOutput`   | HOOK_SPEC_EVENT_TYPE_PRE_OUTPUT                 |