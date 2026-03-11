# RequestCatalogManagementServiceCreateRequestEnrollmentBehavior

Defines how to handle the request policies of the entitlements in the catalog during enrollment.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RequestCatalogManagementServiceCreateRequestEnrollmentBehaviorRequestCatalogEnrollmentBehaviorUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RequestCatalogManagementServiceCreateRequestEnrollmentBehavior("custom_value")
```


## Values

| Name                                                                                                                            | Value                                                                                                                           |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `RequestCatalogManagementServiceCreateRequestEnrollmentBehaviorRequestCatalogEnrollmentBehaviorUnspecified`                     | REQUEST_CATALOG_ENROLLMENT_BEHAVIOR_UNSPECIFIED                                                                                 |
| `RequestCatalogManagementServiceCreateRequestEnrollmentBehaviorRequestCatalogEnrollmentBehaviorBypassEntitlementRequestPolicy`  | REQUEST_CATALOG_ENROLLMENT_BEHAVIOR_BYPASS_ENTITLEMENT_REQUEST_POLICY                                                           |
| `RequestCatalogManagementServiceCreateRequestEnrollmentBehaviorRequestCatalogEnrollmentBehaviorEnforceEntitlementRequestPolicy` | REQUEST_CATALOG_ENROLLMENT_BEHAVIOR_ENFORCE_ENTITLEMENT_REQUEST_POLICY                                                          |