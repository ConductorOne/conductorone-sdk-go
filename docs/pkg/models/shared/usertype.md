# UserType

The type of the user.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.UserTypeUserTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.UserType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `UserTypeUserTypeUnspecified` | USER_TYPE_UNSPECIFIED         |
| `UserTypeUserTypeSystem`      | USER_TYPE_SYSTEM              |
| `UserTypeUserTypeHuman`       | USER_TYPE_HUMAN               |
| `UserTypeUserTypeService`     | USER_TYPE_SERVICE             |
| `UserTypeUserTypeAgent`       | USER_TYPE_AGENT               |