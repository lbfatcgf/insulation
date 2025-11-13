package logger

import (
	"context"
	"log/slog"
	"os"
	"time"

	"insulation/server/base/pkg/config"
	"insulation/tools"

	fp "path/filepath"
)

type logMsg struct {
	Level   slog.Level
	Message string
	arg     []any
}
type Logger struct {
	file     *os.File
	grouping string
	log      *slog.Logger
	consloe  *slog.Logger
	wirteCh  chan *logMsg
	cancel   context.CancelFunc
	level    slog.Level
}

func NewLogger(grouping string, logConsloe bool) *Logger {
	logger := &Logger{
		wirteCh: make(chan *logMsg, 1024),
	}
	err := logger.logFileHandler()
	if err != nil {
		panic(err)
	}
	if logConsloe {
		logger.consloe = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: logger.level,
		}))
	}
	go logger.start()
	return logger
}

func (l *Logger) logFileHandler() error {
	l.level = slog.LevelDebug
	switch config.Global().Log.Level {
	case "debug":
		l.level = slog.LevelDebug
	case "info":
		l.level = slog.LevelInfo
	case "warn":
		l.level = slog.LevelWarn
	case "error":
		l.level = slog.LevelError
	}
	err := l.genLogger()
	if err != nil {
		return err
	}
	return nil
}

func (l *Logger) genLogger() error {
	if len(config.Global().Log.Path) == 0 {
		exepath, err := os.Executable()
		if err != nil {
			panic(err)
		}
		logdir := fp.Join(fp.Dir(exepath), tools.SafeFilePath("logs"))
		openFile, err := os.OpenFile(fp.Join(logdir, tools.SafeFilePath(l.grouping+"-"+time.Now().Format("2006-01-02")+".log")), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
		if err != nil {
			return err
		}
		if l.file != nil {
			l.file.Close()
		}
		l.file = openFile
	} else {
		logdir := fp.Join(fp.Dir(config.Global().Log.Path), tools.SafeFilePath("logs"))
		openFile, err := os.OpenFile(fp.Join(logdir, tools.SafeFilePath(l.grouping+"-"+time.Now().Format("2006-01-02")+".log")), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
		if err != nil {
			return err
		}
		if l.file != nil {
			l.file.Close()
		}
		l.file = openFile
	}
	logHandler := slog.NewJSONHandler(l.file, &slog.HandlerOptions{
		Level: l.level,
	})
	l.log = slog.New(logHandler)
	return nil
}

func (l *Logger) checkFileNeedChange() bool {
	fPath := l.file.Name()
	if len(config.Global().Log.Path) == 0 {
		exepath, err := os.Executable()
		if err != nil {
			panic(err)
		}
		logdir := fp.Join(fp.Dir(exepath), tools.SafeFilePath("logs"))

		return fp.Join(logdir, tools.SafeFilePath(l.grouping+"-"+time.Now().Format("2006-01-02")+".log")) != fPath
	} else {
		logdir := fp.Join(fp.Dir(config.Global().Log.Path), tools.SafeFilePath("logs"))
		return fp.Join(logdir, tools.SafeFilePath(l.grouping+"-"+time.Now().Format("2006-01-02")+".log")) != fPath
	}
}

func (l *Logger) Info(message string, args ...any) {
	l.wirteCh <- &logMsg{Level: slog.LevelInfo, Message: message, arg: args}
}

func (l *Logger) Error(message string, args ...any) {
	l.wirteCh <- &logMsg{Level: slog.LevelError, Message: message, arg: args}
}

func (l *Logger) Debug(message string, args ...any) {
	l.wirteCh <- &logMsg{Level: slog.LevelDebug, Message: message, arg: args}
}

func (l *Logger) Warn(message string, args ...any) {
	l.wirteCh <- &logMsg{Level: slog.LevelWarn, Message: message, arg: args}
}

func (l *Logger) Close() {
	l.cancel()
}

func (l *Logger) start() {
	ctx, cancel := context.WithCancel(context.Background())
	l.cancel = cancel
	for {
		select {
		case <-ctx.Done():
			l.file.Close()
			return
		case logMsg := <-l.wirteCh:
			if l.checkFileNeedChange() {
				l.genLogger()
			}
			if l.consloe != nil {
				l.consloe.Log(ctx, logMsg.Level, logMsg.Message, logMsg.arg...)
			}
			if l.log != nil {
				l.log.Log(ctx, logMsg.Level, logMsg.Message, logMsg.arg...)
			}
		}
	}
}
