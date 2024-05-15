/*
 Navicat Premium Data Transfer

 Source Server         : mysql8
 Source Server Type    : MySQL
 Source Server Version : 80300 (8.3.0)
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 80300 (8.3.0)
 File Encoding         : 65001

 Date: 15/05/2024 00:39:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_profiles
-- ----------------------------
DROP TABLE IF EXISTS `user_profile`;
CREATE TABLE `user_profile` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `gender` tinyint(1) NOT NULL COMMENT '性别：0保密 1 男 2女 3未知',
  `level` tinyint(1) NOT NULL DEFAULT '0' COMMENT '等级',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `salt` char(30) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '盐',
  `birthday` date DEFAULT NULL COMMENT '出生日期',
  `signature` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '个性签名',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id` DESC),
  UNIQUE KEY `idx_user` (`user_id`) USING BTREE COMMENT '用户id索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '?',
  `phone` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `wechat_unionid` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '微信unionid',
  `wechat_applet_openid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '微信小程序id',
  `wechat_official_openid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '微信公众号id',
  `name` varchar(15) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `nickname` varchar(30) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id` DESC) USING BTREE,
  UNIQUE KEY `idx_phone` (`phone` DESC) USING BTREE COMMENT '手机号索引',
  UNIQUE KEY `idx_email` (`email` DESC) USING BTREE COMMENT '邮箱索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
