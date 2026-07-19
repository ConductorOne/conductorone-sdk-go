# MCPServerAuthGoogleServiceAccount

MCPServerAuthGoogleServiceAccount provides Google service account authentication.
 The admin uploads the raw JSON key file from the GCP console; the backend parses
 it to extract the private key, client email (issuer), and token URI.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `CredentialsJSON`                                                     | `*string`                                                             | :heavy_minus_sign:                                                    | Raw JSON content of the Google service account key file.              |
| `Scopes`                                                              | []`string`                                                            | :heavy_minus_sign:                                                    | OAuth2 scopes to request when exchanging the JWT for an access token. |