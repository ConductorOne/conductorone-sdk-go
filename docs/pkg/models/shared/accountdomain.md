# AccountDomain

The accountDomain field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccountDomainAppUserDomainUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountDomain("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `AccountDomainAppUserDomainUnspecified` | APP_USER_DOMAIN_UNSPECIFIED             |
| `AccountDomainAppUserDomainExternal`    | APP_USER_DOMAIN_EXTERNAL                |
| `AccountDomainAppUserDomainTrusted`     | APP_USER_DOMAIN_TRUSTED                 |