# UserStatus

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.UserStatusUnknown

// Open enum: custom values can be created with a direct type cast
custom := shared.UserStatus("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `UserStatusUnknown`  | UNKNOWN              |
| `UserStatusEnabled`  | ENABLED              |
| `UserStatusDisabled` | DISABLED             |
| `UserStatusDeleted`  | DELETED              |