package utils

import (
	"AD/dto"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// 调用远程服务器 API 获取预测结果
func CallAIPrediction(imagePath string) (*dto.PredictionResponse, error) {
	// 构造远程命令
	remoteCommand := fmt.Sprintf(
		"source /home/xyc/bigtiao/softvote/venv/bin/activate && python /home/xyc/bigtiao/softvote/judge.py %s",
		imagePath,
	)

	// 通过 SSH 执行远程命令
	cmd := exec.Command("ssh", "xyc@183.6.97.121", remoteCommand)

	// 捕获输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to execute remote command: %v, output: %s", err, string(output))
	}

	// 查找 JSON 的起始和结束位置
	start := strings.Index(string(output), "{")
	end := strings.LastIndex(string(output), "}")

	// 检查是否找到 JSON 数据
	if start == -1 || end == -1 {
		return nil, fmt.Errorf("no JSON data found in output")
	}

	// 提取 JSON 部分
	jsonStr := string(output)[start : end+1]

	// 将单引号替换为双引号
	jsonStr = strings.ReplaceAll(jsonStr, "'", "\"")

	// 解析 JSON
	var predictionResponse dto.PredictionResponse
	err = json.Unmarshal([]byte(jsonStr), &predictionResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse prediction response: %v", err)
	}

	return &predictionResponse, nil

	return &predictionResponse, nil
}
