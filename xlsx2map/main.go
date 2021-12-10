package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/zhs007/block7"
	block7game "github.com/zhs007/block7/game"
	"github.com/zhs007/goutils"
	"go.uber.org/zap"
)

func main() {
	goutils.InitLogger("xlsx2map", block7.Version, "debug", true, "./")
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".xlsx" {
			stage, err := block7game.LoadExcel(path)
			if err != nil {
				goutils.Error("loadExcel",
					zap.String("fn", path),
					zap.Error(err))

				return err
			}

			fn := strings.Split(info.Name(), ".")[0]
			fd, err := json.Marshal(stage)
			if err != nil {
				goutils.Error("Marshal",
					goutils.JSON("stage", stage),
					zap.Error(err))

				return err
			}

			ioutil.WriteFile(fmt.Sprintf("%v.json", fn), fd, 0644)

			stage.Analyze(fn)

			return nil
		}

		return nil
	})
	if err != nil {
		goutils.Error("Walk",
			zap.Error(err))
	}
}
