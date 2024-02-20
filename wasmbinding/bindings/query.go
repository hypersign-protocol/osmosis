package bindings

// OsmosisQuery contains osmosis custom queries.
// See https://github.com/osmosis-labs/osmosis-bindings/blob/main/packages/bindings/src/query.rs
type OsmosisQuery struct {
	/// Given a subdenom minted by a contract via `OsmosisMsg::MintTokens`,
	/// returns the full denom as used by `BankMsg::Send`.
	FullDenom *FullDenom `json:"full_denom,omitempty"`
	/// Returns the admin of a denom, if the denom is a Token Factory denom.
	DenomAdmin *DenomAdmin `json:"denom_admin,omitempty"`
	/// Hypersign: Check if user has position
	UserPositionExists *UserPositionExists `json:"user_position_exists,omitempty"`
}

type UserPositionExists struct {
	Address string `json:"address"`
	PoolId  uint64 `json:"pool_id"`
}

type FullDenom struct {
	CreatorAddr string `json:"creator_addr"`
	Subdenom    string `json:"subdenom"`
}

type DenomAdmin struct {
	Subdenom string `json:"subdenom"`
}

type DenomAdminResponse struct {
	Admin string `json:"admin"`
}

type FullDenomResponse struct {
	Denom string `json:"denom"`
}

type UserPositionExistsResponse struct {
	Result bool `json:"result"`
}
