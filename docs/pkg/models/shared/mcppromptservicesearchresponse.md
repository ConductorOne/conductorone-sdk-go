# MCPPromptServiceSearchResponse

MCPPromptServiceSearchResponse contains matching prompt metadata.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `List`                                                        | [][shared.MCPPrompt](../../../pkg/models/shared/mcpprompt.md) | :heavy_minus_sign:                                            | Matching prompts.                                             |
| `NextPageToken`                                               | `*string`                                                     | :heavy_minus_sign:                                            | Token for the next page, or empty for the final page.         |