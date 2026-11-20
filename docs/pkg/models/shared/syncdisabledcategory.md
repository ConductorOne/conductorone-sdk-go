# SyncDisabledCategory

Current sync-disabled category; not an exact transition history.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SyncDisabledCategoryDisabledCategoryUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SyncDisabledCategory("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `SyncDisabledCategoryDisabledCategoryUnspecified`    | DISABLED_CATEGORY_UNSPECIFIED                        |
| `SyncDisabledCategoryDisabledCategoryOperations`     | DISABLED_CATEGORY_OPERATIONS                         |
| `SyncDisabledCategoryDisabledCategoryCustomerOptOut` | DISABLED_CATEGORY_CUSTOMER_OPT_OUT                   |
| `SyncDisabledCategoryDisabledCategoryDeployment`     | DISABLED_CATEGORY_DEPLOYMENT                         |
| `SyncDisabledCategoryDisabledCategoryDataAnomaly`    | DISABLED_CATEGORY_DATA_ANOMALY                       |