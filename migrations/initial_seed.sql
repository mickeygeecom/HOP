-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               11.3.2-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             12.6.0.6765
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for quiz
CREATE DATABASE IF NOT EXISTS `quiz` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */;
USE `quiz`;

-- Dumping structure for table quiz.questions
CREATE TABLE IF NOT EXISTS `questions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `quiz_id` int(11) DEFAULT NULL,
  `question_text` text DEFAULT NULL,
  `option1` varchar(255) NOT NULL,
  `option2` varchar(255) NOT NULL,
  `option3` varchar(255) NOT NULL,
  `option4` varchar(255) NOT NULL,
  `correct_option` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `quiz_id` (`quiz_id`),
  CONSTRAINT `questions_ibfk_1` FOREIGN KEY (`quiz_id`) REFERENCES `quiz` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=87 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table quiz.questions: ~24 rows (approximately)
INSERT INTO `questions` (`id`, `quiz_id`, `question_text`, `option1`, `option2`, `option3`, `option4`, `correct_option`) VALUES
	(7, 13, 'Who is the king of Odense', 'gEE', 'Toober', 'Faustiman', 'Everyone', 1),
	(10, 1, 'Which number is the highest?', '1', '2', '3', '4', 4),
	(11, 1, 'Who has the most swag in LootLocker?', 'Erla (right one)', 'Erla', 'Erla', 'Erla', 1),
	(66, 15, 'What is the chemical symbol for gold?', 'Au', 'Ag', 'Fe', 'Cu', 1),
	(67, 15, 'What is the closest planet to the Sun?', 'Earth', 'Mars', 'Venus', 'Mercury', 4),
	(68, 15, 'Who discovered penicillin?', 'Marie Curie', 'Alexander Fleming', 'Louis Pasteur', 'Albert Einstein', 2),
	(69, 15, 'What is the powerhouse of the cell?', 'Nucleus', 'Mitochondria', 'Chloroplast', 'Ribosome', 2),
	(70, 15, 'What is the atomic number of oxygen?', '6', '7', '8', '9', 3),
	(71, 15, 'Who proposed the theory of relativity?', 'Isaac Newton', 'Albert Einstein', 'Galileo Galilei', 'Stephen Hawking', 2),
	(72, 15, 'Which element is the most abundant in the Earth’s crust?', 'Oxygen', 'Silicon', 'Aluminum', 'Carbon', 2),
	(73, 15, 'What is the chemical formula for water?', 'H2O', 'CO2', 'O2', 'NaCl', 1),
	(74, 15, 'What is the largest organ in the human body?', 'Heart', 'Brain', 'Skin', 'Liver', 3),
	(75, 15, 'What is the SI unit of electric current?', 'Volt', 'Ampere', 'Ohm', 'Watt', 2),
	(76, 16, 'In which year did World War I begin?', '1914', '1918', '1939', '1945', 1),
	(77, 16, 'Who was the first President of the United States?', 'George Washington', 'Thomas Jefferson', 'Abraham Lincoln', 'John Adams', 1),
	(78, 16, 'Where was the Magna Carta signed?', 'France', 'Germany', 'Italy', 'England', 4),
	(79, 16, 'Who wrote "The Communist Manifesto"?', 'Karl Marx', 'Friedrich Engels', 'Vladimir Lenin', 'Leon Trotsky', 1),
	(80, 16, 'Which ancient civilization built the Great Pyramid of Giza?', 'Greek', 'Roman', 'Egyptian', 'Mayan', 3),
	(81, 16, 'Who painted the Mona Lisa?', 'Pablo Picasso', 'Vincent van Gogh', 'Leonardo da Vinci', 'Michelangelo', 3),
	(82, 16, 'Which battle marked the end of Napoleon’s rule as Emperor of France?', 'Battle of Waterloo', 'Battle of Austerlitz', 'Battle of Trafalgar', 'Battle of Leipzig', 1),
	(83, 16, 'Who was the first female Prime Minister of the United Kingdom?', 'Theresa May', 'Margaret Thatcher', 'Angela Merkel', 'Jacinda Ardern', 2),
	(84, 16, 'Which US President signed the Emancipation Proclamation?', 'Abraham Lincoln', 'Thomas Jefferson', 'George Washington', 'Ulysses S. Grant', 1),
	(85, 16, 'In which year did the Berlin Wall fall?', '1989', '1991', '1990', '1987', 1),
	(86, 1, 'When Andreas is working, how often does he wear a cap?', 'Never', 'Rarely', 'When he eats', 'Always', 4);

-- Dumping structure for table quiz.quiz
CREATE TABLE IF NOT EXISTS `quiz` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text DEFAULT NULL,
  `join_code` varchar(6) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `join_code` (`join_code`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table quiz.quiz: ~6 rows (approximately)
INSERT INTO `quiz` (`id`, `name`, `description`, `join_code`, `created_at`) VALUES
	(1, 'LootLocker', 'A LootLocker test quiz containing weird a funny questions', 'PY2SW8', '2024-04-15 09:52:06'),
	(13, '🇩🇰', 'All the daninsh Q\'s', 'BHD0NZ', '2024-04-29 14:07:19'),
	(15, 'Science Quiz', 'Test your knowledge about science', 'RODQ7L', '2024-04-30 08:00:00'),
	(16, 'History Quiz', 'Explore historical events and figures', 'J9RX2V', '2024-04-30 09:00:00'),
	(18, 'Empty Quiz', 'For testing purpose', '9WV0UO', '2024-05-02 18:48:39'),
	(26, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.', 'O8EBZ5', '2024-05-21 13:39:50');

-- Dumping structure for table quiz.quiz_sessions
CREATE TABLE IF NOT EXISTS `quiz_sessions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `quiz_id` int(11) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `quiz_id` (`quiz_id`),
  CONSTRAINT `quiz_sessions_ibfk_1` FOREIGN KEY (`quiz_id`) REFERENCES `quiz` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table quiz.quiz_sessions: ~0 rows (approximately)

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
