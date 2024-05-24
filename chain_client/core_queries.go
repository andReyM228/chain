package chain_client

import (
	"context"
	coremoduletypes "github.com/andReyM228/one/x/core/types"
)

func (c chainClient) GetStats(ctx context.Context, requestAllStats *coremoduletypes.QueryAllStatsRequest) (*coremoduletypes.QueryAllStatsResponse, error) {
	stats, err := c.coreQueryClient.StatsAll(ctx, requestAllStats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}
