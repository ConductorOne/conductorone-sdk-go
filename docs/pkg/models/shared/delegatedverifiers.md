# DelegatedVerifiers

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DelegatedVerifiersDelegatedVerifierTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DelegatedVerifiers("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `DelegatedVerifiersDelegatedVerifierTypeUnspecified` | DELEGATED_VERIFIER_TYPE_UNSPECIFIED                  |
| `DelegatedVerifiersDelegatedVerifierTypeGoogle`      | DELEGATED_VERIFIER_TYPE_GOOGLE                       |
| `DelegatedVerifiersDelegatedVerifierTypeMicrosoft`   | DELEGATED_VERIFIER_TYPE_MICROSOFT                    |
| `DelegatedVerifiersDelegatedVerifierTypeGithub`      | DELEGATED_VERIFIER_TYPE_GITHUB                       |