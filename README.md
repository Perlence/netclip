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


## Example

Start `netclip`:

```bash
$ netclip &
2015/10/06 20:56:00 netclip server is listening on :2547
```

Send copy request:

```bash
$ echo -n "beep" | nc -q0 127.0.0.1 2547
2015/10/06 20:56:12 "beep"
```

Send paste request:

```bash
$ echo -n "" | nc -q0 127.0.0.1 2547
2015/10/06 20:56:36 ""
```
