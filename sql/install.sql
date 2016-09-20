-- Adminer 4.2.3 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';

CREATE TABLE `a_client` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT '名称',
  `recv_url` varchar(200) NOT NULL COMMENT '接收通知的网址（支持占位符：{type},{event},{id}）',
  `type_magazine` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否(1/0)接收杂志通知',
  `type_book` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否(1/0)接收图书通知',
  `type_article` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否(1/0)接收文章通知',
  `event_add` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否(1/0)接收新增通知',
  `event_edit` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否(1/0)接收修改通知',
  `event_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否(1/0)接收删除通知',
  `pre_hook` varchar(30) NOT NULL DEFAULT '' COMMENT '通知之前的处理钩子',
  `hook_param` varchar(500) NOT NULL DEFAULT '' COMMENT '钩子参数',
  `disabled` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '禁用',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='接收通知的客户端';


CREATE TABLE `a_event` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(5) NOT NULL COMMENT '名称:add,del,edit',
  `target` varchar(8) NOT NULL COMMENT '目标:magazine,book,article',
  `target_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '目标id',
  `memo` varchar(500) NOT NULL DEFAULT '' COMMENT '说明',
  `created` int(10) unsigned NOT NULL COMMENT '`omitempty`创建时间',
  `finished` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '`omitempty`完成时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='事件';


CREATE TABLE `a_notice` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `client_id` int(10) unsigned NOT NULL,
  `event_id` bigint(20) unsigned NOT NULL,
  `created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '`omitempty`创建时间',
  `retry` smallint(2) unsigned NOT NULL DEFAULT '0' COMMENT '`omitempty`重试次数',
  `finished` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '`omitempty`完成时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `client_id_event_id` (`client_id`,`event_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='通知记录';


-- 2016-09-20 04:09:32