# EffectiveUserPolicySource

Why the policy applies to the user.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EffectiveUserPolicySourceEffectiveSessionPolicySourceUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EffectiveUserPolicySource("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `EffectiveUserPolicySourceEffectiveSessionPolicySourceUnspecified`       | EFFECTIVE_SESSION_POLICY_SOURCE_UNSPECIFIED                              |
| `EffectiveUserPolicySourceEffectiveSessionPolicySourceDirect`            | EFFECTIVE_SESSION_POLICY_SOURCE_DIRECT                                   |
| `EffectiveUserPolicySourceEffectiveSessionPolicySourceGroup`             | EFFECTIVE_SESSION_POLICY_SOURCE_GROUP                                    |
| `EffectiveUserPolicySourceEffectiveSessionPolicySourceTenantDefault`     | EFFECTIVE_SESSION_POLICY_SOURCE_TENANT_DEFAULT                           |
| `EffectiveUserPolicySourceEffectiveSessionPolicySourceTenantDefaultNone` | EFFECTIVE_SESSION_POLICY_SOURCE_TENANT_DEFAULT_NONE                      |