package def

import (
	"context"
	"os"
	"path"
	"strings"

	"github.com/powerman/structlog"
)

// Constants.
var (
	ProgName    = strings.TrimSuffix(path.Base(os.Args[0]), ".test")
	Hostname    = "localhost"
	HostnameInt = "127.0.0.1"
)

// NewContext returns context.Background() which contains logger
// configured for given service.
func NewContext(service string) context.Context {
	return structlog.NewContext(context.Background(), structlog.New(structlog.KeyApp, service))
}
