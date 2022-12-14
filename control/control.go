package control

import (
	"fmt"
	"gyu/mapper"
	"gyu/messageparser"
	"gyu/model"
	customrouter "gyu/route"
	"net/http"

	"github.com/gorilla/mux"
)

type Control struct {
	mapper mapper.Mapper
	parser messageparser.Parser
	router *mux.Router
}

func (c *Control) SetMapper(mapper mapper.Mapper) {
	c.mapper = mapper
}

func (c *Control) SetParser(parser messageparser.Parser) {
	c.parser = parser
}

func (c *Control) SetRouter(mux *mux.Router, routerSetter customrouter.CustomRouter) {
	routerSetter.SetRouteWithHandler(mux)
	c.router = mux
	c.startServer()
}

func (c *Control) startServer() {
	port := ":8081"
	fmt.Println("start server", port)
	http.ListenAndServe(":8081", c.router)
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
