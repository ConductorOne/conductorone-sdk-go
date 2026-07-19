# DefaultVisibility

System-managed default visibility from MCP config (set during discovery).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DefaultVisibilityToolVisibilityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DefaultVisibility("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `DefaultVisibilityToolVisibilityUnspecified` | TOOL_VISIBILITY_UNSPECIFIED                  |
| `DefaultVisibilityToolVisibilityFeatured`    | TOOL_VISIBILITY_FEATURED                     |
| `DefaultVisibilityToolVisibilityAvailable`   | TOOL_VISIBILITY_AVAILABLE                    |
| `DefaultVisibilityToolVisibilityBypassed`    | TOOL_VISIBILITY_BYPASSED                     |