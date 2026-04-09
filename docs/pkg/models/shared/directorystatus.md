# DirectoryStatus

The status of the user in the directory.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DirectoryStatusUnknown

// Open enum: custom values can be created with a direct type cast
custom := shared.DirectoryStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `DirectoryStatusUnknown`  | UNKNOWN                   |
| `DirectoryStatusEnabled`  | ENABLED                   |
| `DirectoryStatusDisabled` | DISABLED                  |
| `DirectoryStatusDeleted`  | DELETED                   |