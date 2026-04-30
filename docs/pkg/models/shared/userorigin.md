# UserOrigin

The origin of the user, describing who owns the user's lifecycle.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.UserOriginUserOriginUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.UserOrigin("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `UserOriginUserOriginUnspecified` | USER_ORIGIN_UNSPECIFIED           |
| `UserOriginUserOriginDirectory`   | USER_ORIGIN_DIRECTORY             |
| `UserOriginUserOriginLocal`       | USER_ORIGIN_LOCAL                 |
| `UserOriginUserOriginSystem`      | USER_ORIGIN_SYSTEM                |