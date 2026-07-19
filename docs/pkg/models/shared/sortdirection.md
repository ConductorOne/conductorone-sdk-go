# SortDirection

Direction for sort_by. UNSPECIFIED means ascending.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SortDirectionSortDirectionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SortDirection("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `SortDirectionSortDirectionUnspecified` | SORT_DIRECTION_UNSPECIFIED              |
| `SortDirectionSortDirectionAsc`         | SORT_DIRECTION_ASC                      |
| `SortDirectionSortDirectionDesc`        | SORT_DIRECTION_DESC                     |