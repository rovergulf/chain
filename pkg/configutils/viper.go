package configutils

import (
	"github.com/rovergulf/chain/node"
	"github.com/rovergulf/chain/params"
	"github.com/rovergulf/chain/pkg/traceutils"
	"github.com/spf13/viper"
	"os"
)

func SetDefaultConfigValues() {
	viper.SetDefault("metrics", true)
	viper.SetDefault("metrics_addr", ":8088")
	viper.SetDefault(traceutils.CollectorUrlConfigKey, os.Getenv("JAEGER_TRACE"))

	// chain settings

	// storage
	viper.SetDefault("db", "")
	viper.SetDefault("data_dir", "tmp")
	viper.SetDefault("keystore", "")

	// process id
	viper.SetDefault("pid_file", "/var/run/rbn/pidfile")

	// TBD dgraphdb connection settings
	// !!! Database interface needs to be implemented to use that
	viper.SetDefault("dgraphdb.enabled", false)
	viper.SetDefault("dgraphdb.host", "127.0.0.1")
	viper.SetDefault("dgraphdb.port", "9080")
	viper.SetDefault("dgraphdb.user", "")
	viper.SetDefault("dgraphdb.password", "")
	viper.SetDefault("dgraphdb.tls.enabled", false)
	viper.SetDefault("dgraphdb.tls.cert", "")
	viper.SetDefault("dgraphdb.tls.key", "")
	viper.SetDefault("dgraphdb.tls.verify", false)
	viper.SetDefault("dgraphdb.tls.auth", "")

	// chain network setup
	viper.SetDefault("network.id", params.MainNetworkId)

	// p2p settings
	viper.SetDefault("node.max_peers", 256)
	viper.SetDefault("node.addr", "127.0.0.1")
	viper.SetDefault("node.port", 9420)
	viper.SetDefault("node.sync_mode", node.SyncModeDefault)
	viper.SetDefault("node.sync_interval", 5)
	viper.SetDefault("node.cache_dir", "")
	viper.SetDefault("node.no_discovery", false)

	// http server
	viper.SetDefault("http.disabled", false)
	viper.SetDefault("http.addr", "127.0.0.1")
	viper.SetDefault("http.port", 9469)
	viper.SetDefault("http.dial_timeout", 30)
	viper.SetDefault("http.read_timeout", 30)
	viper.SetDefault("http.write_timeout", 30)
	viper.SetDefault("http.ssl.enabled", false)
	viper.SetDefault("http.ssl.cert", "")
	viper.SetDefault("http.ssl.key", "")
	viper.SetDefault("http.ssl.verify", false)

	// ws server

	// json rpc server
	// TBD

	//

	// TBD
	// Cache
	//viper.SetDefault("cache.enabled", false)
	viper.SetDefault("cache.size", 256<<20) // 256mb

	// Runtime configuration
	//viper.SetDefault("runtime.max_cpu", runtime.NumCPU())
	//viper.SetDefault("runtime.max_mem", getAvailableOSMemory())
}
