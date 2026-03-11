# FilterType

The filterType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FilterTypeAccessProfileFilterTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FilterType("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `FilterTypeAccessProfileFilterTypeUnspecified`     | ACCESS_PROFILE_FILTER_TYPE_UNSPECIFIED             |
| `FilterTypeAccessProfileFilterTypeIncludeAll`      | ACCESS_PROFILE_FILTER_TYPE_INCLUDE_ALL             |
| `FilterTypeAccessProfileFilterTypeExcludeAll`      | ACCESS_PROFILE_FILTER_TYPE_EXCLUDE_ALL             |
| `FilterTypeAccessProfileFilterTypeExcludeSpecific` | ACCESS_PROFILE_FILTER_TYPE_EXCLUDE_SPECIFIC        |
| `FilterTypeAccessProfileFilterTypeIncludeSpecific` | ACCESS_PROFILE_FILTER_TYPE_INCLUDE_SPECIFIC        |