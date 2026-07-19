# SignInPolicyServiceCreateRequestAllowedMfaTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SignInPolicyServiceCreateRequestAllowedMfaTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SignInPolicyServiceCreateRequestAllowedMfaTypes("custom_value")
```


## Values

| Name                                                                              | Value                                                                             |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `SignInPolicyServiceCreateRequestAllowedMfaTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                                                       |
| `SignInPolicyServiceCreateRequestAllowedMfaTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                                                           |
| `SignInPolicyServiceCreateRequestAllowedMfaTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                                                          |
| `SignInPolicyServiceCreateRequestAllowedMfaTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                                                              |
| `SignInPolicyServiceCreateRequestAllowedMfaTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                                                         |
| `SignInPolicyServiceCreateRequestAllowedMfaTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                                                     |
| `SignInPolicyServiceCreateRequestAllowedMfaTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                                                  |
| `SignInPolicyServiceCreateRequestAllowedMfaTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT                                               |
| `SignInPolicyServiceCreateRequestAllowedMfaTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                                                      |