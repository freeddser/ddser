-- MySQL dump 10.13  Distrib 5.7.24, for Linux (x86_64)
--
-- Host: localhost    Database: ddser
-- ------------------------------------------------------
-- Server version	5.7.24-0ubuntu0.18.04.1

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
-- Table structure for table `resource`
--

DROP TABLE IF EXISTS `resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `resource_name` varchar(255) DEFAULT NULL,
  `resource_desc` varchar(255) DEFAULT NULL,
  `resource_parent` varchar(255) DEFAULT NULL,
  `resource_url` varchar(255) DEFAULT NULL,
  `create_date` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10030 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `resource`
--

LOCK TABLES `resource` WRITE;
/*!40000 ALTER TABLE `resource` DISABLE KEYS */;
INSERT INTO `resource` VALUES (11,'Manager_privileges','Manager_privileges','0','/admin/manager','2015-06-01 10:19:55'),(8,'User','User','11','/admin/manager/user','2015-11-16 13:41:13'),(10014,'Task','Task','0','/task','2015-11-17 12:19:17'),(10015,'AddTask','AddTask','10014','/task/add','2015-11-17 12:20:21'),(10016,'ShowTask','ShowTask','10014','/task/show','2015-11-17 12:20:55'),(14,'Resource','Resource','11','/admin/manager/resource','2015-05-31 16:50:33'),(10017,'tasklogs','tasklogs','10014','/task/tasklog','2015-11-17 12:22:39'),(30,'Role','Role','11','/admin/manager/role','2015-11-16 13:40:59'),(34,'UserRole','UserRole','11','/admin/manager/userrole','2015-06-04 21:52:21'),(10000,'Default','default','0','/default','2015-06-04 21:52:21'),(10001,'RoleResource','RoleResource','11','/admin/manager/roleresource','2015-06-04 21:52:21'),(10018,'Upload','Upload','0','/upload','2015-11-30 13:39:37'),(10019,'UploadSql','UploadSql','10018','/upload/sql','2015-11-30 13:39:54'),(10023,'RoleTask','RoleTask','11','/admin/manager/roletasktype','2016-05-14 23:12:27');
/*!40000 ALTER TABLE `resource` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(255) DEFAULT NULL,
  `create_date` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=131 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'Admin','2015-05-29 10:27:57'),(100,'user','2015-05-29 10:28:20'),(129,'projectA','2017-09-25 11:08:58'),(130,'projectB','2017-09-25 11:09:09');
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_resource`
--

DROP TABLE IF EXISTS `role_resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) DEFAULT NULL,
  `resource_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=515 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_resource`
--

LOCK TABLES `role_resource` WRITE;
/*!40000 ALTER TABLE `role_resource` DISABLE KEYS */;
INSERT INTO `role_resource` VALUES (62,1,10000),(92,1,10000),(70,1,10000),(100,1,10000),(180,1,10000),(107,1,10000),(36,100,10000),(223,1,10000),(347,1,10000),(242,1,10000),(422,1,10000),(396,1,10000),(380,1,10000),(365,1,10000),(485,100,10000),(412,1,10000),(432,1,10000),(476,1,10000),(512,1,10023),(511,1,10001),(510,1,34),(509,1,30),(508,1,14),(507,1,11),(506,1,8),(514,130,10000),(513,129,10000),(486,100,10014),(487,100,10015),(488,100,10016),(489,100,10017),(490,100,10018),(491,100,10019);
/*!40000 ALTER TABLE `role_resource` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `task`
--

DROP TABLE IF EXISTS `task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `task` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task_type` int(11) NOT NULL,
  `taskname` varchar(200) NOT NULL,
  `task_shell` varchar(545) NOT NULL,
  `shell_params` varchar(545) DEFAULT NULL,
  `create_time` varchar(64) NOT NULL,
  `task_type_desc` varchar(64) NOT NULL,
  `token_id` int(11) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '1' COMMENT 'default 1\n\n',
  `task_bz` varchar(45) NOT NULL DEFAULT 'Y',
  `pusher` varchar(45) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `task_status_idx` (`status`)
) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task`
--

