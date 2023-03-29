create table sys_time_task
(
    id                         varchar(256) not null
        primary key,

    name_                      varchar(256) null,
    type_                      varchar(256) null,

    once_at_                   datetime     null,

    frequency_                 varchar(256) null, -- 频率：每天、每周、每月
    frequency_day_             varchar(256) null,
    frequency_day_repeat_      int          null,
    frequency_day_repeat_unit_ varchar(256) null,
    frequency_day_start_       time         null,
    frequency_day_end_         time         null,
    frequency_start_at_        datetime     null,
    frequency_end_at_          datetime     null,

    description_               text         null,

    is_start_                  tinyint      null,
    source_                    text         null,

    order_                     bigint       null,
    create_at_                 datetime     null,
    update_at_                 datetime     null
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