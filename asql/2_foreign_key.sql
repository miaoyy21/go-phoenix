
------ 主外键约束 ------
ALTER TABLE sys_depart ADD CONSTRAINT sys_depart_ForeignKey_parent_id FOREIGN KEY (parent_id_) REFERENCES sys_depart(id);
ALTER TABLE sys_dict_item ADD CONSTRAINT sys_dict_item_ForeignKey_kind_id FOREIGN KEY (kind_id_) REFERENCES sys_dict_kind(id);
ALTER TABLE sys_menu ADD CONSTRAINT sys_menu_ForeignKey_parent_id FOREIGN KEY (parent_id_) REFERENCES sys_menu(id);
ALTER TABLE sys_organization_role ADD CONSTRAINT sys_organization_role_ForeignKey_role_id FOREIGN KEY (role_id_) REFERENCES sys_role(id);
ALTER TABLE sys_role_menu ADD CONSTRAINT sys_role_menu_ForeignKey_role_id FOREIGN KEY (role_id_) REFERENCES sys_role(id);
ALTER TABLE sys_role_menu ADD CONSTRAINT sys_role_menu_ForeignKey_menu_id FOREIGN KEY (menu_id_) REFERENCES sys_menu(id);
ALTER TABLE sys_table_column ADD CONSTRAINT sys_table_column_ForeignKey_table_id FOREIGN KEY (table_id_) REFERENCES sys_table(id);
ALTER TABLE sys_user ADD CONSTRAINT sys_user_ForeignKey_depart_id FOREIGN KEY (depart_id_) REFERENCES sys_depart(id);

ALTER TABLE sys_auto_no ADD CONSTRAINT sys_auto_no_ForeignKey_kind_id FOREIGN KEY (kind_id_) REFERENCES sys_auto_no_kind(id);
ALTER TABLE sys_auto_no_item ADD CONSTRAINT sys_auto_no_item_ForeignKey_kind_id FOREIGN KEY (kind_id_) REFERENCES sys_auto_no_kind(id);

ALTER TABLE sys_data_service ADD CONSTRAINT sys_data_service_ForeignKey_table_id FOREIGN KEY (table_id_) REFERENCES sys_table(id);

ALTER TABLE wf_options_diagram ADD CONSTRAINT wf_options_diagram_ForeignKey_diagram_id FOREIGN KEY (diagram_id_) REFERENCES wf_diagram(id);
ALTER TABLE wf_options_node ADD CONSTRAINT wf_options_node_ForeignKey_diagram_id FOREIGN KEY (diagram_id_) REFERENCES wf_diagram(id);
ALTER TABLE wf_options_link ADD CONSTRAINT wf_options_link_ForeignKey_diagram_id FOREIGN KEY (diagram_id_) REFERENCES wf_diagram(id);

ALTER TABLE wf_flow ADD CONSTRAINT wf_flow_ForeignKey_diagram_id FOREIGN KEY (diagram_id_) REFERENCES wf_options_diagram(diagram_id_);
ALTER TABLE wf_flow_task ADD CONSTRAINT wf_flow_task_ForeignKey_diagram_id FOREIGN KEY (diagram_id_) REFERENCES wf_options_diagram(diagram_id_);
ALTER TABLE wf_flow_executors ADD CONSTRAINT wf_flow_executors_ForeignKey_diagram_id FOREIGN KEY (diagram_id_) REFERENCES wf_options_diagram(diagram_id_);
