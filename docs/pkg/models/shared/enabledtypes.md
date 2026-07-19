# EnabledTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EnabledTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EnabledTypes("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `EnabledTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                    |
| `EnabledTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                        |
| `EnabledTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                       |
| `EnabledTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                           |
| `EnabledTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                      |
| `EnabledTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                  |
| `EnabledTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE               |
| `EnabledTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT            |
| `EnabledTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                   |