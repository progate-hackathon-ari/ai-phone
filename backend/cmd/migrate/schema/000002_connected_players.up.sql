CREATE TABLE `connected_players` (
  `room_id` varchar(255) NOT NULL,
  `connection_id` varchar(255) NOT NULL,
  `player_index` int NOT NULL,
  `username` varchar(255) NOT NULL,
  CONSTRAINT `fk_room_id` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`room_id`)
);