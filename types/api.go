package types

import (
	"database/sql/driver"
	"encoding/json"
	"math/big"
	"time"

	"github.com/pkg/errors"
)

type ApiResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type StatsSystem struct {
	CPUCores                      uint64 `mapstructure:"cpu_cores"`
	CPUThreads                    uint64 `mapstructure:"cpu_threads"`
	CPUNodeSystemSecondsTotal     uint64 `mapstructure:"cpu_node_system_seconds_total"`
	CPUNodeUserSecondsTotal       uint64 `mapstructure:"cpu_node_user_seconds_total"`
	CPUNodeIowaitSecondsTotal     uint64 `mapstructure:"cpu_node_iowait_seconds_total"`
	CPUNodeIdleSecondsTotal       uint64 `mapstructure:"cpu_node_idle_seconds_total"`
	MemoryNodeBytesTotal          uint64 `mapstructure:"memory_node_bytes_total"`
	MemoryNodeBytesFree           uint64 `mapstructure:"memory_node_bytes_free"`
	MemoryNodeBytesCached         uint64 `mapstructure:"memory_node_bytes_cached"`
	MemoryNodeBytesBuffers        uint64 `mapstructure:"memory_node_bytes_buffers"`
	DiskNodeBytesTotal            uint64 `mapstructure:"disk_node_bytes_total"`
	DiskNodeBytesFree             uint64 `mapstructure:"disk_node_bytes_free"`
	DiskNodeIoSeconds             uint64 `mapstructure:"disk_node_io_seconds"`
	DiskNodeReadsTotal            uint64 `mapstructure:"disk_node_reads_total"`
	DiskNodeWritesTotal           uint64 `mapstructure:"disk_node_writes_total"`
	NetworkNodeBytesTotalReceive  uint64 `mapstructure:"network_node_bytes_total_receive"`
	NetworkNodeBytesTotalTransmit uint64 `mapstructure:"network_node_bytes_total_transmit"`
	MiscNodeBootTsSeconds         uint64 `mapstructure:"misc_node_boot_ts_seconds"`
	MiscOS                        string `mapstructure:"misc_os"`
}

type StatsProcess struct {
	CPUProcessSecondsTotal     uint64 `mapstructure:"cpu_process_seconds_total"`
	MemoryProcessBytes         uint64 `mapstructure:"memory_process_bytes"`
	ClientName                 string `mapstructure:"client_name"`
	ClientVersion              string `mapstructure:"client_version"`
	ClientBuild                uint64 `mapstructure:"client_build"`
	SyncEth2FallbackConfigured bool   `mapstructure:"sync_eth2_fallback_configured"`
	SyncEth2FallbackConnected  bool   `mapstructure:"sync_eth2_fallback_connected"`
}

type StatsAdditionalsValidator struct {
	ValidatorTotal  uint64 `mapstructure:"validator_total"`
	ValidatorActive uint64 `mapstructure:"validator_active"`
}

type StatsAdditionalsBeaconnode struct {
	DiskBeaconchainBytesTotal       uint64 `mapstructure:"disk_beaconchain_bytes_total"`
	NetworkLibp2pBytesTotalReceive  uint64 `mapstructure:"network_libp2p_bytes_total_receive"`
	NetworkLibp2pBytesTotalTransmit uint64 `mapstructure:"network_libp2p_bytes_total_transmit"`
	NetworkPeersConnected           uint64 `mapstructure:"network_peers_connected"`
	SyncEth1Connected               bool   `mapstructure:"sync_eth1_connected"`
	SyncEth2Synced                  bool   `mapstructure:"sync_eth2_synced"`
	SyncBeaconHeadSlot              uint64 `mapstructure:"sync_beacon_head_slot"`
	SyncEth1FallbackConfigured      bool   `mapstructure:"sync_eth1_fallback_configured"`
	SyncEth1FallbackConnected       bool   `mapstructure:"sync_eth1_fallback_connected"`
}

type StatsMeta struct {
	Version         uint64 `mapstructure:"version"`
	Timestamp       uint64 `mapstructure:"timestamp"`
	Process         string `mapstructure:"process"`
	Machine         string
	ExporterVersion string `mapstructure:"exporter_version"`
}

type StatsDataStruct struct {
	Validator interface{} `json:"validator"`
	Node      interface{} `json:"node"`
	System    interface{} `json:"system"`
}

type WidgetResponse struct {
	Eff             any   `json:"efficiency"`
	Validator       any   `json:"validator"`
	Epoch           int64 `json:"epoch"`
	RocketpoolStats any   `json:"rocketpool_network_stats"`
}

type UsersNotificationsRequest struct {
	EventNames    []string `json:"event_names"`
	EventFilters  []string `json:"event_filters"`
	Search        string   `json:"search"`
	Limit         uint64   `json:"limit"`
	Offset        uint64   `json:"offset"`
	JoinValidator bool     `json:"join_validator"`
}

type DashboardRequest struct {
	IndicesOrPubKey string `json:"indicesOrPubkey"`
}

