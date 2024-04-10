package provider

import (
	"strings"

	"github.com/fzdy-zz/proxypool/pkg/tool"
)

type VlessSub struct {
	Base
}

func (sub VlessSub) Provide() string {
	sub.Types = "vless"
	sub.preFilter()
	var resultBuilder strings.Builder
	for _, p := range *sub.Proxies {
		resultBuilder.WriteString(p.Link() + "\n")
	}
	return tool.Base64EncodeString(resultBuilder.String(), false)
}
