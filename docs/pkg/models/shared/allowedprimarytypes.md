# AllowedPrimaryTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AllowedPrimaryTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AllowedPrimaryTypes("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `AllowedPrimaryTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                           |
| `AllowedPrimaryTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                               |
| `AllowedPrimaryTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                              |
| `AllowedPrimaryTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                                  |
| `AllowedPrimaryTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                             |
| `AllowedPrimaryTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                         |
| `AllowedPrimaryTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                      |
| `AllowedPrimaryTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT                   |
| `AllowedPrimaryTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                          |