# SessionPolicyStepUpRequiredTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SessionPolicyStepUpRequiredTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SessionPolicyStepUpRequiredTypes("custom_value")
```


## Values

| Name                                                               | Value                                                              |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `SessionPolicyStepUpRequiredTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                                        |
| `SessionPolicyStepUpRequiredTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                                            |
| `SessionPolicyStepUpRequiredTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                                           |
| `SessionPolicyStepUpRequiredTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                                               |
| `SessionPolicyStepUpRequiredTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                                          |
| `SessionPolicyStepUpRequiredTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                                      |
| `SessionPolicyStepUpRequiredTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                                   |
| `SessionPolicyStepUpRequiredTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT                                |
| `SessionPolicyStepUpRequiredTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                                       |