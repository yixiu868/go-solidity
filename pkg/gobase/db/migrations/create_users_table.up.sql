CREATE TABLE `users` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) DEFAULT NULL,
                         `updated_at` datetime(3) DEFAULT NULL,
                         `deleted_at` datetime(3) DEFAULT NULL,
                         `name` longtext,
                         `email` longtext,
                         `age` tinyint unsigned DEFAULT NULL,
                         `birthday` datetime(3) DEFAULT NULL,
                         `member_number` longtext,
                         `activated_at` datetime(3) DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;