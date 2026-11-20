# RequestCatalogManagementServiceCreateRequestType

The type of access profile to create: CATALOG (shown as Catalog in the
 UI), BUNDLE (Role), or CATALOG_AND_BUNDLE (Hybrid). Omit it or send
 UNSPECIFIED to create a Hybrid profile.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RequestCatalogManagementServiceCreateRequestTypeRequestCatalogTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RequestCatalogManagementServiceCreateRequestType("custom_value")
```


## Values

| Name                                                                                 | Value                                                                                |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `RequestCatalogManagementServiceCreateRequestTypeRequestCatalogTypeUnspecified`      | REQUEST_CATALOG_TYPE_UNSPECIFIED                                                     |
| `RequestCatalogManagementServiceCreateRequestTypeRequestCatalogTypeCatalog`          | REQUEST_CATALOG_TYPE_CATALOG                                                         |
| `RequestCatalogManagementServiceCreateRequestTypeRequestCatalogTypeProfile`          | REQUEST_CATALOG_TYPE_PROFILE                                                         |
| `RequestCatalogManagementServiceCreateRequestTypeRequestCatalogTypeCatalogAndBundle` | REQUEST_CATALOG_TYPE_CATALOG_AND_BUNDLE                                              |
| `RequestCatalogManagementServiceCreateRequestTypeRequestCatalogTypeBundle`           | REQUEST_CATALOG_TYPE_BUNDLE                                                          |