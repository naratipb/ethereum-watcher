package ethereum_watcher

import (
	"context"
	"testing"

	"github.com/naratipb/ethereum-watcher/plugin"
	"github.com/naratipb/ethereum-watcher/structs"
	"github.com/sirupsen/logrus"
)

func TestNewBlockNumPlugin(t *testing.T) {
	logrus.SetLevel(logrus.InfoLevel)

	api := "https://mainnet.infura.io/v3/19d753b2600445e292d54b1ef58d4df4"
	w := NewHttpBasedEthWatcher(context.Background(), api)

	logrus.Println("waiting for new block...")
	w.RegisterBlockPlugin(plugin.NewBlockNumPlugin(func(i uint64, b bool) {
		logrus.Printf(">> found new block: %d, is removed: %t", i, b)
	}))

	w.RunTillExit()
}

func TestSimpleBlockPlugin(t *testing.T) {
	api := "https://mainnet.infura.io/v3/19d753b2600445e292d54b1ef58d4df4"
	w := NewHttpBasedEthWatcher(context.Background(), api)

	w.RegisterBlockPlugin(plugin.NewSimpleBlockPlugin(func(block *structs.RemovableBlock) {
		logrus.Infof(">> %+v", block.Block)
	}))

	w.RunTillExit()
}
