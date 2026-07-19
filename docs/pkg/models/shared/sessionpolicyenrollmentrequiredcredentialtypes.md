# SessionPolicyEnrollmentRequiredCredentialTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SessionPolicyEnrollmentRequiredCredentialTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SessionPolicyEnrollmentRequiredCredentialTypes("custom_value")
```


## Values

| Name                                                                             | Value                                                                            |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `SessionPolicyEnrollmentRequiredCredentialTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                                                      |
| `SessionPolicyEnrollmentRequiredCredentialTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                                                          |
| `SessionPolicyEnrollmentRequiredCredentialTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                                                         |
| `SessionPolicyEnrollmentRequiredCredentialTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                                                             |
| `SessionPolicyEnrollmentRequiredCredentialTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                                                        |
| `SessionPolicyEnrollmentRequiredCredentialTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                                                    |
| `SessionPolicyEnrollmentRequiredCredentialTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                                                 |
| `SessionPolicyEnrollmentRequiredCredentialTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT                                              |
| `SessionPolicyEnrollmentRequiredCredentialTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                                                     |