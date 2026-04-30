# TargetedAppUserTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TargetedAppUserTypesAppUserTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TargetedAppUserTypes("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `TargetedAppUserTypesAppUserTypeUnspecified`    | APP_USER_TYPE_UNSPECIFIED                       |
| `TargetedAppUserTypesAppUserTypeUser`           | APP_USER_TYPE_USER                              |
| `TargetedAppUserTypesAppUserTypeServiceAccount` | APP_USER_TYPE_SERVICE_ACCOUNT                   |
| `TargetedAppUserTypesAppUserTypeSystemAccount`  | APP_USER_TYPE_SYSTEM_ACCOUNT                    |