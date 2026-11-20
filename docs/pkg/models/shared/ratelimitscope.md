# RateLimitScope

Whose requests the limit counts.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RateLimitScopeScopeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RateLimitScope("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `RateLimitScopeScopeUnspecified`          | SCOPE_UNSPECIFIED                         |
| `RateLimitScopeScopeOrganization`         | SCOPE_ORGANIZATION                        |
| `RateLimitScopeScopeUser`                 | SCOPE_USER                                |
| `RateLimitScopeScopeAPIKey`               | SCOPE_API_KEY                             |
| `RateLimitScopeScopeUnauthenticated`      | SCOPE_UNAUTHENTICATED                     |
| `RateLimitScopeScopeUnauthenticatedPerIP` | SCOPE_UNAUTHENTICATED_PER_IP              |