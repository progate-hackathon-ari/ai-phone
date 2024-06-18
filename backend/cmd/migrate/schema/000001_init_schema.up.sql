CREATE TABLE `rooms` (
  `room_id` varchar(255) PRIMARY KEY,
  `extra_prompt` text,
  `is_started` TINYINT(1) NOT NULL,
  `game_size` int NOT NULL,
  `current_game` int NOT NULL
);
