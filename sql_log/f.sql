USE `test`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(63) CHARACTER SET utf8 NOT NULL,
  `email` varchar(127) CHARACTER SET utf8 NOT NULL,
  `password` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `first_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `phone` varchar(31) CHARACTER SET utf8 DEFAULT NULL,
  `birthday` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `creator_id` int(10) unsigned DEFAULT '0',
  `class` varchar(63) COLLATE utf8_unicode_ci DEFAULT NULL,
  `school` varchar(63) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name_UNIQUE` (`user_name`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


CREATE TABLE `news` (
  `id` INT(11) UNSIGNED NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `pretty_url` VARCHAR(255) NOT NULL,
  `content` TEXT NULL,
  `creator_id` INT(11) UNSIGNED NOT NULL DEFAULT 0,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `pretty_url_UNIQUE` (`pretty_url` ASC));

CREATE TABLE `news_category` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(45) NULL,
  `prettyurl` VARCHAR(45) NULL,
  `description` VARCHAR(2000) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `prettyurl_UNIQUE` (`prettyurl` ASC));

CREATE TABLE `document` (
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(127) NOT NULL,
  `path_store` VARCHAR(255) NOT NULL,
  `class_id` INT(11) UNSIGNED NULL,
  `subject_id` INT(11) UNSIGNED NULL,
  `creator_id` INT(11) UNSIGNED NOT NULL DEFAULT 0,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `question` (
  `id` INT(11) UNSIGNED NOT NULL,
  `title` VARCHAR(255) NULL,
  `content` TEXT NULL,
  `type_id` INT(11) UNSIGNED NULL,
  `class_id` INT(11) UNSIGNED NULL,
  `difficult_id` INT(11) UNSIGNED NULL,
  `subject_id` INT(11) UNSIGNED NULL,
  `creator_id` INT(11) UNSIGNED NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  `answer` TEXT NULL,
  PRIMARY KEY (`id`));



