# MCPPromptServiceListResponse

MCPPromptServiceListResponse contains one page of stored prompts.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `NextPageToken`                                               | `*string`                                                     | :heavy_minus_sign:                                            | Token for the next page, or empty for the final page.         |
| `Prompts`                                                     | [][shared.MCPPrompt](../../../pkg/models/shared/mcpprompt.md) | :heavy_minus_sign:                                            | Stored prompts.                                               |