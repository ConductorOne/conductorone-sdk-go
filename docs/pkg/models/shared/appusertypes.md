# AppUserTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AppUserTypesAppUserTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AppUserTypes("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `AppUserTypesAppUserTypeUnspecified`    | APP_USER_TYPE_UNSPECIFIED               |
| `AppUserTypesAppUserTypeUser`           | APP_USER_TYPE_USER                      |
| `AppUserTypesAppUserTypeServiceAccount` | APP_USER_TYPE_SERVICE_ACCOUNT           |
| `AppUserTypesAppUserTypeSystemAccount`  | APP_USER_TYPE_SYSTEM_ACCOUNT            |