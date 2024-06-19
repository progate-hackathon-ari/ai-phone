CREATE TABLE `in_game_prompts` (
  `room_id` VARCHAR(255) NOT NULL,
  `connection_id` VARCHAR(255) NOT NULL,
  `game_index` INT NOT NULL,
  `prompt` TEXT NOT NULL,
  FOREIGN KEY (`room_id`) REFERENCES `rooms` (`room_id`),
  FOREIGN KEY (`connection_id`) REFERENCES `connected_player` (`connection_id`)
);