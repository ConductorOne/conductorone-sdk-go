# RequiredAgeSuite

Exact Age suite required for this submission. UNSPECIFIED preserves legacy X25519 behavior.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RequiredAgeSuiteAgeSuiteUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RequiredAgeSuite("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `RequiredAgeSuiteAgeSuiteUnspecified`    | AGE_SUITE_UNSPECIFIED                    |
| `RequiredAgeSuiteAgeSuiteX25519`         | AGE_SUITE_X25519                         |
| `RequiredAgeSuiteAgeSuiteMlkem768X25519` | AGE_SUITE_MLKEM768X25519                 |