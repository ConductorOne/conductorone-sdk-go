# Types

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Types("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `TypesCredentialTypeUnspecified`        | CREDENTIAL_TYPE_UNSPECIFIED             |
| `TypesCredentialTypePasskey`            | CREDENTIAL_TYPE_PASSKEY                 |
| `TypesCredentialTypePassword`           | CREDENTIAL_TYPE_PASSWORD                |
| `TypesCredentialTypeTotp`               | CREDENTIAL_TYPE_TOTP                    |
| `TypesCredentialTypeEmailOtp`           | CREDENTIAL_TYPE_EMAIL_OTP               |
| `TypesCredentialTypeRecoveryCode`       | CREDENTIAL_TYPE_RECOVERY_CODE           |
| `TypesCredentialTypeDelegatedGoogle`    | CREDENTIAL_TYPE_DELEGATED_GOOGLE        |
| `TypesCredentialTypeDelegatedMicrosoft` | CREDENTIAL_TYPE_DELEGATED_MICROSOFT     |
| `TypesCredentialTypeUpstreamIdp`        | CREDENTIAL_TYPE_UPSTREAM_IDP            |