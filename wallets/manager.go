package wallets

import (
	"errors"
	"github.com/dgraph-io/badger/v3"
	"github.com/rovergulf/chain/pkg/logutils"
	"github.com/rovergulf/chain/storage/badgerdb"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"path"
)

const DbWalletFile = "wallets.db"

var (
	ErrAccountNotExists = errors.New("account not exists")
	ErrInvalidAuth      = errors.New("invalid authentication code")
	ErrAccountIsLocked  = errors.New("account is locked")
)

type Manager struct {
	db     *badger.DB
	logger *zap.SugaredLogger
	tracer trace.Tracer
	quit   chan struct{}
}

// NewManager returns wallets Manager instance
func NewManager() (*Manager, error) {
	logger, err := logutils.NewLogger()
	if err != nil {
		return nil, err
	}

	walletsDbPath := path.Join(viper.GetString("data_dir"), "keystore")
	badgerOpts := badger.DefaultOptions(walletsDbPath)
	db, err := badgerdb.OpenDB(walletsDbPath, badgerOpts)
	if err != nil {
		return nil, err
	}

	return &Manager{
		db:     db,
		logger: logger,
	}, err
}

func (m *Manager) DbSize() (int64, int64) {
	return m.db.Size()
}

func (m *Manager) Shutdown() {
	if m.db != nil {
		if err := m.db.Close(); err != nil {
			m.logger.Errorf("Unable to close wallets db: %s", err)
		}
	}
}
