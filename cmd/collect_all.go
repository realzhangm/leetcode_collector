package main

import (
	"flag"
	collector "github.com/realzhangm/leetcode_collector/pkg/collector"
	"github.com/realzhangm/leetcode_collector/pkg/doa"
)

var (
	noFetch = flag.Bool("noFetch", false, "do not fetch from LeetCode")
)

// 实现简单的功能
// 将 LeetCode 上的提交的代码，同步到本地，并生产一个 Markdown 汇总文件。
func run() {
	c := collector.NewCollector(collector.GetConfig())
	doa.MustOK(c.LoadInfo())
	if !*noFetch {
		doa.MustOK(c.FetchAllFromLeetCode())
	}
	doa.MustOK(c.JsonToMarkDown())
	doa.MustOK(c.OutputSolutionsCode())
	doa.MustOK(c.OutputTagsMarkDown())
}

func main() {
	flag.Parse()
	run()
}
