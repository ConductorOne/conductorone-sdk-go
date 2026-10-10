# AuthRejectionReason

Keeps only rejections with this reason. Unset keeps all.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

value := operations.AuthRejectionReasonTbAuthRejectionReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := operations.AuthRejectionReason("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AuthRejectionReasonTbAuthRejectionReasonUnspecified`           | TB_AUTH_REJECTION_REASON_UNSPECIFIED                            |
| `AuthRejectionReasonTbAuthRejectionReasonMissingCredential`     | TB_AUTH_REJECTION_REASON_MISSING_CREDENTIAL                     |
| `AuthRejectionReasonTbAuthRejectionReasonInvalidToken`          | TB_AUTH_REJECTION_REASON_INVALID_TOKEN                          |
| `AuthRejectionReasonTbAuthRejectionReasonExpiredToken`          | TB_AUTH_REJECTION_REASON_EXPIRED_TOKEN                          |
| `AuthRejectionReasonTbAuthRejectionReasonUnknownTenant`         | TB_AUTH_REJECTION_REASON_UNKNOWN_TENANT                         |
| `AuthRejectionReasonTbAuthRejectionReasonAudienceMismatch`      | TB_AUTH_REJECTION_REASON_AUDIENCE_MISMATCH                      |
| `AuthRejectionReasonTbAuthRejectionReasonInsufficientScope`     | TB_AUTH_REJECTION_REASON_INSUFFICIENT_SCOPE                     |
| `AuthRejectionReasonTbAuthRejectionReasonInvalidProof`          | TB_AUTH_REJECTION_REASON_INVALID_PROOF                          |
| `AuthRejectionReasonTbAuthRejectionReasonCredentialUnavailable` | TB_AUTH_REJECTION_REASON_CREDENTIAL_UNAVAILABLE                 |