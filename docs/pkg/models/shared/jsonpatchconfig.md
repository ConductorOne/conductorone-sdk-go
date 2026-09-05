# JSONPatchConfig

JSONPatchConfig adds, overwrites, or removes fields on a tool call's JSON
 input, with no function invocation. Only valid on
 HOOK_EVENT_TYPE_PRE_TOOL_USE. cel_expression is evaluated against
 ctx/input/caller and must produce a map; static_overlay is a fixed map.
 Either result is shallow-merged onto the input under RFC 7396 merge patch
 semantics: a key overwrites or adds that key, a null value removes it, and a
 nested object replaces rather than merging into the existing one.

This message contains a oneof named source. Only a single field of the following list may be set at a time:
  - celExpression
  - staticOverlay



## Fields

| Field                                                                                                                                            | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `CelExpression`                                                                                                                                  | `*string`                                                                                                                                        | :heavy_minus_sign:                                                                                                                               | The celExpression field.<br/>This field is part of the `source` oneof.<br/>See the documentation for `c1.api.hooks.v1.JSONPatchConfig` for more details. |
| `StaticOverlay`                                                                                                                                  | map[string]`any`                                                                                                                                 | :heavy_minus_sign:                                                                                                                               | N/A                                                                                                                                              |