-- TABLES: 

-- USER
-- 	id
-- 	first_name
-- 	last_name
-- 	email
-- 	username? //user email for username
-- 	password
-- 	role
-- 	created
-- 	modified
-- 	last_logged_in
-- PRODUCTS
-- 	id
-- 	name
-- 	type
-- 	catergory
-- 	price
-- 	inventory
-- 	created
-- 	modified
-- USER_CONTACT_INFO //meta?
-- 	address? //meta?
--	- address
--	- address_2
--	- city
--	- zip
-- 	phone
--	- home
--	- cell
-- 	created
-- 	modified
-- META
-- 	id
-- 	table
-- 	table_id
-- 	meta_key
-- 	meta_value
-- 	created
-- 	modified
-- ORDERS
-- 	id
-- 	user_id?
-- 	total
-- 	created
-- 	modified

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
-- Database: `pizza`
--
