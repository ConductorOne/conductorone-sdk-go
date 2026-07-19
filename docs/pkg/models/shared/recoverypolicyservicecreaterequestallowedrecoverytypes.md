# RecoveryPolicyServiceCreateRequestAllowedRecoveryTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RecoveryPolicyServiceCreateRequestAllowedRecoveryTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RecoveryPolicyServiceCreateRequestAllowedRecoveryTypes("custom_value")
```


## Values

| Name                                                                                     | Value                                                                                    |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `RecoveryPolicyServiceCreateRequestAllowedRecoveryTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                                                              |
| `RecoveryPolicyServiceCreateRequestAllowedRecoveryTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                                                                  |
| `RecoveryPolicyServiceCreateRequestAllowedRecoveryTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                                                                 |
| `RecoveryPolicyServiceCreateRequestAllowedRecoveryTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                                                                     |
| `RecoveryPolicyServiceCreateRequestAllowedRecoveryTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                                                                |
| `RecoveryPolicyServiceCreateRequestAllowedRecoveryTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                                                            |
| `RecoveryPolicyServiceCreateRequestAllowedRecoveryTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                                                         |
| `RecoveryPolicyServiceCreateRequestAllowedRecoveryTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT                                                      |
| `RecoveryPolicyServiceCreateRequestAllowedRecoveryTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                                                             |