# UnenrollmentEntitlementBehavior

Defines how to handle the revoke policies of the entitlements in the catalog during unenrollment.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.UnenrollmentEntitlementBehaviorRequestCatalogUnenrollmentEntitlementBehaviorUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.UnenrollmentEntitlementBehavior("custom_value")
```


## Values

| Name                                                                                      | Value                                                                                     |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `UnenrollmentEntitlementBehaviorRequestCatalogUnenrollmentEntitlementBehaviorUnspecified` | REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_UNSPECIFIED                             |
| `UnenrollmentEntitlementBehaviorRequestCatalogUnenrollmentEntitlementBehaviorBypass`      | REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_BYPASS                                  |
| `UnenrollmentEntitlementBehaviorRequestCatalogUnenrollmentEntitlementBehaviorEnforce`     | REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_ENFORCE                                 |