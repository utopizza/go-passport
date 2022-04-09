CREATE TABLE `account_info`
(
    `id`               bigint(20) unsigned              NOT NULL COMMENT '主键，无业务意义',
    `user_id`          bigint(20) unsigned              NOT NULL COMMENT '用户id',
    `account_group_id` bigint(20) unsigned              NOT NULL COMMENT '账号组',
    `screen_name`      varchar(127) COLLATE utf8mb4_bin NOT NULL COMMENT '展示名',
    `avatar_url`       varchar(511) COLLATE utf8mb4_bin          DEFAULT NULL COMMENT '头像',
    `description`      varchar(1023) COLLATE utf8mb4_bin         DEFAULT NULL COMMENT '描述/签名儿',
    `password`         varchar(127) COLLATE utf8mb4_bin NOT NULL COMMENT '密码',
    `status`           smallint(5) unsigned                      DEFAULT NULL COMMENT '状态',
    `gender`           tinyint(4)                                DEFAULT NULL COMMENT '性别',
    `extra`            text COLLATE utf8mb4_bin COMMENT '附加信息',
    `idc`              char(16) COLLATE utf8mb4_bin     NOT NULL DEFAULT '' COMMENT 'idc',
    `create_time`      datetime                         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time`      timestamp                        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_uid_agid` (`user_id`, `account_group_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='账号基础信息'