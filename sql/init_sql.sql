CREATE TABLE IF NOT EXISTS `user` (
    `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `age` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '年龄',
    `balance` decimal(11,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '余额',
    `updated_at` datetime NOT NULL COMMENT '更新时间',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `address` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `uid` int unsigned NOT NULL,
    `province` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `city` varchar(20) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `update_time` int unsigned NOT NULL,
    `create_time` int unsigned NOT NULL,
    `delete_time` int unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `uid` (`uid`)
    ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE IF NOT EXISTS `hobby` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `updated_at` int unsigned NOT NULL,
    `created_at` int unsigned NOT NULL,
    `deleted_at` int unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE
    ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

s
