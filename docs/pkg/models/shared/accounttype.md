# AccountType

The accountType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccountTypeAppUserTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountType("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `AccountTypeAppUserTypeUnspecified`    | APP_USER_TYPE_UNSPECIFIED              |
| `AccountTypeAppUserTypeUser`           | APP_USER_TYPE_USER                     |
| `AccountTypeAppUserTypeServiceAccount` | APP_USER_TYPE_SERVICE_ACCOUNT          |
| `AccountTypeAppUserTypeSystemAccount`  | APP_USER_TYPE_SYSTEM_ACCOUNT           |