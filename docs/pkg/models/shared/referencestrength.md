# ReferenceStrength

GrantReasonReferenceStrength is used to indicate the strength of the reference to the reason.
 This is used to determine if a grant should be removed when all strong reasons are removed.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ReferenceStrengthGrantReasonReferenceStrengthUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ReferenceStrength("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ReferenceStrengthGrantReasonReferenceStrengthUnspecified` | GRANT_REASON_REFERENCE_STRENGTH_UNSPECIFIED                |
| `ReferenceStrengthGrantReasonReferenceStrengthWeak`        | GRANT_REASON_REFERENCE_STRENGTH_WEAK                       |
| `ReferenceStrengthGrantReasonReferenceStrengthStrong`      | GRANT_REASON_REFERENCE_STRENGTH_STRONG                     |