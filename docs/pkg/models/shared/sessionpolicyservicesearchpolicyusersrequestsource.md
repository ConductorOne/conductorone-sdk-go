# SessionPolicyServiceSearchPolicyUsersRequestSource

When set, restrict results to this source. UNSPECIFIED returns all users.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SessionPolicyServiceSearchPolicyUsersRequestSourceEffectiveSessionPolicySourceFilterUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SessionPolicyServiceSearchPolicyUsersRequestSource("custom_value")
```


## Values

| Name                                                                                              | Value                                                                                             |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `SessionPolicyServiceSearchPolicyUsersRequestSourceEffectiveSessionPolicySourceFilterUnspecified` | EFFECTIVE_SESSION_POLICY_SOURCE_FILTER_UNSPECIFIED                                                |
| `SessionPolicyServiceSearchPolicyUsersRequestSourceEffectiveSessionPolicySourceFilterDirect`      | EFFECTIVE_SESSION_POLICY_SOURCE_FILTER_DIRECT                                                     |
| `SessionPolicyServiceSearchPolicyUsersRequestSourceEffectiveSessionPolicySourceFilterGroup`       | EFFECTIVE_SESSION_POLICY_SOURCE_FILTER_GROUP                                                      |