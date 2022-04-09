CREATE TABLE `account_mobile`
(
    `id`               bigint(20)          NOT NULL PRIMARY KEY,
    `user_id`          bigint(20)          NOT NULL,
    `mobile`           varchar(31)         NOT NULL,
    `mobile_id`        bigint(20)          NOT NULL,
    `extra`            text
)