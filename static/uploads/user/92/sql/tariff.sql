-- MySQL dump 10.13  Distrib 5.7.24, for Linux (x86_64)
--
-- Host: localhost    Database: tariff
-- ------------------------------------------------------
-- Server version	5.7.24-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `mst_meter`
--

DROP TABLE IF EXISTS `mst_meter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mst_meter` (
  `service_id` int(11) NOT NULL,
  `drop_distance` int(11) NOT NULL COMMENT 'Distance per drop, in meter, eg. 1 drop = 100 meter',
  `drop_time` int(11) NOT NULL COMMENT 'Time per drop, in seconds, eg. 1 drop = 35 seconds',
  `speed_limit` int(11) NOT NULL COMMENT 'Speed limit between distance and time pricing, in km/h',
  `created_by_msb` bigint(20) unsigned NOT NULL,
  `created_by_lsb` bigint(20) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_by_msb` bigint(20) unsigned NOT NULL,
  `updated_by_lsb` bigint(20) unsigned NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mst_meter`
--

LOCK TABLES `mst_meter` WRITE;
/*!40000 ALTER TABLE `mst_meter` DISABLE KEYS */;
/*!40000 ALTER TABLE `mst_meter` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mst_minimum_payment`
--

DROP TABLE IF EXISTS `mst_minimum_payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mst_minimum_payment` (
  `service_id` int(11) NOT NULL,
  `area_msb` bigint(20) unsigned NOT NULL,
  `area_lsb` bigint(20) unsigned NOT NULL,
  `hail_type` int(11) NOT NULL COMMENT '0 - On street\n1 - Booking',
  `currency` char(3) NOT NULL DEFAULT 'IDR',
  `min_payment` decimal(15,5) NOT NULL,
  `created_by_msb` bigint(20) unsigned NOT NULL,
  `created_by_lsb` bigint(20) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_by_msb` bigint(20) unsigned NOT NULL,
  `updated_by_lsb` bigint(20) unsigned NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`service_id`,`area_msb`,`area_lsb`,`hail_type`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mst_minimum_payment`
--

LOCK TABLES `mst_minimum_payment` WRITE;
/*!40000 ALTER TABLE `mst_minimum_payment` DISABLE KEYS */;
/*!40000 ALTER TABLE `mst_minimum_payment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mst_tariff`
--

DROP TABLE IF EXISTS `mst_tariff`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mst_tariff` (
  `tariff_id` int(11) NOT NULL,
  `service_id` int(11) NOT NULL,
  `area_msb` bigint(20) unsigned NOT NULL,
  `area_lsb` bigint(20) unsigned NOT NULL,
  `tariff_type` int(11) NOT NULL COMMENT '0 - normal tariff',
  `currency` char(3) NOT NULL DEFAULT 'IDR',
  `flag_fall` decimal(15,5) NOT NULL COMMENT 'Fixed start fee',
  `distance_rate` decimal(15,5) NOT NULL COMMENT 'tariff per km',
  `time_rate` decimal(15,5) NOT NULL COMMENT 'tariff per hour',
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `is_default` tinyint(1) NOT NULL,
  `priority` int(11) NOT NULL,
  `status` int(11) NOT NULL COMMENT '0 - inactive\n1 - active',
  `created_by_msb` bigint(20) unsigned NOT NULL,
  `created_by_lsb` bigint(20) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_by_msb` bigint(20) unsigned NOT NULL,
  `updated_by_lsb` bigint(20) unsigned NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`tariff_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mst_tariff`
--

LOCK TABLES `mst_tariff` WRITE;
/*!40000 ALTER TABLE `mst_tariff` DISABLE KEYS */;
/*!40000 ALTER TABLE `mst_tariff` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-12-07  9:28:26
