package provider

import (
	"strings"

	"github.com/fzdy-zz/proxypool/pkg/tool"
)

type Hysteria2Sub struct {
	Base
}

func (sub Hysteria2Sub) Provide() string {
	sub.Types = "hysteria2"
	sub.preFilter()
	var resultBuilder strings.Builder
	for _, p := range *sub.Proxies {
		resultBuilder.WriteString(p.Link() + "\n")
	}
	return tool.Base64EncodeString(resultBuilder.String(), false)
}
