# SignInPolicyServiceCreateRequestAllowedPrimaryTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SignInPolicyServiceCreateRequestAllowedPrimaryTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SignInPolicyServiceCreateRequestAllowedPrimaryTypes("custom_value")
```


## Values

| Name                                                                                  | Value                                                                                 |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `SignInPolicyServiceCreateRequestAllowedPrimaryTypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED                                                           |
| `SignInPolicyServiceCreateRequestAllowedPrimaryTypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                                                               |
| `SignInPolicyServiceCreateRequestAllowedPrimaryTypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                                                              |
| `SignInPolicyServiceCreateRequestAllowedPrimaryTypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                                                                  |
| `SignInPolicyServiceCreateRequestAllowedPrimaryTypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP                                                             |
| `SignInPolicyServiceCreateRequestAllowedPrimaryTypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE                                                         |
| `SignInPolicyServiceCreateRequestAllowedPrimaryTypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE                                                      |
| `SignInPolicyServiceCreateRequestAllowedPrimaryTypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT                                                   |
| `SignInPolicyServiceCreateRequestAllowedPrimaryTypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP                                                          |