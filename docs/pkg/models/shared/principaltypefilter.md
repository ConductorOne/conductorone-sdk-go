# PrincipalTypeFilter

Filters principals included in the scope. Unspecified is treated as users.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PrincipalTypeFilterPrincipalTypeFilterUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PrincipalTypeFilter("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `PrincipalTypeFilterPrincipalTypeFilterUnspecified`       | PRINCIPAL_TYPE_FILTER_UNSPECIFIED                         |
| `PrincipalTypeFilterPrincipalTypeFilterUsers`             | PRINCIPAL_TYPE_FILTER_USERS                               |
| `PrincipalTypeFilterPrincipalTypeFilterResources`         | PRINCIPAL_TYPE_FILTER_RESOURCES                           |
| `PrincipalTypeFilterPrincipalTypeFilterUsersAndResources` | PRINCIPAL_TYPE_FILTER_USERS_AND_RESOURCES                 |