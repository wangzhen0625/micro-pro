/*
Navicat MySQL Data Transfer

Source Server         : 84
Source Server Version : 50715
Source Host           : 192.168.74.84:3306
Source Database       : topsecgw

Target Server Type    : MYSQL
Target Server Version : 50715
File Encoding         : 65001

Date: 2018-01-05 15:34:27
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for 1231231WHsFUp
-- ----------------------------
DROP TABLE IF EXISTS `1231231WHsFUp`;
CREATE TABLE `1231231WHsFUp` (
  `user_id` int(8) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL,
  `account_name` varchar(255) NOT NULL,
  `account_pwd` varchar(255) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for aadas6nIJvP
-- ----------------------------
DROP TABLE IF EXISTS `aadas6nIJvP`;
CREATE TABLE `aadas6nIJvP` (
  `user_id` int(8) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL,
  `account_name` varchar(255) NOT NULL,
  `account_pwd` varchar(255) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for ad_access
-- ----------------------------
DROP TABLE IF EXISTS `ad_access`;
CREATE TABLE `ad_access` (
  `role_id` smallint(6) unsigned NOT NULL,
  `node_id` varchar(256) NOT NULL,
  `level` tinyint(1) NOT NULL,
  `module` varchar(50) DEFAULT NULL,
  KEY `groupId` (`role_id`),
  KEY `nodeId` (`node_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for ad_node
-- ----------------------------
DROP TABLE IF EXISTS `ad_node`;
CREATE TABLE `ad_node` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `title` varchar(50) DEFAULT NULL,
  `child` tinyint(1) unsigned DEFAULT '0',
  `status` tinyint(1) DEFAULT '0',
  `remark` varchar(255) DEFAULT NULL,
  `sort` smallint(6) unsigned DEFAULT NULL,
  `pid` smallint(6) unsigned NOT NULL,
  `level` tinyint(1) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `level` (`level`),
  KEY `pid` (`pid`),
  KEY `status` (`status`),
  KEY `name` (`name`)
) ENGINE=MyISAM AUTO_INCREMENT=242 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for ad_node_copy
-- ----------------------------
DROP TABLE IF EXISTS `ad_node_copy`;
CREATE TABLE `ad_node_copy` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `title` varchar(50) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  `remark` varchar(255) DEFAULT NULL,
  `sort` smallint(6) unsigned DEFAULT NULL,
  `pid` smallint(6) unsigned NOT NULL,
  `level` tinyint(1) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `level` (`level`),
  KEY `pid` (`pid`),
  KEY `status` (`status`),
  KEY `name` (`name`)
) ENGINE=MyISAM AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for ad_role
-- ----------------------------
DROP TABLE IF EXISTS `ad_role`;
CREATE TABLE `ad_role` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `status` tinyint(1) unsigned DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `pid` smallint(6) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`),
  KEY `status` (`status`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for admin_operation_audit
-- ----------------------------
DROP TABLE IF EXISTS `admin_operation_audit`;
CREATE TABLE `admin_operation_audit` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `description` varchar(255) DEFAULT NULL,
  `operator` varchar(255) DEFAULT NULL,
  `target_name` varchar(255) DEFAULT NULL,
  `module_name` varchar(255) DEFAULT NULL,
  `operation_type` varchar(11) DEFAULT NULL,
  `level` varchar(255) DEFAULT NULL,
  `level_name` varchar(255) DEFAULT NULL,
  `channel` varchar(255) DEFAULT NULL,
  `operation_result` varchar(255) DEFAULT NULL,
  `operation_time` int(11) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for admin_operation_log
-- ----------------------------
DROP TABLE IF EXISTS `admin_operation_log`;
CREATE TABLE `admin_operation_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `operator` varchar(255) DEFAULT '',
  `opt_ip` varchar(255) DEFAULT '',
  `operation` varchar(255) DEFAULT '',
  `opt_object` varchar(255) DEFAULT '',
  `opt_content` varchar(255) DEFAULT '',
  `opt_result` varchar(255) DEFAULT '',
  `opt_time` int(11) DEFAULT '0',
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6777 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for administrator
-- ----------------------------
DROP TABLE IF EXISTS `administrator`;
CREATE TABLE `administrator` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `certb64` blob,
  `certmd5` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `loginip` varchar(255) DEFAULT NULL,
  `logintime` int(11) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `id` int(8) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
  `alias` varchar(255) DEFAULT '',
  `appBaseAddr` varchar(255) DEFAULT '',
  `appType` int(2) DEFAULT '0',
  `access_mode` tinyint(1) DEFAULT '0',
  `port` int(10) DEFAULT '0',
  `ssl_offload` tinyint(1) DEFAULT '0',
  `access_addr` varchar(255) DEFAULT '',
  `browserName` varchar(255) DEFAULT '',
  `browserVersion` varchar(255) DEFAULT '',
  `clientUrl` varchar(255) DEFAULT '',
  `ssoId` int(2) DEFAULT '0',
  `status` int(11) DEFAULT '0',
  `appImg` blob,
  `displayOrder` int(11) DEFAULT '0',
  `desc` varchar(255) DEFAULT '' COMMENT '描述',
  `create_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_name` (`name`),
  UNIQUE KEY `uni_port` (`port`) USING BTREE,
  UNIQUE KEY `uni_acc_addr` (`access_addr`),
  UNIQUE KEY `uni_alias` (`alias`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for app_browser
-- ----------------------------
DROP TABLE IF EXISTS `app_browser`;
CREATE TABLE `app_browser` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(11) unsigned DEFAULT NULL,
  `browser_id` int(11) DEFAULT NULL,
  `brower_version` varchar(11) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `create_time` int(11) unsigned DEFAULT NULL,
  `update_time` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for app_browser_type
-- ----------------------------
DROP TABLE IF EXISTS `app_browser_type`;
CREATE TABLE `app_browser_type` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for app_image
-- ----------------------------
DROP TABLE IF EXISTS `app_image`;
CREATE TABLE `app_image` (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `img` mediumblob,
  `type` varchar(255) DEFAULT NULL,
  `mime` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for app_rule
-- ----------------------------
DROP TABLE IF EXISTS `app_rule`;
CREATE TABLE `app_rule` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT '',
  `type` tinyint(1) DEFAULT '0',
  `app_id` varchar(255) DEFAULT '',
  `app_uris` varchar(255) DEFAULT '',
  `week` varchar(255) DEFAULT '',
  `start_time` int(11) DEFAULT '0',
  `end_time` int(11) DEFAULT '0',
  `network_seg` varchar(255) DEFAULT '',
  `mask` varchar(255) DEFAULT '',
  `create_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for app_uri
-- ----------------------------
DROP TABLE IF EXISTS `app_uri`;
CREATE TABLE `app_uri` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_id` int(11) DEFAULT '0',
  `uri` varchar(255) DEFAULT '',
  `create_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config` (
  `name` varchar(255) NOT NULL,
  `value` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for config_mgr
-- ----------------------------
DROP TABLE IF EXISTS `config_mgr`;
CREATE TABLE `config_mgr` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `cfg` text,
  `desc` varchar(255) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for log_column
-- ----------------------------
DROP TABLE IF EXISTS `log_column`;
CREATE TABLE `log_column` (
  `id` int(11) NOT NULL,
  `column_name` varchar(255) DEFAULT NULL COMMENT '字段名称',
  `option` varchar(255) DEFAULT NULL COMMENT '可选项内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for mail_service
-- ----------------------------
DROP TABLE IF EXISTS `mail_service`;
CREATE TABLE `mail_service` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT '' COMMENT '模板名称',
  `desc` varchar(255) DEFAULT '',
  `status` int(1) DEFAULT NULL,
  `host` varchar(255) DEFAULT '' COMMENT '邮件服务器域名',
  `port` varchar(255) DEFAULT '' COMMENT '邮件服务器域名端口',
  `smtp_secure` tinyint(1) unsigned DEFAULT '0' COMMENT '是否使用ssl加密 0不使用',
  `sender_name` varchar(255) DEFAULT '' COMMENT '发送者的邮件地址',
  `sender_pwd` varchar(255) DEFAULT '' COMMENT '发送者的密码',
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for mail_template
-- ----------------------------
DROP TABLE IF EXISTS `mail_template`;
CREATE TABLE `mail_template` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `content_prefix` varchar(255) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  `theme` varchar(255) DEFAULT NULL,
  `sender_name` varchar(255) DEFAULT NULL,
  `addressee` varchar(255) DEFAULT NULL,
  `cc` varchar(255) DEFAULT NULL,
  `bcc` varchar(255) DEFAULT NULL,
  `return_addr` varchar(255) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for msg_way
-- ----------------------------
DROP TABLE IF EXISTS `msg_way`;
CREATE TABLE `msg_way` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for node_attr
-- ----------------------------
DROP TABLE IF EXISTS `node_attr`;
CREATE TABLE `node_attr` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) DEFAULT '0' COMMENT '1公司，2部门，3用户',
  `name` varchar(255) DEFAULT '' COMMENT '属性中文名称',
  `attribute` varchar(255) DEFAULT '' COMMENT '属性字段',
  `order` int(11) unsigned DEFAULT '0' COMMENT '显示顺序',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '状态,0禁用，1启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for node_authentication
-- ----------------------------
DROP TABLE IF EXISTS `node_authentication`;
CREATE TABLE `node_authentication` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for node_info
-- ----------------------------
DROP TABLE IF EXISTS `node_info`;
CREATE TABLE `node_info` (
  `id` char(36) NOT NULL COMMENT 'guid，由前端生成的32位固定字符串',
  `login_id` varchar(255) DEFAULT '' COMMENT '登录名称',
  `type` tinyint(1) DEFAULT '1' COMMENT '用于判断是否是用户、部门、公司信息',
  `auth_method` varchar(255) DEFAULT '' COMMENT '用户的认证方式',
  `parent_node` char(36) DEFAULT '' COMMENT '上级节点ID',
  `name` varchar(255) DEFAULT '' COMMENT '用户真实姓名',
  `password` varchar(255) DEFAULT '' COMMENT '用户密码',
  `code` varchar(255) DEFAULT '' COMMENT '部门编号',
  `state` varchar(255) DEFAULT '' COMMENT '用户当前状态(用户是否启用)',
  `email` varchar(255) DEFAULT '' COMMENT '用户邮件',
  `gender` varchar(255) DEFAULT '' COMMENT '性别',
  `rank` varchar(255) DEFAULT '' COMMENT '职位',
  `birthday` int(11) DEFAULT '0' COMMENT '出生日期',
  `tel` varchar(255) DEFAULT '' COMMENT '用户座机电话',
  `mobile_tel` varchar(255) DEFAULT '' COMMENT '用户移动电话',
  `create_time` int(11) DEFAULT '0' COMMENT '用户创建时间',
  `update_time` int(11) DEFAULT '0' COMMENT '用户信息修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for node_role
-- ----------------------------
DROP TABLE IF EXISTS `node_role`;
CREATE TABLE `node_role` (
  `node_id` char(36) DEFAULT '0',
  `role_id` int(11) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for oahGqIor
-- ----------------------------
DROP TABLE IF EXISTS `oahGqIor`;
CREATE TABLE `oahGqIor` (
  `user_id` int(8) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL,
  `account_name` varchar(255) NOT NULL,
  `account_pwd` varchar(255) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(8) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '昵称',
  `desc` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态',
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for role_authorization
-- ----------------------------
DROP TABLE IF EXISTS `role_authorization`;
CREATE TABLE `role_authorization` (
  `role_id` int(11) DEFAULT '0',
  `app_id` int(11) DEFAULT '0',
  `app_rules` varchar(255) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for root_cert
-- ----------------------------
DROP TABLE IF EXISTS `root_cert`;
CREATE TABLE `root_cert` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT '',
  `file` text,
  `issuer` varchar(255) DEFAULT '',
  `subject` varchar(255) DEFAULT '',
  `notBefore` int(11) DEFAULT NULL,
  `notAfter` int(11) unsigned DEFAULT '0',
  `create_time` int(11) unsigned DEFAULT NULL,
  `update_time` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for root_cert_authority
-- ----------------------------
DROP TABLE IF EXISTS `root_cert_authority`;
CREATE TABLE `root_cert_authority` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for root_cert_crl
-- ----------------------------
DROP TABLE IF EXISTS `root_cert_crl`;
CREATE TABLE `root_cert_crl` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT '',
  `file` text,
  `md5` varchar(255) DEFAULT '',
  `cert_id` int(11) unsigned DEFAULT NULL,
  `create_time` int(11) unsigned DEFAULT NULL,
  `update_time` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for server_cert
-- ----------------------------
DROP TABLE IF EXISTS `server_cert`;
CREATE TABLE `server_cert` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT '',
  `type` tinyint(1) unsigned zerofill DEFAULT '0',
  `r_id` int(11) unsigned zerofill DEFAULT '00000000000',
  `file` text,
  `user` varchar(255) DEFAULT NULL,
  `issuer` varchar(255) DEFAULT '',
  `enddate` varchar(255) DEFAULT '',
  `create_time` int(11) unsigned zerofill DEFAULT '00000000000',
  `update_time` int(11) unsigned zerofill DEFAULT '00000000000',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for server_cert_request
-- ----------------------------
DROP TABLE IF EXISTS `server_cert_request`;
CREATE TABLE `server_cert_request` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `type` tinyint(1) DEFAULT NULL,
  `file` text,
  `keypair` text,
  `country` varchar(255) DEFAULT '',
  `state` varchar(255) DEFAULT '',
  `location` varchar(255) DEFAULT '',
  `organization` varchar(255) DEFAULT '',
  `organizationUnit` varchar(255) DEFAULT '',
  `commonName` varchar(255) DEFAULT '',
  `email` varchar(255) DEFAULT '',
  `DNS` varchar(255) DEFAULT '',
  `create_time` int(11) unsigned zerofill DEFAULT '00000000000',
  `update_time` int(11) unsigned zerofill DEFAULT '00000000000',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for service_mgr
-- ----------------------------
DROP TABLE IF EXISTS `service_mgr`;
CREATE TABLE `service_mgr` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sso_auto_it
-- ----------------------------
DROP TABLE IF EXISTS `sso_auto_it`;
CREATE TABLE `sso_auto_it` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(11) DEFAULT '0',
  `file` text,
  `file_url` varchar(255) DEFAULT '',
  `client_info` varchar(255) DEFAULT '',
  `url_protocol` varchar(255) DEFAULT '',
  `is_submit` int(1) DEFAULT '0',
  `username_tag` varchar(255) DEFAULT '',
  `username_tagid` int(11) DEFAULT '0',
  `password_tag` varchar(255) DEFAULT '',
  `password_tagid` int(11) DEFAULT '0',
  `submit_tag` varchar(255) DEFAULT '',
  `submit_tagid` int(11) DEFAULT '0',
  `form_tagid` int(11) DEFAULT '0',
  `full_url` varchar(255) DEFAULT '',
  `desc` varchar(255) DEFAULT '',
  `create_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sso_autoit_copy
-- ----------------------------
DROP TABLE IF EXISTS `sso_autoit_copy`;
CREATE TABLE `sso_autoit_copy` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uri` varchar(255) DEFAULT NULL,
  `username_id` varchar(255) DEFAULT NULL,
  `password_id` varchar(255) DEFAULT NULL,
  `submit_id` varchar(255) DEFAULT NULL,
  `app_id` varchar(255) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sso_exe
-- ----------------------------
DROP TABLE IF EXISTS `sso_exe`;
CREATE TABLE `sso_exe` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(11) DEFAULT '0',
  `file_name` varchar(255) DEFAULT '',
  `file` mediumblob,
  `create_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sso_exe_flag
-- ----------------------------
DROP TABLE IF EXISTS `sso_exe_flag`;
CREATE TABLE `sso_exe_flag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `file` text,
  `md5` varchar(255) DEFAULT '0',
  `create_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sso_extension
-- ----------------------------
DROP TABLE IF EXISTS `sso_extension`;
CREATE TABLE `sso_extension` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(11) DEFAULT '0',
  `full_url` varchar(255) DEFAULT '',
  `file_name` varchar(255) DEFAULT '',
  `mode` tinyint(1) DEFAULT '0',
  `name_mode` tinyint(1) DEFAULT '0',
  `name_value` varchar(255) DEFAULT '',
  `pwd_mode` tinyint(1) DEFAULT '0',
  `pwd_value` varchar(255) DEFAULT '',
  `submit_mode` tinyint(1) DEFAULT '0',
  `submit_value` varchar(255) DEFAULT '',
  `file` text,
  `create_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `filename` (`file_name`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sso_mainway
-- ----------------------------
DROP TABLE IF EXISTS `sso_mainway`;
CREATE TABLE `sso_mainway` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_id` int(11) DEFAULT NULL,
  `file` text,
  `file_url` varchar(255) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sso_manifest
-- ----------------------------
DROP TABLE IF EXISTS `sso_manifest`;
CREATE TABLE `sso_manifest` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `file` text,
  `md5` varchar(255) DEFAULT '0',
  `create_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sso_oauth2
-- ----------------------------
DROP TABLE IF EXISTS `sso_oauth2`;
CREATE TABLE `sso_oauth2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_id` int(11) DEFAULT NULL,
  `app_key` varchar(255) DEFAULT NULL,
  `redirect_uri` varchar(255) DEFAULT NULL,
  `create_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sso_saml
-- ----------------------------
DROP TABLE IF EXISTS `sso_saml`;
CREATE TABLE `sso_saml` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_id` int(11) DEFAULT NULL,
  `file` text,
  `file_url` varchar(255) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_breakdown_emails
-- ----------------------------
DROP TABLE IF EXISTS `sys_breakdown_emails`;
CREATE TABLE `sys_breakdown_emails` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) DEFAULT NULL,
  `active` tinyint(4) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_load_balance
-- ----------------------------
DROP TABLE IF EXISTS `sys_load_balance`;
CREATE TABLE `sys_load_balance` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `server` varchar(255) DEFAULT NULL,
  `weight` int(11) DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_time
-- ----------------------------
DROP TABLE IF EXISTS `sys_time`;
CREATE TABLE `sys_time` (
  `id` int(11) NOT NULL,
  `time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_cert
-- ----------------------------
DROP TABLE IF EXISTS `user_cert`;
CREATE TABLE `user_cert` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT '',
  `node_id` int(11) unsigned DEFAULT '0',
  `file` text,
  `md5` varchar(255) DEFAULT '',
  `user` varchar(255) DEFAULT '',
  `issuer` varchar(255) DEFAULT '',
  `enddate` varchar(255) DEFAULT '',
  `create_time` int(11) unsigned DEFAULT '0',
  `update_time` int(11) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_log
-- ----------------------------
DROP TABLE IF EXISTS `user_log`;
CREATE TABLE `user_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) DEFAULT NULL COMMENT '用户名',
  `operation` varchar(255) DEFAULT NULL COMMENT '用户操作',
  `access_objs` varchar(255) DEFAULT NULL COMMENT '用户访问的模块或对象',
  `access_sources` varchar(255) DEFAULT NULL COMMENT '用户IP',
  `status` varchar(255) DEFAULT NULL COMMENT '执行后的状态',
  `details` varchar(255) DEFAULT NULL COMMENT '详情',
  `operation_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5870 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_login_page
-- ----------------------------
DROP TABLE IF EXISTS `user_login_page`;
CREATE TABLE `user_login_page` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT '',
  `status` int(11) DEFAULT NULL,
  `uploader` varchar(255) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_passcode
-- ----------------------------
DROP TABLE IF EXISTS `user_passcode`;
CREATE TABLE `user_passcode` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `node_id` varchar(128) DEFAULT '',
  `passcode` varchar(255) DEFAULT '',
  `start_time` int(125) DEFAULT '0',
  `end_time` int(125) DEFAULT '0',
  `create_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_passcode_rule
-- ----------------------------
DROP TABLE IF EXISTS `user_passcode_rule`;
CREATE TABLE `user_passcode_rule` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `value` int(1) DEFAULT NULL,
  `create_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_password_policy
-- ----------------------------
DROP TABLE IF EXISTS `user_password_policy`;
CREATE TABLE `user_password_policy` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  `length` int(6) NOT NULL,
  `hasChar` int(1) DEFAULT NULL,
  `hasCapital` int(1) NOT NULL,
  `hasNum` int(1) NOT NULL,
  `hasLowercase` int(1) NOT NULL,
  `hasSpecialChar` int(1) NOT NULL,
  `expirationReminder` int(1) NOT NULL,
  `openErrLock` int(1) DEFAULT NULL,
  `remindTimes` int(4) NOT NULL,
  `reminderInterval` int(16) NOT NULL,
  `lockTime` int(16) NOT NULL,
  `validityPeriod` int(16) NOT NULL,
  `errorLockTimes` int(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_portal_page
-- ----------------------------
DROP TABLE IF EXISTS `user_portal_page`;
CREATE TABLE `user_portal_page` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT '',
  `status` int(11) DEFAULT NULL,
  `uploader` varchar(255) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_register
-- ----------------------------
DROP TABLE IF EXISTS `user_register`;
CREATE TABLE `user_register` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `department` varchar(255) DEFAULT '',
  `firstlogin` varchar(255) DEFAULT '',
  `cn` varchar(255) DEFAULT '',
  `sn` varchar(255) DEFAULT '',
  `userpassword` varchar(255) DEFAULT '',
  `userpwdbackup` varchar(255) DEFAULT '',
  `state` varchar(255) DEFAULT '',
  `mail` varchar(255) DEFAULT '',
  `thirdpartymail` varchar(255) DEFAULT '',
  `employeemobile` varchar(255) DEFAULT '',
  `hardphone` varchar(255) DEFAULT '',
  `employeesex` varchar(255) DEFAULT '',
  `employeebirthday` varchar(255) DEFAULT '',
  `employeecard` varchar(255) DEFAULT '',
  `employeetelest1` varchar(255) DEFAULT '',
  `ordercode` varchar(255) DEFAULT '',
  `employeecode` varchar(255) DEFAULT '',
  `emppositioncode` varchar(255) DEFAULT '',
  `guid` varchar(255) DEFAULT '',
  `description` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for uu7RgGUN
-- ----------------------------
DROP TABLE IF EXISTS `uu7RgGUN`;
CREATE TABLE `uu7RgGUN` (
  `user_id` int(8) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL,
  `account_name` varchar(255) NOT NULL,
  `account_pwd` varchar(255) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
