# SSOApplicationServiceCreateRequestSubjectType

How the user's identifier reaches this application. Leave unset to use the
 tenant default.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SSOApplicationServiceCreateRequestSubjectTypeSsoSubjectTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SSOApplicationServiceCreateRequestSubjectType("custom_value")
```


## Values

| Name                                                                       | Value                                                                      |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `SSOApplicationServiceCreateRequestSubjectTypeSsoSubjectTypeUnspecified`   | SSO_SUBJECT_TYPE_UNSPECIFIED                                               |
| `SSOApplicationServiceCreateRequestSubjectTypeSsoSubjectTypePairwise`      | SSO_SUBJECT_TYPE_PAIRWISE                                                  |
| `SSOApplicationServiceCreateRequestSubjectTypeSsoSubjectTypePublic`        | SSO_SUBJECT_TYPE_PUBLIC                                                    |
| `SSOApplicationServiceCreateRequestSubjectTypeSsoSubjectTypeCompatibility` | SSO_SUBJECT_TYPE_COMPATIBILITY                                             |