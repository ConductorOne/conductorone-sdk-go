# Cause

The cause field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CauseEdgeEgressRuleEntitlementIssueCauseUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Cause("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `CauseEdgeEgressRuleEntitlementIssueCauseUnspecified` | EDGE_EGRESS_RULE_ENTITLEMENT_ISSUE_CAUSE_UNSPECIFIED  |
| `CauseEdgeEgressRuleEntitlementIssueCauseDeleted`     | EDGE_EGRESS_RULE_ENTITLEMENT_ISSUE_CAUSE_DELETED      |
| `CauseEdgeEgressRuleEntitlementIssueCauseMissing`     | EDGE_EGRESS_RULE_ENTITLEMENT_ISSUE_CAUSE_MISSING      |
| `CauseEdgeEgressRuleEntitlementIssueCauseOverCap`     | EDGE_EGRESS_RULE_ENTITLEMENT_ISSUE_CAUSE_OVER_CAP     |