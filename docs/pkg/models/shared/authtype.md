# AuthType

Authentication type for the paper vault recipient (Paper Vault only)

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AuthTypeStoreCredentialAuthTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AuthType("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `AuthTypeStoreCredentialAuthTypeUnspecified` | STORE_CREDENTIAL_AUTH_TYPE_UNSPECIFIED       |
| `AuthTypeStoreCredentialAuthTypeSsoInternal` | STORE_CREDENTIAL_AUTH_TYPE_SSO_INTERNAL      |
| `AuthTypeStoreCredentialAuthTypeVerifyEmail` | STORE_CREDENTIAL_AUTH_TYPE_VERIFY_EMAIL      |