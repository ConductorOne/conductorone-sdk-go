# Visibility

Admin-settable visibility override (how this tool is surfaced to users).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.VisibilityToolVisibilityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Visibility("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `VisibilityToolVisibilityUnspecified` | TOOL_VISIBILITY_UNSPECIFIED           |
| `VisibilityToolVisibilityFeatured`    | TOOL_VISIBILITY_FEATURED              |
| `VisibilityToolVisibilityAvailable`   | TOOL_VISIBILITY_AVAILABLE             |
| `VisibilityToolVisibilityBypassed`    | TOOL_VISIBILITY_BYPASSED              |