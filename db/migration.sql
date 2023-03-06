CREATE USER IF NOT EXISTS 'root'@'%' IDENTIFIED BY 'snmsung1';
CREATE USER IF NOT EXISTS 'root'@'localhost' IDENTIFIED BY 'snmsung1';

GRANT ALL PRIVILEGES ON *.* TO 'root'@'%';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost';

FLUSH PRIVILEGES;

USE writedream;

CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(30) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `title_UNIQUE` (`title`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `issues` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(50) NOT NULL,
  `content` varchar(255) NOT NULL,
  `view_count` int DEFAULT '0',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `memos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `issue_id` int NOT NULL,
  `text` varchar(255) NOT NULL,
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `issue_id` (`issue_id`),
  CONSTRAINT `memos_ibfk_1` FOREIGN KEY (`issue_id`) REFERENCES `issues` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `issue_category` (
  `issue_id` int NOT NULL,
  `category_id` int NOT NULL,
  PRIMARY KEY (`issue_id`,`category_id`),
  KEY `idx_issue` (`category_id`),
  KEY `idx_category` (`category_id`),
  CONSTRAINT `iss_cate_ibfk_1` FOREIGN KEY (`issue_id`) REFERENCES `issues` (`id`) ON DELETE CASCADE,
  CONSTRAINT `iss_cate_ibfk_2` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO writedream.categories (`title`) VALUES ('Default');