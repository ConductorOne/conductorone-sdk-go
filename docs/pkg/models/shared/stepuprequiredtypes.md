# StepUpRequiredTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StepUpRequiredTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.StepUpRequiredTypes("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `StepUpRequiredTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                           |
| `StepUpRequiredTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                               |
| `StepUpRequiredTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                              |
| `StepUpRequiredTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                                  |
| `StepUpRequiredTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                             |
| `StepUpRequiredTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                         |
| `StepUpRequiredTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                      |
| `StepUpRequiredTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT                   |
| `StepUpRequiredTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                          |