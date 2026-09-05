# BlockToolCallConfig

BlockToolCallConfig unconditionally denies the tool call when its hook's
 filter matches. Only valid for HOOK_EVENT_TYPE_POST_TOOL_USE.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Message`                                                                           | `*string`                                                                           | :heavy_minus_sign:                                                                  | Message shown when the tool call is denied. Empty falls back to a<br/> generic default. |