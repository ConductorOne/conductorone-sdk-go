# PaperSecretServiceCreateInternalRequestRequiredAgeSuite

Exact Age suite required for this submission. UNSPECIFIED preserves legacy X25519 behavior.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PaperSecretServiceCreateInternalRequestRequiredAgeSuiteAgeSuiteUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PaperSecretServiceCreateInternalRequestRequiredAgeSuite("custom_value")
```


## Values

| Name                                                                            | Value                                                                           |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `PaperSecretServiceCreateInternalRequestRequiredAgeSuiteAgeSuiteUnspecified`    | AGE_SUITE_UNSPECIFIED                                                           |
| `PaperSecretServiceCreateInternalRequestRequiredAgeSuiteAgeSuiteX25519`         | AGE_SUITE_X25519                                                                |
| `PaperSecretServiceCreateInternalRequestRequiredAgeSuiteAgeSuiteMlkem768X25519` | AGE_SUITE_MLKEM768X25519                                                        |