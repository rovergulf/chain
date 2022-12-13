package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/mitchellh/go-homedir"
	"github.com/rovergulf/chain/params"
	"github.com/rovergulf/chain/pkg/configutils"
	"github.com/rovergulf/chain/pkg/logutils"
	"github.com/rovergulf/chain/pkg/resutils"
	"github.com/rovergulf/chain/wallets"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var (
	logger *zap.SugaredLogger

	cfgFile, dataDir string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chain",
	Short: "Rovergulf Chain application",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $ROVERGULF_CHAIN_HOME/config.yaml)")
	// logger & debug opts
	rootCmd.PersistentFlags().Bool("log_json", false, "Enable JSON formatted logs output")
	rootCmd.PersistentFlags().Int("log_level", int(zapcore.DebugLevel), "Log level")
	rootCmd.PersistentFlags().Bool("log_stacktrace", false, "Log stacktrace verbose")
	rootCmd.PersistentFlags().Bool("dev", true, "Enable development/testing environment") // always true during initial development

	// main flags
	rootCmd.PersistentFlags().StringVar(&dataDir, "data_dir", os.Getenv("DATA_DIR"), "BlockChain data directory")

	rootCmd.Flags().BoolP("version", "v", false, "Show application version")

	bindViperPersistentFlag(rootCmd, "app.dev", "dev")
	bindViperPersistentFlag(rootCmd, "log_json", "log_json")
	bindViperPersistentFlag(rootCmd, "log_level", "log_level")
	bindViperPersistentFlag(rootCmd, "log_stacktrace", "log_stacktrace")
	bindViperPersistentFlag(rootCmd, "data_dir", "data_dir")

	rootCmd.AddCommand(walletsCmd())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	configutils.SetDefaultConfigValues()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".chain" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(os.Getenv("ROVERGULF_CHAIN_HOME"))
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	lg, err := logutils.NewLogger()
	if err != nil {
		log.Fatalf("Unable to init zap logger: %s", err)
	}
	logger = lg
}

func writeOutput(cmd *cobra.Command, v interface{}) error {
	outputFormat, _ := cmd.Flags().GetString("output")
	if outputFormat == "json" {
		return resutils.WriteJSON(os.Stdout, logger, v)
	} else {
		return resutils.WriteYAML(os.Stdout, logger, v)
	}
}

func bindViperFlag(cmd *cobra.Command, viperVal, flagName string) {
	if err := viper.BindPFlag(viperVal, cmd.Flags().Lookup(flagName)); err != nil {
		log.Printf("Failed to bind viper flag: %s", err)
	}
}

func bindViperPersistentFlag(cmd *cobra.Command, viperVal, flagName string) {
	if err := viper.BindPFlag(viperVal, cmd.PersistentFlags().Lookup(flagName)); err != nil {
		log.Printf("Failed to bind viper flag: %s", err)
	}
}

func addOutputFormatFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("output", "o", "yaml", "specify output format (yaml/json)")
}

func addNetworkIdFlag(cmd *cobra.Command) {
	cmd.Flags().String("network-id", hexutil.EncodeUint64(params.MainNetworkId), "Chain network id")
	bindViperFlag(cmd, "network-id", "network-id")
}

func addAddressFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("address", "a", "", "Specify wallet address")
	cmd.MarkFlagRequired("address")
	bindViperFlag(cmd, "address", "address")
}

func prepareWalletsManager(cmd *cobra.Command, args []string) error {
	wm, err := wallets.NewManager()
	if err != nil {
		return err
	}

	accountManager = wm
	return nil
}
