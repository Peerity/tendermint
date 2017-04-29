package tendermint

type Config struct {
	Node       NodeConfig       `mapstructure:"node"`
	Chain      ChainConfig      `mapstructure:"chain"`
	ABCI       ABCIConfig       `mapstructure:"abci"`
	Network    NetworkConfig    `mapstructure:"network"`
	Blockchain BlockchainConfig `mapstructure:"blockchain"`
	Consensus  ConsensusConfig  `mapstructure:"consensus"`
	Block      BlockConfig      `mapstructure:"block"`
	Mempool    MempoolConfig    `mapstructure:"mempool"`
	RPC        RPCConfig        `mapstructure:"rpc"`
	DB         DBConfig         `mapstructure:"db"`
}

type NodeConfig struct {
	Moniker           string `mapstructure:"moniker"`             // "anonymous"
	PrivValidatorFile string `mapstructure:"priv_validator_file"` // rootDir+"/priv_validator.json")

	LogLevel       string `mapstructure:"log_level"`  // info
	ProfListenAddr string `mapstructure:"prof_laddr"` // ""
}

type ChainConfig struct {
	ChainID     string `mapstructure:"chain_id"`
	GenesisFile string `mapstructure:"genesis_file"` // rootDir/genesis.json
}

type ABCIConfig struct {
	ProxyApp string `mapstructure:"proxy_app"` //  tcp://0.0.0.0:46658
	Mode     string `mapstructure:"mode"`      // socket

	FilterPeers bool `mapstructure:"filter_peers"` // false
}

type NetworkConfig struct {
	ListenAddr     string `mapstructure:"listen_adddr"` // "tcp://0.0.0.0:46656")
	Seeds          string `mapstructure:"seeds"`        // []string ...
	SkipUPNP       bool   `mapstructure:"skip_upnp"`
	AddrBookFile   string `mapstructure:"addr_book_file"`   // rootDir+"/addrbook.json")
	AddrBookString bool   `mapstructure:"addr_book_string"` // true
	PexReactor     bool   `mapstructure:"pex_reactor"`      // false
}

type BlockchainConfig struct {
	FastSync bool `mapstructure:"fast_sync"` // true
}

type ConsensusConfig struct {
	WalFile  string //rootDir+"/data/cs.wal/wal")
	WalLight bool   // false

	// all timeouts are in ms
	TimeoutPropose        int // 3000
	TimeoutProposeDelta   int // 500
	TimeoutPrevote        int // 1000
	TimeoutPrevoteDelta   int // 500
	TimeoutPrecommit      int // 1000
	TimeoutPrecommitDelta int // 500
	TimeoutCommit         int // 1000

	// make progress asap (no `timeout_commit`) on full precommit votes
	SkipTimeoutCommit bool // false
}

type BlockConfig struct {
	MaxTxs          int  // 10000
	PartSize        int  // 65536
	DisableDataHash bool // false
}

type MempoolConfig struct {
	Recheck      bool   // true
	RecheckEmpty bool   // true
	Broadcast    bool   // true
	WalDir       string // rootDir+"/data/mempool.wal")
}

type RPCConfig struct {
	RPCListenAddress  string // "tcp://0.0.0.0:46657")
	GRPCListenAddress string // ""
}

type DBConfig struct {
	Backend string `mapstructure:"backend"` // leveldb
	Dir     string `mapstructure:"dir"`     // rootDir/data

	TxIndex string `mapstructure:"tx_index"` // "kv"
}
