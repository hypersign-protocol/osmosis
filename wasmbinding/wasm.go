package wasmbinding

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	tokenfactorykeeper "github.com/osmosis-labs/osmosis/v23/x/tokenfactory/keeper"
	concentratedliquidity "github.com/osmosis-labs/osmosis/v23/x/concentrated-liquidity"
)

func RegisterCustomPlugins(
	bank *bankkeeper.BaseKeeper,
	tokenFactory *tokenfactorykeeper.Keeper,
	concentratedLiquidity *concentratedliquidity.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin(tokenFactory, concentratedLiquidity)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(bank, tokenFactory),
	)

	return []wasmkeeper.Option{
		queryPluginOpt,
		messengerDecoratorOpt,
	}
}

func RegisterStargateQueries(queryRouter baseapp.GRPCQueryRouter, codec codec.Codec) []wasmkeeper.Option {
	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Stargate: StargateQuerier(queryRouter, codec),
	})

	return []wasmkeeper.Option{
		queryPluginOpt,
	}
}
