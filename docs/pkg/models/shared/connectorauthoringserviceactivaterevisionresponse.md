# ConnectorAuthoringServiceActivateRevisionResponse

The ConnectorAuthoringServiceActivateRevisionResponse message.


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ActivationEpoch`                                                                         | `*string`                                                                                 | :heavy_minus_sign:                                                                        | Committed activation epoch, or 0 when requested but not active yet.                       |
| `DeploymentEventID`                                                                       | `*string`                                                                                 | :heavy_minus_sign:                                                                        | Non-empty while instance rotation is pending. Observe completion through GetInstallation. |
| `RevisionID`                                                                              | `*string`                                                                                 | :heavy_minus_sign:                                                                        | The requested revision; published only after cutover commits.                             |