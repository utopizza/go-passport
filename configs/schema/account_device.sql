CREATE TABLE `account_device`
(
    `id`               bigint(20) unsigned NOT NULL COMMENT '主键，无业务意义',
    `user_id`          bigint(20) unsigned NOT NULL COMMENT '用户id',
    `app_id`           int(11)             NOT NULL,
    `account_group_id` int(11)             NOT NULL COMMENT '账号组',
    `upgraded`         int(11)             NOT NULL DEFAULT '0' COMMENT '是否升级为正式帐号',
    `device_id`        bigint(20) unsigned NOT NULL COMMENT '设备id',
    `extra`            text COMMENT '附加信息',
    `idc`              char(16)            NOT NULL DEFAULT '' COMMENT 'idc',
    `create_time`      datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time`      timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_uid_aid_agid` (`user_id`, `app_id`, `account_group_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='账号设备绑定关系'
