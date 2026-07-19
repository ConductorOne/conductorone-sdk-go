# MCPAccessProfileToolBindingHistoryItemChangeKind

Whether this binding was added or removed.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MCPAccessProfileToolBindingHistoryItemChangeKindListChangeKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MCPAccessProfileToolBindingHistoryItemChangeKind("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `MCPAccessProfileToolBindingHistoryItemChangeKindListChangeKindUnspecified` | LIST_CHANGE_KIND_UNSPECIFIED                                                |
| `MCPAccessProfileToolBindingHistoryItemChangeKindListChangeKindAdded`       | LIST_CHANGE_KIND_ADDED                                                      |
| `MCPAccessProfileToolBindingHistoryItemChangeKindListChangeKindRemoved`     | LIST_CHANGE_KIND_REMOVED                                                    |