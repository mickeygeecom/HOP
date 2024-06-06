-- --------------------------------------------------------
-- Host:                         box.mickeygee.com
-- Server version:               11.3.2-MariaDB-1:11.3.2+maria~ubu2004 - mariadb.org binary distribution
-- Server OS:                    debian-linux-gnu
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
) ENGINE=InnoDB AUTO_INCREMENT=783 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table quiz.questions: ~97 rows (approximately)
INSERT INTO `questions` (`id`, `quiz_id`, `question_text`, `option1`, `option2`, `option3`, `option4`, `correct_option`) VALUES
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
	(86, 1, 'When Andreas is working, how often does he wear a cap?', 'Never', 'Rarely', 'When he eats', 'Always', 4),
	(516, 77, 'What is the capital city of Denmark?', 'Stockholm', 'Copenhagen', 'Helsinki', 'Oslo', 2),
	(517, 77, 'Which Danish physicist won the Nobel Prize in Physics in 1922 for his work on atomic structure and quantum theory?', 'Hans Christian Ørsted', 'Anders Celsius', 'Tycho Brahe', 'Niels Bohr', 4),
	(518, 77, 'Which Danish author is known for writing fairy tales such as \'The Little Mermaid\' and \'The Emperor\'s New Clothes\'?', 'Søren Kierkegaard', 'Karen Blixen', 'Hans Christian Andersen', 'Victor Borge', 3),
	(519, 77, 'In which year did Denmark become a member of the European Union?', '1973', '1986', '2000', '1992', 1),
	(520, 77, 'What is the highest point in Denmark, reaching only 170.86 meters above sea level?', 'Ejer Bavnehøj', 'Møllehøj', 'Yding Skovhøj', 'Kongeåen', 2),
	(521, 77, 'Which Danish architect is known for designing the Sydney Opera House in Australia?', 'Jørn Utzon', 'Bjarke Ingels', 'Arne Jacobsen', 'Henning Larsen', 1),
	(522, 77, 'What is the currency used in Denmark?', 'Euro', 'Krona', 'Danish Mark', 'Danish Krone', 4),
	(523, 77, 'Which traditional Danish dish consists of pickled herring served in a sour cream sauce with onions, apples, and spices?', 'Stegt flæsk med persillesovs', 'Sild i eddike', 'Frikadeller', 'Smørrebrød', 2),
	(524, 77, 'Which Danish director and screenwriter won an Academy Award for Best Foreign Language Film for the movie \'In a Better World\' in 2010?', 'Susanne Bier', 'Thomas Vinterberg', 'Nicolas Winding Refn', 'Lars von Trier', 1),
	(525, 77, 'What is the name of the bridge that connects Denmark and Sweden?', 'Great Belt Bridge', 'Kapellbrücke', 'Millau Viaduct', 'Øresund Bridge', 4),
	(526, 77, 'Which Danish footballer, known for his time at Tottenham Hotspur and Real Madrid, is the all-time leading goal scorer for the Denmark national team?', 'Michael Laudrup', 'Jon Dahl Tomasson', 'Christian Eriksen', 'Nicklas Bendtner', 4),
	(527, 77, 'Which Danish company is known for its colorful building bricks that are popular with children worldwide?', 'LEGO', 'Bang & Olufsen', 'Maersk', 'IKEA', 1),
	(528, 77, 'Which Danish astronomer, known for his accurate astronomical observations, constructed the Uraniborg castle and Tycho Brahe Observatory on the island of Ven?', 'Ole Rømer', 'Anders Celsius', 'Niels Bohr', 'Tycho Brahe', 4),
	(529, 77, 'What major sporting event was held in Denmark in 1992, where the Danish football team won against all odds?', 'FIFA World Cup', 'UEFA Champions League', 'Euro 92', 'Olympic Games', 3),
	(530, 77, 'Which Danish actor starred as Hannibal Lecter in the movie \'The Silence of the Lambs\' and won an Academy Award for Best Actor?', 'Viggo Mortensen', 'Mads Mikkelsen', 'Jesper Christensen', 'Ulrich Thomsen', 2),
	(561, 80, 'What is the capital city of Sweden?', 'Helsinki', 'Copenhagen', 'Oslo', 'Stockholm', 4),
	(562, 80, 'Which of the following is not a famous Swedish furniture company?', 'Husqvarna', 'Volvo', 'H&M', 'IKEA', 2),
	(563, 80, 'What is the currency used in Sweden?', 'Euro', 'Dollar', 'Krona', 'Yen', 3),
	(564, 80, 'The Nobel Prize ceremony takes place in which city of Sweden?', 'Malmö', 'Gothenburg', 'Stockholm', 'Uppsala', 3),
	(565, 80, 'Which Swedish band is famous for hits like \'Dancing Queen\'?', 'Coldplay', 'Nirvana', 'The Beatles', 'ABBA', 4),
	(566, 80, 'What natural phenomenon is commonly observed in the northern parts of Sweden?', 'Hurricanes', 'Tornadoes', 'Desertification', 'Aurora Borealis', 4),
	(567, 80, 'Which Swedish inventor is known for patenting the dynamite?', 'Alfred Nobel', 'Nikola Tesla', 'Thomas Edison', 'Marie Curie', 1),
	(568, 80, 'What is the national dish of Sweden?', 'Köttbullar', 'Raggmunk', 'Surströmming', 'Smörgåsbord', 1),
	(569, 80, 'Which Swedish actress starred in \'The Girl with the Dragon Tattoo\'?', 'Scarlett Johansson', 'Noomi Rapace', 'Alicia Vikander', 'Mia Wasikowska', 2),
	(570, 80, 'What Swedish car manufacturer is known for its safety features?', 'Toyota', 'BMW', 'Volvo', 'Volkswagen', 3),
	(571, 80, 'Which Swedish scientist proposed the three laws of planetary motion?', 'Johannes Kepler', 'Carl Linnaeus', 'Anders Celsius', 'Tycho Brahe', 1),
	(572, 80, 'In which year did Sweden officially abolish its monarchy and become a republic?', '1920', '1895', '1905', '1945', 3),
	(573, 80, 'Which Swedish progressive metal band is known for their album \'Blackwater Park\'?', 'Meshuggah', 'In Flames', 'Dark Tranquillity', 'Opeth', 4),
	(574, 80, 'What is the name of the Swedish traditional Midsummer celebration?', 'Valborgsmässoafton', 'Midsommar', 'Luciafest', 'Walpurgis Night', 2),
	(575, 80, 'Which Swedish director won an Academy Award for Best Foreign Language Film in 1983?', 'Roy Andersson', 'Lukas Moodysson', 'Ruben Östlund', 'Ingmar Bergman', 4),
	(710, 86, 'What is 5 + 3?', '8', '10', '2', '7', 1),
	(711, 86, 'What is 12 - 4?', '10', '6', '7', '8', 4),
	(712, 86, 'What is 9 + 6?', '18', '15', '13', '21', 2),
	(713, 86, 'What is 20 - 9?', '13', '9', '11', '7', 3),
	(714, 86, 'What is 14 + 3?', '20', '17', '15', '22', 2),
	(715, 86, 'What is 25 - 11?', '16', '18', '12', '14', 4),
	(716, 86, 'What is 8 + 7?', '17', '21', '12', '15', 4),
	(717, 86, 'What is 18 - 5?', '14', '11', '12', '13', 4),
	(718, 86, 'What is 23 + 9?', '36', '32', '34', '30', 2),
	(719, 86, 'What is 33 - 15?', '14', '18', '20', '16', 2),
	(720, 86, 'What is 11 + 12?', '27', '23', '25', '21', 2),
	(721, 86, 'What is 30 - 8?', '22', '20', '24', '19', 1),
	(722, 86, 'What is 6 + 9?', '17', '12', '15', '14', 3),
	(723, 86, 'What is 17 - 6?', '10', '11', '8', '9', 2),
	(724, 86, 'What is 22 + 14?', '34', '38', '32', '36', 4),
	(725, 87, 'Which type of flowers do women commonly enjoy receiving?', 'Tulips', 'Daisies', 'Sunflowers', 'Roses', 4),
	(726, 87, 'What is a popular choice for a romantic dinner setting?', 'Cooking together at home', 'Fast food restaurant', 'Candlelit dinner at a fancy restaurant', 'Picnic in the park', 3),
	(727, 87, 'Which type of jewelry is often considered a perfect gift for women?', 'Necklace', 'Anklet', 'Earrings', 'Bracelet', 1),
	(728, 87, 'What type of surprise gesture do women appreciate?', 'Ignoring special occasions', 'Buying expensive gifts', 'Surprise spa day', 'Doing household chores without being asked', 4),
	(729, 87, 'Which activity is likely to be enjoyed on a girls\' night out?', 'Watching movies at home', 'Bar hopping', 'Attending a book club meeting', 'Gardening', 2),
	(730, 87, 'What type of music do many women enjoy listening to?', 'Jazz', 'Classical', 'Heavy metal', 'Hip-hop', 1),
	(731, 87, 'What drink is often preferred when celebrating a special occasion?', 'Soda', 'Tea', 'Milk', 'Champagne', 4),
	(732, 87, 'Which color is commonly associated with love and romance?', 'Red', 'Blue', 'Yellow', 'Green', 1),
	(733, 87, 'What type of vacation is often ideal for a relaxing getaway?', 'Beach resort', 'Safari adventure', 'Ski trip', 'Backpacking in the mountains', 1),
	(734, 87, 'Which fashion accessory do many women consider essential?', 'Top hat', 'Handbag', 'Spurs', 'Fanny pack', 2),
	(735, 87, 'What type of dessert is frequently seen as an indulgent treat for women?', 'Plain rice', 'Chocolate cake', 'Bowl of peas', 'Sliced cucumber', 2),
	(736, 87, 'Which type of movie genre appeals to many women?', 'Gore-filled horror', 'Romantic comedy', 'Sci-fi action', 'Documentaries', 2),
	(737, 87, 'What kind of perfume fragrance is often favored by women?', 'Floral', 'Bacon', 'Fish', 'Freshly cut grass', 1),
	(738, 87, 'Which type of gift wrap decoration is typically favored by women?', 'Hay bale', 'Lace ribbon', 'Toothpicks', 'Barbed wire', 2),
	(754, 89, 'What is a commonly used method to authenticate users by requiring two pieces of unique information, typically something the user knows and something the user possesses?', 'Username and Password', 'IP Address', 'Biometric Scan', 'RSA Token', 4),
	(755, 89, 'Which type of attack aims to make a computer system or network unavailable for its intended users, by disrupting services such as servers or networks?', 'Malware Attack', 'Denial of Service (DoS) Attack', 'Phishing Attack', 'SQL Injection Attack', 2),
	(756, 89, 'Which encryption protocol provides secure communication over a computer network by encrypting the data that is transmitted between systems?', 'SSL/TLS', 'SHA-256', 'RSA', 'AES', 1),
	(757, 89, 'What security measure involves verifying the identity of a user or process attempting to access a system by using a password or other authentication factors?', 'Firewall', 'Intrusion Detection System', 'Two-Factor Authentication', 'Antivirus Software', 3),
	(758, 89, 'In cybersecurity, what does \'Phishing\' refer to?', 'Tricking users into revealing sensitive information', 'Encrypting data files', 'Overloading a network', 'Stealing physical assets', 1),
	(759, 89, 'Which type of malware is designed to block access to computer systems or files until a ransom is paid?', 'Adware', 'Spyware', 'Ransomware', 'Rootkit', 3),
	(760, 89, 'What is the practice of protecting information by converting it into a code that can only be deciphered with a key known as?', 'Obfuscation', 'Decryption', 'Authentication', 'Cryptography', 4),
	(761, 89, 'What is a social engineering technique that involves manipulating individuals into performing actions or divulging confidential information?', 'Phishing', 'Whaling Attack', 'Zero-Day Exploit', 'SQL Injection', 1),
	(762, 89, 'Which cybersecurity concept includes the mechanisms, policies, procedures, and technologies that work together to protect information assets?', 'Information Security', 'End-to-End Encryption', 'Data Loss Prevention (DLP)', 'Access Control', 1),
	(763, 89, 'What is a program that appears to be legitimate software but is actually designed to carry out harmful activities on a computer system?', 'Exploit Kit', 'Firewall', 'Trojan Horse', 'Backdoor', 3),
	(764, 89, 'Which of the following is an example of a physical security control?', 'Firewall', 'Network Intrusion Detection System', 'Antivirus Software', 'Biometric Door Lock', 4),
	(765, 89, 'What is a technique used to gain unauthorized access to information in an information system by changing the system\'s code or data in unexpected ways?', 'Denial-of-Service (DoS) Attack', 'Cross-Site Scripting (XSS)', 'Man-in-the-Middle Attack', 'SQL Injection', 4),
	(766, 89, 'Which security measure hides the true identity and location of network communications to protect data during transmission?', 'Digital Signature', 'Intrusion Detection System', 'Virtual Private Network (VPN)', 'Packet Filtering', 3),
	(767, 89, 'In IT security, what does \'Patch Management\' involve?', 'Ensuring software updates are applied', 'Software testing process', 'Removing obsolete hardware', 'Securing physical locations', 1),
	(768, 89, 'What is the term for the process of analyzing a software application\'s code to identify and correct programming errors, vulnerabilities, and malicious code?', 'Bug Bounty', 'Code Review', 'Penetration Testing', 'Social Engineering', 2);

