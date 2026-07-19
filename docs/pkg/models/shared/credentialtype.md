# CredentialType

The credentialType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CredentialTypeCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CredentialType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `CredentialTypeCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                      |
| `CredentialTypeCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                          |
| `CredentialTypeCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                         |
| `CredentialTypeCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                             |
| `CredentialTypeCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                        |
| `CredentialTypeCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                    |
| `CredentialTypeCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                 |
| `CredentialTypeCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT              |
| `CredentialTypeCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                     |