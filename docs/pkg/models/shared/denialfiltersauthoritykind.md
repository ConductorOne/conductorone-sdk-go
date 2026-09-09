# DenialFiltersAuthorityKind

Optional authority type to match. Set authority_id with this field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DenialFiltersAuthorityKindAuthorityKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DenialFiltersAuthorityKind("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `DenialFiltersAuthorityKindAuthorityKindUnspecified`        | AUTHORITY_KIND_UNSPECIFIED                                  |
| `DenialFiltersAuthorityKindAuthorityKindFundPolicy`         | AUTHORITY_KIND_FUND_POLICY                                  |
| `DenialFiltersAuthorityKindAuthorityKindFundAssignment`     | AUTHORITY_KIND_FUND_ASSIGNMENT                              |
| `DenialFiltersAuthorityKindAuthorityKindFundRule`           | AUTHORITY_KIND_FUND_RULE                                    |
| `DenialFiltersAuthorityKindAuthorityKindTenantAppCap`       | AUTHORITY_KIND_TENANT_APP_CAP                               |
| `DenialFiltersAuthorityKindAuthorityKindSubjectAppLimit`    | AUTHORITY_KIND_SUBJECT_APP_LIMIT                            |
| `DenialFiltersAuthorityKindAuthorityKindEntitlementBinding` | AUTHORITY_KIND_ENTITLEMENT_BINDING                          |
| `DenialFiltersAuthorityKindAuthorityKindAppUser`            | AUTHORITY_KIND_APP_USER                                     |
| `DenialFiltersAuthorityKindAuthorityKindUser`               | AUTHORITY_KIND_USER                                         |