# CatalogMode

How the apply pass treats catalog (access profile) bindings on matched
 entitlements. UNSPECIFIED reads as legacy list management; UNMANAGED means
 the rule does not manage the list and only the app's default request
 catalog follows the self-service switch.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CatalogModeAppAccessRequestDefaultsCatalogModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CatalogMode("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CatalogModeAppAccessRequestDefaultsCatalogModeUnspecified` | APP_ACCESS_REQUEST_DEFAULTS_CATALOG_MODE_UNSPECIFIED        |
| `CatalogModeAppAccessRequestDefaultsCatalogModeManaged`     | APP_ACCESS_REQUEST_DEFAULTS_CATALOG_MODE_MANAGED            |
| `CatalogModeAppAccessRequestDefaultsCatalogModeUnmanaged`   | APP_ACCESS_REQUEST_DEFAULTS_CATALOG_MODE_UNMANAGED          |