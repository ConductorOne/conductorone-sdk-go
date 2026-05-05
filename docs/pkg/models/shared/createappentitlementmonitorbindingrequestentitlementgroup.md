# CreateAppEntitlementMonitorBindingRequestEntitlementGroup

Which side of the conflict monitor (A or B) to place this entitlement in.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CreateAppEntitlementMonitorBindingRequestEntitlementGroupEntitlementGroupUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CreateAppEntitlementMonitorBindingRequestEntitlementGroup("custom_value")
```


## Values

| Name                                                                                   | Value                                                                                  |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `CreateAppEntitlementMonitorBindingRequestEntitlementGroupEntitlementGroupUnspecified` | ENTITLEMENT_GROUP_UNSPECIFIED                                                          |
| `CreateAppEntitlementMonitorBindingRequestEntitlementGroupEntitlementGroupA`           | ENTITLEMENT_GROUP_A                                                                    |
| `CreateAppEntitlementMonitorBindingRequestEntitlementGroupEntitlementGroupB`           | ENTITLEMENT_GROUP_B                                                                    |