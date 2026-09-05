# PolicyUserSource

Why the policy applies to this user. DIRECT or GROUP.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PolicyUserSourceEffectiveSessionPolicySourceUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PolicyUserSource("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `PolicyUserSourceEffectiveSessionPolicySourceUnspecified`       | EFFECTIVE_SESSION_POLICY_SOURCE_UNSPECIFIED                     |
| `PolicyUserSourceEffectiveSessionPolicySourceDirect`            | EFFECTIVE_SESSION_POLICY_SOURCE_DIRECT                          |
| `PolicyUserSourceEffectiveSessionPolicySourceGroup`             | EFFECTIVE_SESSION_POLICY_SOURCE_GROUP                           |
| `PolicyUserSourceEffectiveSessionPolicySourceTenantDefault`     | EFFECTIVE_SESSION_POLICY_SOURCE_TENANT_DEFAULT                  |
| `PolicyUserSourceEffectiveSessionPolicySourceTenantDefaultNone` | EFFECTIVE_SESSION_POLICY_SOURCE_TENANT_DEFAULT_NONE             |