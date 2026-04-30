# JustificationVisibility

Controls whether the justification field is shown or hidden on the request form.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.JustificationVisibilityJustificationVisibilityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.JustificationVisibility("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `JustificationVisibilityJustificationVisibilityUnspecified` | JUSTIFICATION_VISIBILITY_UNSPECIFIED                        |
| `JustificationVisibilityJustificationVisibilityShow`        | JUSTIFICATION_VISIBILITY_SHOW                               |
| `JustificationVisibilityJustificationVisibilityHide`        | JUSTIFICATION_VISIBILITY_HIDE                               |