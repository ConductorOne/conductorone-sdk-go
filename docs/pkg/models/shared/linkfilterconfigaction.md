# LinkFilterConfigAction

Action taken on a disallowed link. Unspecified = REDACT.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.LinkFilterConfigActionLinkFilterActionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.LinkFilterConfigAction("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `LinkFilterConfigActionLinkFilterActionUnspecified` | LINK_FILTER_ACTION_UNSPECIFIED                      |
| `LinkFilterConfigActionLinkFilterActionRedact`      | LINK_FILTER_ACTION_REDACT                           |
| `LinkFilterConfigActionLinkFilterActionAnnotate`    | LINK_FILTER_ACTION_ANNOTATE                         |