# AgeSuite

Exact Age suite used by the stored ciphertext.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AgeSuiteAgeSuiteUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AgeSuite("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `AgeSuiteAgeSuiteUnspecified`    | AGE_SUITE_UNSPECIFIED            |
| `AgeSuiteAgeSuiteX25519`         | AGE_SUITE_X25519                 |
| `AgeSuiteAgeSuiteMlkem768X25519` | AGE_SUITE_MLKEM768X25519         |