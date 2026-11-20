# AccessReviewServiceCreateRequestScopeType

How the campaign's scope is defined, such as by entitlements or by users. Treated as
 ACCESS_REVIEW_SCOPE_TYPE_BY_ENTITLEMENTS when unset. Some scope types require a feature to be
 enabled for your tenant. See AccessReviewScopeType for the available values.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccessReviewServiceCreateRequestScopeType("custom_value")
```


## Values

| Name                                                                              | Value                                                                             |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeUnspecified`       | ACCESS_REVIEW_SCOPE_TYPE_UNSPECIFIED                                              |
| `AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeByEntitlements`    | ACCESS_REVIEW_SCOPE_TYPE_BY_ENTITLEMENTS                                          |
| `AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeByAccessConflicts` | ACCESS_REVIEW_SCOPE_TYPE_BY_ACCESS_CONFLICTS                                      |
| `AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeByResource`        | ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE                                              |
| `AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeByInheritance`     | ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE                                           |
| `AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeByUsers`           | ACCESS_REVIEW_SCOPE_TYPE_BY_USERS                                                 |
| `AccessReviewServiceCreateRequestScopeTypeAccessReviewScopeTypeByAccessProfiles`  | ACCESS_REVIEW_SCOPE_TYPE_BY_ACCESS_PROFILES                                       |