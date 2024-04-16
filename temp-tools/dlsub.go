package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config 结构体标签应正确使用 yaml:"字段名" 的形式
type Config struct {
	URL string `yaml:"url"`
	Dir string `yaml:"dir"`
}

func main() {
	// 读取配置文件
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}

	// 下载并保存为 a.yaml 文件
	data, err := download(config.URL, "a.yaml")
	if err != nil {
		log.Fatalf("下载文件失败: %v", err)
	}

	err = os.WriteFile("a.yaml", data, 0644)
	if err != nil {
		log.Fatalf("无法保存文件: %v", err)
	}

	// 去除第一行
	lines := strings.Split(string(data), "\n")
	if len(lines) > 1 {
		data = []byte(strings.Join(lines[1:], "\n"))
	}

	// 拆分文件
	lines = strings.Split(string(data), "\n")
	if len(lines) > 500 {
		chunkSize := 500
		for i := 0; i < len(lines); i += chunkSize {
			end := i + chunkSize
			if end > len(lines) {
				end = len(lines)
			}

			chunk := strings.Join(lines[i:end], "\n")
			chunk = "proxies:\n" + chunk

			filename := fmt.Sprintf("%schunk%d.yaml", config.Dir, i/chunkSize)
			//filename := fmt.Sprintf("%s/%s_chunk%d.yaml", config.Dir, filepath.Base(config.URL), i/chunkSize)
			err = os.WriteFile(filename, []byte(chunk), 0644)
			if err != nil {
				log.Fatalf("无法保存拆分文件: %v", err)
			}
		}
	}
}

func download(url string, outputFilePath string) ([]byte, error) {
	cmd := exec.Command("wget", "-O", outputFilePath, url)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("无法使用wget下载文件: %v", err)
	}

	// 读取下载的文件内容并返回
	data, err := os.ReadFile(outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("无法读取下载的文件内容: %v", err)
	}

	return data, nil
}
