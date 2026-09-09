# RequestCatalogType

The type of this access profile. Reports CATALOG_AND_BUNDLE for a profile
 created before the type was recorded; UNSPECIFIED only for a tenant whose
 backfill has not been run. Updates require the access profile types feature
 and an update mask containing "type".

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