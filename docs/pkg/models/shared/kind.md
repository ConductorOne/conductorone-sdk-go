# Kind

The kind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.KindAuthorityKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Kind("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `KindAuthorityKindUnspecified`        | AUTHORITY_KIND_UNSPECIFIED            |
| `KindAuthorityKindFundPolicy`         | AUTHORITY_KIND_FUND_POLICY            |
| `KindAuthorityKindFundAssignment`     | AUTHORITY_KIND_FUND_ASSIGNMENT        |
| `KindAuthorityKindFundRule`           | AUTHORITY_KIND_FUND_RULE              |
| `KindAuthorityKindTenantAppCap`       | AUTHORITY_KIND_TENANT_APP_CAP         |
| `KindAuthorityKindSubjectAppLimit`    | AUTHORITY_KIND_SUBJECT_APP_LIMIT      |
| `KindAuthorityKindEntitlementBinding` | AUTHORITY_KIND_ENTITLEMENT_BINDING    |
| `KindAuthorityKindAppUser`            | AUTHORITY_KIND_APP_USER               |
| `KindAuthorityKindUser`               | AUTHORITY_KIND_USER                   |