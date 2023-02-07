package jsvm

import (
	"go-phoenix/handle"
)

type Context struct {
	ctx *handle.Context
}

func NewContext(ctx *handle.Context) *Context {
	return &Context{ctx}
}

func (b *Context) Params() map[string]string {
	return b.ctx.GetParams()
}

func (b *Context) Values() map[string]string {
	return b.ctx.GetValues()
}

func (b *Context) DepartId() string {
	return b.ctx.GetDepartId()
}

func (b *Context) DepartCode() string {
	return b.ctx.GetDepartCode()
}

func (b *Context) DepartName() string {
	return b.ctx.GetDepartName()
}

func (b *Context) UserId() string {
	return b.ctx.GetUserId()
}

func (b *Context) UserCode() string {
	return b.ctx.GetUserCode()
}

func (b *Context) UserName() string {
	return b.ctx.GetUserName()
}
