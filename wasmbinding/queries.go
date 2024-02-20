package wasmbinding

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v23/wasmbinding/bindings"
	tokenfactorykeeper "github.com/osmosis-labs/osmosis/v23/x/tokenfactory/keeper"
	concentratedliquidity "github.com/osmosis-labs/osmosis/v23/x/concentrated-liquidity"
)

type QueryPlugin struct {
	tokenFactoryKeeper          *tokenfactorykeeper.Keeper
	concentratedLiquidityKeeper *concentratedliquidity.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(tfk *tokenfactorykeeper.Keeper, clk *concentratedliquidity.Keeper) *QueryPlugin {
	return &QueryPlugin{
		tokenFactoryKeeper:          tfk,
		concentratedLiquidityKeeper: clk,
	}
}

// GetDenomAdmin is a query to get denom admin.
func (qp QueryPlugin) GetDenomAdmin(ctx sdk.Context, denom string) (*bindings.DenomAdminResponse, error) {
	metadata, err := qp.tokenFactoryKeeper.GetAuthorityMetadata(ctx, denom)
	if err != nil {
		return nil, fmt.Errorf("failed to get admin for denom: %s", denom)
	}

	return &bindings.DenomAdminResponse{Admin: metadata.Admin}, nil
}

func (qp QueryPlugin) GetUserPostitions(ctx sdk.Context, address string, poolId uint64) (*bindings.UserPositionExistsResponse, error) {
	sdkAddr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return nil, err
	}

	positions, err := qp.concentratedLiquidityKeeper.GetUserPositions(ctx, sdkAddr, poolId)
	if err != nil {
		return nil, err
	}

	if len(positions) > 0 {
		return &bindings.UserPositionExistsResponse{
			Result: true,
		}, nil
	} else {
		return &bindings.UserPositionExistsResponse{
			Result: false,
		}, nil
	}
}
