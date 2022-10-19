DROP DATABASE IF EXISTS `ibank`;
CREATE DATABASE `ibank` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `ibank`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `birthdate` timestamp NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `document` varchar(11) NOT NULL,
  `city` varchar(255) NOT NULL,
  `zipcode` varchar(8) NOT NULL,
  `phone` varchar(16) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`)
) ENGINE=INNODB;

INSERT INTO `user` VALUES 
	(1, 'Ramon Ferreira', NOW(), '123', 'rfnascimento@ibm.com', '12196183067', 'Brasilia', '70000000', '5561999991111', 1);

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11),
  `opening_date` timestamp NOT NULL DEFAULT NOW(),
  `agency` varchar(4) NOT NULL,
  `number` varchar(8) NOT NULL,
  `check_digit` varchar(1) NOT NULL,
  `pin` varchar(8) NOT NULL,
  `balance` decimal(15,2) NOT NULL DEFAULT '0.00',
  `status` tinyint(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=INNODB;

INSERT INTO `account` VALUES 
	(1, 1, NOW(), '0001', '10001234', '7', '00000000', 0.00, 1);

DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bank_id` int(11) NOT NULL,
  `agency` varchar(4) NOT NULL,
  `number` varchar(8) NOT NULL,
  `check_digit` varchar(1) NOT NULL,
  `type` varchar(16) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `value` decimal(15,2) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=INNODB;