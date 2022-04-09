CREATE TABLE `account_email`
(
    `id`               bigint(20)   NOT NULL PRIMARY KEY,
    `user_id`          bigint(20)   NOT NULL,
    `email`            varchar(255) NOT NULL,
    `email_id`         bigint(20) DEFAULT NULL,
    `extra`            text
)