-- Dumping structure for table quiz.quiz
CREATE TABLE IF NOT EXISTS `quiz` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text DEFAULT NULL,
  `join_code` varchar(6) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `join_code` (`join_code`)
) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table quiz.quiz: ~8 rows (approximately)
INSERT INTO `quiz` (`id`, `name`, `description`, `join_code`, `created_at`) VALUES
	(1, 'LootLocker', 'A LootLocker test quiz containing weird a funny questions', 'PY2SW8', '2024-04-15 09:52:06'),
	(15, 'Science Quiz', 'Test your knowledge about science', 'RODQ7L', '2024-04-30 08:00:00'),
	(16, 'History Quiz', 'Explore historical events and figures', 'J9RX2V', '2024-04-30 09:00:00'),
	(77, 'Denmark Quiz', 'AI-quiz on Denmark', '1R136G', '2024-06-05 15:33:06'),
	(80, 'Sweden Quiz', 'AI-quiz generated on Sweden as subject', 'HXAUZO', '2024-06-05 15:41:55'),
	(86, 'Simple Addiction and Substraction math', 'AI-quiz generated on Simple Addiction and Substraction math as subject', 'Y0GPDV', '2024-06-05 16:27:57'),
	(87, 'What do women like the most?', 'AI-quiz generated on What do women like the most? as subject', 'NY67A7', '2024-06-05 16:37:05'),
	(89, 'IT Security', 'AI-quiz generated on IT Security', 'NX7KFW', '2024-06-05 16:51:20');

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
