package main_test

import (
	"testing"

	"github.com/fgeck/iptv-m3u-organizer/config"
	"github.com/fgeck/iptv-m3u-organizer/organizer"
)

func Test_Organizer_Run(t *testing.T) {
	organizer := organizer.NewOrganizer("./config.yaml", "", "", "./downloads/result.m3u", string(config.BasicAuth))
	organizer.Run()
}
