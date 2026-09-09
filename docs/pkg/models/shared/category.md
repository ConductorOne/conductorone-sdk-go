# Category

How the target type treats the current state.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CategoryRequestCatalogTypeChangeImpactCategoryUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Category("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `CategoryRequestCatalogTypeChangeImpactCategoryUnspecified`    | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_UNSPECIFIED        |
| `CategoryRequestCatalogTypeChangeImpactCategoryBlocksChange`   | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_BLOCKS_CHANGE      |
| `CategoryRequestCatalogTypeChangeImpactCategoryWillDisable`    | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_WILL_DISABLE       |
| `CategoryRequestCatalogTypeChangeImpactCategoryWillRemove`     | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_WILL_REMOVE        |
| `CategoryRequestCatalogTypeChangeImpactCategoryInformational`  | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_INFORMATIONAL      |
| `CategoryRequestCatalogTypeChangeImpactCategoryNeedsAttention` | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_NEEDS_ATTENTION    |
| `CategoryRequestCatalogTypeChangeImpactCategoryNotSupported`   | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_NOT_SUPPORTED      |
| `CategoryRequestCatalogTypeChangeImpactCategoryIgnored`        | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_IGNORED            |