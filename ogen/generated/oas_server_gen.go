// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// BlockFromTimestamp implements BlockFromTimestamp operation.
	//
	// Returns block on or before the specified block timestamp.
	//
	// GET /v1/utilities/chains/{chainId}/block/from-timestamp
	BlockFromTimestamp(ctx context.Context, params BlockFromTimestampParams) (BlockFromTimestampRes, error)
	// BlockToTimestamp implements BlockToTimestamp operation.
	//
	// Returns block timestamp of the specified block number.
	//
	// GET /v1/utilities/chains/{chainId}/block/to-timestamp
	BlockToTimestamp(ctx context.Context, params BlockToTimestampParams) (BlockToTimestampRes, error)
	// CheckHealth implements CheckHealth operation.
	//
	// Healthcheck status.
	//
	// GET /health
	CheckHealth(ctx context.Context) (CheckHealthRes, error)
	// CrossChainBlock implements CrossChainBlock operation.
	//
	// Returns cross chain's block on or before the timestamp of input chain's block number.
	//
	// GET /v1/utilities/chains/{chainId}/block/cross-chain
	CrossChainBlock(ctx context.Context, params CrossChainBlockParams) (CrossChainBlockRes, error)
	// GetCachedModelResults implements GetCachedModelResults operation.
	//
	// Returns cached run results for a slug.<p>This endpoint is for analyzing model runs. To run a model
	// and get results, use `POST /v1/model/run`.
	//
	// GET /v1/model/results
	GetCachedModelResults(ctx context.Context, params GetCachedModelResultsParams) (*ModelRuntimeStatsResponse, error)
	// GetChains implements GetChains operation.
	//
	// Returns metadata for the list of blockchains supported by Credmark platform.
	//
	// GET /v1/utilities/chains
	GetChains(ctx context.Context) (GetChainsRes, error)
	// GetDailyModelUsage implements GetDailyModelUsage operation.
	//
	// Returns a list of daily model request statistics, either for a specific requester or for everyone.
	//
	// GET /v1/usage/requests
	GetDailyModelUsage(ctx context.Context, params GetDailyModelUsageParams) ([]GetDailyModelUsageOKItem, error)
	// GetLatestBlock implements GetLatestBlock operation.
	//
	// Returns latest block of the specified chain.
	//
	// GET /v1/utilities/chains/{chainId}/block/latest
	GetLatestBlock(ctx context.Context, params GetLatestBlockParams) (GetLatestBlockRes, error)
	// GetModelBySlug implements GetModelBySlug operation.
	//
	// Returns the metadata for the specified model.
	//
	// GET /v1/models/{slug}
	GetModelBySlug(ctx context.Context, params GetModelBySlugParams) (*ModelMetadata, error)
	// GetModelDeploymentsBySlug implements GetModelDeploymentsBySlug operation.
	//
	// Returns the deployments for a model.
	//
	// GET /v1/models/{slug}/deployments
	GetModelDeploymentsBySlug(ctx context.Context, params GetModelDeploymentsBySlugParams) ([]ModelDeployment, error)
	// GetModelRuntimeStats implements GetModelRuntimeStats operation.
	//
	// Returns runtime stats for all models.
	//
	// GET /v1/model/runtime-stats
	GetModelRuntimeStats(ctx context.Context) (*ModelRuntimeStatsResponse, error)
	// GetPositions implements GetPositions operation.
	//
	// Returns positions for a list of accounts.
	//
	// GET /v1/portfolio/{chainId}/{accounts}/positions
	GetPositions(ctx context.Context, params GetPositionsParams) (GetPositionsRes, error)
	// GetPositionsHistorical implements GetPositionsHistorical operation.
	//
	// Returns positions for a list of accounts over a series of blocks.
	//
	// GET /v1/portfolio/{chainId}/{accounts}/positions/historical
	GetPositionsHistorical(ctx context.Context, params GetPositionsHistoricalParams) (GetPositionsHistoricalRes, error)
	// GetReturns implements GetReturns operation.
	//
	// Returns PnL of portfolio for a list of accounts.
	//
	// GET /v1/portfolio/{chainId}/{accounts}/returns
	GetReturns(ctx context.Context, params GetReturnsParams) (GetReturnsRes, error)
	// GetTokenAbi implements GetTokenAbi operation.
	//
	// Returns ABI of a token.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/abi
	GetTokenAbi(ctx context.Context, params GetTokenAbiParams) (GetTokenAbiRes, error)
	// GetTokenBalance implements GetTokenBalance operation.
	//
	// Returns token balance for an account.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/balance
	GetTokenBalance(ctx context.Context, params GetTokenBalanceParams) (GetTokenBalanceRes, error)
	// GetTokenBalanceHistorical implements GetTokenBalanceHistorical operation.
	//
	// Returns historical token balance for an account.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/balance/historical
	GetTokenBalanceHistorical(ctx context.Context, params GetTokenBalanceHistoricalParams) (GetTokenBalanceHistoricalRes, error)
	// GetTokenCreationBlock implements GetTokenCreationBlock operation.
	//
	// Returns creation block number of a token.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/creation-block
	GetTokenCreationBlock(ctx context.Context, params GetTokenCreationBlockParams) (GetTokenCreationBlockRes, error)
	// GetTokenDecimals implements GetTokenDecimals operation.
	//
	// Returns decimals of a token.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/decimals
	GetTokenDecimals(ctx context.Context, params GetTokenDecimalsParams) (GetTokenDecimalsRes, error)
	// GetTokenHolders implements GetTokenHolders operation.
	//
	// Returns holders of a token at a block or time.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/holders
	GetTokenHolders(ctx context.Context, params GetTokenHoldersParams) (GetTokenHoldersRes, error)
	// GetTokenHoldersCount implements GetTokenHoldersCount operation.
	//
	// Returns total number of holders of a token at a block or time.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/holders/count
	GetTokenHoldersCount(ctx context.Context, params GetTokenHoldersCountParams) (GetTokenHoldersCountRes, error)
	// GetTokenHoldersCountHistorical implements GetTokenHoldersCountHistorical operation.
	//
	// Returns historical total number of holders of a token at a block or time.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/holders/count/historical
	GetTokenHoldersCountHistorical(ctx context.Context, params GetTokenHoldersCountHistoricalParams) (GetTokenHoldersCountHistoricalRes, error)
	// GetTokenLogo implements GetTokenLogo operation.
	//
	// Returns logo of a token.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/logo
	GetTokenLogo(ctx context.Context, params GetTokenLogoParams) (GetTokenLogoRes, error)
	// GetTokenMetadata implements GetTokenMetadata operation.
	//
	// Returns metadata for a token.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}
	GetTokenMetadata(ctx context.Context, params GetTokenMetadataParams) (GetTokenMetadataRes, error)
	// GetTokenName implements GetTokenName operation.
	//
	// Returns name of a token.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/name
	GetTokenName(ctx context.Context, params GetTokenNameParams) (GetTokenNameRes, error)
	// GetTokenPrice implements GetTokenPrice operation.
	//
	// Returns price data for a token.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/price
	GetTokenPrice(ctx context.Context, params GetTokenPriceParams) (GetTokenPriceRes, error)
	// GetTokenPriceHistorical implements GetTokenPriceHistorical operation.
	//
	// Returns historical price data for a token.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/price/historical
	GetTokenPriceHistorical(ctx context.Context, params GetTokenPriceHistoricalParams) (GetTokenPriceHistoricalRes, error)
	// GetTokenSymbol implements GetTokenSymbol operation.
	//
	// Returns symbol of a token.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/symbol
	GetTokenSymbol(ctx context.Context, params GetTokenSymbolParams) (GetTokenSymbolRes, error)
	// GetTokenTotalSupply implements GetTokenTotalSupply operation.
	//
	// Returns total supply of a token.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/total-supply
	GetTokenTotalSupply(ctx context.Context, params GetTokenTotalSupplyParams) (GetTokenTotalSupplyRes, error)
	// GetTokenTotalSupplyHistorical implements GetTokenTotalSupplyHistorical operation.
	//
	// Returns historical total supply for a token.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/total-supply/historical
	GetTokenTotalSupplyHistorical(ctx context.Context, params GetTokenTotalSupplyHistoricalParams) (GetTokenTotalSupplyHistoricalRes, error)
	// GetTokenVolume implements GetTokenVolume operation.
	//
	// Returns traded volume for a token over a period of blocks or time.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/volume
	GetTokenVolume(ctx context.Context, params GetTokenVolumeParams) (GetTokenVolumeRes, error)
	// GetTokenVolumeHistorical implements GetTokenVolumeHistorical operation.
	//
	// Returns traded volume for a token over a period of blocks or time divided by intervals.
	//
	// GET /v1/tokens/{chainId}/{tokenAddress}/volume/historical
	GetTokenVolumeHistorical(ctx context.Context, params GetTokenVolumeHistoricalParams) (GetTokenVolumeHistoricalRes, error)
	// GetTopModels implements GetTopModels operation.
	//
	// Returns a list of the top used models.
	//
	// GET /v1/usage/top
	GetTopModels(ctx context.Context) ([]GetTopModelsOKItem, error)
	// GetTotalModelUsage implements GetTotalModelUsage operation.
	//
	// Returns a count of model runs.
	//
	// GET /v1/usage/total
	GetTotalModelUsage(ctx context.Context) ([]GetTotalModelUsageOKItem, error)
	// GetValue implements GetValue operation.
	//
	// Returns value of portfolio for a list of accounts.
	//
	// GET /v1/portfolio/{chainId}/{accounts}/value
	GetValue(ctx context.Context, params GetValueParams) (GetValueRes, error)
	// GetValueHistorical implements GetValueHistorical operation.
	//
	// Returns portfolio value for a list of accounts over a series of blocks.
	//
	// GET /v1/portfolio/{chainId}/{accounts}/value/historical
	GetValueHistorical(ctx context.Context, params GetValueHistoricalParams) (GetValueHistoricalRes, error)
	// ListModels implements ListModels operation.
	//
	// Returns a list of metadata for available models.
	//
	// GET /v1/models
	ListModels(ctx context.Context) ([]ModelMetadata, error)
	// RunModel implements RunModel operation.
	//
	// Runs a model and returns result object.
	//
	// POST /v1/model/run
	RunModel(ctx context.Context, req *RunModelDto) (RunModelRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
