package params

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"io"
)

// Options represents configuration options for whole package
// TODO deprecate and remove Options usage
type Options struct {
	DbFilePath      string `json:"db_file_path" yaml:"db_file_path"`
	WalletsFilePath string `json:"wallets_file_path" yaml:"wallets_file_path"`
	NodeFilePath    string `json:"node_file_path" yaml:"node_file_path"`
	Address         string `json:"address" yaml:"address"`
	NodeId          string `json:"node_id" yaml:"node_id"`
	Miner           string `json:"miner" yaml:"miner"`

	Badger badger.Options     `json:"-" yaml:"-"`
	Logger *zap.SugaredLogger `json:"-" yaml:"-"`
	Tracer opentracing.Tracer `json:"-" yaml:"-"`
	Closer io.Closer          `json:"-" yaml:"-"`
}
