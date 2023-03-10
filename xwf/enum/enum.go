package enum

type Category string

var (
	CategoryStart   Category = "Start"
	CategoryExecute Category = "Execute"
	CategoryBranch  Category = "Branch"
	CategoryEnd     Category = "End"
)

type ExecutorPolicy string

var (
	ExecutorPolicyNone                ExecutorPolicy = "None"
	ExecutorPolicyAllUsers            ExecutorPolicy = "AllUsers"
	ExecutorPolicyStartDepartLeader   ExecutorPolicy = "StartDepartLeader"
	ExecutorPolicyForwardDepartLeader ExecutorPolicy = "ForwardDepartLeader"
)

type FlowStatus string

var (
	FlowStatusDraft     FlowStatus = "Draft"
	FlowStatusRevoked   FlowStatus = "Revoked"
	FlowStatusSuspended FlowStatus = "Suspended"
	FlowStatusExecuting FlowStatus = "Executing"
	FlowStatusRejected  FlowStatus = "Rejected"
	FlowStatusFinished  FlowStatus = "Finished"
	FlowStatusDiscarded FlowStatus = "Discarded"
)

type FlowNodeStatus string

var (
	FlowNodeStatusExecuting        FlowStatus = "Executing"
	FlowNodeStatusCanceled         FlowStatus = "Canceled"
	FlowNodeStatusExecutedAuto     FlowStatus = "Executed Auto"
	FlowNodeStatusExecutedAccepted FlowStatus = "Executed Accepted"
	FlowNodeStatusExecutedRejected FlowStatus = "Executed Rejected"
)
