package wallet

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
)

func NewAPI(s *Service) *API {
	return &API{s}
}

// API is class with methods available over RPC.
type API struct {
	s *Service
}

// GetTransfers returns transfers in range of blocks. If `end` is nil all transfers from `start` will be returned.
// TODO(dshulyak) benchmark loading many transfers from database. We can avoid json unmarshal/marshal if we will
// read header, tx and receipt as a raw json.
func (api *API) GetTransfers(ctx context.Context, start, end *hexutil.Big) ([]Transfer, error) {
	log.Debug("call to get transfers", "start", start, "end", end)
	if start == nil {
		return nil, errors.New("start of the query must be provided. use 0 if you want to load all transfers")
	}
	if api.s.db == nil {
		return nil, errors.New("wallet service is not initialized")
	}
	rst, err := api.s.db.GetTransfers((*big.Int)(start), (*big.Int)(end))
	if err != nil {
		return nil, err
	}
	log.Debug("result from database for transfers", "start", start, "end", end, "len", len(rst))
	return rst, nil
}

// GetTransfersByAddress returns transfers for a single address between two blocks.
func (api *API) GetTransfersByAddress(ctx context.Context, address common.Address, start, end *hexutil.Big) ([]Transfer, error) {
	log.Debug("call to get transfers for an address", "address", address, "start", start, "end", end)
	if start == nil {
		return nil, errors.New("start of the query must be provided. use 0 if you want to load all transfers")
	}
	if api.s.db == nil {
		return nil, errors.New("wallet service is not initialized")
	}
	rst, err := api.s.db.GetTransfersByAddress(address, (*big.Int)(start), (*big.Int)(end))
	if err != nil {
		return nil, err
	}
	log.Debug("result from database for address", "address", address, "start", start, "end", end, "len", len(rst))
	return rst, nil
}