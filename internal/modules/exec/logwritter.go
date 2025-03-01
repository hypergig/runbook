package exec

import (
	"io"
	"log/slog"
	"strings"
)

var _ io.Writer = (*logWritter)(nil)

func (l *logWritter) Write(p []byte) (n int, err error) {
	for _, line := range strings.Split(string(p), "\n") {
		slog.Log(l.ctx, l.level, line)
	}
	return len(p), nil
}
