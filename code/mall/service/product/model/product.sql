CREATE TABLE `product` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `name` varchar(255) NOT NULL DEFAULT '' COMMENT '产品名称',
                           `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '产品描述',
                           `stock` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '产品库存',
                           `amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '产品金额',
                           `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '产品状态 -1已删除',
                           `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                           `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;