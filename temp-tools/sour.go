package main

//已经完成，测试可用
import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config 结构体用于保存从sour.yaml中读取的配置
type Config struct {
	Dir   string `yaml:"dir"`
	URL   string `yaml:"url"`
	Dizhi string `yaml:"dizhi"`
}

func main() {
	// 读取sour.yaml中的配置
	configData, err := os.ReadFile("sour.yaml")
	if err != nil {
		fmt.Println("读取配置文件时出错:", err)
		return
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println("配置解析时出错:", err)
		return
	}

	// 列出目标目录中所有的YAML文件
	files, err := os.ReadDir(config.Dir)
	if err != nil {
		fmt.Println("读取目录时出错:", err)
		return
	}

	// 创建一个结构体来保存指定格式的YAML文件名
	type Change struct {
		Type    string `yaml:"type"`
		Options struct {
			URL string `yaml:"url"`
		} `yaml:"options"`
	}

	var changes []Change

	// 为chang.yaml生成数据
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".yaml" {
			change := Change{
				Type: "clash",
			}
			change.Options.URL = config.URL + file.Name()

			changes = append(changes, change)
		}
	}

	// 将结构体编组为YAML
	data, err := yaml.Marshal(changes)
	if err != nil {
		fmt.Println("将数据编组为YAML时出错:", err)
		return
	}

	// 将数据写入chang.yaml
	err = os.WriteFile("chang.yaml", data, 0644)
	if err != nil {
		fmt.Println("写入到chang.yaml时出错:", err)
		return
	}

	// 将chang.yaml保存到“dizhi”变量指定的路径
	err = os.WriteFile(config.Dizhi+"chang.yaml", data, 0644)
	if err != nil {
		fmt.Println("将chang.yaml保存到指定目录时出错:", err)
		return
	}

	fmt.Println("成功生成了包含所需数据的chang.yaml，并将其保存到指定目录中！🚀")
}
