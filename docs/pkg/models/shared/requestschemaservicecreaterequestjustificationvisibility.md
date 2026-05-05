# RequestSchemaServiceCreateRequestJustificationVisibility

Controls whether the justification field is shown or hidden on the request form.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RequestSchemaServiceCreateRequestJustificationVisibilityJustificationVisibilityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RequestSchemaServiceCreateRequestJustificationVisibility("custom_value")
```


## Values

| Name                                                                                         | Value                                                                                        |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `RequestSchemaServiceCreateRequestJustificationVisibilityJustificationVisibilityUnspecified` | JUSTIFICATION_VISIBILITY_UNSPECIFIED                                                         |
| `RequestSchemaServiceCreateRequestJustificationVisibilityJustificationVisibilityShow`        | JUSTIFICATION_VISIBILITY_SHOW                                                                |
| `RequestSchemaServiceCreateRequestJustificationVisibilityJustificationVisibilityHide`        | JUSTIFICATION_VISIBILITY_HIDE                                                                |