LOCK TABLES `task` WRITE;
/*!40000 ALTER TABLE `task` DISABLE KEYS */;
INSERT INTO `task` VALUES (1,131,'deploy_to_projectB','/tmp/projectB.sh','dev','2018-12-29 10:54:04','deploy_to_projectB',37,1003,'Y','admin'),(2,131,'deploy_to_projectB','/tmp/projectB.sh','dev','2018-12-29 10:05:11','deploy_to_projectB',37,1003,'Y','admin'),(3,130,'deploy_to_projectA','/tmp/deployA.sh','master','2018-12-28 10:46:53','deploy_to_projectA',36,501,'Y','admin'),(4,130,'deploy_to_projectA','/tmp/deployA.sh','master','2018-12-28 15:58:30','deploy_to_projectA',36,501,'N','admin'),(5,131,'deploy_to_projectB','/tmp/projectB.sh','dev','2018-12-29 09:47:44','deploy_to_projectB',37,1003,'Y','admin'),(6,131,'deploy_to_projectB','/tmp/projectB.sh','dev','2018-12-29 10:06:45','deploy_to_projectB',37,501,'Y','admin'),(7,131,'deploy_to_projectB','/tmp/projectB.sh','dev','2018-12-29 10:15:48','deploy_to_projectB',37,1003,'Y','unknow');
/*!40000 ALTER TABLE `task` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `task_log`
--

DROP TABLE IF EXISTS `task_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `task_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task_logs` longtext,
  `task_id` int(11) NOT NULL,
  `token_id` int(11) DEFAULT NULL,
  `client_infos` varchar(200) DEFAULT NULL,
  `create_time` varchar(30) DEFAULT NULL,
  `update_time` varchar(30) DEFAULT NULL,
  `task_status` varchar(30) DEFAULT NULL,
  `pusher` varchar(20) DEFAULT NULL,
  `log_bz` int(3) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task_log`
--

LOCK TABLES `task_log` WRITE;
/*!40000 ALTER TABLE `task_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `task_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `task_status`
--

DROP TABLE IF EXISTS `task_status`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `task_status` (
  `status_id` int(11) NOT NULL,
  `status_desc` varchar(145) NOT NULL,
  PRIMARY KEY (`status_id`),
  UNIQUE KEY `id_UNIQUE` (`status_id`),
  KEY `task_status_status_id_idx` (`status_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task_status`
--

LOCK TABLES `task_status` WRITE;
/*!40000 ALTER TABLE `task_status` DISABLE KEYS */;
INSERT INTO `task_status` VALUES (1001,'Client_received'),(1002,'Client_executing'),(1003,'Client_execting _done'),(102,'Client_script_does_not_exist'),(101,'Client_unkonw_task_type'),(10,'Client_unreceived'),(501,'task_writing'),(502,'task_confirmation'),(1004,'Client_exec_Error');
/*!40000 ALTER TABLE `task_status` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `task_type`
--

DROP TABLE IF EXISTS `task_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `task_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task_type_desc` varchar(200) NOT NULL,
  `task_shell` varchar(128) DEFAULT NULL,
  `task_shell_params` varchar(128) DEFAULT NULL,
  `role_id` int(11) NOT NULL,
  `token_id` varchar(45) DEFAULT NULL,
  `create_time` varchar(45) DEFAULT NULL,
  `type` varchar(11) DEFAULT 'web',
  PRIMARY KEY (`id`),
  KEY `task_type_role_id_idx` (`role_id`),
  KEY `task_type_token_id_idx` (`token_id`)
) ENGINE=MyISAM AUTO_INCREMENT=139 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task_type`
--

LOCK TABLES `task_type` WRITE;
/*!40000 ALTER TABLE `task_type` DISABLE KEYS */;
INSERT INTO `task_type` VALUES (134,'restart_nginx','/etc/init.d/nginx ','restart',129,'36','2017-09-25 12:02:56','api'),(132,'psjob','ps',' auxwww',129,'36','2017-09-25 11:59:50','api'),(131,'deploy_to_projectB','/tmp/projectB.sh','dev',130,'37','2017-09-25 11:16:12','web'),(130,'deploy_to_projectA','/tmp/deployA.sh','master',129,'36','2017-09-25 11:15:45','web'),(129,'hostname','hostname','',0,'0','2017-09-25 11:13:10','api'),(135,'test','test','',129,'36','2017-09-26 09:38:30','api'),(128,'check_load','uptime','',129,'36','2017-09-25 11:12:45','api'),(136,'xx','xx','x',129,'36','2018-12-28 10:22:37','api'),(137,'x','x','x',129,'36','2018-12-28 13:54:10','web'),(138,'131','1','1',129,'36','2018-12-28 14:01:57','web');
/*!40000 ALTER TABLE `task_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `token`
--

DROP TABLE IF EXISTS `token`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `token` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `token` varchar(45) NOT NULL,
  `role_name` varchar(145) NOT NULL,
  `role_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `token_UNIQUE` (`token`)
) ENGINE=MyISAM AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `token`
--

LOCK TABLES `token` WRITE;
/*!40000 ALTER TABLE `token` DISABLE KEYS */;
INSERT INTO `token` VALUES (36,'4fb4371df8692d10ace005dcf339f424','projectA',129),(37,'aa795812eb89cf18fa9483f7c5ca05f2','projectB',130);
/*!40000 ALTER TABLE `token` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `upload_file`
--

DROP TABLE IF EXISTS `upload_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `upload_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `file_name` varchar(60) COLLATE utf8_bin NOT NULL,
  `file_path` varchar(120) COLLATE utf8_bin NOT NULL,
  `upload_date` varchar(60) COLLATE utf8_bin NOT NULL,
  `uf_bz` varchar(6) COLLATE utf8_bin NOT NULL DEFAULT 'Y',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `upload_file`
--

LOCK TABLES `upload_file` WRITE;
/*!40000 ALTER TABLE `upload_file` DISABLE KEYS */;
/*!40000 ALTER TABLE `upload_file` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `api_key` varchar(255) DEFAULT NULL,
  `create_date` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=94 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (91,'ddser','eebd290ecc914f669f01cbecc67c01b8','9e3bca027b8ed405624c5f9edc0bce0e','2017-09-25 11:08:04'),(92,'admin','21232f297a57a5a743894a0e4a801fc3','144bf8f01ff5afc702ec34d6e20bf2a3','2017-09-25 11:08:25'),(93,'g1','e10adc3949ba59abbe56e057f20f883e','adee894b374d23b3c8e7b211ea3917c4','2018-12-28 14:23:50');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_role`
--

DROP TABLE IF EXISTS `user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_role_role_id_idx` (`role_id`)
) ENGINE=MyISAM AUTO_INCREMENT=1249 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_role`
--

LOCK TABLES `user_role` WRITE;
/*!40000 ALTER TABLE `user_role` DISABLE KEYS */;
INSERT INTO `user_role` VALUES (1240,92,130),(1246,91,129),(1239,92,129),(1131,81,1),(1238,92,100),(1186,57,1),(1237,92,1),(1245,91,100),(1244,91,1),(1247,91,130),(1248,93,100);
/*!40000 ALTER TABLE `user_role` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-12-29 12:02:04
