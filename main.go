package main

import (
	"log"
	"os"

	"github.com/fgeck/iptv-m3u-organizer/organizer"
	"github.com/spf13/cobra"
)

var (
	url            string
	configFilePath string
	outputFilePath string
	filter         string
	authType       string
)

func main() {
	logger := log.New(os.Stdout, "main: ", log.Ldate|log.Ltime|log.Lshortfile)
	rootCmd := &cobra.Command{
		Use:   "iptv-m3u-organizer",
		Short: "A CLI tool to organize IPTV M3U playlists",
		Long:  `iptv-m3u-organizer is a command-line tool that helps you organize IPTV M3U playlists. Either a configuration file or a URL, filter and output file path must be provided.`,
	}

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the IPTV M3U organizer",
		Run:   runOrganizer,
	}

	runCmd.Flags().StringVarP(&url, "url", "u", "", "URL for the M3U playlist")
	runCmd.Flags().StringVarP(&configFilePath, "config", "c", "", "Path to the YAML configuration file")
	runCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "Path to the resulting filtered m3u file")
	runCmd.Flags().StringVarP(&filter, "filter", "f", "", "JSON or YAML filter configuration")
	runCmd.Flags().StringVarP(&authType, "auth", "a", "", "authentication type. Must be basic or urlParam.")

	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}

func runOrganizer(cmd *cobra.Command, args []string) {
	organizer := organizer.NewOrganizer(configFilePath, url, filter, outputFilePath, authType)
	organizer.Run()
}
