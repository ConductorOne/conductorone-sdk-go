# CurrentAccountType

The currentAccountType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CurrentAccountTypeAppUserTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CurrentAccountType("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CurrentAccountTypeAppUserTypeUnspecified`    | APP_USER_TYPE_UNSPECIFIED                     |
| `CurrentAccountTypeAppUserTypeUser`           | APP_USER_TYPE_USER                            |
| `CurrentAccountTypeAppUserTypeServiceAccount` | APP_USER_TYPE_SERVICE_ACCOUNT                 |
| `CurrentAccountTypeAppUserTypeSystemAccount`  | APP_USER_TYPE_SYSTEM_ACCOUNT                  |