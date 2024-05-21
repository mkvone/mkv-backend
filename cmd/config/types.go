package config

import "context"

const (
	showBLocks = 100
	staleHours = 24
)

type Config struct {
	Swagger   SwaggerConfig   `toml:"swagger"`
	Databases DatabasesConfig `toml:"databases"`
	Alert     AlertConfig     `toml:"alert"`
	Chains    []ChainConfig   `toml:"chains"`
	ctx       context.Context
	cancel    context.CancelFunc
}

type SwaggerConfig struct {
	Enable bool `toml:"enable"`
	Port   int  `toml:"port"`
}

type DatabasesConfig struct {
	MongoDB MongoDBConfig `toml:"mongodb"`
}

type MongoDBConfig struct {
	URL string `toml:"url"`
}

type AlertConfig struct {
	Slack    SlackConfig    `toml:"slack"`
	Discord  DiscordConfig  `toml:"discord"`
	Telegram TelegramConfig `toml:"telegram"`
}

type SlackConfig struct {
	Enable  bool   `toml:"enable"`
	Webhook string `toml:"webhook"`
}

type DiscordConfig struct {
	Enable  bool   `toml:"enable"`
	Webhook string `toml:"webhook"`
}

type TelegramConfig struct {
	Enable bool   `toml:"enable"`
	Token  string `toml:"token"`
	ChatID string `toml:"chat_id"`
}

type ChainConfig struct {
	Name      string `toml:"name"`
	ChainID   string
	Appname   string
	Version   string
	GoVersion string
	CosmosSDK string
	Endpoints Endpoints `toml:"endpoints"`
	Symbol    Symbol    `toml:"symbol"`
	ImgURL    string    `toml:"img_url"`
	Snapshot  Snapshot  `toml:"snapshot"`
	Validator Validator `toml:"validator"`
	Block     Block
}
type Block struct {
	Height string
	Time   string
}
type Symbol struct {
	Ticker string `json:"ticker"`
	Price  float32
}
type Endpoints struct {
	Rpc  string `toml:"rpc"`
	Api  string `toml:"api"`
	Grpc string `toml:"grpc"`
}
type Snapshot struct {
	Enable      bool   `toml:"enable"`
	SnapshotURL string `toml:"snapshot_url"`
	Files       []File `json:"files"`
}
type File struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Size   string `json:"size"`
	Date   string `json:"date"`
	Height string `json:"height"`
}
type Uptime struct {
	Percent     float32
	MissedBlock int
	TotalBlock  int
	Tombstoned  bool
}
type Validator struct {
	Enable                bool   `toml:"enable"`
	OperatorAddr          string `toml:"operator_addr"`
	WalletAddress         string `toml:"wallet_address"`
	ValconAddress         string
	Rank                  int
	VotingPower           int
	Uptime                Uptime
	ConsensusPubkey       ConsensusPubkey
	Jailed                bool
	Status                string
	Tokens                string
	DelegatorShares       string
	Description           Description
	Commission            Commission
	TotalDelegationCounts string
}
type Commission struct {
	Rate          string
	MaxRate       string
	MaxChangeRate string
}
type Description struct {
	Moniker  string `toml:"moniker"`
	Identity string `toml:"identity"`
	Website  string `toml:"website"`
	Details  string `toml:"details"`
}
type ConsensusPubkey struct {
	Type  string `toml:"type_url"`
	Value string `toml:"value"`
}
