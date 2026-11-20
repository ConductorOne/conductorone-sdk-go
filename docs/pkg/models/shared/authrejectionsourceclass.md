# AuthRejectionSourceClass

Set only when outcome is AUTH_REJECTED. Class of the peer address.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AuthRejectionSourceClassTbAuthRejectionSourceClassUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AuthRejectionSourceClass("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AuthRejectionSourceClassTbAuthRejectionSourceClassUnspecified` | TB_AUTH_REJECTION_SOURCE_CLASS_UNSPECIFIED                      |
| `AuthRejectionSourceClassTbAuthRejectionSourceClassLoopback`    | TB_AUTH_REJECTION_SOURCE_CLASS_LOOPBACK                         |
| `AuthRejectionSourceClassTbAuthRejectionSourceClassPrivate`     | TB_AUTH_REJECTION_SOURCE_CLASS_PRIVATE                          |
| `AuthRejectionSourceClassTbAuthRejectionSourceClassPublic`      | TB_AUTH_REJECTION_SOURCE_CLASS_PUBLIC                           |
| `AuthRejectionSourceClassTbAuthRejectionSourceClassUnknown`     | TB_AUTH_REJECTION_SOURCE_CLASS_UNKNOWN                          |