CREATE TABLE `tbl_msg` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `content` varchar(50) NOT NULL DEFAULT '' COMMENT '内容',
    `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_uid_ctime` (`uid`,`ctime`),
    KEY `ix_mtime` (`mtime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='消息表';