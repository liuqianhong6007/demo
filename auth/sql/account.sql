CREATE TABLE IF NOT EXISTS `account` (
      `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '账号ID',
      `account` varchar(50) NOT NULL COMMENT '账户',
      `password` varchar(50) NOT NULL COMMENT '密码',
      `create_time` bigint(20) NOT NULL COMMENT '创建时间',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账号';