type DiscordEmbed struct {
	Color       string              `json:"color,omitempty"`
	Description string              `json:"description,omitempty"`
	Fields      []DiscordEmbedField `json:"fields,omitempty"`
	Title       string              `json:"title,omitempty"`
	Type        string              `json:"type,omitempty"`
}

type DiscordEmbedField struct {
	Inline bool   `json:"inline"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

type DiscordComponent struct {
	Type       uint64                   `json:"type"`
	Components []DiscordComponentButton `json:"components"`
}

type DiscordComponentButton struct {
	Style    uint64 `json:"style"`
	CustomID string `json:"custom_id"`
	Label    string `json:"label"`
	URL      string `json:"url"`
	Disabled bool   `json:"disabled"`
	Type     uint64 `json:"type"`
}

type DiscordReq struct {
	Content         string             `json:"content,omitempty"`
	Username        string             `json:"username,omitempty"`
	Avatar_url      string             `json:"avatar_url,omitempty"`
	Tts             bool               `json:"tts,omitempty"`
	Embeds          []DiscordEmbed     `json:"embeds,omitempty"`
	AllowedMentions []interface{}      `json:"allowedMentions,omitempty"`
	Components      []DiscordComponent `json:"components,omitempty"`
	Files           interface{}        `json:"files,omitempty"`
	Payload         string             `json:"payload,omitempty"`
	Attachments     interface{}        `json:"attachments,omitempty"`
	Flags           int                `json:"flags,omitempty"`
}

type ExecutionPerformanceResponse struct {
	Performance1d  *big.Int `json:"performance1d"`
	Performance7d  *big.Int `json:"performance7d"`
	Performance31d *big.Int `json:"performance31d"`
	ValidatorIndex uint64   `json:"validatorindex"`
}

type ExecutionBlockApiResponse struct {
	Hash               string                `json:"blockHash"`
	BlockNumber        uint64                `json:"blockNumber"`
	Timestamp          uint64                `json:"timestamp"`
	BlockReward        *big.Int              `json:"blockReward"`
	BlockMevReward     *big.Int              `json:"blockMevReward"`
	FeeRecipientReward *big.Int              `json:"producerReward"`
	FeeRecipient       string                `json:"feeRecipient"`
	GasLimit           uint64                `json:"gasLimit"`
	GasUsed            uint64                `json:"gasUsed"`
	BaseFee            *big.Int              `json:"baseFee"`
	TxCount            uint64                `json:"txCount"`
	InternalTxCount    uint64                `json:"internalTxCount"`
	UncleCount         uint64                `json:"uncleCount"`
	ParentHash         string                `json:"parentHash"`
	UncleHash          string                `json:"uncleHash"`
	Difficulty         *big.Int              `json:"difficulty"`
	PoSData            *ExecBlockProposer    `json:"posConsensus"`
	RelayData          *RelayDataApiResponse `json:"relay"`
	ConsensusAlgorithm string                `json:"consensusAlgorithm"`
}

type RelayDataApiResponse struct {
	TagID                string `json:"tag"`
	BuilderPubKey        string `json:"builderPubkey"`
	ProposerFeeRecipient string `json:"producerFeeRecipient"`
}

type AddressIndexOrPubkey struct {
	Address []byte
	Index   uint64
	Pubkey  []byte
}

type RelaysData struct {
	MevRecipient  []byte    `db:"proposer_fee_recipient"`
	MevBribe      WeiString `db:"value"`
	ExecBlockHash []byte    `db:"exec_block_hash"`
	TagID         string    `db:"tag_id"`
	BuilderPubKey []byte    `db:"builder_pubkey"`
}

type ExecBlockProposer struct {
	ExecBlock uint64 `db:"exec_block_number" json:"executionBlockNumber"`
	Proposer  uint64 `db:"proposer" json:"proposerIndex"`
	Slot      uint64 `db:"slot" json:"slot"`
	Epoch     uint64 `db:"epoch" json:"epoch"`
	Finalized bool   `json:"finalized"`
}

func (e *DiscordReq) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &e)
}

func (a DiscordReq) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type ApiEth1AddressResponse struct {
	Address string `json:"address"`
	Ether   string `json:"ether"`
	Tokens  []struct {
		Address  string  `json:"address"`
		Balance  string  `json:"balance"`
		Symbol   string  `json:"symbol"`
		Decimals string  `json:"decimals,omitempty"`
		Price    float64 `json:"price,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"tokens"`
}

type APIEth1AddressTxResponse struct {
	Transactions []Eth1TransactionParsed
	Page         string `json:"page"`
}

type Eth1TransactionParsed struct {
	Hash               string    `json:"hash,omitempty"`
	BlockNumber        uint64    `json:"block,omitempty"`
	Time               time.Time `json:"time,omitempty"`
	MethodId           string    `json:"method,omitempty"`
	From               string    `json:"from,omitempty"`
	To                 string    `json:"to,omitempty"`
	Value              string    `json:"value,omitempty"`
	TxFee              string    `json:"fee,omitempty"`
	GasPrice           string    `json:"gasPrice,omitempty"`
	IsContractCreation bool      `json:"is_contract_creation,omitempty"`
	InvokesContract    bool      `json:"invokes_contract,omitempty"`
}
