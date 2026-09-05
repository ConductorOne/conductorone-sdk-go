# TargetType

Requested type. UNSPECIFIED and the deprecated PROFILE value are rejected.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TargetTypeRequestCatalogTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TargetType("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `TargetTypeRequestCatalogTypeUnspecified`      | REQUEST_CATALOG_TYPE_UNSPECIFIED               |
| `TargetTypeRequestCatalogTypeCatalog`          | REQUEST_CATALOG_TYPE_CATALOG                   |
| `TargetTypeRequestCatalogTypeProfile`          | REQUEST_CATALOG_TYPE_PROFILE                   |
| `TargetTypeRequestCatalogTypeCatalogAndBundle` | REQUEST_CATALOG_TYPE_CATALOG_AND_BUNDLE        |
| `TargetTypeRequestCatalogTypeBundle`           | REQUEST_CATALOG_TYPE_BUNDLE                    |