# GrantFilterType

The grantFilterType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GrantFilterTypeGrantFilterTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GrantFilterType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `GrantFilterTypeGrantFilterTypeUnspecified` | GRANT_FILTER_TYPE_UNSPECIFIED               |
| `GrantFilterTypeGrantFilterTypePermanent`   | GRANT_FILTER_TYPE_PERMANENT                 |
| `GrantFilterTypeGrantFilterTypeTemporary`   | GRANT_FILTER_TYPE_TEMPORARY                 |