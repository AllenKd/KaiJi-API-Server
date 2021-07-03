package configs

import (
	"github.com/spf13/viper"
	"testing"
)

func TestNew(t *testing.T) {
	viper.AddConfigPath("../../../configs")
	c := New()
	t.Log(c.Mode)

}
