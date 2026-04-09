# ValidationMode

Validation approach. See MicrosoftValidationMode for details on each mode.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ValidationModeMicrosoftValidationModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ValidationMode("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `ValidationModeMicrosoftValidationModeUnspecified` | MICROSOFT_VALIDATION_MODE_UNSPECIFIED              |
| `ValidationModeMicrosoftValidationModeAcrs`        | MICROSOFT_VALIDATION_MODE_ACRS                     |
| `ValidationModeMicrosoftValidationModeOidc`        | MICROSOFT_VALIDATION_MODE_OIDC                     |