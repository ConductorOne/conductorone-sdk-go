# SubjectType

How the user's identifier reaches this application.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SubjectTypeSsoSubjectTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SubjectType("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `SubjectTypeSsoSubjectTypeUnspecified`   | SSO_SUBJECT_TYPE_UNSPECIFIED             |
| `SubjectTypeSsoSubjectTypePairwise`      | SSO_SUBJECT_TYPE_PAIRWISE                |
| `SubjectTypeSsoSubjectTypePublic`        | SSO_SUBJECT_TYPE_PUBLIC                  |
| `SubjectTypeSsoSubjectTypeCompatibility` | SSO_SUBJECT_TYPE_COMPATIBILITY           |