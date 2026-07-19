# CredentialTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CredentialTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CredentialTypes("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `CredentialTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                       |
| `CredentialTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                           |
| `CredentialTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                          |
| `CredentialTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                              |
| `CredentialTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                         |
| `CredentialTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                     |
| `CredentialTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                  |
| `CredentialTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT               |
| `CredentialTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                      |