CREATE TABLE `account_info`
(
    `id`               bigint(20) NOT NULL PRIMARY KEY,
    `user_id`          bigint(20) NOT NULL,
    `screen_name`      varchar(127) NOT NULL,
    `avatar_url`       varchar(511) DEFAULT NULL,
    `description`      varchar(1023)  DEFAULT NULL,
    `password`         varchar(127) NOT NULL,
    `status`           smallint(5)  DEFAULT NULL,
    `gender`           tinyint(4) DEFAULT NULL,
    `extra`            text DEFAULT NULL
)