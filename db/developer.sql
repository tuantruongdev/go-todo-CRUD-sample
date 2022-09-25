-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Máy chủ: 127.0.0.1
-- Thời gian đã tạo: Th9 25, 2022 lúc 02:35 PM
-- Phiên bản máy phục vụ: 10.4.24-MariaDB
-- Phiên bản PHP: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Cơ sở dữ liệu: `developer`
--

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `todo_items`
--

CREATE TABLE `todo_items` (
  `id` int(11) NOT NULL,
  `title` varchar(150) NOT NULL,
  `status` enum('Doing','Finished') DEFAULT 'Doing',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Đang đổ dữ liệu cho bảng `todo_items`
--

INSERT INTO `todo_items` (`id`, `title`, `status`, `created_at`, `updated_at`) VALUES
(1, 'do homework', 'Doing', '2022-09-24 18:52:21', '2022-09-24 18:52:21'),
(5, 'iphone 14', 'Doing', '2022-09-24 19:23:47', '2022-09-24 19:23:47'),
(6, 'homework', 'Doing', '2022-09-24 19:24:13', '2022-09-24 19:24:13'),
(7, 'homework', 'Doing', '2022-09-24 19:24:25', '2022-09-24 19:24:25'),
(8, 'homework22', 'Doing', '2022-09-25 08:41:40', '2022-09-25 08:41:40'),
(9, '', 'Doing', '2022-09-25 08:43:45', '2022-09-25 08:43:45'),
(10, '', 'Doing', '2022-09-25 08:43:49', '2022-09-25 08:43:49'),
(11, 'homework69', 'Finished', '2022-09-25 09:37:18', '2022-09-25 11:13:30');

--
-- Chỉ mục cho các bảng đã đổ
--

--
-- Chỉ mục cho bảng `todo_items`
--
ALTER TABLE `todo_items`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT cho các bảng đã đổ
--

--
-- AUTO_INCREMENT cho bảng `todo_items`
--
ALTER TABLE `todo_items`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
