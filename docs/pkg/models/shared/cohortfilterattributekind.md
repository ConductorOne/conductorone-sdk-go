# CohortFilterAttributeKind

The kind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CohortFilterAttributeKindCohortFilterAttributeKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CohortFilterAttributeKind("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `CohortFilterAttributeKindCohortFilterAttributeKindUnspecified` | COHORT_FILTER_ATTRIBUTE_KIND_UNSPECIFIED                        |
| `CohortFilterAttributeKindCohortFilterAttributeKindBuiltin`     | COHORT_FILTER_ATTRIBUTE_KIND_BUILTIN                            |
| `CohortFilterAttributeKindCohortFilterAttributeKindCustom`      | COHORT_FILTER_ATTRIBUTE_KIND_CUSTOM                             |