# ExcludeTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ExcludeTypesUserTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ExcludeTypes("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ExcludeTypesUserTypeUnspecified` | USER_TYPE_UNSPECIFIED             |
| `ExcludeTypesUserTypeSystem`      | USER_TYPE_SYSTEM                  |
| `ExcludeTypesUserTypeHuman`       | USER_TYPE_HUMAN                   |
| `ExcludeTypesUserTypeService`     | USER_TYPE_SERVICE                 |
| `ExcludeTypesUserTypeAgent`       | USER_TYPE_AGENT                   |