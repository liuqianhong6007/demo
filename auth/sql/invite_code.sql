CREATE TABLE IF NOT EXISTS `invite_code` (
      `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '邀请码ID',
      `invite_code` varchar(50) NOT NULL COMMENT '邀请码',
      `used` smallint NOT NULL COMMENT '是否已使用',
      `create_time` bigint(20) NOT NULL COMMENT '创建时间',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='邀请码';