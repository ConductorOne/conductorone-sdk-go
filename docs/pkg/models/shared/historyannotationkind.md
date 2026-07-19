# HistoryAnnotationKind

The kind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.HistoryAnnotationKindAnnotationKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.HistoryAnnotationKind("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `HistoryAnnotationKindAnnotationKindUnspecified` | ANNOTATION_KIND_UNSPECIFIED                      |
| `HistoryAnnotationKindAnnotationKindGeneric`     | ANNOTATION_KIND_GENERIC                          |
| `HistoryAnnotationKindAnnotationKindTicket`      | ANNOTATION_KIND_TICKET                           |
| `HistoryAnnotationKindAnnotationKindReason`      | ANNOTATION_KIND_REASON                           |
| `HistoryAnnotationKindAnnotationKindWorkflow`    | ANNOTATION_KIND_WORKFLOW                         |
| `HistoryAnnotationKindAnnotationKindBatch`       | ANNOTATION_KIND_BATCH                            |
| `HistoryAnnotationKindAnnotationKindCorrelation` | ANNOTATION_KIND_CORRELATION                      |
| `HistoryAnnotationKindAnnotationKindAutomation`  | ANNOTATION_KIND_AUTOMATION                       |