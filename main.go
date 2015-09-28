package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/segmentio/go-prompt"
	"net/http"
)

func startUp() {
	articleTitle := prompt.StringRequired("文章标题")
	articleUrl := prompt.StringRequired("文章地址")
	recommandReason := prompt.StringRequired("推荐理由")
	keywords := prompt.StringRequired("文章标签")
	author := prompt.StringRequired("推荐人")

	info := map[string]interface{}{
		"title":       articleTitle,
		"url":         articleUrl,
		"description": recommandReason,
		"tags":        keywords,
		"provider":    author,
	}

	mJson, _ := json.Marshal(info)
	contentReader := bytes.NewReader(mJson)
	req, _ := http.NewRequest("POST", "http://www.75team.com/weekly/admin/article.php?action=add", contentReader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notes", "go version 75weekly is comming")
	client := &http.Client{}
	resp, _ := client.Do(req)

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("感谢您的投稿%s。您的文章已经投递成功，文章标题为：%s", author, articleTitle)
		fmt.Println("按键^C退出命令行投稿")
		startUp()
	}
}

func main() {
	fmt.Println(ansi.Color("欢迎使用奇舞周刊命令行, 请按照提示信息输入周刊信息", "red+b"))
	startUp()
}
