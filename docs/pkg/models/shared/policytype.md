# PolicyType

The type of policy to create (grant, revoke, or certify).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PolicyTypePolicyTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PolicyType("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `PolicyTypePolicyTypeUnspecified`   | POLICY_TYPE_UNSPECIFIED             |
| `PolicyTypePolicyTypeGrant`         | POLICY_TYPE_GRANT                   |
| `PolicyTypePolicyTypeRevoke`        | POLICY_TYPE_REVOKE                  |
| `PolicyTypePolicyTypeCertify`       | POLICY_TYPE_CERTIFY                 |
| `PolicyTypePolicyTypeAccessRequest` | POLICY_TYPE_ACCESS_REQUEST          |
| `PolicyTypePolicyTypeProvision`     | POLICY_TYPE_PROVISION               |