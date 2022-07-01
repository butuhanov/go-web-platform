package basic

import (
    "web-platform/pipeline"
    "web-platform/services"
)

type ServicesComponent struct {}

func (c *ServicesComponent) Init() {}

func (c *ServicesComponent)  ProcessRequest(ctx *pipeline.ComponentContext,
        next func(*pipeline.ComponentContext))  {
    reqContext := ctx.Request.Context()
    ctx.Request.WithContext(services.NewServiceContext(reqContext))
    next(ctx)
}