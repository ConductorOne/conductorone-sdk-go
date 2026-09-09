# AppUserNhiType

NHI classification when this app user carries the non-human-identity trait.
 Read-only; translated from the model's nhi_trait at the API boundary.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AppUserNhiTypeAppUserNhiTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AppUserNhiType("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `AppUserNhiTypeAppUserNhiTypeUnspecified`     | APP_USER_NHI_TYPE_UNSPECIFIED                 |
| `AppUserNhiTypeAppUserNhiTypeAppRegistration` | APP_USER_NHI_TYPE_APP_REGISTRATION            |
| `AppUserNhiTypeAppUserNhiTypeAssumableRole`   | APP_USER_NHI_TYPE_ASSUMABLE_ROLE              |
| `AppUserNhiTypeAppUserNhiTypeManagedIdentity` | APP_USER_NHI_TYPE_MANAGED_IDENTITY            |