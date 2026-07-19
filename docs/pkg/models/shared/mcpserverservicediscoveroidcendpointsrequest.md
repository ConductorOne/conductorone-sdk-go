# MCPServerServiceDiscoverOIDCEndpointsRequest

MCPServerServiceDiscoverOIDCEndpointsRequest fetches the OpenID Connect discovery
 document for a given issuer URL.


## Fields

| Field                                                                                                                   | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `IssuerURL`                                                                                                             | `*string`                                                                                                               | :heavy_minus_sign:                                                                                                      | The issuer URL (e.g. "https://accounts.google.com"). The server appends<br/> /.well-known/openid-configuration to this URL. |