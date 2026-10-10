# AuthzenRef

AuthzenRef names an AuthzenServer. An empty app_id means the Edge's own app
 for rules saved before cross-app references were supported.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `AppID`                                                                                      | `*string`                                                                                    | :heavy_minus_sign:                                                                           | App that owns the AuthzenServer. Empty resolves to the Edge's app.                           |
| `ServerID`                                                                                   | `*string`                                                                                    | :heavy_minus_sign:                                                                           | ID of the AuthzenServer. It must exist in the target app and be in OBSERVE<br/> or ENFORCE mode. |