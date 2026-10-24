# ChallengeRequiredTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ChallengeRequiredTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ChallengeRequiredTypes("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `ChallengeRequiredTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                              |
| `ChallengeRequiredTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                                  |
| `ChallengeRequiredTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                                 |
| `ChallengeRequiredTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                                     |
| `ChallengeRequiredTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                                |
| `ChallengeRequiredTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                            |
| `ChallengeRequiredTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                         |
| `ChallengeRequiredTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT                      |
| `ChallengeRequiredTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                             |