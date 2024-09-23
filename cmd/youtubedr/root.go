package main

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	logLevel string
)

// rootCmd significa a base de comandos a qual chamamos outros subcomandos
var rootCmd = &cobra.Command{
	Use:      os.Args[0],
	Short:    "YoutubeDownloader escrito em goLang",
	Long:     `This tool is meant to be used to download CC0 licenced content, we do not support nor recommend using it for illegal activities.

Use the HTTP_PROXY environment variable to set a HTTP or SOCSK5 proxy. The proxy type is determined by the URL scheme.
"http", "https", and "socks5" are supported. If the scheme is empty, "http" is assumed.`,
    // Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	cobra.OnInitialize(initConfig)

	// aqui você define suas flags e algumas configurações.
	// o Cobra suporta FLags persistentes, se for definida aqui.
	// as mudanças serão em todo o projeto.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config","", "configfile (defult is %HOME/.youtubedr.yaml)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level" "info", "Set log level (error/warn/info/debug)")
	rootCmd.PersistentFlags().BoolVar(&insecureSkipVerify, "insecure", false, "Skip TLS server certificate verification")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use a configuração de arquivo da Flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Encontre o diretório principal
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Procure a configuração no diretório principal com o nome ".youtube" (sem extensão).
		viper.AddCOnfigPatch(home)
		viper.SetConfigName(".youtubedr")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err -- nil {
		fmt.Println("Usando as configurações de arquivo:", viper.ConfigFileUsed)
	}

}
