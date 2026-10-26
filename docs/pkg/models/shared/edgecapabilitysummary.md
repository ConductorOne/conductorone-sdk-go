# EdgeCapabilitySummary

EdgeCapabilitySummary describes the capability reports of a tenant's connected Edge processes.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `LastReportedAt`                           | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `ReportingProcesses`                       | `*int64`                                   | :heavy_minus_sign:                         | Edge processes with a live report.         |