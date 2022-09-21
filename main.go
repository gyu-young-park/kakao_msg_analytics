package main

import (
	"fmt"
	"gyu/control"
	"gyu/mapper"
	"gyu/messageparser"
)

const INPUT_FILE_PATH = "./input.csv"

func main() {
	control := control.Control{}
	control.SetMapper(make(mapper.MessageDataMapper))
	control.SetParser(&messageparser.CSVParser{})
	control.Analyze(INPUT_FILE_PATH)
	messageCountRankResp := control.ExecMessageCountRankList()
	for _, v := range messageCountRankResp {
		fmt.Printf("%d, %s:%d\n", v.Rank, v.Name, v.Count)
	}
	messageContentCountRankResp := control.ExecMessageContentCountRankList(5)
	for _, v := range messageContentCountRankResp {
		for i, count := range v.Counts {
			fmt.Printf("%d,%s %s:%d\n", i+1, v.Name, v.Contents[i], count)
		}
	}
}
