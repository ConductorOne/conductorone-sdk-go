# Purpose

The purpose of this entitlement (e.g., assignment, permission, ownership).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PurposeAppEntitlementPurposeValueUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Purpose("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `PurposeAppEntitlementPurposeValueUnspecified` | APP_ENTITLEMENT_PURPOSE_VALUE_UNSPECIFIED      |
| `PurposeAppEntitlementPurposeValueAssignment`  | APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT       |
| `PurposeAppEntitlementPurposeValuePermission`  | APP_ENTITLEMENT_PURPOSE_VALUE_PERMISSION       |
| `PurposeAppEntitlementPurposeValueOwnership`   | APP_ENTITLEMENT_PURPOSE_VALUE_OWNERSHIP        |