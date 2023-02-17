package flow

import (
	"go-phoenix/xwf/enum"
)

type Flowable interface {
	DiagramId() string
	Key() int
	Code() string
	Name() string
	Category() enum.Category

	Backwards(values string) ([]Flowable, error)
}

// StartFlowable 开始
type StartFlowable interface {
	Flowable

	Revocable() bool

	Start(flowId string, values string, comment string) error
	Revoke(flowId string, values string) error
}

// ExecuteFlowable 执行环节
type ExecuteFlowable interface {
	Flowable

	Rejectable() bool
	RequireRejectComment() bool

	ScopeExecutors(flowId string, values string) (scope []string, err error) // 默认执行者的组织ID
	CustomExecutors(scope []string) (executors map[string]string, err error) // 执行者范围

	ExecuteStart(flowId string, executors map[string]string) error               // 启动执行环节
	ExecuteAccept(id string, values string, comment string) error                // 结束执行环节
	ExecuteReject(id string, flowId string, values string, comment string) error // 驳回执行环节
}

// BranchFlowable 分支
type BranchFlowable interface {
	Flowable

	Branch(flowId string) error
}

// EndFlowable 结束
type EndFlowable interface {
	Flowable

	End(flowId string, values string) error
}
