# Action

Action represents what happens when a component is activated (e.g., button click).

This message contains a oneof named action_type. Only a single field of the following list may be set at a time:
  - event
  - functionCall



## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `FunctionCall`                                                     | [*shared.FunctionCall](../../../pkg/models/shared/functioncall.md) | :heavy_minus_sign:                                                 | FunctionCall represents a client-side function invocation.         |
| `ServerEvent`                                                      | [*shared.ServerEvent](../../../pkg/models/shared/serverevent.md)   | :heavy_minus_sign:                                                 | ServerEvent triggers a server-side action.                         |