# EncodedContentGuardConfig

EncodedContentGuardConfig detects encoded/obfuscated smuggling in tool input:
 long base64 blobs, long hex runs, and invisible/zero-width unicode.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `FlagOnly`                                                                 | `*bool`                                                                    | :heavy_minus_sign:                                                         | When true, detection records the finding but does not deny (observe-only). |
| `MinBase64Run`                                                             | `*int`                                                                     | :heavy_minus_sign:                                                         | Minimum contiguous base64 run length to flag. <= 0 = default (256).        |
| `MinHexRun`                                                                | `*int`                                                                     | :heavy_minus_sign:                                                         | Minimum contiguous hex run length to flag. <= 0 = default (128).           |