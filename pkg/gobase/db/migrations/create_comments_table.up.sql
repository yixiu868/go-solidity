CREATE TABLE `comments` (
                            `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) DEFAULT NULL,
                            `updated_at` datetime(3) DEFAULT NULL,
                            `deleted_at` datetime(3) DEFAULT NULL,
                            `content` text,
                            `post_id` bigint unsigned DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `idx_comments_deleted_at` (`deleted_at`),
                            KEY `fk_posts_comments` (`post_id`),
                            CONSTRAINT `fk_posts_comments` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;