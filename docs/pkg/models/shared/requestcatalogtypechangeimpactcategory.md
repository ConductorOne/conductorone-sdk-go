# RequestCatalogTypeChangeImpactCategory

How the target type treats the current state.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RequestCatalogTypeChangeImpactCategoryRequestCatalogTypeChangeImpactCategoryUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RequestCatalogTypeChangeImpactCategory("custom_value")
```


## Values

| Name                                                                                         | Value                                                                                        |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `RequestCatalogTypeChangeImpactCategoryRequestCatalogTypeChangeImpactCategoryUnspecified`    | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_UNSPECIFIED                                      |
| `RequestCatalogTypeChangeImpactCategoryRequestCatalogTypeChangeImpactCategoryBlocksChange`   | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_BLOCKS_CHANGE                                    |
| `RequestCatalogTypeChangeImpactCategoryRequestCatalogTypeChangeImpactCategoryWillDisable`    | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_WILL_DISABLE                                     |
| `RequestCatalogTypeChangeImpactCategoryRequestCatalogTypeChangeImpactCategoryWillRemove`     | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_WILL_REMOVE                                      |
| `RequestCatalogTypeChangeImpactCategoryRequestCatalogTypeChangeImpactCategoryInformational`  | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_INFORMATIONAL                                    |
| `RequestCatalogTypeChangeImpactCategoryRequestCatalogTypeChangeImpactCategoryNeedsAttention` | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_NEEDS_ATTENTION                                  |
| `RequestCatalogTypeChangeImpactCategoryRequestCatalogTypeChangeImpactCategoryNotSupported`   | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_NOT_SUPPORTED                                    |
| `RequestCatalogTypeChangeImpactCategoryRequestCatalogTypeChangeImpactCategoryIgnored`        | REQUEST_CATALOG_TYPE_CHANGE_IMPACT_CATEGORY_IGNORED                                          |