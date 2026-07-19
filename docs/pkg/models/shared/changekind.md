# ChangeKind

Storage-model enum re-exported here for wire compatibility with the
 storage row. UNSPECIFIED should never appear on the wire.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ChangeKindChangeKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ChangeKind("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ChangeKindChangeKindUnspecified` | CHANGE_KIND_UNSPECIFIED           |
| `ChangeKindChangeKindCreate`      | CHANGE_KIND_CREATE                |
| `ChangeKindChangeKindPut`         | CHANGE_KIND_PUT                   |
| `ChangeKindChangeKindHardDelete`  | CHANGE_KIND_HARD_DELETE           |