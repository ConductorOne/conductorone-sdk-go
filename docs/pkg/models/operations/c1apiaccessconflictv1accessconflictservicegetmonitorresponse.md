# C1APIAccessconflictV1AccessConflictServiceGetMonitorResponse


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ConflictMonitor`                                                        | [*shared.ConflictMonitor](../../../pkg/models/shared/conflictmonitor.md) | :heavy_minus_sign:                                                       | Successful response                                                      |
| `ContentType`                                                            | *string*                                                                 | :heavy_check_mark:                                                       | HTTP response content type for this operation                            |
| `StatusCode`                                                             | *int*                                                                    | :heavy_check_mark:                                                       | HTTP response status code for this operation                             |
| `RawResponse`                                                            | [*http.Response](https://pkg.go.dev/net/http#Response)                   | :heavy_check_mark:                                                       | Raw HTTP response; suitable for custom response parsing                  |