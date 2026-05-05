# MatchType

The matchType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MatchTypeAccessProfileMatchTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MatchType("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `MatchTypeAccessProfileMatchTypeUnspecified` | ACCESS_PROFILE_MATCH_TYPE_UNSPECIFIED        |
| `MatchTypeAccessProfileMatchTypeExact`       | ACCESS_PROFILE_MATCH_TYPE_EXACT              |
| `MatchTypeAccessProfileMatchTypeSuperset`    | ACCESS_PROFILE_MATCH_TYPE_SUPERSET           |
| `MatchTypeAccessProfileMatchTypePartial`     | ACCESS_PROFILE_MATCH_TYPE_PARTIAL            |