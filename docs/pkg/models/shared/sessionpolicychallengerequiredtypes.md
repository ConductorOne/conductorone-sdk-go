# SessionPolicyChallengeRequiredTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SessionPolicyChallengeRequiredTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SessionPolicyChallengeRequiredTypes("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `SessionPolicyChallengeRequiredTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                                           |
| `SessionPolicyChallengeRequiredTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                                               |
| `SessionPolicyChallengeRequiredTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                                              |
| `SessionPolicyChallengeRequiredTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                                                  |
| `SessionPolicyChallengeRequiredTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                                             |
| `SessionPolicyChallengeRequiredTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                                         |
| `SessionPolicyChallengeRequiredTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                                      |
| `SessionPolicyChallengeRequiredTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT                                   |
| `SessionPolicyChallengeRequiredTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                                          |