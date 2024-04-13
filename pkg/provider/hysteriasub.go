package provider

import (
	"strings"

	"github.com/fzdy-zz/proxypool/pkg/tool"
)

type HysteriaSub struct {
	Base
}

func (sub HysteriaSub) Provide() string {
	sub.Types = "hysteria"
	sub.preFilter()
	var resultBuilder strings.Builder
	for _, p := range *sub.Proxies {
		resultBuilder.WriteString(p.Link() + "\n")
	}
	return tool.Base64EncodeString(resultBuilder.String(), false)
}
