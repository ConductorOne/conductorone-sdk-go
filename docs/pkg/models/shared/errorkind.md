# ErrorKind

Classification of last_error. Unspecified when no classified error is present.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ErrorKindErrorKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ErrorKind("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `ErrorKindErrorKindUnspecified`                 | ERROR_KIND_UNSPECIFIED                          |
| `ErrorKindErrorKindAuthoredRuntimeServeRefused` | ERROR_KIND_AUTHORED_RUNTIME_SERVE_REFUSED       |