-- phpMyAdmin SQL Dump
-- version 3.4.11.1deb2+deb7u1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Oct 12, 2016 at 03:32 PM
-- Server version: 5.6.23
-- PHP Version: 5.5.24-1~dotdeb+wheezy.1

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `orb`
--

-- --------------------------------------------------------

--
-- Table structure for table `attendee_categories`
--

CREATE TABLE IF NOT EXISTS `attendee_categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `event_id` int(10) unsigned NOT NULL,
  `creator_id` int(10) unsigned NOT NULL,
  `modifier_id` int(10) unsigned NOT NULL,
  `created` datetime NOT NULL,
  `modified` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `event_id` (`event_id`),
  KEY `creator_id` (`creator_id`),
  KEY `modifier_id` (`modifier_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `attendee_groups`
--

CREATE TABLE IF NOT EXISTS `attendee_groups` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `convention_group_id` int(10) unsigned DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `created` datetime DEFAULT NULL,
  `creator_id` int(10) unsigned DEFAULT NULL,
  `modified` datetime DEFAULT NULL,
  `modifier_id` int(10) unsigned DEFAULT NULL,
  `contact_name` varchar(255) DEFAULT NULL,
  `contact_title` varchar(255) DEFAULT NULL,
  `contact_email` varchar(255) DEFAULT NULL,
  `contact_phone` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `creator_id` (`creator_id`),
  KEY `modifier_id` (`modifier_id`),
  KEY `convention_group_id` (`convention_group_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `audits`
--

CREATE TABLE IF NOT EXISTS `audits` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(255) NOT NULL,
  `table_used` varchar(255) NOT NULL,
  `old_value` text NOT NULL,
  `new_value` text NOT NULL,
  `created` datetime NOT NULL,
  `modifier_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `modifier_id` (`modifier_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `available_blocks`
--

CREATE TABLE IF NOT EXISTS `available_blocks` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `block_id` int(10) unsigned NOT NULL,
  `quantity` int(10) unsigned NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `created` datetime NOT NULL,
  `creator_id` int(10) unsigned NOT NULL,
  `modified` datetime NOT NULL,
  `modifier_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `block_id` (`block_id`),
  KEY `creator_id` (`creator_id`),
  KEY `modifier_id` (`modifier_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `available_slots`
--

CREATE TABLE IF NOT EXISTS `available_slots` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `section_id` int(11) UNSIGNED NOT NULL,
  `date` date NOT NULL,
  `start` time NOT NULL,
  `end` time NOT NULL,
  `creator_id` int(10) unsigned NOT NULL,
  `created` datetime NOT NULL,
  `modifier_id` int(10) unsigned NOT NULL,
  `modified` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `section_id` (`section_id`),
  KEY `creator_id` (`creator_id`),
  KEY `modifier_id` (`modifier_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `blocks`
--

CREATE TABLE IF NOT EXISTS `blocks` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `room_type_id` int(10) unsigned NOT NULL,
  `rate` decimal(19,4) unsigned NOT NULL,
  `created` datetime DEFAULT NULL,
  `creator_id` int(10) unsigned DEFAULT NULL,
  `modified` datetime DEFAULT NULL,
  `modifier_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `room_type_id` (`room_type_id`),
  KEY `creator_id` (`creator_id`),
  KEY `modifier_id` (`modifier_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `events`
--

CREATE TABLE IF NOT EXISTS `events` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `org_id` int(11) unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `start` datetime NOT NULL,
  `end` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `org_id` (`org_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `hotels`
--

CREATE TABLE IF NOT EXISTS `hotels` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `zone_id` int(10) unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `contact_email` varchar(255) NOT NULL,
  `address` varchar(255) DEFAULT NULL,
  `address2` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `state` varchar(255) DEFAULT NULL,
  `country` varchar(3) DEFAULT NULL,
  `postalcode` varchar(25) DEFAULT NULL,
  `phone` varchar(25) DEFAULT NULL,
  `fax` varchar(25) DEFAULT NULL,
  `notes` mediumtext,
  `is_union` tinyint(3) unsigned NOT NULL COMMENT 'Whether the hotel is a Union hotel or not.',
  `created` datetime DEFAULT NULL,
  `creator_id` int(10) unsigned DEFAULT NULL,
  `modified` datetime DEFAULT NULL,
  `modifier_id` int(10) unsigned DEFAULT NULL,
  `contact_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `zone_id` (`zone_id`),
  KEY `creator_id` (`creator_id`),
  KEY `modifier_id` (`modifier_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `hotel_metas`
--

CREATE TABLE IF NOT EXISTS `hotel_metas` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `hotel_id` int(10) unsigned DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `modified` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `hotel_id` (`hotel_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `meeting_spaces`
--

CREATE TABLE IF NOT EXISTS `meeting_spaces` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `venue_id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `creator_id` int(10) unsigned NOT NULL,
  `created` datetime NOT NULL,
  `sq_ft` int(11) NOT NULL,
  `capacity` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `venue_id` (`venue_id`),
  KEY `creator_id` (`creator_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `organizations`
--

CREATE TABLE IF NOT EXISTS `organizations` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created` datetime NOT NULL,
  `modified` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `room_types`
--

CREATE TABLE IF NOT EXISTS `room_types` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `hotel_id` int(11) unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `created` datetime DEFAULT NULL,
  `creator_id` int(10) unsigned DEFAULT NULL,
  `modified` datetime DEFAULT NULL,
  `modifier_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `creator_id` (`creator_id`),
  KEY `modifier_id` (`modifier_id`),
  KEY `hotel_id` (`hotel_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `sections`
--

CREATE TABLE IF NOT EXISTS `sections` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `meeting_space_id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `capacity` int(11) NOT NULL,
  `creator_id` int(10) unsigned NOT NULL,
  `created` datetime NOT NULL,
  `modifier_id` int(10) unsigned NOT NULL,
  `modified` datetime NOT NULL,
  `square_footage` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `meeting_space_id` (`meeting_space_id`),
  KEY `creator_id` (`creator_id`),
  KEY `modifier_id` (`modifier_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `slots`
--

CREATE TABLE IF NOT EXISTS `slots` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `status` varchar(255) NOT NULL,
  `available_slot_id` int(11) UNSIGNED NOT NULL,
  `attendee_group_id` int(11) unsigned NOT NULL,
  `section_id` int(11) UNSIGNED NOT NULL,
  `start` time NOT NULL,
  `end` time NOT NULL,
  `size` int(11) NOT NULL,
  `creator_id` int(10) unsigned NOT NULL,
  `created` datetime NOT NULL,
  `modifier_id` int(10) unsigned NOT NULL,
  `modified` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `available_slot_id` (`available_slot_id`),
  KEY `meeter_id` (`attendee_group_id`),
  KEY `section_id` (`section_id`),
  KEY `creator_id` (`creator_id`),
  KEY `modifier_id` (`modifier_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `subblocks`
--

CREATE TABLE IF NOT EXISTS `subblocks` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `block_id` int(10) unsigned NOT NULL,
  `attendee_group_id` int(10) unsigned NOT NULL,
  `quantity` int(11) NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `created` datetime DEFAULT NULL,
  `creator_id` int(10) unsigned DEFAULT NULL,
  `modified` datetime DEFAULT NULL,
  `modifier_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `block_id` (`block_id`),
  KEY `attendee_id` (`attendee_group_id`),
  KEY `creator_id` (`creator_id`),
  KEY `modifier_id` (`modifier_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE IF NOT EXISTS `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_slug` varchar(40) NOT NULL DEFAULT 'viewer',
  `email` varchar(255) NOT NULL,
  `password` varchar(60) NOT NULL,
  `active` tinyint(4) NOT NULL DEFAULT '1',
  `created` datetime DEFAULT NULL,
  `creator_id` int(10) unsigned DEFAULT NULL,
  `modified` datetime DEFAULT NULL,
  `modifier_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `creator_id` (`creator_id`),
  KEY `modifier_id` (`modifier_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `users_orgs`
--

CREATE TABLE IF NOT EXISTS `users_orgs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `org_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `org_id` (`org_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `venues`
--

CREATE TABLE IF NOT EXISTS `venues` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `event_id` int(11) unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `creator_id` int(10) unsigned NOT NULL,
  `created` datetime NOT NULL,
  `zone_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `event_id` (`event_id`),
  KEY `creator_id` (`creator_id`),
  KEY `zone_id` (`zone_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `zones`
--

CREATE TABLE IF NOT EXISTS `zones` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id` int(11) unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT '',
  `created` datetime DEFAULT NULL,
  `creator_id` int(10) unsigned DEFAULT NULL,
  `modified` datetime DEFAULT NULL,
  `modifier_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `creator_id` (`creator_id`),
  KEY `modifier_id` (`modifier_id`),
  KEY `event_id` (`event_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `attendee_categories`
--
ALTER TABLE `attendee_categories`
  ADD CONSTRAINT `attendee_categories_ibfk_1` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`),
  ADD CONSTRAINT `attendee_categories_ibfk_2` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `attendee_categories_ibfk_3` FOREIGN KEY (`modifier_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `attendee_groups`
--
ALTER TABLE `attendee_groups`
  ADD CONSTRAINT `attendee_groups_ibfk_1` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `attendee_groups_ibfk_2` FOREIGN KEY (`modifier_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `attendee_groups_ibfk_3` FOREIGN KEY (`convention_group_id`) REFERENCES `attendee_categories` (`id`);

--
-- Constraints for table `audits`
--
ALTER TABLE `audits`
  ADD CONSTRAINT `audits_ibfk_1` FOREIGN KEY (`modifier_id`) REFERENCES `events` (`id`);

--
-- Constraints for table `available_blocks`
--
ALTER TABLE `available_blocks`
  ADD CONSTRAINT `available_blocks_ibfk_2` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `available_blocks_ibfk_3` FOREIGN KEY (`modifier_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `available_blocks_ibfk_4` FOREIGN KEY (`block_id`) REFERENCES `blocks` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `available_slots`
--
ALTER TABLE `available_slots`
  ADD CONSTRAINT `available_slots_ibfk_1` FOREIGN KEY (`section_id`) REFERENCES `sections` (`id`),
  ADD CONSTRAINT `available_slots_ibfk_2` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `available_slots_ibfk_3` FOREIGN KEY (`modifier_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `blocks`
--
ALTER TABLE `blocks`
  ADD CONSTRAINT `blocks_ibfk_2` FOREIGN KEY (`room_type_id`) REFERENCES `room_types` (`id`),
  ADD CONSTRAINT `blocks_ibfk_3` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `blocks_ibfk_4` FOREIGN KEY (`modifier_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `events`
--
ALTER TABLE `events`
  ADD CONSTRAINT `events_ibfk_1` FOREIGN KEY (`org_id`) REFERENCES `organizations` (`id`);

--
-- Constraints for table `hotels`
--
ALTER TABLE `hotels`
  ADD CONSTRAINT `hotels_ibfk_1` FOREIGN KEY (`zone_id`) REFERENCES `zones` (`id`),
  ADD CONSTRAINT `hotels_ibfk_2` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `hotels_ibfk_3` FOREIGN KEY (`modifier_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `hotel_metas`
--
ALTER TABLE `hotel_metas`
  ADD CONSTRAINT `hotel_metas_ibfk_1` FOREIGN KEY (`hotel_id`) REFERENCES `hotels` (`id`);

--
-- Constraints for table `meeting_spaces`
--
ALTER TABLE `meeting_spaces`
  ADD CONSTRAINT `meeting_spaces_ibfk_1` FOREIGN KEY (`venue_id`) REFERENCES `venues` (`id`),
  ADD CONSTRAINT `meeting_spaces_ibfk_2` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `room_types`
--
ALTER TABLE `room_types`
  ADD CONSTRAINT `room_types_ibfk_1` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `room_types_ibfk_2` FOREIGN KEY (`modifier_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `room_types_ibfk_3` FOREIGN KEY (`hotel_id`) REFERENCES `hotels` (`id`);

--
-- Constraints for table `sections`
--
ALTER TABLE `sections`
  ADD CONSTRAINT `sections_ibfk_1` FOREIGN KEY (`meeting_space_id`) REFERENCES `meeting_spaces` (`id`),
  ADD CONSTRAINT `sections_ibfk_2` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `sections_ibfk_3` FOREIGN KEY (`modifier_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `slots`
--
ALTER TABLE `slots`
  ADD CONSTRAINT `slots_ibfk_1` FOREIGN KEY (`available_slot_id`) REFERENCES `available_slots` (`id`),
  ADD CONSTRAINT `slots_ibfk_2` FOREIGN KEY (`section_id`) REFERENCES `sections` (`id`),
  ADD CONSTRAINT `slots_ibfk_3` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `slots_ibfk_4` FOREIGN KEY (`modifier_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `subblocks`
--
ALTER TABLE `subblocks`
  ADD CONSTRAINT `subblocks_ibfk_2` FOREIGN KEY (`attendee_group_id`) REFERENCES `attendee_groups` (`id`),
  ADD CONSTRAINT `subblocks_ibfk_3` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `subblocks_ibfk_4` FOREIGN KEY (`modifier_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `subblocks_ibfk_5` FOREIGN KEY (`block_id`) REFERENCES `blocks` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `users_ibfk_2` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `users_ibfk_3` FOREIGN KEY (`modifier_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `users_orgs`
--
ALTER TABLE `users_orgs`
  ADD CONSTRAINT `users_orgs_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `users_orgs_ibfk_2` FOREIGN KEY (`org_id`) REFERENCES `organizations` (`id`);

--
-- Constraints for table `venues`
--
ALTER TABLE `venues`
  ADD CONSTRAINT `venues_ibfk_1` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`),
  ADD CONSTRAINT `venues_ibfk_2` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `venues_ibfk_3` FOREIGN KEY (`zone_id`) REFERENCES `zones` (`id`);

--
-- Constraints for table `zones`
--
ALTER TABLE `zones`
  ADD CONSTRAINT `zones_ibfk_1` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `zones_ibfk_2` FOREIGN KEY (`modifier_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `zones_ibfk_3` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
