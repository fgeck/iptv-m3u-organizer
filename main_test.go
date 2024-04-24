package main_test

import (
	"os"
	"testing"

	"github.com/fgeck/iptv-m3u-organizer/organizer"
)

func Test_Organizer_Run_Yaml(t *testing.T) {
	organizer := organizer.NewOrganizer("./config.yaml", "https://someurl.com/streams", "", "./downloads/result.m3u", "urlParam")
	organizer.Run()
}

func Test_Organizer_Run_CLI_Params(t *testing.T) {
	os.Setenv("USER", "test")
	os.Setenv("PASSWORD", "testpassword")
	organizer := organizer.NewOrganizer("", "https://someurl.com/streams", "{\"fullmatch\": {\"group\": [\"|DE| SPORT\"], \"name\": [\"DE - DAS ERSTE HD\"]}, \"softmatch\": {\"group\": [\"|DE| SPORT\"], \"name\": [\"DE - ZDF HD\"]}}", "./downloads/result.m3u", "urlParam")
	organizer.Run()
}
