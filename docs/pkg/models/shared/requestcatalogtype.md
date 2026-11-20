# RequestCatalogType

The type of this access profile: CATALOG (shown as Catalog in the UI),
 BUNDLE (Role), or CATALOG_AND_BUNDLE (Hybrid). UNSPECIFIED means the
 profile has not been assigned a type yet; it behaves as Hybrid.

 To change the type, call Update with "type" in the update mask and send
 CATALOG, BUNDLE, or CATALOG_AND_BUNDLE. Settings the new type doesn't
 support must be turned off first; PlanTypeChange lists them. Changing the
 type requires access profile types to be enabled for your tenant.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RequestCatalogTypeRequestCatalogTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RequestCatalogType("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `RequestCatalogTypeRequestCatalogTypeUnspecified`      | REQUEST_CATALOG_TYPE_UNSPECIFIED                       |
| `RequestCatalogTypeRequestCatalogTypeCatalog`          | REQUEST_CATALOG_TYPE_CATALOG                           |
| `RequestCatalogTypeRequestCatalogTypeProfile`          | REQUEST_CATALOG_TYPE_PROFILE                           |
| `RequestCatalogTypeRequestCatalogTypeCatalogAndBundle` | REQUEST_CATALOG_TYPE_CATALOG_AND_BUNDLE                |
| `RequestCatalogTypeRequestCatalogTypeBundle`           | REQUEST_CATALOG_TYPE_BUNDLE                            |