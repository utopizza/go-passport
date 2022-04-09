CREATE TABLE `account_connect`
(
    `id`                         bigint(20) NOT NULL PRIMARY KEY,
    `user_id`                    bigint(20) NOT NULL,
    `platform_app_id`            int(11)    NOT NULL,
    `platform_account_id`        varchar(255)           DEFAULT NULL,
    `platform_screen_name`       varchar(127) NOT NULL,
    `platform_profile_image_url` varchar(1023) NOT NULL,
    `description`                varchar(1023) DEFAULT NULL,
    `gender`                     tinyint(4)             DEFAULT NULL,
    `unionid`                    varchar(255)           DEFAULT '',
    `openid`                     varchar(255)           DEFAULT '',
    `access_token`               varchar(255)  NOT NULL,
    `access_token_secret`        varchar(255)  NOT NULL,
    `refresh_token`              varchar(255)           DEFAULT NULL,
    `expired_time`               datetime               DEFAULT NULL,
    `extra`                      text
)