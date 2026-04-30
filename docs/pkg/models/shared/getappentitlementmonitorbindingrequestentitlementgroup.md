# GetAppEntitlementMonitorBindingRequestEntitlementGroup

Which side of the conflict monitor (A or B) this binding belongs to.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GetAppEntitlementMonitorBindingRequestEntitlementGroupEntitlementGroupUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GetAppEntitlementMonitorBindingRequestEntitlementGroup("custom_value")
```


## Values

| Name                                                                                | Value                                                                               |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `GetAppEntitlementMonitorBindingRequestEntitlementGroupEntitlementGroupUnspecified` | ENTITLEMENT_GROUP_UNSPECIFIED                                                       |
| `GetAppEntitlementMonitorBindingRequestEntitlementGroupEntitlementGroupA`           | ENTITLEMENT_GROUP_A                                                                 |
| `GetAppEntitlementMonitorBindingRequestEntitlementGroupEntitlementGroupB`           | ENTITLEMENT_GROUP_B                                                                 |