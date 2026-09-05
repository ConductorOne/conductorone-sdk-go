# SessionPolicyServiceGetEffectiveSessionPolicyResponseSource

Why the policy applies.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SessionPolicyServiceGetEffectiveSessionPolicyResponseSourceEffectiveSessionPolicySourceUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SessionPolicyServiceGetEffectiveSessionPolicyResponseSource("custom_value")
```


## Values

| Name                                                                                                       | Value                                                                                                      |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `SessionPolicyServiceGetEffectiveSessionPolicyResponseSourceEffectiveSessionPolicySourceUnspecified`       | EFFECTIVE_SESSION_POLICY_SOURCE_UNSPECIFIED                                                                |
| `SessionPolicyServiceGetEffectiveSessionPolicyResponseSourceEffectiveSessionPolicySourceDirect`            | EFFECTIVE_SESSION_POLICY_SOURCE_DIRECT                                                                     |
| `SessionPolicyServiceGetEffectiveSessionPolicyResponseSourceEffectiveSessionPolicySourceGroup`             | EFFECTIVE_SESSION_POLICY_SOURCE_GROUP                                                                      |
| `SessionPolicyServiceGetEffectiveSessionPolicyResponseSourceEffectiveSessionPolicySourceTenantDefault`     | EFFECTIVE_SESSION_POLICY_SOURCE_TENANT_DEFAULT                                                             |
| `SessionPolicyServiceGetEffectiveSessionPolicyResponseSourceEffectiveSessionPolicySourceTenantDefaultNone` | EFFECTIVE_SESSION_POLICY_SOURCE_TENANT_DEFAULT_NONE                                                        |