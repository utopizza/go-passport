CREATE TABLE `account_login_name`
(
    `id`               bigint(20) unsigned NOT NULL COMMENT '主键，无业务意义',
    `user_id`          bigint(20) unsigned NOT NULL COMMENT '用户id',
    `account_group_id` int(11)             NOT NULL COMMENT '账号组',
    `login_name`       varchar(63)         NOT NULL COMMENT '登录名',
    `extra`            text COMMENT '附加信息',
    `idc`              char(16)            NOT NULL DEFAULT '' COMMENT 'idc',
    `create_time`      datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time`      timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_uid_agid` (`user_id`, `account_group_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='账号登录名'