package plugins

import (
	"errors"
	"log"
	"os"
	"path"
	"plugin"
	"reflect"
	"runtime"
	"strings"

	"github.com/tecnologer/uno/src/engine"
)

func Load() ([]engine.Game, error) {
	plugins, err := getPluginsList()
	if err != nil {
		return nil, err
	}

	games := make([]engine.Game, 0)

	for _, mod := range plugins {
		// load module
		// 1. open the so file to load the symbols
		plug, err := plugin.Open(mod)
		if err != nil {
			return nil, err
		}

		// 2. look up a symbol (an exported function or variable)
		// in this case, variable Game
		symGame, err := plug.Lookup("Game")
		if err != nil {
			return nil, err
		}

		// 3. Assert that loaded symbol is of a desired type
		// in this case interface type Greeter (defined above)
		var game engine.Game
		game, ok := symGame.(engine.Game)
		if !ok {
			log.Printf("unexpected type from module symbol: %s", reflect.TypeOf(symGame))
			continue
		}

		games = append(games, game)
	}

	return games, nil
}

func getPluginsList() (plugins []string, err error) {
	plugins = make([]string, 0)

	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, errors.New("get dir for connection file")
	}

	dirPath := path.Dir(filename)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".so") {
			continue
		}

		plugins = append(plugins, path.Join(dirPath, e.Name()))
	}

	return
}
