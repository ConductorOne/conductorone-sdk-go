# DetectedAccountType

What the detector thinks the account actually is.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DetectedAccountTypeAppUserTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DetectedAccountType("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `DetectedAccountTypeAppUserTypeUnspecified`    | APP_USER_TYPE_UNSPECIFIED                      |
| `DetectedAccountTypeAppUserTypeUser`           | APP_USER_TYPE_USER                             |
| `DetectedAccountTypeAppUserTypeServiceAccount` | APP_USER_TYPE_SERVICE_ACCOUNT                  |
| `DetectedAccountTypeAppUserTypeSystemAccount`  | APP_USER_TYPE_SYSTEM_ACCOUNT                   |