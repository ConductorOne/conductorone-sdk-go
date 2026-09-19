# InstantiateClassifierTemplateRequest

The InstantiateClassifierTemplateRequest message.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `DisplayName`                                                             | `*string`                                                                 | :heavy_minus_sign:                                                        | Optional display name override for the created Classifier.                |
| `Params`                                                                  | map[string]`string`                                                       | :heavy_minus_sign:                                                        | Bind-time param values keyed by TemplateParam.key (e.g. grant_policy_id). |
| `TemplateVersion`                                                         | `*string`                                                                 | :heavy_minus_sign:                                                        | Optional; empty instantiates the latest version.                          |