# Format

Output format for the report. When unspecified, programmatic public-API
 callers (REST gateway and MCP) get JSON and the in-app UI gets XLSX. JSON
 and CSV return the per-decision certification rows; XLSX returns the full
 multi-sheet Excel workbook.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FormatAccessReviewReportFormatUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Format("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `FormatAccessReviewReportFormatUnspecified` | ACCESS_REVIEW_REPORT_FORMAT_UNSPECIFIED     |
| `FormatAccessReviewReportFormatXlsx`        | ACCESS_REVIEW_REPORT_FORMAT_XLSX            |
| `FormatAccessReviewReportFormatJSON`        | ACCESS_REVIEW_REPORT_FORMAT_JSON            |
| `FormatAccessReviewReportFormatCsv`         | ACCESS_REVIEW_REPORT_FORMAT_CSV             |