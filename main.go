package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Penguin-71630/meme-bot-frontend-dc/bot"
	"github.com/Penguin-71630/meme-bot-frontend-dc/tracing"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use: "dcbot",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
		viper.BindPFlags(cmd.PersistentFlags())
		viper.BindPFlags(cmd.Flags())
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		appname := "go2025-dcbot"
		tracing.InitTracer(appname)
		if viper.GetString("uptrace-dsn") != "" {
			tracing.InitUptrace(appname)
			defer tracing.DeferUptrace(ctx)
		}

		// Initialize bot
		discordBot, err := bot.New()
		if err != nil {
			tracing.Logger.Ctx(ctx).
				Panic("failed to create bot",
					zap.Error(err))
			panic(err)
		}

		// Start bot
		if err := discordBot.Start(); err != nil {
			tracing.Logger.Ctx(ctx).
				Panic("failed to start bot",
					zap.Error(err))
			panic(err)
		}
		fmt.Println("Bot is now running. Press CTRL-C to exit.")
		defer discordBot.Stop()

		// Wait for interrupt signal
		sc := make(chan os.Signal, 1)
		signal.Notify(sc,
			syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc
	},
}

func init() {
	cobra.EnableTraverseRunHooks = true
	rootCmd.Flags().
		String("discord-bot-token", "", "discord bot token")
	rootCmd.Flags().
		String("api-endpoint", "http://localhost:8080", "api endpoint")

	rootCmd.Flags().
		Bool("zap-production", true, "Toggle production log format")
	rootCmd.Flags().
		String("uptrace-dsn", "", "Uptrace DSN (disabled by default)")
}

func main() {
	rootCmd.Execute()
}
