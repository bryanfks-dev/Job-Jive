-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 02, 2024 at 05:23 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `testing_webdev`
--

-- --------------------------------------------------------

--
-- Table structure for table `attendances`
--

CREATE TABLE `attendances` (
  `User_ID` int(10) UNSIGNED NOT NULL,
  `Date` datetime NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `attendances`
--

INSERT INTO `attendances` (`User_ID`, `Date`, `created_at`, `updated_at`) VALUES
(6, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(4, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(1, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(1, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(6, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(6, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(5, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(5, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(4, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(1, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(6, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(5, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(1, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(2, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(5, '2024-05-02 15:11:24', '2024-05-02 08:11:24', '2024-05-02 08:11:24');

-- --------------------------------------------------------

--
-- Table structure for table `divisions`
--

CREATE TABLE `divisions` (
  `Department_ID` int(10) UNSIGNED NOT NULL,
  `Department_Name` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `divisions`
--

INSERT INTO `divisions` (`Department_ID`, `Department_Name`, `created_at`, `updated_at`) VALUES
(1, 'dolores', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(2, 'vitae', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(3, 'voluptatem', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(4, 'odio', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(5, 'ut', '2024-05-02 08:11:24', '2024-05-02 08:11:24');

-- --------------------------------------------------------

--
-- Table structure for table `employee_work_assessments`
--

CREATE TABLE `employee_work_assessments` (
  `User_ID` int(10) UNSIGNED NOT NULL,
  `Feedback` text NOT NULL,
  `Date` date NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `employee_work_assessments`
--

INSERT INTO `employee_work_assessments` (`User_ID`, `Feedback`, `Date`, `created_at`, `updated_at`) VALUES
(6, 'Et eius sit est quia numquam.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(6, 'Iusto ab adipisci nihil.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(3, 'Molestias dignissimos repellendus voluptatibus.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(6, 'Laudantium ab quisquam nemo aut.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(6, 'Vitae voluptatum explicabo quis.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(4, 'At ut eligendi laboriosam.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(5, 'Tempore neque debitis ducimus asperiores.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(6, 'Repellat aliquam laudantium qui provident.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(5, 'Voluptatem asperiores omnis.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(3, 'Est consequuntur a illum.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(4, 'Nam quia qui aliquam deleniti.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(3, 'Maxime hic animi.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(6, 'Deserunt pariatur quos laborum.', '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24');

-- --------------------------------------------------------

--
-- Table structure for table `migrations`
--

CREATE TABLE `migrations` (
  `id` int(10) UNSIGNED NOT NULL,
  `migration` varchar(255) NOT NULL,
  `batch` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `migrations`
--

INSERT INTO `migrations` (`id`, `migration`, `batch`) VALUES
(1, '2024_05_02_141503_divisions', 1),
(2, '2024_05_02_141522_users', 1),
(3, '2024_05_02_141544_salaries', 1),
(4, '2024_05_02_141602_attendances', 1),
(5, '2024_05_02_141622_employee_work_assessments', 1);

-- --------------------------------------------------------

--
-- Table structure for table `salaries`
--

CREATE TABLE `salaries` (
  `User_ID` int(10) UNSIGNED NOT NULL,
  `Initial_Salary` int(11) DEFAULT NULL,
  `Final_Salary` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `salaries`
--

INSERT INTO `salaries` (`User_ID`, `Initial_Salary`, `Final_Salary`, `created_at`, `updated_at`) VALUES
(6, 19491, 19699, '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(1, 14980, 14886, '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(2, 14676, 16680, '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(2, 14206, 11860, '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(3, 15657, 18846, '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(3, 15529, 10673, '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(6, 16050, 12134, '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(3, 14074, 11565, '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(1, 18584, 11433, '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(2, 16510, 17701, '2024-05-02 08:11:24', '2024-05-02 08:11:24');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `User_ID` int(10) UNSIGNED NOT NULL,
  `Full_Name` varchar(255) NOT NULL,
  `Email` varchar(255) NOT NULL,
  `Password` varchar(255) NOT NULL,
  `Manager_ID` int(10) UNSIGNED DEFAULT NULL,
  `Address` varchar(255) NOT NULL,
  `NIK` varchar(255) NOT NULL,
  `Gender` varchar(255) NOT NULL,
  `Phone_Number` varchar(255) NOT NULL,
  `Department_ID` int(10) UNSIGNED NOT NULL,
  `First_Login` date NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`User_ID`, `Full_Name`, `Email`, `Password`, `Manager_ID`, `Address`, `NIK`, `Gender`, `Phone_Number`, `Department_ID`, `First_Login`, `created_at`, `updated_at`) VALUES
(1, 'Admin', 'admin@localhost', '$2y$12$cFAc6toyDmUGAlwfJEmOROSuBw5Kn0HCJ6WjZMrJsOiLIeVIFAtTi', NULL, 'Jl. Cangcimen', '123456789', 'Laki-laki', '08123456789', 1, '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(2, 'Kanz', 'Kanz@localhost', '$2y$12$FJreDDbO/w7FFjOLv4BluelDJpdxXS/6PRkVJw42QeDQkK1Ncl/K.', NULL, 'Jl. Cangcimen', '123456789', 'Laki-laki', '08123456789', 1, '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(3, 'Mr. Cyril Heaney', 'humberto33@example.org', '$2y$12$2DMjOX.IUXNkuVEo308TpOpHF6bHLGwjLgnU35aIdf2Ba7mDs2Uk.', 2, '88064 Parker Underpass Suite 863\nKianaberg, MO 55298', '3062193791', 'Perempuan', '234-703-8597', 2, '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(4, 'Alta Kovacek', 'hlowe@example.com', '$2y$12$2DMjOX.IUXNkuVEo308TpOpHF6bHLGwjLgnU35aIdf2Ba7mDs2Uk.', 2, '1257 Ike Glens\nPort Ezramouth, CT 89188-7158', '5872222178', 'Laki-laki', '+1.423.640.5940', 3, '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(5, 'Laurie Cremin', 'ttoy@example.org', '$2y$12$2DMjOX.IUXNkuVEo308TpOpHF6bHLGwjLgnU35aIdf2Ba7mDs2Uk.', 1, '896 Kuhlman Pine\nLake Colt, VT 62958', '6627805376', 'Perempuan', '917.397.3689', 5, '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24'),
(6, 'Evert Friesen', 'robb.fisher@example.net', '$2y$12$2DMjOX.IUXNkuVEo308TpOpHF6bHLGwjLgnU35aIdf2Ba7mDs2Uk.', 1, '3035 Sabrina Motorway Suite 498\nLake Twila, MA 76478-5938', '4044701073', 'Laki-laki', '971.674.8449', 4, '2024-05-02', '2024-05-02 08:11:24', '2024-05-02 08:11:24');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `attendances`
--
ALTER TABLE `attendances`
  ADD KEY `attendances_user_id_foreign` (`User_ID`);

--
-- Indexes for table `divisions`
--
ALTER TABLE `divisions`
  ADD PRIMARY KEY (`Department_ID`);

--
-- Indexes for table `employee_work_assessments`
--
ALTER TABLE `employee_work_assessments`
  ADD KEY `employee_work_assessments_user_id_foreign` (`User_ID`);

--
-- Indexes for table `migrations`
--
ALTER TABLE `migrations`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `salaries`
--
ALTER TABLE `salaries`
  ADD KEY `salaries_user_id_foreign` (`User_ID`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`User_ID`),
  ADD UNIQUE KEY `users_email_unique` (`Email`),
  ADD KEY `users_manager_id_foreign` (`Manager_ID`),
  ADD KEY `users_department_id_foreign` (`Department_ID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `divisions`
--
ALTER TABLE `divisions`
  MODIFY `Department_ID` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `migrations`
--
ALTER TABLE `migrations`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `User_ID` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `attendances`
--
ALTER TABLE `attendances`
  ADD CONSTRAINT `attendances_user_id_foreign` FOREIGN KEY (`User_ID`) REFERENCES `users` (`User_ID`) ON DELETE CASCADE;

--
-- Constraints for table `employee_work_assessments`
--
ALTER TABLE `employee_work_assessments`
  ADD CONSTRAINT `employee_work_assessments_user_id_foreign` FOREIGN KEY (`User_ID`) REFERENCES `users` (`User_ID`) ON DELETE CASCADE;

--
-- Constraints for table `salaries`
--
ALTER TABLE `salaries`
  ADD CONSTRAINT `salaries_user_id_foreign` FOREIGN KEY (`User_ID`) REFERENCES `users` (`User_ID`) ON DELETE CASCADE;

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `users_department_id_foreign` FOREIGN KEY (`Department_ID`) REFERENCES `divisions` (`Department_ID`) ON DELETE CASCADE,
  ADD CONSTRAINT `users_manager_id_foreign` FOREIGN KEY (`Manager_ID`) REFERENCES `users` (`User_ID`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
