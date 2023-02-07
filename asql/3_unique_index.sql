
------ 唯一索引 ------
ALTER TABLE sys_depart ADD UNIQUE INDEX sys_user_UniqueIndex_code (code_);
ALTER TABLE sys_dict_item ADD UNIQUE INDEX sys_dict_item_UniqueIndex_kind_id_code (kind_id_,code_);
ALTER TABLE sys_dict_kind ADD UNIQUE INDEX sys_dict_kind_UniqueIndex_code (code_);
ALTER TABLE sys_organization_role ADD UNIQUE INDEX sys_organization_role_UniqueIndex_organization_id_role_id (organization_id_,role_id_);
ALTER TABLE sys_role ADD UNIQUE INDEX sys_role_UniqueIndex_code (code_);
ALTER TABLE sys_role_menu ADD UNIQUE INDEX sys_role_menu_UniqueIndex_role_id_menu_id (role_id_,menu_id_);
ALTER TABLE sys_table ADD UNIQUE INDEX sys_table_UniqueIndex_code (code_);
ALTER TABLE sys_table_column ADD UNIQUE INDEX sys_table_column_UniqueIndex_table_id_code (table_id_,code_);
ALTER TABLE sys_user ADD UNIQUE INDEX sys_user_UniqueIndex_depart_id_account_id (depart_id_,account_id_);

ALTER TABLE sys_auto_no ADD UNIQUE INDEX sys_auto_no_UniqueIndex_code (kind_id_,prefix_);
ALTER TABLE sys_auto_no_kind ADD UNIQUE INDEX sys_auto_no_kind_UniqueIndex_code (code_);

ALTER TABLE sys_data_service ADD UNIQUE INDEX sys_data_service_UniqueIndex_table_id_method_code (table_id_,method_,code_);

ALTER TABLE wf_diagram ADD UNIQUE INDEX wf_diagram_UniqueIndex_code (code_);
ALTER TABLE wf_options_diagram ADD UNIQUE INDEX wf_options_diagram_UniqueIndex_diagram_id (diagram_id_);
ALTER TABLE wf_options_node ADD UNIQUE INDEX wf_options_node_UniqueIndex_diagram_id_key (diagram_id_,key_);
ALTER TABLE wf_options_link ADD UNIQUE INDEX wf_options_link_UniqueIndex_diagram_id_from_key_to_key (diagram_id_,from_key_,to_key_);

ALTER TABLE wf_flow ADD UNIQUE INDEX wf_flow_UniqueIndex_instance_id (instance_id_);




