# WorkloadFederationServiceTestTokenRequest

The WorkloadFederationServiceTestTokenRequest message.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `SourceIP`                                                                              | **string*                                                                               | :heavy_minus_sign:                                                                      | Optional: override source IP for CIDR testing.<br/> If empty, uses the request's source IP. |
| `SubjectToken`                                                                          | **string*                                                                               | :heavy_minus_sign:                                                                      | The raw JWT to validate (the subject_token from a CI job).                              |