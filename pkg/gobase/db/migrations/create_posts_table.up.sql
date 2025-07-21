CREATE TABLE `posts` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) DEFAULT NULL,
                         `updated_at` datetime(3) DEFAULT NULL,
                         `deleted_at` datetime(3) DEFAULT NULL,
                         `title` longtext,
                         `body` longtext,
                         `author` longtext,
                         `user_id` bigint unsigned DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         KEY `idx_posts_deleted_at` (`deleted_at`),
                         KEY `fk_users_posts` (`user_id`),
                         CONSTRAINT `fk_users_posts` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;