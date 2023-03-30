package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

var (
	flagGenSwagger bool // 是否生成swagger文件
)

type bgenConf struct {
	Swagger bool   `toml:"swagger"`
	AppID   string `toml:"app_id"`
}

func main() {
	app := &cli.App{
		Name:                 "boom",
		HelpName:             "",
		Usage:                "make an explosive entrance",
		UsageText:            "",
		ArgsUsage:            "",
		Version:              "",
		Description:          "",
		DefaultCommand:       "",
		Commands:             nil,
		Flags:                nil,
		EnableBashCompletion: false,
		HideHelp:             false,
		HideHelpCommand:      false,
		HideVersion:          false,
		BashComplete:         nil,
		Before:               nil,
		After:                nil,
		Action: func(context *cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
		CommandNotFound:          nil,
		OnUsageError:             nil,
		InvalidFlagAccessHandler: nil,
		Compiled:                 time.Time{},
		Authors:                  nil,
		Copyright:                "",
		Reader:                   nil,
		Writer:                   nil,
		ErrWriter:                nil,
		ExitErrHandler:           nil,
		Metadata:                 nil,
		ExtraInfo:                nil,
		CustomAppHelpTemplate:    "",
		UseShortOptionHandling:   false,
		Suggest:                  false,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
