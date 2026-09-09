# DefaultSubjectType

The subject type materialized onto new SSO applications that do not choose
 one. Changing this default does not change existing applications. When
 unset, the server uses pairwise subjects.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DefaultSubjectTypeSsoSubjectTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DefaultSubjectType("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `DefaultSubjectTypeSsoSubjectTypeUnspecified`   | SSO_SUBJECT_TYPE_UNSPECIFIED                    |
| `DefaultSubjectTypeSsoSubjectTypePairwise`      | SSO_SUBJECT_TYPE_PAIRWISE                       |
| `DefaultSubjectTypeSsoSubjectTypePublic`        | SSO_SUBJECT_TYPE_PUBLIC                         |
| `DefaultSubjectTypeSsoSubjectTypeCompatibility` | SSO_SUBJECT_TYPE_COMPATIBILITY                  |