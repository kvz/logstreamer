logstreamer [![Build Status][BuildStatusIMGURL]][BuildStatusURL]
===============
[![Flattr][FlattrIMGURL]][FlattrURL]

[BuildStatusIMGURL]:        https://secure.travis-ci.org/kvz/logstreamer.png?branch=master
[BuildStatusURL]:           //travis-ci.org/kvz/logstreamer  "Build Status"
[FlattrIMGURL]:             http://api.flattr.com/button/flattr-badge-large.png
[FlattrURL]:                https://flattr.com/submit/auto?user_id=kvz&url=github.com/kvz/logstreamer&title=logstreamer&language=&tags=github&category=software

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

Here I issue two local commands, `ls -al` and `ls nonexisting`:

![screen shot 2013-07-02 at 2 48 33 pm](https://f.cloud.github.com/assets/26752/736371/16177cf0-e316-11e2-8dc6-320f52f71442.png)

Over at [Transloadit](http://transloadit.com) we use it for streaming remote commands.
Servers stream command output over SSH back to me, and every line is prefixed with a date, their hostname & marked red in case they
wrote to stderr.

## License

This project is licensed under the MIT license, see `LICENSE.txt`.
