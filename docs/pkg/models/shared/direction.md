# Direction

Direction to sort in. Unspecified falls back to ASC when sort_field is set.
 No defined_only validation here: protoc-gen-validate mis-resolves the
 cross-package enum name map to this file's c1.models.app.v1 import alias
 instead of c1.api.search.v1, which fails to compile. The query builder
 already treats any unrecognized value as ASC, so this is safe to omit.

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