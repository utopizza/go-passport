CREATE TABLE `account_connect`
(
    `id`                         bigint(20) unsigned               NOT NULL COMMENT '主键，无业务意义',
    `user_id`                    bigint(20)                        NOT NULL COMMENT '用户id',
    `platform_app_id`            int(11)                           NOT NULL COMMENT '第三方平台和app组合编号',
    `platform_account_id`        varchar(255) COLLATE utf8mb4_bin           DEFAULT NULL COMMENT '第三方账号唯一标识',
    `platform_screen_name`       varchar(127) COLLATE utf8mb4_bin  NOT NULL COMMENT '第三方平台账号展示名',
    `platform_profile_image_url` varchar(1023) COLLATE utf8mb4_bin NOT NULL COMMENT '第三方平台账号头像url',
    `description`                varchar(1023) COLLATE utf8mb4_bin          DEFAULT NULL COMMENT '第三方平台账号描述',
    `gender`                     tinyint(4)                                 DEFAULT NULL COMMENT '第三方平台账号性别',
    `unionid`                    varchar(255) COLLATE utf8mb4_bin           DEFAULT '',
    `openid`                     varchar(255) COLLATE utf8mb4_bin           DEFAULT '',
    `access_token`               varchar(255) COLLATE utf8mb4_bin  NOT NULL,
    `access_token_secret`        varchar(255) COLLATE utf8mb4_bin  NOT NULL,
    `refresh_token`              varchar(255) COLLATE utf8mb4_bin           DEFAULT NULL,
    `expired_time`               datetime                                   DEFAULT NULL,
    `extra`                      text COLLATE utf8mb4_bin COMMENT '附加信息',
    `idc`                        char(16) COLLATE utf8mb4_bin      NOT NULL DEFAULT '' COMMENT 'idc',
    `create_time`                datetime                          NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time`                timestamp                         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_uid_paid` (`user_id`, `platform_app_id`),
    KEY `idx_platform_account_id` (`platform_account_id`(128))
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='账号第三方登录绑定关系'