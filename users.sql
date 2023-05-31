CREATE DATABASE `crud`;
USE `crud`;
CREATE TABLE `users`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(100)    NOT NULL,
    `created_at` timestamp       NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp       NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp       NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 7
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
