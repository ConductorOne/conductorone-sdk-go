# Direction

Direction to sort in. Unspecified falls back to ASC when sort_field is set;
 when sort_field is also unspecified, the server default order (created_at
 DESC) applies.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DirectionSortDirectionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Direction("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `DirectionSortDirectionUnspecified` | SORT_DIRECTION_UNSPECIFIED          |
| `DirectionSortDirectionAsc`         | SORT_DIRECTION_ASC                  |
| `DirectionSortDirectionDesc`        | SORT_DIRECTION_DESC                 |