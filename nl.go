package main

import (
	"cloud.google.com/go/language/apiv1"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
	"log"
)

func main() {
	doAnalyze()
}

func doAnalyze() {
	// コンテキスト生成
	ctx := context.Background()

	// クライアント生成
	client, err := language.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// 分析するテキストを宣言
	text := "Natural Language API Client LibrariesクライアントをGoで実装してみる。"

	// AnalyzeSyntaxRequest構造体を生成
	req := &languagepb.AnalyzeSyntaxRequest{
		Document: &languagepb.Document{
			// 指定なし 0
			// Plain text 1
			// HTML 2
			Type: 1,
			// 分析する内容
			Source: &languagepb.Document_Content{
				Content: text,
			},
		},
		EncodingType: languagepb.EncodingType_UTF8,
	}

	// 構文解析を行う
	resp, err := client.AnalyzeSyntax(ctx, req)
	if err != nil {
		log.Fatalf("Failed to AnalyzeSyntax: %v", err)
	}

	// jsonにする
	rawJson, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Failed to Marshal json: %v", err)
	}

	jsonStr := string(rawJson)

	fmt.Println("構文解析の結果\n" + jsonStr)

}
