# PreToolBlockConfig

PreToolBlockConfig unconditionally denies the tool call before it executes
 when its hook's filter matches. Only valid for HOOK_EVENT_TYPE_PRE_TOOL_USE.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Message`                                                                           | `*string`                                                                           | :heavy_minus_sign:                                                                  | Message shown when the tool call is denied. Empty falls back to a<br/> generic default. |