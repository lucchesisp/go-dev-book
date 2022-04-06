-- CREATE DATABASE AND TABLES
CREATE DATABASE IF NOT EXISTS `devbook`;
USE `devbook`;

DROP TABLE IF EXISTS `publications`;
DROP TABLE IF EXISTS `followers`;
DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int auto_increment primary key,
  `name` varchar(255) NOT NULL,
  `nickname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `followers` (
  PRIMARY KEY (`user_id`, `follower_id`),
  `user_id` int NOT NULL, FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  `follower_id` int NOT NULL, FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `publications` (
  `id` int auto_increment primary key,
  `content` text NOT NULL,
  `author_id` int NOT NULL, FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE,
  `like_count` int NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- INSERT DEV DATA

INSERT INTO `users` (`name`, `nickname`, `email`, `password`) VALUES
('John Doe 1', 'johndoe1', 'johndoe1@gmail.com', '$2a$10$xc4P8JnOAReL4qNAyTmGOuR.5Dr9.4bQxKhpp/dH/GxpVlNOnw9Ka'),
('John Doe 2', 'johndoe2', 'johndoe2@gmail.com', '$2a$10$xc4P8JnOAReL4qNAyTmGOuR.5Dr9.4bQxKhpp/dH/GxpVlNOnw9Ka'),
('John Doe 3', 'johndoe3', 'johndoe3@gmail.com', '$2a$10$xc4P8JnOAReL4qNAyTmGOuR.5Dr9.4bQxKhpp/dH/GxpVlNOnw9Ka'),
('John Doe 4', 'johndoe4', 'johndoe4@gmail.com', '$2a$10$xc4P8JnOAReL4qNAyTmGOuR.5Dr9.4bQxKhpp/dH/GxpVlNOnw9Ka');

INSERT INTO `followers` (`user_id`, `follower_id`) VALUES
(1, 2),
(3, 1),
(1, 3);

INSERT INTO `publications` (`content`, `author_id`) VALUES
('Publication User 1', 1),
('Publication User 2', 2),
('Publication User 3', 3),
('Publication User 4', 4);