# Environment Variables

This document describes the environment variables used by `<app>`.

The application may consume other undocumented environment variables; this
document only shows those variables defined using [Ferrite].

## Index

- [`DEBUG`](#DEBUG) — enable or disable debugging features

## Specification

### `DEBUG`

> enable or disable debugging features

This variable **MAY** be set to one of the values below.
If it is undefined or empty a default value of `false` is used.

```bash
export DEBUG=true
export DEBUG=false # (default)
```

<!-- references -->

[ferrite]: https://github.com/dogmatiq/ferrite
