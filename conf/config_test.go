package conf

import (
	"fmt"
	"testing"
)

func TestOnParseConfig(t *testing.T) {
	t.Run("Test on ParseConfig", func(t *testing.T) {
		err := Parse()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(GlobalConfig)
		}
	})
}
