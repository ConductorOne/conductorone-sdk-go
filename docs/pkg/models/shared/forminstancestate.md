# FormInstanceState

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FormInstanceStateFormInstanceStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FormInstanceState("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `FormInstanceStateFormInstanceStateUnspecified` | FORM_INSTANCE_STATE_UNSPECIFIED                 |
| `FormInstanceStateFormInstanceStateWaiting`     | FORM_INSTANCE_STATE_WAITING                     |
| `FormInstanceStateFormInstanceStateDone`        | FORM_INSTANCE_STATE_DONE                        |