package mapper

import "gyu/model"

type Mapper interface {
	MappingData([][]string)
	GetMessageCountRankList() []*model.MessageCountRankResp
	GetMessageContentCountRankList(int) []*model.MessageContentCountRankResp
}
