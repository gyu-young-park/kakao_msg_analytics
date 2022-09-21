package control

import (
	"gyu/mapper"
	"gyu/messageparser"
	"gyu/model"
)

type Control struct {
	mapper mapper.Mapper
	parser messageparser.Parser
}

func (c *Control) SetMapper(mapper mapper.Mapper) {
	c.mapper = mapper
}

func (c *Control) SetParser(parser messageparser.Parser) {
	c.parser = parser
}

func (c *Control) Analyze(inputFile string) error {
	rows, err := c.parser.Parse(inputFile)
	if err != nil {
		return err
	}
	c.mapper.MappingData(rows)
	return nil
}

func (c *Control) ExecMessageCountRankList() []*model.MessageCountRankResp {
	return c.mapper.GetMessageCountRankList()
}

func (c *Control) ExecMessageContentCountRankList(limit int) []*model.MessageContentCountRankResp {
	return c.mapper.GetMessageContentCountRankList(limit)
}
