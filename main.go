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
	filter         string
)

func main() {
	logger := log.New(os.Stdout, "main: ", log.Ldate|log.Ltime|log.Lshortfile)
	rootCmd := &cobra.Command{
		Use:   "iptv-m3u-organizer",
		Short: "A CLI tool to organize IPTV M3U playlists",
		Long:  `iptv-m3u-organizer is a command-line tool that helps you organize IPTV M3U playlists.`,
	}

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the IPTV M3U organizer",
		Run:   runOrganizer,
	}

	runCmd.Flags().StringVarP(&url, "url", "u", "", "URL for the M3U playlist")
	runCmd.Flags().StringVarP(&configFilePath, "config", "c", "", "Path to the YAML configuration file")
	runCmd.Flags().StringVarP(&filter, "filter", "f", "", "JSON or YAML filter configuration")

	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}

func runOrganizer(cmd *cobra.Command, args []string) {
	organizer := organizer.NewOrganizer(configFilePath, url, filter)
	organizer.Run()
}
