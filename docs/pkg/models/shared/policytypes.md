# PolicyTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PolicyTypesPolicyTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PolicyTypes("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `PolicyTypesPolicyTypeUnspecified`   | POLICY_TYPE_UNSPECIFIED              |
| `PolicyTypesPolicyTypeGrant`         | POLICY_TYPE_GRANT                    |
| `PolicyTypesPolicyTypeRevoke`        | POLICY_TYPE_REVOKE                   |
| `PolicyTypesPolicyTypeCertify`       | POLICY_TYPE_CERTIFY                  |
| `PolicyTypesPolicyTypeAccessRequest` | POLICY_TYPE_ACCESS_REQUEST           |
| `PolicyTypesPolicyTypeProvision`     | POLICY_TYPE_PROVISION                |