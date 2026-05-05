# CreateAppEntitlementRequestPurpose

The purpose of the entitlement (e.g., assignment, permission, ownership).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CreateAppEntitlementRequestPurposeAppEntitlementPurposeValueUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CreateAppEntitlementRequestPurpose("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `CreateAppEntitlementRequestPurposeAppEntitlementPurposeValueUnspecified` | APP_ENTITLEMENT_PURPOSE_VALUE_UNSPECIFIED                                 |
| `CreateAppEntitlementRequestPurposeAppEntitlementPurposeValueAssignment`  | APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT                                  |
| `CreateAppEntitlementRequestPurposeAppEntitlementPurposeValuePermission`  | APP_ENTITLEMENT_PURPOSE_VALUE_PERMISSION                                  |
| `CreateAppEntitlementRequestPurposeAppEntitlementPurposeValueOwnership`   | APP_ENTITLEMENT_PURPOSE_VALUE_OWNERSHIP                                   |