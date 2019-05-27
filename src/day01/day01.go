package main

import (
	"log"
	"os"
	"sync"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main 是整个程序的入口
func main() {
	// 使用特定的项做搜索
	search.Run("president")
}

// Run执行搜索逻辑
func Run(searchTerm string) {
	// 获取需要搜索的数据源列表

	feeds, err := RetrieveFeeds()

	if err != nil {

		log.Fatal(err)

	}

	// 创建一个无缓冲的通道，接收匹配后的结果

	results := make(chan *Result)

	// 构造一个waitGroup，以便处理所有的数据源

	var waitGroup sync.WaitGroup

	// 设置需要等待处理
	// 每个数据源的goroutine的数量

	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个goroutine来查找结果

	for _, feed := range feeds {
		// 获取一个匹配器用于查找

		matcher, exists := matchers[feed.Type]

		if !exists {

			matcher = matchers["default"]

		}

		// 启动一个goroutine来执行搜索

		go func(matcher Matcher, feed *Feed) {

			Match(matcher, feed, searchTerm, results)

			waitGroup.Done()

		}(matcher, feed)

	}

	// 启动一个goroutine来监控是否所有的工作都做完了

	go func() {
		// 等候所有任务完成

		waitGroup.Wait()

		// 用关闭通道的方式，通知Display函数
		// 可以退出程序了

		close(results)

	}()

	// 启动函数，显示返回的结果，并且
	// 在最后一个结果显示完后返回

	Display(results)

}
