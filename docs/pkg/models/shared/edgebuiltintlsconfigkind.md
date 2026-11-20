# EdgeBuiltinTLSConfigKind

The kind of managed material. Only EDGE_BUILTIN_TLS_KIND_EDGE_CA is
 supported.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EdgeBuiltinTLSConfigKindEdgeBuiltinTLSKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EdgeBuiltinTLSConfigKind("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `EdgeBuiltinTLSConfigKindEdgeBuiltinTLSKindUnspecified` | EDGE_BUILTIN_TLS_KIND_UNSPECIFIED                       |
| `EdgeBuiltinTLSConfigKindEdgeBuiltinTLSKindEdgeCa`      | EDGE_BUILTIN_TLS_KIND_EDGE_CA                           |