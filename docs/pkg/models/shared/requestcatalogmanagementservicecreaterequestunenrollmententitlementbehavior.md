# RequestCatalogManagementServiceCreateRequestUnenrollmentEntitlementBehavior

Defines how to handle the revoke policies of the entitlements in the catalog during unenrollment.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RequestCatalogManagementServiceCreateRequestUnenrollmentEntitlementBehaviorRequestCatalogUnenrollmentEntitlementBehaviorUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RequestCatalogManagementServiceCreateRequestUnenrollmentEntitlementBehavior("custom_value")
```


## Values

| Name                                                                                                                                  | Value                                                                                                                                 |
| ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `RequestCatalogManagementServiceCreateRequestUnenrollmentEntitlementBehaviorRequestCatalogUnenrollmentEntitlementBehaviorUnspecified` | REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_UNSPECIFIED                                                                         |
| `RequestCatalogManagementServiceCreateRequestUnenrollmentEntitlementBehaviorRequestCatalogUnenrollmentEntitlementBehaviorBypass`      | REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_BYPASS                                                                              |
| `RequestCatalogManagementServiceCreateRequestUnenrollmentEntitlementBehaviorRequestCatalogUnenrollmentEntitlementBehaviorEnforce`     | REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_ENFORCE                                                                             |