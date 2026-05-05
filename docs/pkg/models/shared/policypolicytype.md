# PolicyPolicyType

The type of this policy (grant, revoke, or certify). The lowercased type
 name (e.g., "grant") is also the key for the baseline entry in policy_steps.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PolicyPolicyTypePolicyTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PolicyPolicyType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `PolicyPolicyTypePolicyTypeUnspecified`   | POLICY_TYPE_UNSPECIFIED                   |
| `PolicyPolicyTypePolicyTypeGrant`         | POLICY_TYPE_GRANT                         |
| `PolicyPolicyTypePolicyTypeRevoke`        | POLICY_TYPE_REVOKE                        |
| `PolicyPolicyTypePolicyTypeCertify`       | POLICY_TYPE_CERTIFY                       |
| `PolicyPolicyTypePolicyTypeAccessRequest` | POLICY_TYPE_ACCESS_REQUEST                |
| `PolicyPolicyTypePolicyTypeProvision`     | POLICY_TYPE_PROVISION                     |