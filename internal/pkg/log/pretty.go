package log

import (
	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
	"log/slog"
	"os"
)

var (
	pretty *slog.Logger
	level  = &slog.LevelVar{}
)

func init() {
	setup(os.Stderr)
}

func setup(output *os.File) {
	// windows terminal support via go-colorable
	writer := colorable.NewColorable(output)
	pretty = slog.New(
		tint.NewHandler(writer, &tint.Options{
			ReplaceAttr: rewrite,
			Level:       level,
		}),
	)
}

func rewrite(_ []string, attr slog.Attr) slog.Attr {
	// remove time
	if attr.Key == slog.TimeKey {
		return slog.Attr{}
	}
	// tint errors red
	if err, ok := attr.Value.Any().(error); ok {
		a := tint.Err(err)
		a.Key = attr.Key
		return a
	}
	return attr
}

func SetOutput(target *os.File) {
	setup(target)
}

func EnableDebug(debug bool) {
	if debug {
		level.Set(slog.LevelDebug)
	}
}

func Error(msg string, args ...any) {
	pretty.Error(msg, args...)
}

func Warn(msg string, args ...any) {
	pretty.Warn(msg, args...)
}

func Info(msg string, args ...any) {
	pretty.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	pretty.Debug(msg, args...)
}
