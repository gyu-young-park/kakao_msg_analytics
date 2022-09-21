package mapper

import (
	"gyu/model"
	"sort"
	"strings"
)

type MessageDataList []*model.MessageData
type MessageDataMapper map[string]MessageDataList

func (m MessageDataMapper) getAllMessageCount(key string) int {
	return len(m[key])
}

func (m MessageDataMapper) MappingData(rows [][]string) {
	for _, row := range rows {
		date := row[0]
		if strings.Contains(date, "Date") {
			continue
		}
		name := row[1]
		content := row[2]
		m[name] = append(m[name], model.NewMessageData(date, name, content))
	}
}

func (m MessageDataMapper) GetMessageCountRankList() []*model.MessageCountRankResp {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m.getAllMessageCount(keys[i]) > m.getAllMessageCount(keys[j])
	})
	var resp []*model.MessageCountRankResp
	for i, key := range keys {
		resp = append(resp, model.NewMessageCountRankResp(key, m.getAllMessageCount(key), i+1))
	}
	return resp
}

func (m MessageDataMapper) GetMessageContentCountRankList(limit int) []*model.MessageContentCountRankResp {
	var resp []*model.MessageContentCountRankResp
	for k, v := range m {
		contentMap := make(map[string]int)
		for j, _ := range v {
			trimmedContent := strings.TrimSpace(m[k][j].GetContent())
			contents := strings.Split(trimmedContent, " ")
			for _, content := range contents {
				if strings.Contains(content, "사진") || content == "" {
					continue
				}
				contentMap[content]++
			}
		}
		keys := make([]string, 0, len(m))
		for key := range contentMap {
			keys = append(keys, key)
		}
		sort.SliceStable(keys, func(i, j int) bool {
			return contentMap[keys[i]] > contentMap[keys[j]]
		})
		var rankedContent []string
		var countContent []int
		for i, key := range keys {
			if i >= limit {
				break
			}
			countContent = append(countContent, contentMap[key])
			rankedContent = append(rankedContent, key)
			//fmt.Printf("%d, %s %s:%d\n", i+1, k, key, contentMap[key])
		}
		resp = append(resp, model.NewMessageContentCountRankResp(k, countContent, rankedContent))
	}
	return resp
}
