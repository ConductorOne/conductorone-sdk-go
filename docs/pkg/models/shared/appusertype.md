# AppUserType

The appplication user type. Type can be user, system or service.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AppUserTypeAppUserTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AppUserType("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `AppUserTypeAppUserTypeUnspecified`    | APP_USER_TYPE_UNSPECIFIED              |
| `AppUserTypeAppUserTypeUser`           | APP_USER_TYPE_USER                     |
| `AppUserTypeAppUserTypeServiceAccount` | APP_USER_TYPE_SERVICE_ACCOUNT          |
| `AppUserTypeAppUserTypeSystemAccount`  | APP_USER_TYPE_SYSTEM_ACCOUNT           |