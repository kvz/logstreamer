package logstreamer

import (
	"bytes"
	"io"
	"os"
	"log"
)

type Logstreamer struct {
	Logger    *log.Logger
	buf       *bytes.Buffer
	readLines string
	// If prefix == stdout, colors green
	// If prefix == stderr, colors red
	// Else, prefix is taken as-is, and prepended to anything
	// you throw at Write()
	prefix string
	// if true, saves output in memory
	record  bool
	persist string

	// Adds color to stdout & stderr if terminal supports it
	colorOkay  string
	colorFail  string
	colorReset string
}

func NewLogstreamer(logger *log.Logger, prefix string, record bool) *Logstreamer {
	streamer := &Logstreamer{
		Logger:     logger,
		buf:        bytes.NewBuffer([]byte("")),
		prefix:     prefix,
		record:     record,
		persist:    "",
		colorOkay:  "",
		colorFail:  "",
		colorReset: "",
	}

	if os.Getenv("TERM") == "xterm" {
		streamer.colorOkay  = "\x1b[32m"
		streamer.colorFail  = "\x1b[31m"
		streamer.colorReset = "\x1b[0m"
	}

	return streamer
}

func (l *Logstreamer) Write(p []byte) (n int, err error) {
	if n, err = l.buf.Write(p); err != nil {
		return
	}

	err = l.OutputLines()
	return
}

func (l *Logstreamer) Close() error {
	l.Flush()
	l.buf = bytes.NewBuffer([]byte(""))
	return nil
}

func (l *Logstreamer) Flush() error {
	var p []byte
	if _, err := l.buf.Read(p); err != nil {
		return err
	}

	l.out(string(p))
	return nil
}

func (l *Logstreamer) OutputLines() (err error) {
	for {
		line, err := l.buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		l.readLines += line
		l.out(line)
	}

	return nil
}

func (l *Logstreamer) ResetReadLines() {
	l.readLines = ""
}

func (l *Logstreamer) ReadLines() string {
	return l.readLines
}

func (l *Logstreamer) FlushRecord() string {
	buffer := l.persist
	l.persist = ""
	return buffer
}

func (l *Logstreamer) out(str string) (err error) {
	if l.record == true {
		l.persist = l.persist + str
	}

	if l.prefix == "stdout" {
		str = l.colorOkay + l.prefix + l.colorReset + " " + str
	} else if l.prefix == "stderr" {
		str = l.colorFail + l.prefix + l.colorReset + " " + str
	} else {
		str = l.prefix + str
	}

	l.Logger.Print(str)

	return nil
}

