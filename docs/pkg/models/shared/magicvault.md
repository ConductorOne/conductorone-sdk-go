# MagicVault

MagicVault configures a vault that grants time-limited credential access via magic links.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AllowUnauthedViews`                                                          | `*bool`                                                                       | :heavy_minus_sign:                                                            | Controls whether unauthenticated users can view credentials via a magic link. |
| `AllowedViews`                                                                | `*int64`                                                                      | :heavy_minus_sign:                                                            | The maximum number of times a credential in this vault may be viewed.         |