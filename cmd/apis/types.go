package apis

import "time"

type Status struct {
	Result struct {
		NodeInfo struct {
			Network string `json:"network"`
		} `json:"node_info"`
	} `json:"result"`
}

type NodeInfo struct {
	NodeInfo struct {
		Network string `json:"network"`
	} `json:"default_node_info"`
	ApplicationVersion struct {
		Name             string `json:"name"`
		Appname          string `json:"app_name"`
		Version          string `json:"version"`
		GoVersion        string `json:"go_version"`
		CosmosSDKVersion string `json:"cosmos_sdk_version"`
	} `json:"application_version"`
}

type DelegationCounts struct {
	Pagination Pagination `json:"pagination"`
}
type Pagination struct {
	Total string `json:"total"`
}
type SlashingParams struct {
	Params struct {
		SignedBlocksWindow      string `json:"signed_blocks_window"`
		MinSignedPerWindow      string `json:"min_signed_per_window"`
		DowntimeJailDuration    string `json:"downtime_jail_duration"`
		SlashFractionDoubleSign string `json:"slash_fraction_double_sign"`
		SlashFractionDowntime   string `json:"slash_fraction_downtime"`
	} `json:"params"`
}
type ValidatorSet struct {
	Validator  Validator   `json:"validator,omitempty"`
	Validators []Validator `json:"validators,omitempty"`
}
type Validator struct {
	OperatorAddress         string          `json:"operator_address"`
	ConsensusPubkey         ConsensusPubkey `json:"consensus_pubkey"`
	Jailed                  bool            `json:"jailed"`
	Status                  string          `json:"status"`
	Tokens                  string          `json:"tokens"`
	DelegatorShares         string          `json:"delegator_shares"`
	Description             Description     `json:"description"`
	UnbondingHeight         string          `json:"unbonding_height"`
	UnbondingTime           time.Time       `json:"unbonding_time"`
	Commission              Commission      `json:"commission"`
	MinSelfDelegation       string          `json:"min_self_delegation"`
	UnbondingOnHoldRefCount string          `json:"unbonding_on_hold_ref_count"`
	UnbondingIDS            []interface{}   `json:"unbonding_ids"`
}

type Commission struct {
	CommissionRates CommissionRates `json:"commission_rates"`
	UpdateTime      time.Time       `json:"update_time"`
}

type CommissionRates struct {
	Rate          string `json:"rate"`
	MaxRate       string `json:"max_rate"`
	MaxChangeRate string `json:"max_change_rate"`
}

type ConsensusPubkey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type Description struct {
	Moniker         string `json:"moniker"`
	Identity        string `json:"identity"`
	Website         string `json:"website"`
	SecurityContact string `json:"security_contact"`
	Details         string `json:"details"`
}

type SigningInfo struct {
	ValSigningInfo ValSigningInfo `json:"val_signing_info"`
}

type ValSigningInfo struct {
	Address      string `json:"address"`
	StartHeight  string `json:"start_height"`
	IndexOffset  string `json:"index_offset"`
	Tombstoned   bool   `json:"tombstoned"`
	JailedUntil  string `json:"jailed_until"`
	MissedBlocks string `json:"missed_blocks_counter"`
}
