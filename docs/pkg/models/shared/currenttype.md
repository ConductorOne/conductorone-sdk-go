# CurrentType

The profile's current type. UNSPECIFIED means the profile has not been
 assigned a type yet; it behaves as Hybrid.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CurrentTypeRequestCatalogTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CurrentType("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `CurrentTypeRequestCatalogTypeUnspecified`      | REQUEST_CATALOG_TYPE_UNSPECIFIED                |
| `CurrentTypeRequestCatalogTypeCatalog`          | REQUEST_CATALOG_TYPE_CATALOG                    |
| `CurrentTypeRequestCatalogTypeProfile`          | REQUEST_CATALOG_TYPE_PROFILE                    |
| `CurrentTypeRequestCatalogTypeCatalogAndBundle` | REQUEST_CATALOG_TYPE_CATALOG_AND_BUNDLE         |
| `CurrentTypeRequestCatalogTypeBundle`           | REQUEST_CATALOG_TYPE_BUNDLE                     |