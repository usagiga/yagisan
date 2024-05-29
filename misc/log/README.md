# /misc/log

`/misc/log` sets default logger of slog.
The logger has [slog-context](https://github.com/PumpkinSeed/slog-context) handler.

## Usage

Just import this package, then set default logger of slog.

```go
package main

import (
  _ "github.com/usagiga/yagisan/misc/log"
)
```

- Write a log
  - call slog global functions (e.g. `slog.Infof()`)
- Add context-wide log fields(such as Trace ID, Request ID, ...)
  - call `cslog.WithValue()`
