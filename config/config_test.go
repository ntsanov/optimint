package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestViperAndCobra(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	cmd := &cobra.Command{}
	AddFlags(cmd)

	v := viper.GetViper()
	assert.NoError(v.BindPFlags(cmd.Flags()))

	assert.NoError(cmd.Flags().Set(flagAggregator, "true"))
	assert.NoError(cmd.Flags().Set(flagDALayer, "foobar"))
	assert.NoError(cmd.Flags().Set(flagDAConfig, `{"json":true}`))
	assert.NoError(cmd.Flags().Set(flagBlockTime, "1234s"))
	assert.NoError(cmd.Flags().Set(flagNamespaceID, "0102030405060708"))

	nc := DefaultNodeConfig
	assert.NoError(nc.GetViperConfig(v))

	assert.Equal(true, nc.Aggregator)
	assert.Equal("foobar", nc.DALayer)
	assert.Equal(`{"json":true}`, nc.DAConfig)
	assert.Equal(1234*time.Second, nc.BlockTime)
	assert.Equal([8]byte{1, 2, 3, 4, 5, 6, 7, 8}, nc.NamespaceID)
}
