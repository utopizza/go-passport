CREATE TABLE `account_delete`
(
    `id`               bigint(20) unsigned          NOT NULL COMMENT '主键，无业务意义',
    `user_id`          bigint(20) unsigned          NOT NULL COMMENT '用户id',
    `account_group_id` int(11)                      NOT NULL COMMENT '账号组',
    `app_id`           int(11)                      NOT NULL COMMENT 'app_id',
    `infos`            text COLLATE utf8mb4_bin COMMENT '账号信息，json string格式',
    `extra`            text COLLATE utf8mb4_bin COMMENT '附加信息',
    `idc`              char(16) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'idc',
    `create_time`      datetime                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time`      timestamp                    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    KEY `idx_uid` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='账户物理删除信息存档'