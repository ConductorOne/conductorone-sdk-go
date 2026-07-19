# TaskAuditRequestDefaultsAppliedSource

The source field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TaskAuditRequestDefaultsAppliedSourceAppliedSettingsSourceUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskAuditRequestDefaultsAppliedSource("custom_value")
```


## Values

| Name                                                                             | Value                                                                            |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `TaskAuditRequestDefaultsAppliedSourceAppliedSettingsSourceUnspecified`          | APPLIED_SETTINGS_SOURCE_UNSPECIFIED                                              |
| `TaskAuditRequestDefaultsAppliedSourceAppliedSettingsSourceEntitlementOverride`  | APPLIED_SETTINGS_SOURCE_ENTITLEMENT_OVERRIDE                                     |
| `TaskAuditRequestDefaultsAppliedSourceAppliedSettingsSourceRoutingRule`          | APPLIED_SETTINGS_SOURCE_ROUTING_RULE                                             |
| `TaskAuditRequestDefaultsAppliedSourceAppliedSettingsSourceEntitlementDefault`   | APPLIED_SETTINGS_SOURCE_ENTITLEMENT_DEFAULT                                      |
| `TaskAuditRequestDefaultsAppliedSourceAppliedSettingsSourceAppDefault`           | APPLIED_SETTINGS_SOURCE_APP_DEFAULT                                              |
| `TaskAuditRequestDefaultsAppliedSourceAppliedSettingsSourceAccessProfileDefault` | APPLIED_SETTINGS_SOURCE_ACCESS_PROFILE_DEFAULT                                   |