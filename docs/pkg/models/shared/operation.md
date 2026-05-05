# Operation

Which connector RPC this dispatches to.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.OperationOperationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Operation("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `OperationOperationUnspecified` | OPERATION_UNSPECIFIED           |
| `OperationOperationGrant`       | OPERATION_GRANT                 |