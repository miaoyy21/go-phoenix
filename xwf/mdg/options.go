package mdg

import "go-phoenix/xwf/enum"

type OptionsNode struct {
	Key int `json:"key"`

	Category enum.Category `json:"category"`
	Code     string        `json:"code_"`

	Rejectable           bool `json:"rejectable_,omitempty"`
	RequireRejectComment bool `json:"require_reject_comment_,omitempty"`
	Revocable            bool `json:"revocable_"`

	OnBeforeScript string `json:"on_before_script_,omitempty"`
	OnAfterScript  string `json:"on_after_script_,omitempty"`
	OnRejectScript string `json:"on_reject_script_,omitempty"`
	OnRevokeScript string `json:"on_revoke_script_,omitempty"`
	OnRemoveScript string `json:"on_remove_script_,omitempty"`

	ExecutorCustomNum     string              `json:"executor_custom_num_,omitempty"`
	ExecutorSelectableNum string              `json:"executor_selectable_num_,omitempty"`
	ExecutorSavable       bool                `json:"executor_savable_,omitempty"`
	ExecutorUsers         string              `json:"executor_users_,omitempty"`
	ExecutorNameUsers     string              `json:"executor_name_users_,omitempty"`
	ExecutorDeparts       string              `json:"executor_departs_,omitempty"`
	ExecutorNameDeparts   string              `json:"executor_name_departs_,omitempty"`
	ExecutorRoles         string              `json:"executor_roles_,omitempty"`
	ExecutorNameRoles     string              `json:"executor_name_roles_,omitempty"`
	ExecutorPolicy        enum.ExecutorPolicy `json:"executor_policy_,omitempty"`
	ExecutorScript        string              `json:"executor_script_,omitempty"`
}

type OptionsLink struct {
	From     int    `json:"from"`
	To       int    `json:"to"`
	OnScript string `json:"on_script_"`
}

type Options struct {
	Diagram struct {
		Code        string `json:"code_"`
		Name        string `json:"name_"`
		Icon        string `json:"icon_"`
		Description string `json:"description_"`
		Keyword     string `json:"keyword_"`
		ExceedDays  string `json:"exceed_days_"`
	} `json:"diagram"`

	Nodes []OptionsNode `json:"nodes"`
	Links []OptionsLink `json:"links"`
}
