CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255) UNIQUE NOT NULL,
  `name` varchar(255) NOT NULL,
  `bio` varchar(255),
  `password` varchar(255) NOT NULL,
  `image` varchar(255),
  `imageName` varchar(255),
  `provider` ENUM ('EMAIL', 'GOOGLE') NOT NULL DEFAULT "EMAIL"
);

CREATE TABLE `projects` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `published` bool NOT NULL DEFAULT true,
  `image` varchar(255),
  `imageName` varchar(255),
  `estimatedDuration` int NOT NULL,
  `startDate` datetime,
  `endDate` datetime,
  `status` varchar(255) NOT NULL DEFAULT "OPEN",
  `user_id` int NOT NULL
);

CREATE TABLE `requests` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `project_id` int NOT NULL,
  `user_id` int NOT NULL,
  `status` varchar(255) NOT NULL DEFAULT "PENDING"
);

CREATE TABLE `project_users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `project_id` int NOT NULL,
  `user_id` int NOT NULL
);

CREATE TABLE `notifications` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `description` varchar(255) NOT NULL,
  `initiator_id` int NOT NULL,
  `listener_id` int NOT NULL
);

CREATE TABLE `chats` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `project_id` int NOT NULL,
  `user_id` int NOT NULL,
  `message` text NOT NULL
);

ALTER TABLE `projects` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `requests` ADD FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);

ALTER TABLE `requests` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `project_users` ADD FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);

ALTER TABLE `project_users` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `notifications` ADD FOREIGN KEY (`initiator_id`) REFERENCES `users` (`id`);

ALTER TABLE `notifications` ADD FOREIGN KEY (`listener_id`) REFERENCES `users` (`id`);

ALTER TABLE `chats` ADD FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);

ALTER TABLE `chats` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
