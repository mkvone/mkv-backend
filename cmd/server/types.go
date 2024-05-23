package server

import "github.com/mkvone/mkv-backend/cmd/config"

type APIResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"` // omitempty를 사용하여 데이터가 없는 경우 필드를 생략
}

type ValidatorResponse struct {
	Name        string    `json:"name"`
	ChainID     string    `json:"chain_id"`
	Image       string    `json:"image"`
	BlockHeight string    `json:"block_height"`
	BlockTime   string    `json:"block_time"`
	Price       float32   `json:"price"`
	Ticker      string    `json:"ticker"`
	Validator   Validator `json:"validator"`
}
type Validator struct {
	OperatorAddr          string      `json:"operator_addr"`
	WalletAddress         string      `json:"wallet_address"`
	ValconAddress         string      `json:"valcon_address"`
	VotingPower           any         `json:"voting_power"`
	Uptime                Uptime      `json:"uptime"`
	Jailed                bool        `json:"jailed"`
	Status                string      `json:"status"`
	Tokens                string      `json:"tokens"`
	DelegatorShares       string      `json:"delegator_shares"`
	Description           Description `json:"description"`
	Commission            Commission  `json:"commission"`
	TotalDelegationCounts string      `json:"total_delegation_counts"`
}
type Uptime struct {
	Percent     float32 `json:"percent"`
	MissedBlock int     `json:"missed_block"`
	TotalBlock  int     `json:"total_block"`
	Tombstoned  bool    `json:"tombstoned"`
}
type Commission struct {
	Rate          string `json:"rate"`
	MaxRate       string `json:"max_rate"`
	MaxChangeRate string `json:"max_change_rate"`
}
type Description struct {
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
	Website  string `json:"website"`
	Details  string `json:"details"`
}

type Stats struct {
	Uptime     float32 `json:"uptime_avg"`
	Staked     float32 `json:"staked_value_avg"`
	Chains     int     `json:"chains_operating"`
	Delegators int     `json:"delegators"`
}

type Endpoints struct {
	// Chains []struct {
	Name    string `json:"name"`
	ChainID string `json:"chain_id"`
	Rpc     string `json:"rpc"`
	Rest    string `json:"rest_api"`
	Grpc    string `json:"grpc"`
	Img     string `json:"img_url"`
	// } `json:"chains"`
}

type Snapshots struct {
	Name    string        `json:"name"`
	ChainID string        `json:"chain_id"`
	App     string        `json:"app"`
	Go      string        `json:"go_version"`
	Img     string        `json:"img_url"`
	Base    string        `json:"base_url"`
	Files   []config.File `json:"files"`
}
