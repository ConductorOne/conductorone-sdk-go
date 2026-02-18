# FunctionTestResultLog

A log entry captured during a test.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Level`                                    | **string*                                  | :heavy_minus_sign:                         | The log level (e.g., "info", "error").     |
| `Log`                                      | **string*                                  | :heavy_minus_sign:                         | The log message content.                   |
| `Source`                                   | **string*                                  | :heavy_minus_sign:                         | The log source (e.g., "stdout", "stderr"). |