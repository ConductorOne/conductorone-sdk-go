# AllowedMfaTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AllowedMfaTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AllowedMfaTypes("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `AllowedMfaTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                       |
| `AllowedMfaTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                           |
| `AllowedMfaTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                          |
| `AllowedMfaTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                              |
| `AllowedMfaTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                         |
| `AllowedMfaTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                     |
| `AllowedMfaTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                  |
| `AllowedMfaTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT               |
| `AllowedMfaTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                      |