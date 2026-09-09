# Operation

The operation field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.OperationA2UIProvenanceOperationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Operation("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `OperationA2UIProvenanceOperationUnspecified`   | A2UI_PROVENANCE_OPERATION_UNSPECIFIED           |
| `OperationA2UIProvenanceOperationLookedUp`      | A2UI_PROVENANCE_OPERATION_LOOKED_UP             |
| `OperationA2UIProvenanceOperationCounted`       | A2UI_PROVENANCE_OPERATION_COUNTED               |
| `OperationA2UIProvenanceOperationFetchedRecord` | A2UI_PROVENANCE_OPERATION_FETCHED_RECORD        |
| `OperationA2UIProvenanceOperationSearched`      | A2UI_PROVENANCE_OPERATION_SEARCHED              |
| `OperationA2UIProvenanceOperationReadTrend`     | A2UI_PROVENANCE_OPERATION_READ_TREND            |
| `OperationA2UIProvenanceOperationCreated`       | A2UI_PROVENANCE_OPERATION_CREATED               |
| `OperationA2UIProvenanceOperationUpdated`       | A2UI_PROVENANCE_OPERATION_UPDATED               |
| `OperationA2UIProvenanceOperationDeleted`       | A2UI_PROVENANCE_OPERATION_DELETED               |
| `OperationA2UIProvenanceOperationRanProgram`    | A2UI_PROVENANCE_OPERATION_RAN_PROGRAM           |
| `OperationA2UIProvenanceOperationBuiltReport`   | A2UI_PROVENANCE_OPERATION_BUILT_REPORT          |