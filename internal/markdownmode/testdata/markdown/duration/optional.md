# Environment Variables

This document describes the environment variables used by `<app>`.

Please note that **undefined** variables and **empty strings** are considered
equivalent.

The application may consume other undocumented environment variables; this
document only shows those variables defined using [Ferrite].

## Index

- [`GRPC_TIMEOUT`](#GRPC_TIMEOUT) — gRPC request timeout

## Specification

### `GRPC_TIMEOUT`

> gRPC request timeout

This variable **MAY** be set to `1ns` or greater, or left undefined.

```bash
export GRPC_TIMEOUT=640511h56m49.213693952s # randomly generated example
```

<!-- references -->

[ferrite]: https://github.com/dogmatiq/ferrite
