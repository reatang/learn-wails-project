package controllers

import (
	"learn-wails-project/backend/types"
)

type HelloController struct {
}

func (c *HelloController) Hello(req *types.Request) *types.Response {

	return &types.Response{
		Message: "Hello " + req.Name,
	}
}
