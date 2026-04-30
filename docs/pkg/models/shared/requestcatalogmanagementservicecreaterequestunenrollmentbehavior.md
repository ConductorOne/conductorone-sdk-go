# RequestCatalogManagementServiceCreateRequestUnenrollmentBehavior

Defines how to handle the revocation of the entitlements in the catalog during unenrollment.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RequestCatalogManagementServiceCreateRequestUnenrollmentBehaviorRequestCatalogUnenrollmentBehaviorUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RequestCatalogManagementServiceCreateRequestUnenrollmentBehavior("custom_value")
```


## Values

| Name                                                                                                                  | Value                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `RequestCatalogManagementServiceCreateRequestUnenrollmentBehaviorRequestCatalogUnenrollmentBehaviorUnspecified`       | REQUEST_CATALOG_UNENROLLMENT_BEHAVIOR_UNSPECIFIED                                                                     |
| `RequestCatalogManagementServiceCreateRequestUnenrollmentBehaviorRequestCatalogUnenrollmentBehaviorLeaveAccessAsIs`   | REQUEST_CATALOG_UNENROLLMENT_BEHAVIOR_LEAVE_ACCESS_AS_IS                                                              |
| `RequestCatalogManagementServiceCreateRequestUnenrollmentBehaviorRequestCatalogUnenrollmentBehaviorRevokeAll`         | REQUEST_CATALOG_UNENROLLMENT_BEHAVIOR_REVOKE_ALL                                                                      |
| `RequestCatalogManagementServiceCreateRequestUnenrollmentBehaviorRequestCatalogUnenrollmentBehaviorRevokeUnjustified` | REQUEST_CATALOG_UNENROLLMENT_BEHAVIOR_REVOKE_UNJUSTIFIED                                                              |