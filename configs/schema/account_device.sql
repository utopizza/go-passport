CREATE TABLE `account_device`
(
    `id`               bigint(20) NOT NULL PRIMARY KEY,
    `user_id`          bigint(20) NOT NULL,
    `app_id`           int(11) NOT NULL,
    `device_id`        bigint(20) NOT NULL,
    `extra`            text
)
