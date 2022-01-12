package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/zhs007/block7"
	block7game "github.com/zhs007/block7/game"
	"github.com/zhs007/goutils"
	"go.uber.org/zap"
)

func main() {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	lst := []*block7game.MapState{}
	goutils.InitLogger("xlsx2map", block7.Version, "debug", true, "./")
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".xlsx" {
			if strings.Contains(info.Name(), "$") {
				return nil
			}

			stage, err := block7game.LoadExcel(path)
			if err != nil {
				goutils.Error("loadExcel",
					zap.String("fn", path),
					zap.Error(err))

				return nil
			}

			fn := strings.Split(info.Name(), ".")[0]
			fd, err := json.Marshal(stage)
			if err != nil {
				goutils.Error("Marshal",
					goutils.JSON("stage", stage),
					zap.Error(err))

				return nil
			}

			ioutil.WriteFile(fmt.Sprintf("%v.json", fn), fd, 0644)

			ms, err := stage.Analyze2(fn)
			if err != nil {
				goutils.Error("stage.Analyze2",
					zap.String("fn", fn),
					zap.Error(err))

				return nil
			}

			lst = append(lst, ms)

			return nil
		}

		return nil
	})
	if err != nil {
		goutils.Error("Walk",
			zap.Error(err))
	}

	block7game.SaveMapStateList("allstate.json", lst)
}
