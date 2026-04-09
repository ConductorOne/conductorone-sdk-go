# AccountTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccountTypesAppUserTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountTypes("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `AccountTypesAppUserTypeUnspecified`    | APP_USER_TYPE_UNSPECIFIED               |
| `AccountTypesAppUserTypeUser`           | APP_USER_TYPE_USER                      |
| `AccountTypesAppUserTypeServiceAccount` | APP_USER_TYPE_SERVICE_ACCOUNT           |
| `AccountTypesAppUserTypeSystemAccount`  | APP_USER_TYPE_SYSTEM_ACCOUNT            |