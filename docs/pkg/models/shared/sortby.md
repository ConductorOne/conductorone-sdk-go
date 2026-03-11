# SortBy

Sort order

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SortBySearchSortByUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SortBy("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `SortBySearchSortByUnspecified` | SEARCH_SORT_BY_UNSPECIFIED      |
| `SortBySearchSortByCreatedDesc` | SEARCH_SORT_BY_CREATED_DESC     |
| `SortBySearchSortByCreatedAsc`  | SEARCH_SORT_BY_CREATED_ASC      |
| `SortBySearchSortByExpiresAsc`  | SEARCH_SORT_BY_EXPIRES_ASC      |
| `SortBySearchSortByNameAsc`     | SEARCH_SORT_BY_NAME_ASC         |