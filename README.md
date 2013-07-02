logstreamer
===========

Prefixes streams (e.g. stdout or stderr) in Go.

If you are executing a lot of (remote) commands, you may want to indent all of their
output, prefix the loglines with hostnames, or mark anything that was thrown to stderr
red, so you can spot errors more easily.

For this purpose, Logstreamer was written.

You pass 3 arguments to `NewLogstreamer()`:

 - Your `*log.Logger`
 - Your desired prefix (`"stdout"` and `"stderr"` prefixed have special meaning)
 - If the lines should be recorded `true` or `false`. This is useful if you want to retrieve any errors.

This returns an interface that you can point `exec.Command`'s `cmd.Stderr` and `cmd.Stdout` to.
All bytes that are written to it are split by newline and then prefixed to your specification.

## Test

```bash
$ cd src/pkg/logstreamer/
$ go test
```
