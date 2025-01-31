package ximalaya

import (
	"testing"

	"github.com/135e2/annie/extractors/types"
	"github.com/135e2/annie/test"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "normal test",
			args: test.Args{
				URL:   "https://www.ximalaya.com/sound/211583675",
				Title: "狼的眼睛为什么会发光 - 十万个为什么【宝宝巴士百科故事】",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := New().Extract(tt.args.URL, types.Options{})
			test.CheckError(t, err)
			test.Check(t, tt.args, data[0])
		})
	}
}
