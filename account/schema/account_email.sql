CREATE TABLE `account_email`
(
    `id`               bigint(20) unsigned NOT NULL COMMENT '主键，无业务意义',
    `user_id`          bigint(20) unsigned NOT NULL COMMENT '用户id',
    `account_group_id` int(11)             NOT NULL COMMENT '账号组',
    `email`            varchar(255)        NOT NULL COMMENT '邮箱',
    `extra`            text COMMENT '附加信息',
    `idc`              char(16)            NOT NULL DEFAULT '' COMMENT 'idc',
    `create_time`      datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time`      timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `email_id`         bigint(20)                   DEFAULT NULL COMMENT 'email_id',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_uid_agid` (`user_id`, `account_group_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='账号邮箱绑定关系'
