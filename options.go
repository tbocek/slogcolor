package slogcolor

import (
	"github.com/fatih/color"
	"log/slog"
	"time"
)

var DefaultOptions *Options = &Options{
	Level:         slog.LevelInfo,
	TimeFormat:    time.DateTime,
	SrcFileMode:   ShortFile,
	SrcFileLength: 0,
	MsgPrefix:     color.HiWhiteString("| "),
	MsgLength:     0,
	MsgColor:      color.New(),
}

type Options struct {
	// Level reports the minimum level to log.
	// Levels with lower levels are discarded.
	// If nil, the Handler uses [slog.LevelInfo].
	Level slog.Leveler

	// The time format.
	TimeFormat string

	// SrcFileMode is the source file mode.
	SrcFileMode SourceFileMode

	// SrcFileLength to show fixed length filename to line up the log output, default 0 shows complete filename
	SrcFileLength int

	// MsgPrefix to show prefix before message, default: white colored "| "
	MsgPrefix string

	// MsgColor is the color of the message, default to empty
	MsgColor *color.Color

	// MsgLength to show fixed length message to line up the log output, default 0 shows complete message
	MsgLength int
}
