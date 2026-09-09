# AuthorityKind

Type of configuration that supplied the denying limit.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AuthorityKindAuthorityKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AuthorityKind("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `AuthorityKindAuthorityKindUnspecified`        | AUTHORITY_KIND_UNSPECIFIED                     |
| `AuthorityKindAuthorityKindFundPolicy`         | AUTHORITY_KIND_FUND_POLICY                     |
| `AuthorityKindAuthorityKindFundAssignment`     | AUTHORITY_KIND_FUND_ASSIGNMENT                 |
| `AuthorityKindAuthorityKindFundRule`           | AUTHORITY_KIND_FUND_RULE                       |
| `AuthorityKindAuthorityKindTenantAppCap`       | AUTHORITY_KIND_TENANT_APP_CAP                  |
| `AuthorityKindAuthorityKindSubjectAppLimit`    | AUTHORITY_KIND_SUBJECT_APP_LIMIT               |
| `AuthorityKindAuthorityKindEntitlementBinding` | AUTHORITY_KIND_ENTITLEMENT_BINDING             |
| `AuthorityKindAuthorityKindAppUser`            | AUTHORITY_KIND_APP_USER                        |
| `AuthorityKindAuthorityKindUser`               | AUTHORITY_KIND_USER                            |