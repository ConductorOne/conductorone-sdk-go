# SortField

Column to sort by. Unspecified (0) keeps the server's default order (app, then display name).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SortFieldAppResourceSortFieldUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SortField("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `SortFieldAppResourceSortFieldUnspecified`     | APP_RESOURCE_SORT_FIELD_UNSPECIFIED            |
| `SortFieldAppResourceSortFieldSecretCreatedAt` | APP_RESOURCE_SORT_FIELD_SECRET_CREATED_AT      |
| `SortFieldAppResourceSortFieldSecretExpiresAt` | APP_RESOURCE_SORT_FIELD_SECRET_EXPIRES_AT      |
| `SortFieldAppResourceSortFieldLastUsedAt`      | APP_RESOURCE_SORT_FIELD_LAST_USED_AT           |