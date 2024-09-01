package jsvm

import (
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type Context struct {
	ctx *handle.Context
}

func NewContext(ctx *handle.Context) *Context {
	return &Context{ctx}
}

func (b *Context) Params() map[string]string {
	return b.ctx.Params()
}

func (b *Context) Values() map[string]string {
	return b.ctx.Values()
}

func (b *Context) DepartId() string {
	return b.ctx.DepartId()
}

func (b *Context) DepartCode() string {
	return b.ctx.DepartCode()
}

func (b *Context) DepartName() string {
	return b.ctx.DepartName()
}

func (b *Context) UserId() string {
	return b.ctx.UserId()
}

func (b *Context) UserCode() string {
	return b.ctx.UserCode()
}

func (b *Context) UserName() string {
	return b.ctx.UserName()
}

func (b *Context) DateTime() string {
	return asql.GetNow()
}

func (b *Context) Date() string {
	return asql.GetNow()[:10]
}
