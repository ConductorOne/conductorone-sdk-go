# UserStatuses

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.UserStatusesUserUnknown

// Open enum: custom values can be created with a direct type cast
custom := shared.UserStatuses("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `UserStatusesUserUnknown`  | USER_UNKNOWN               |
| `UserStatusesUserEnabled`  | USER_ENABLED               |
| `UserStatusesUserDisabled` | USER_DISABLED              |
| `UserStatusesUserDeleted`  | USER_DELETED               |