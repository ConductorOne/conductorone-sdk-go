# AllowedRecoveryTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AllowedRecoveryTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AllowedRecoveryTypes("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `AllowedRecoveryTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                            |
| `AllowedRecoveryTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                                |
| `AllowedRecoveryTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                               |
| `AllowedRecoveryTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                                   |
| `AllowedRecoveryTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                              |
| `AllowedRecoveryTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                          |
| `AllowedRecoveryTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                       |
| `AllowedRecoveryTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT                    |
| `AllowedRecoveryTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                           |