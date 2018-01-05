--
CREATE DATABASE IF NOT EXISTS `travel`;
USE `travel`;
-- ----------------------------------------------

CREATE TABLE IF NOT EXISTS `user` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;