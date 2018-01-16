--
CREATE DATABASE IF NOT EXISTS `travel`;
USE `travel`;
-- ----------------------------------------------

CREATE TABLE IF NOT EXISTS `user` (
  `id`    int(16) unsigned NOT NULL AUTO_INCREMENT,
  `name`  char(32) NOT NULL,
  `pass`  char(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ----------------------------------------------
CREATE TABLE IF NOT EXISTS `article` (
  `id`      INT(16) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title`   CHAR(32) UNIQUE NOT NULL,
  `content` TEXT NOT NULL,
  `views`   INT(32) UNSIGNED DEFAULT 0,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status`  BOOL  NOT NULL DEFAULT 1
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;