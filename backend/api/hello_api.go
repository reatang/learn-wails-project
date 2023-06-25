package api

import (
	"learn-wails-project/backend/types"
)

type HelloApi struct {
}

func (c *HelloApi) Hello(req *types.Request) *types.Response {

	return &types.Response{
		Message: "Hello " + req.Name,
	}
}
