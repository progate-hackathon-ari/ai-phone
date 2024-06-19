CREATE TABLE `connected_player` (
  `connection_id` VARCHAR(255) PRIMARY KEY,
  `room_id` VARCHAR(255) NOT NULL,
  `index` INT NOT NULL,
  `username` VARCHAR(255) NOT NULL,
  FOREIGN KEY (`room_id`) REFERENCES `rooms` (`room_id`)
);
