# UserSchemasStatus

The status of the user in the system.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.UserSchemasStatusUnknown

// Open enum: custom values can be created with a direct type cast
custom := shared.UserSchemasStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `UserSchemasStatusUnknown`  | UNKNOWN                     |
| `UserSchemasStatusEnabled`  | ENABLED                     |
| `UserSchemasStatusDisabled` | DISABLED                    |
| `UserSchemasStatusDeleted`  | DELETED                     |