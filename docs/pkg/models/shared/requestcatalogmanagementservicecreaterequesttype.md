# RequestCatalogManagementServiceCreateRequestType

The type of access profile to create. Leave unset for
 REQUEST_CATALOG_TYPE_CATALOG_AND_BUNDLE, which is what every profile
 created before this field existed is. Setting it requires the
 ACCESS_PROFILE_TYPES feature.

 PROFILE is rejected rather than resolved: it is deprecated, has no stored
 counterpart, and shares wire number 2 with the stored BUNDLE, so honoring
 it would silently persist a type the caller did not ask for.

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