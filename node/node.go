package node

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/rovergulf/chain/pkg/logutils"
	"github.com/rovergulf/chain/pkg/traceutils"
	"github.com/rovergulf/chain/wallets"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"os"
)

type Node struct {
	logger *zap.SugaredLogger
	tracer trace.Tracer

	walletsManager *wallets.Manager

	key *keystore.Key

	peer *p2p.Peer
}

func New() (*Node, error) {
	zapLogger, err := logutils.NewLogger()
	if err != nil {
		return nil, err
	}

	var traceProvider trace.TracerProvider
	jaegerTraceUrl := viper.GetString(traceutils.CollectorUrlConfigKey)
	if len(jaegerTraceUrl) > 0 {
		if traceProvider, err = traceutils.NewJaegerProvider(jaegerTraceUrl); err != nil {
			return nil, err
		}
	}

	n := &Node{
		logger: zapLogger,
	}

	if traceProvider != nil {
		n.tracer = traceProvider.Tracer("node")
	}

	wm, err := wallets.NewManager()
	if err != nil {
		zapLogger.Errorw("Unable to init wallets manager", "err", err)
		return nil, err
	}
	n.walletsManager = wm

	//n.peer = p2p.NewPeer(enode.PubkeyToIDV4())

	return n, nil
}

func (n *Node) GracefulShutdown(ctx context.Context, sig string) {
	defer ctx.Done()

	n.logger.Warnw("Graceful shutdown signal received", "sig", sig)

	os.Exit(0)
}
