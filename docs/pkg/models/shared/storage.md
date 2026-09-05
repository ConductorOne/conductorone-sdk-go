# Storage

The storage field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StorageFunctionInvocationResultStorageUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Storage("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `StorageFunctionInvocationResultStorageUnspecified` | FUNCTION_INVOCATION_RESULT_STORAGE_UNSPECIFIED      |
| `StorageFunctionInvocationResultStorageInline`      | FUNCTION_INVOCATION_RESULT_STORAGE_INLINE           |
| `StorageFunctionInvocationResultStorageVfs`         | FUNCTION_INVOCATION_RESULT_STORAGE_VFS              |