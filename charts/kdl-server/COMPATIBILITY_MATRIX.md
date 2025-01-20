## Compatibility matrix

| Release ↓ / Kubernetes → | 1.24 | 1.25 | 1.26 | 1.27 | 1.28 | 1.29 | 1.30 | 1.31 |
|:------------------------:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|
| 6.0.2                    | ✅   | ✅   | ✅   | ✅   | ✅   | ✅   | ✅   | ✅   |
| 6.1.0                    | ❌   | ❌   | ✅   | ✅   | ✅   | ✅   | ✅   | ✅   |
| 6.2.0                    | ❌   | ❌   | ✅   | ✅   | ✅   | ✅   | ✅   | ✅   |

| Release ↓ / kdl-app → | 1.38.0 | 1.38.1 | 1.39.0 | 1.40.0 |
|:---------------------:|:------:|:------:|:------:|:------:|
| 6.0.2                 | ✅     | ✅     | ❌     | ❌     |
| 6.1.0                 | ❌     | ❌     | ✅     | ❌     |
| 6.2.0                 | ❌     | ❌     | ❌     | ✅     |

| Release ↓ / project-operator → | 0.19.0 | 0.20.0 | 0.21.0 |
|:------------------------------:|:------:|:------:|:------:|
| 6.0.2                          | ✅     | ❌     | ❌     |
| 6.1.0                          | ❌     | ✅     | ❌     |
| 6.2.0                          | ❌     | ❌     | ✅     |

| Release ↓ / user-tools-operator → | 0.30.0 | 0.31.0 | 0.32.0 |
|:---------------------------------:|:------:|:------:|:------:|
| 6.0.2                             | ✅     | ❌     | ❌     |
| 6.1.0                             | ❌     | ✅     | ❌     |
| 6.2.0                             | ❌     | ❌     | ✅     |

| Symbol | Description |
|:------:|-------------|
| ✅     | Perfect match: all features are supported. Client and server versions have exactly the same features/APIs. |
| 🟠     | Forward compatibility: the client will work with the server, but not all new server features are supported. The server has features that the client library cannot use. |
| ❌     | Backward compatibility/Not applicable: the client has features that may not be present in the server. Common features will work, but some client APIs might not be available in the server. |
| -      | Not tested: this combination has not been verified or is not applicable. |
