create table wf_flow
(
    id                  varchar(256)  not null
        primary key,

    instance_id_        varchar(256)  null,
    values_             text          null,
    keyword_            varchar(1024) null,

    diagram_id_         varchar(256)  null,
    start_key_          int           null,

    executed_keys_      varchar(1024) null, -- 当前已执行节点 []int
    activated_keys_     varchar(1024) null, -- 当前激活节点 []int
    status_             varchar(256)  null, -- 流程状态：Draft(草稿) Revoked(撤回) Suspended（挂起） Executing(执行中) Rejected(驳回) Finished(已结束)
    status_text_        varchar(1024) null,

    create_at_          datetime      null, -- 流程创建时间
    start_at_           datetime      null, -- 流程发起时间
    active_at_          datetime      null, -- 流程活动时间
    end_at_             datetime      null, -- 流程结束时间

    order_              bigint        null,
    create_depart_id_   varchar(256)  null,
    create_depart_code_ varchar(256)  null,
    create_depart_name_ varchar(256)  null,
    create_user_id_     varchar(256)  null,
    create_user_code_   varchar(256)  null,
    create_user_name_   varchar(256)  null
);

create table wf_flow_node
(
    id                    varchar(256) not null
        primary key,

    instance_id_          varchar(256) null,
    executed_id_          varchar(256) null,

    diagram_id_           varchar(256) null,
    key_                  int          null,
    category_             varchar(256) null,
    code_                 varchar(256) null,
    name_                 varchar(256) null,

    executor_user_id_     varchar(256) null,
    executor_user_name_   varchar(256) null,

    activated_at_         datetime     null, -- 激活时间
    canceled_at_          datetime     null, -- 作废时间
    executed_at_          datetime     null, -- 执行时间
    comment_              text         null, -- 流转意见

    status_               varchar(256) null, -- 流转状态：Executing(执行中) Canceled(作废) ExecutedAdvanced(已执行) ExecutedRejected(已驳回)

    order_                bigint       null,
    executed_depart_id_   varchar(256) null,
    executed_depart_code_ varchar(256) null,
    executed_depart_name_ varchar(256) null,
    executed_user_id_     varchar(256) null,
    executed_user_code_   varchar(256) null,
    executed_user_name_   varchar(256) null
);

create table wf_flow_executors
(
    id                  varchar(256) not null
        primary key,

    diagram_id_         varchar(256) null,
    key_                int          null,

    order_              bigint       null,
    executor_user_id_   varchar(256) null,
    executor_user_name_ varchar(256) null,

    create_at_          datetime     null,
    create_user_id_     varchar(256) null,
    create_user_name_   varchar(256) null
);

-- 查询索引
SHOW
    INDEX FROM sys_role;


-- 查询外键
SELECT C.table_name             AS table_,
       C.constraint_name        AS code_,
       C.column_name            AS column_,
       C.referenced_table_name  AS referenced_table_,
       C.referenced_column_name AS referenced_column_
FROM information_schema.key_column_usage C
         INNER JOIN information_schema.tables T ON T.TABLE_NAME = C.TABLE_NAME
         INNER JOIN information_schema.referential_constraints R
                    ON R.table_name = C.table_name AND R.constraint_name = C.constraint_name AND
                       R.referenced_table_name = C.referenced_table_name
WHERE C.referenced_table_name IS NOT NULL
  AND C.table_schema = 'phoenix';