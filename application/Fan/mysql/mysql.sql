--
CREATE DATABASE IF NOT EXISTS `Fan`;
USE `Fan`;
-- ----------------------------------------------

CREATE TABLE IF NOT EXISTS `article` (
  `id`      int(16) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title`   char(32) NOT NULL,
  `auther`  char(16),
  `class`   char(16),
  `content` text NOT NULL,
  `picture` char,
  `views`   INT unsigned DEFAULT 0,
  `datetime`  DATETIME
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------------------------

CREATE TABLE IF NOT EXISTS `admin`  (
  `id`  INT(16) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name`  CHAR(16)  NOT NULL,
  `password`  CHAR(16) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;