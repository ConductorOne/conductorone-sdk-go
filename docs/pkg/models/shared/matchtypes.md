# MatchTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MatchTypesAccessProfileMatchTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MatchTypes("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `MatchTypesAccessProfileMatchTypeUnspecified` | ACCESS_PROFILE_MATCH_TYPE_UNSPECIFIED         |
| `MatchTypesAccessProfileMatchTypeExact`       | ACCESS_PROFILE_MATCH_TYPE_EXACT               |
| `MatchTypesAccessProfileMatchTypeSuperset`    | ACCESS_PROFILE_MATCH_TYPE_SUPERSET            |
| `MatchTypesAccessProfileMatchTypePartial`     | ACCESS_PROFILE_MATCH_TYPE_PARTIAL             |