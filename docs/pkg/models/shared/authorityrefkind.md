# AuthorityRefKind

Type of configuration.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AuthorityRefKindAuthorityKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AuthorityRefKind("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `AuthorityRefKindAuthorityKindUnspecified`        | AUTHORITY_KIND_UNSPECIFIED                        |
| `AuthorityRefKindAuthorityKindFundPolicy`         | AUTHORITY_KIND_FUND_POLICY                        |
| `AuthorityRefKindAuthorityKindFundAssignment`     | AUTHORITY_KIND_FUND_ASSIGNMENT                    |
| `AuthorityRefKindAuthorityKindFundRule`           | AUTHORITY_KIND_FUND_RULE                          |
| `AuthorityRefKindAuthorityKindTenantAppCap`       | AUTHORITY_KIND_TENANT_APP_CAP                     |
| `AuthorityRefKindAuthorityKindSubjectAppLimit`    | AUTHORITY_KIND_SUBJECT_APP_LIMIT                  |
| `AuthorityRefKindAuthorityKindEntitlementBinding` | AUTHORITY_KIND_ENTITLEMENT_BINDING                |
| `AuthorityRefKindAuthorityKindAppUser`            | AUTHORITY_KIND_APP_USER                           |
| `AuthorityRefKindAuthorityKindUser`               | AUTHORITY_KIND_USER                               |