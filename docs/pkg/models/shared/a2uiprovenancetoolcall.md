# A2UIProvenanceToolCall

A2UIProvenanceToolCall is one tool call extracted from the transcript.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CalledAt`                                                  | [*time.Time](https://pkg.go.dev/time#Time)                  | :heavy_minus_sign:                                          | N/A                                                         |
| `InputDigest`                                               | `*string`                                                   | :heavy_minus_sign:                                          | Leading characters of the tool input, whitespace-collapsed. |
| `ToolName`                                                  | `*string`                                                   | :heavy_minus_sign:                                          | The toolName field.                                         |