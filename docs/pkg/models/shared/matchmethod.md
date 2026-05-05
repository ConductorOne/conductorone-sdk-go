# MatchMethod

How the upstream subject was resolved to a ConductorOne user.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MatchMethodSsfSubjectMatchMethodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MatchMethod("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `MatchMethodSsfSubjectMatchMethodUnspecified`   | SSF_SUBJECT_MATCH_METHOD_UNSPECIFIED            |
| `MatchMethodSsfSubjectMatchMethodIdpUser`       | SSF_SUBJECT_MATCH_METHOD_IDP_USER               |
| `MatchMethodSsfSubjectMatchMethodEmail`         | SSF_SUBJECT_MATCH_METHOD_EMAIL                  |
| `MatchMethodSsfSubjectMatchMethodNotFound`      | SSF_SUBJECT_MATCH_METHOD_NOT_FOUND              |
| `MatchMethodSsfSubjectMatchMethodNotApplicable` | SSF_SUBJECT_MATCH_METHOD_NOT_APPLICABLE         |