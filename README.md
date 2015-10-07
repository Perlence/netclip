# netclip

`netclip` is a TCP server that listens for copy/paste requests. When client
comes in and sends non-empty request, e.g. `beep`, server will put `beep` into
clipboard. Likewise, when client sends an empty string, server will get text
from clipboard and send it back to client.

There is also a client that mirrors the interface of `xclip`.

Both sides are configured via environment variable `NETCLIP_ADDR`. By default
it's `:2547`.


## Installation

To get `netclip` and `xclip`:

```bash
$ go get github.com/Perlence/netclip/...
```
