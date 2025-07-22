package log

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

const CRITICAL = log.Level(10)

const FatalLevel = log.FatalLevel
const CriticalLevel = CRITICAL
const ErrorLevel = log.ErrorLevel
const WarnLevel = log.WarnLevel
const InfoLevel = log.InfoLevel
const DebugLevel = log.DebugLevel

type Logger struct {
	*log.Logger
}

var LOG *Logger = &Logger{log.Default()}

var Fatal = LOG.Fatal
var Fatalf = LOG.Fatalf
var Critical = LOG.Critical
var Criticalf = LOG.Criticalf
var Error = LOG.Error
var Errorf = LOG.Errorf
var Warn = LOG.Warn
var Warnf = LOG.Warnf
var Info = LOG.Info
var Infof = LOG.Infof
var Debug = LOG.Debug
var Debugf = LOG.Debugf
var Print = LOG.Print
var Printf = LOG.Printf
var Log = LOG.Log
var Logf = LOG.Logf

var Helper = LOG.Helper
var GetLevel = LOG.GetLevel
var With = LOG.With
var WithPrefix = LOG.WithPrefix
var GetPrefix = LOG.GetPrefix
var SetPrefix = LOG.SetPrefix
var SetLevel = LOG.SetLevel
var SetLogLevel = LOG.SetLogLevel

func init() {
	os.Setenv("TERM", "xterm-256color")

	styles := log.DefaultStyles()
	styles.Levels[CRITICAL] = lipgloss.NewStyle().
		SetString("CRITICAL").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("201")).
		Foreground(lipgloss.Color("0"))

	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Values["err"] = lipgloss.NewStyle().Bold(true)

	styles.Values["accepted"] = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	styles.Values["invalid"] = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	styles.Values["old"] = lipgloss.NewStyle().Foreground(lipgloss.Color("208"))
	styles.Values["yours"] = lipgloss.NewStyle().Foreground(lipgloss.Color("11"))
	styles.Values["duplicate"] = lipgloss.NewStyle().Foreground(lipgloss.Color("171"))
	styles.Values["nop"] = lipgloss.NewStyle().Foreground(lipgloss.Color("11"))
	styles.Values["unavailable"] = lipgloss.NewStyle().Foreground(lipgloss.Color("88"))
	styles.Values["dispatch_err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("88"))
	styles.Values["young"] = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	styles.Values["critical"] = lipgloss.NewStyle().Foreground(lipgloss.Color("201"))

	styles.Values["flags"] = lipgloss.NewStyle().Foreground(lipgloss.Color("33"))
	styles.Values["expired"] = lipgloss.NewStyle().Foreground(lipgloss.Color("208"))

	LOG.SetStyles(styles)
	LOG.SetTimeFormat("15:04:05")
}

func (l *Logger) SetLevel(lvl log.Level) {
	l.Logger.SetLevel(lvl)
	l.SetReportCaller(lvl <= DebugLevel)
}

func (l *Logger) SetLogLevel(level string) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		l.Fatalf("Invalid log level '%s': %v", level, err)
	}
	l.SetLevel(lvl)
}

func (l *Logger) With(keyvals ...interface{}) *Logger {
	return &Logger{l.Logger.With(keyvals...)}
}
func (l *Logger) WithPrefix(prefix string) *Logger {
	return &Logger{l.Logger.WithPrefix(prefix)}
}

func (l *Logger) Critical(msg interface{}, keyvals ...interface{}) {
	l.Helper()
	l.Log(CRITICAL, msg, keyvals...)
}

func (l *Logger) Criticalf(format string, keyvals ...interface{}) {
	l.Helper()
	l.Logf(CRITICAL, format, keyvals...)
}
