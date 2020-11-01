CREATE DATABASE
IF
	NOT EXISTS blog_service DEFAULT CHARACTER
	SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

use blog_service;
-- 标签表

DROP table if exists `blog_tag`;
CREATE TABLE `blog_tag` (
	`id` int(10) UNSIGNED not NULL AUTO_INCREMENT,
	`created_on` int(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
	`created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
	`modified_on` int(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
	`modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
	`delete_on` int(10) UNSIGNED DEFAULT '0' COMMENT '删除时间',
	`is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否删除 0-未删除 1-已删除',
	`name` VARCHAR(100) DEFAULT '' COMMENT '标签名称',
	`state` TINYINT(3) UNSIGNED DEFAULT '1' COMMENT '状态 0-禁用 1-启动',
	PRIMARY KEY (`id`)
)ENGINE = INNODB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';


-- 文章表
DROP table if exists `blog_article`;
CREATE TABLE `blog_article` (
	`id` int(10) UNSIGNED not NULL AUTO_INCREMENT,
	`title` VARCHAR(100) DEFAULT '' COMMENT '文章标题',
	`desc` VARCHAR(255) DEFAULT '' COMMENT '文章简述',
	`cover_image_url` VARCHAR(255) DEFAULT '' COMMENT '封面图片地址',
	`content` LONGTEXT COMMENT '文章内容',
	`created_on` int(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
	`created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
	`modified_on` int(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
	`modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
	`delete_on` int(10) UNSIGNED DEFAULT '0' COMMENT '删除时间',
	`is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否删除 0-未删除 1-已删除',
	`name` VARCHAR(100) DEFAULT '' COMMENT '标签名称',
	`state` TINYINT(3) UNSIGNED DEFAULT '1' COMMENT '状态 0-禁用 1-启动',
	PRIMARY KEY (`id`)
)ENGINE = INNODB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';


-- 文章标签关联表
DROP table if exists `blog_article_tag`;
CREATE TABLE `blog_article_tag` (
	`id` int(10) UNSIGNED not NULL AUTO_INCREMENT,
	`article_id` int(10) UNSIGNED not NULL COMMENT '文章id',
	`tag_id` int(10) UNSIGNED not NULL COMMENT '标签id',
	`created_on` int(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
	`created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
	`modified_on` int(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
	`modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
	`delete_on` int(10) UNSIGNED DEFAULT '0' COMMENT '删除时间',
	`is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否删除 0-未删除 1-已删除',
	`name` VARCHAR(100) DEFAULT '' COMMENT '标签名称',
	`state` TINYINT(3) UNSIGNED DEFAULT '1' COMMENT '状态 0-禁用 1-启动',
	PRIMARY KEY (`id`)
)ENGINE = INNODB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联表';

