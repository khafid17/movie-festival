-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.4.11-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             12.3.0.6589
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for movie_festival
CREATE DATABASE IF NOT EXISTS `movie_festival` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `movie_festival`;

-- Dumping structure for table movie_festival.movies
CREATE TABLE IF NOT EXISTS `movies` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `description` text DEFAULT NULL,
  `duration` int(11) NOT NULL,
  `artists` varchar(255) NOT NULL,
  `genres` varchar(255) NOT NULL,
  `watch_url` varchar(255) NOT NULL,
  `views` int(11) DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table movie_festival.movies: ~2 rows (approximately)
INSERT INTO `movies` (`id`, `title`, `description`, `duration`, `artists`, `genres`, `watch_url`, `views`, `created_at`, `updated_at`) VALUES
	(1, 'Harry Poter', 'Updated description', 160, 'Piter', 'Action, Fantasy', 'http://example.com/watch/123456', 5, '2024-12-28 16:49:01', '2024-12-28 20:12:37'),
	(2, 'Agak Lain', 'Deskripsi film', 150, 'Ernes', 'Komedi', 'http://example.com/film-url-new', 1, '2024-12-28 17:03:10', '2024-12-28 19:43:51');

-- Dumping structure for table movie_festival.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table movie_festival.users: ~2 rows (approximately)
INSERT INTO `users` (`id`, `username`, `password`) VALUES
	(1, 'exampleUsername', 'examplePassword'),
	(2, 'khafid', 'khafid');

-- Dumping structure for table movie_festival.votes
CREATE TABLE IF NOT EXISTS `votes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `movie_id` int(11) NOT NULL,
  `voted` tinyint(1) DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table movie_festival.votes: ~4 rows (approximately)
INSERT INTO `votes` (`id`, `user_id`, `movie_id`, `voted`, `created_at`) VALUES
	(2, 0, 0, 1, '2024-12-28 19:27:58'),
	(3, 0, 0, 1, '2024-12-28 20:07:56'),
	(4, 0, 0, 1, '2024-12-28 20:08:16'),
	(5, 2, 2, 1, '2024-12-28 20:08:49');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
