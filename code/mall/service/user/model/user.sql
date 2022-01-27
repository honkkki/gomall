CREATE TABLE `user` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                        `name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户姓名',
                        `gender` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '用户性别',
                        `mobile` varchar(255) NOT NULL DEFAULT '' COMMENT '用户电话',
                        `password` varchar(255) NOT NULL DEFAULT '' COMMENT '用户密码',
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `udx_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;