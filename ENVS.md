# Environments

| Name                   | Required | Secret | Default value  | Usage                                     | Example        |
|------------------------|----------|--------|----------------|-------------------------------------------|----------------|
| `ENV_CI`               | ✅        |        |                |                                           | `dev`          |
| `LOG_ENCODING_CONSOLE` |          |        | `false`        | allows to set user-friendly formatting    | `false`        |
| `LOG_LEVEL`            |          |        | `info`         | allows to set custom logger level         | `info`         |
| `LOG_TRACE`            |          |        | `fatal`        | allows to set custom trace level          | `fatal`        |
| `LOG_WITH_CALLER`      |          |        | `false`        | allows to show stack trace                | `false`        |
| `LOG_WITH_STACK_TRACE` |          |        | `false`        | allows to show stack trace                | `false`        |
| `OPS_ENABLED`          |          |        | `false`        | allows to enable ops server               | `false`        |
| `OPS_NETWORK`          | ✅        |        | `tcp`          | allows to set ops listen network: tcp/udp | `tcp`          |
| `OPS_TRACING_ENABLED`  |          |        | `false`        | allows to enable tracing                  | `false`        |
| `OPS_METRICS_ENABLED`  |          |        | `true`         | allows to enable metrics                  | `true`         |
| `OPS_METRICS_PATH`     | ✅        |        | `/metrics`     | allows to set custom metrics path         | `/metrics`     |
| `OPS_METRICS_PORT`     | ✅        |        | `10000`        | allows to set custom metrics port         | `10000`        |
| `OPS_HEALTHY_ENABLED`  |          |        | `true`         | allows to enable health checker           | `true`         |
| `OPS_HEALTHY_PATH`     | ✅        |        | `/healthy`     | allows to set custom healthy path         | `/healthy`     |
| `OPS_HEALTHY_PORT`     | ✅        |        | `10000`        | allows to set custom healthy port         | `10000`        |
| `OPS_PROFILER_ENABLED` |          |        | `false`        | allows to enable profiler                 | `false`        |
| `OPS_PROFILER_PATH`    | ✅        |        | `/debug/pprof` | allows to set custom profiler path        | `/debug/pprof` |
| `OPS_PROFILER_PORT`    | ✅        |        | `10000`        | allows to set custom profiler port        | `10000`        |
| `STATIC_SERVER_ADDR`   | ✅        |        | `:8080`        |                                           |                |
