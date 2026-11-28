# EdgeInferenceVertexChatConfig

EdgeInferenceVertexChatConfig selects the Google Cloud project and location for Vertex AI MaaS Chat.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Location`                                                                      | `*string`                                                                       | :heavy_minus_sign:                                                              | Serving location. Only global is currently supported.                           |
| `ProjectID`                                                                     | `*string`                                                                       | :heavy_minus_sign:                                                              | ID of the serving project, which may differ from the service account's project. |