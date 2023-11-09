# C1APIPolicyV1PoliciesGetResponse


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ContentType`                                                                | *string*                                                                     | :heavy_check_mark:                                                           | HTTP response content type for this operation                                |
| `GetPolicyResponse`                                                          | [*shared.GetPolicyResponse](../../../pkg/models/shared/getpolicyresponse.md) | :heavy_minus_sign:                                                           | The GetPolicyResponse message contains the policy object.                    |
| `StatusCode`                                                                 | *int*                                                                        | :heavy_check_mark:                                                           | HTTP response status code for this operation                                 |
| `RawResponse`                                                                | [*http.Response](https://pkg.go.dev/net/http#Response)                       | :heavy_minus_sign:                                                           | Raw HTTP response; suitable for custom response parsing                      |