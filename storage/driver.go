package storage

import "context"

type DriverType string

const BadgerDriver DriverType = "badgerdb"
const DgraphDriver DriverType = "dgraphdb"

func (dt DriverType) String() string {
	return string(dt)
}

type Driver struct {
	Env DriverType `json:"env"`
}

type Storage interface {
	SaveTx(ctx context.Context) error
	FindTx(ctx context.Context)
	SearchTxs(ctx context.Context)
	SaveBlock(ctx context.Context)
	FindBlock(ctx context.Context)
	SearchBlocks(ctx context.Context)
	SaveEvent(ctx context.Context)
	FindEvent(ctx context.Context)
	SearchEvents(ctx context.Context)
}
