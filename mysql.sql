-- MySQL dump 10.13  Distrib 8.0.28, for macos11 (x86_64)
--
-- Host: localhost    Database: phoenix
-- ------------------------------------------------------
-- Server version	8.0.28

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `school`
--

DROP TABLE IF EXISTS `school`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `school` (
                          `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                          `parent_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                          `order_` bigint DEFAULT NULL,
                          `create_depart_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                          `create_depart_name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                          `create_user_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                          `create_user_code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                          `create_user_name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                          `create_at_` datetime DEFAULT NULL,
                          `update_at_` datetime DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `school`
--

LOCK TABLES `school` WRITE;
/*!40000 ALTER TABLE `school` DISABLE KEYS */;
/*!40000 ALTER TABLE `school` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `students`
--

DROP TABLE IF EXISTS `students`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `students` (
                            `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                            `code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `sex_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `birth_` date DEFAULT NULL,
                            `age_` tinyint DEFAULT NULL,
                            `province_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `pay_` decimal(13,2) DEFAULT NULL,
                            `score_` decimal(13,2) DEFAULT NULL,
                            `course_` int DEFAULT NULL,
                            `is_full_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `order_` bigint DEFAULT NULL,
                            `description_` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
                            `parent_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `create_depart_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `create_depart_name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `create_user_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `create_user_code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `create_user_name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `create_at_` datetime DEFAULT NULL,
                            `update_at_` datetime DEFAULT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `students`
--

LOCK TABLES `students` WRITE;
/*!40000 ALTER TABLE `students` DISABLE KEYS */;
INSERT INTO `students` VALUES ('111',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `students` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_auto_no`
--

DROP TABLE IF EXISTS `sys_auto_no`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_auto_no` (
                               `id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
                               `kind_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `prefix_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `value_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `create_at_` datetime DEFAULT NULL,
                               `update_at_` datetime DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `sys_auto_no_UniqueIndex_code` (`kind_id_`,`prefix_`),
                               CONSTRAINT `sys_auto_no_ForeignKey_kind_id` FOREIGN KEY (`kind_id_`) REFERENCES `sys_auto_no_kind` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_auto_no`
--

LOCK TABLES `sys_auto_no` WRITE;
/*!40000 ALTER TABLE `sys_auto_no` DISABLE KEYS */;
INSERT INTO `sys_auto_no` VALUES ('dg35yy95ddc1ndjc9abu7f6qr8no427s','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC222221128333','1','2022-11-28 14:04:41',NULL),('dgutg68oyaw75k6haxacwckgoyc6cvkd','gf4pyt7zs4jciuf3dscc2puvfwgcb6cf','CGHT-2023-','9','2023-01-04 23:11:16','2023-02-22 18:15:02'),('dhp4bbwdmassg2xn9nvbz63sz8asspbw','gf4qaxll5apogijed57ubfe22ejqzbw7','HTR20230213','7','2023-02-13 21:03:39','2023-02-13 21:13:34'),('dhp4bjajmf6oohow9zw97b57ojghzyfg','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXCAAAA230213BBB','1','2023-02-13 21:04:01',NULL),('dhp4bm8d8xyr4h5k9favjpf7vt4mvdv5','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXCaaa230213bbb','2','2023-02-13 21:04:08','2023-02-13 21:04:18'),('dhp7zy824n6dmk5t9oqcf9z1rr1bc97g','gf4qaxll5apogijed57ubfe22ejqzbw7','HTR20230214','2','2023-02-14 01:14:32','2023-02-14 01:14:39'),('dhuyq66jfkaff91va45byddfw9nnuy3f','gf4qaxll5apogijed57ubfe22ejqzbw7','HTR20230222','4','2023-02-22 18:14:44','2023-02-22 18:15:06'),('gf4wizlok4zxlbbmde5u66wusttlhzge','gf4pyt7zs4jciuf3dscc2puvfwgcb6cf','CGHT-2022-','89','2022-07-07 01:33:51','2022-08-04 21:08:58'),('gf4wj73cn5sy22okdn4up56xqplu4ykl','gf4qaxll5apogijed57ubfe22ejqzbw7','HTR20220707','15','2022-07-07 01:34:13','2022-07-07 08:40:25'),('gf4wja3gpya3k2thc5ts7i67jt7s6tkd','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC220707','6','2022-07-07 01:34:17','2022-07-07 08:40:23'),('gf7sssexyifcsddodejfc5hab2jln4t7','gf4qaxll5apogijed57ubfe22ejqzbw7','HTR20220711','5','2022-07-11 10:35:13','2022-07-11 21:33:57'),('gf7ssuq7iyluedumcgoftkndsqkbmxqc','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC220711','4','2022-07-11 10:35:23','2022-07-11 21:33:59'),('gfbjscgtxawltw6ydpita374d5ywpurf','gf4qaxll5apogijed57ubfe22ejqzbw7','HTR20220714','40','2022-07-14 01:08:50','2022-07-14 14:01:59'),('gfbjscxg2l5bxo7fc34fuwsyekvdqj4y','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC220714','3','2022-07-14 01:08:52','2022-07-14 14:52:09'),('gfnnxzqhipyzc3d5cdveybcjhbpayvza','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC220801','2','2022-08-01 10:48:30','2022-08-01 10:48:32'),('gfnofyoule6agfgbcauf3oyicgnurr6v','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC344555220801','4','2022-08-01 11:18:18','2022-08-01 11:20:37'),('gfnog4a3bfwnq4nzdeztxmdefqkaebmz','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC345220801','2','2022-08-01 11:18:32','2022-08-01 11:20:20'),('gfnogycacnyh2yqbcccf7o6ub7osqnzk','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC3456220801','1','2022-08-01 11:20:25',NULL),('gfnohcokvkiao76tciqvcfr43esrntuk','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC344555220801123','2','2022-08-01 11:21:06','2022-08-01 11:21:17'),('gfnohhpw22cmjtnoddocqw2totunljlm','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC344555220801121','1','2022-08-01 11:21:26',NULL),('gfnohuhjyoilfmezcx2uc3i4vqwqtq5e','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC12208012','13','2022-08-01 11:22:17','2022-08-01 11:33:47'),('gfnoiptskft3ctkycxyfwmb57h74qiup','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC22208013','1','2022-08-01 11:24:07',NULL),('gfnoirq6ouyjdgfvcngvxstlmwoylg4o','gf4qaxll5apogijed57ubfe22ejqzbw7','HTR20220801','11','2022-08-01 11:24:14','2022-08-01 11:33:35'),('gfnol5qvzgzhlo3ncucevjzfoz2o4ard','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC32208014','3','2022-08-01 11:29:18','2022-08-01 11:29:30'),('gfpwdscin7x7hgfhd4wtp3p7k74iw2j7','gf4qaxll5apogijed57ubfe22ejqzbw7','HTR20220804','2','2022-08-04 21:08:49','2022-08-04 21:08:51'),('gfpwdxz4txbv2ag6cg3cvbrkx6ecnmfe','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','LXC123220804456','2','2022-08-04 21:09:11','2022-08-04 21:09:36'),('ggls7m4mfdyrhzxdcurcmos6uxpv3krg','gf4qaxll5apogijed57ubfe22ejqzbw7','','4','2022-09-16 03:53:44','2022-09-16 03:53:50'),('ggls7qtjmhvllntbcwmvpab3xapgjf5c','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','','6','2022-09-16 03:54:03','2022-11-10 10:13:45');
/*!40000 ALTER TABLE `sys_auto_no` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_auto_no_item`
--

DROP TABLE IF EXISTS `sys_auto_no_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_auto_no_item` (
                                    `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                    `kind_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `value_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                                    `order_` bigint DEFAULT NULL,
                                    `create_at_` datetime DEFAULT NULL,
                                    `update_at_` datetime DEFAULT NULL,
                                    PRIMARY KEY (`id`),
                                    KEY `sys_auto_no_item_ForeignKey_kind_id` (`kind_id_`),
                                    CONSTRAINT `sys_auto_no_item_ForeignKey_kind_id` FOREIGN KEY (`kind_id_`) REFERENCES `sys_auto_no_kind` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_auto_no_item`
--

LOCK TABLES `sys_auto_no_item` WRITE;
/*!40000 ALTER TABLE `sys_auto_no_item` DISABLE KEYS */;
INSERT INTO `sys_auto_no_item` VALUES ('gf4pyt5wxt2tdw4mcj4uh7koxor3rms2','gf4pyt7zs4jciuf3dscc2puvfwgcb6cf','STRING','CGHT',1657102180142897,'2022-07-06 18:09:40','2022-07-07 01:32:49'),('gf4pyz7n2uopkhfldgvt554xpkmwupf2','gf4pyt7zs4jciuf3dscc2puvfwgcb6cf','DATETIME','2006',1657102204718234,'2022-07-06 18:10:05','2022-07-07 01:33:30'),('gf4pzc4himmijqwzdbjszy2d6epdjhlo','gf4pyt7zs4jciuf3dscc2puvfwgcb6cf','SEQ','0001',1657102240390706,'2022-07-06 18:10:40','2022-07-06 18:28:04'),('gf4qaxmwyb4yixdycuzedakar6ct52iv','gf4qaxll5apogijed57ubfe22ejqzbw7','STRING','HTR',1657103222424960,'2022-07-06 18:27:02','2022-07-06 18:28:22'),('gf4qbopfdeyyhbe5doqfggv52g2dhoks','gf4qaxll5apogijed57ubfe22ejqzbw7','DATETIME','20060102',1657103314824670,'2022-07-06 18:28:35',NULL),('gf4qbqplvdj5voybdjldqe63zi5pfkpd','gf4qaxll5apogijed57ubfe22ejqzbw7','SEQ','0001',1657103322356627,'2022-07-06 18:28:42','2022-07-06 18:29:42'),('gf4w6nncd2znffxlcdssyhevt7jedrqe','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','STRING','LXC',1657127502951800,'2022-07-07 01:11:43','2022-07-07 01:11:52'),('gf4w6v7ncanm6ltrdp7dkngwkjo5rsod','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','DATETIME','060102',1657127525343773,'2022-07-07 01:12:13','2022-07-07 01:13:26'),('gf4w6wsnrdypbnkzdxovyledb4osb2d6','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','SEQ','0001',1657127539556660,'2022-07-07 01:12:20','2022-08-01 11:30:10'),('gf4wibj2zfzq6q6zdvnfbxxljy3ktfab','gf4pyt7zs4jciuf3dscc2puvfwgcb6cf','STRING','-',1657102240388657,'2022-07-07 01:32:13','2022-07-07 01:32:16'),('gf4wihvudgoqc4amck5uxwknwpzqnt4m','gf4pyt7zs4jciuf3dscc2puvfwgcb6cf','STRING','-',1657102204718233,'2022-07-07 01:32:40','2022-07-07 01:32:42'),('gfnnxlkrykeuowsbcpdf4yvkltvagznf','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','VALUES','xxxx',1657127517879782,'2022-08-01 10:47:34','2022-08-01 11:20:15'),('gfnoh5c76h7pesd7cnrvfbglv53mcbnn','gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','VALUES','yyyy',1657127536182212,'2022-08-01 11:20:45',NULL);
/*!40000 ALTER TABLE `sys_auto_no_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_auto_no_kind`
--

DROP TABLE IF EXISTS `sys_auto_no_kind`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_auto_no_kind` (
                                    `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                    `code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `description_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `order_` bigint DEFAULT NULL,
                                    `create_at_` datetime DEFAULT NULL,
                                    `update_at_` datetime DEFAULT NULL,
                                    PRIMARY KEY (`id`),
                                    UNIQUE KEY `sys_auto_no_kind_UniqueIndex_code` (`code_`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_auto_no_kind`
--

LOCK TABLES `sys_auto_no_kind` WRITE;
/*!40000 ALTER TABLE `sys_auto_no_kind` DISABLE KEYS */;
INSERT INTO `sys_auto_no_kind` VALUES ('gf4pyt7zs4jciuf3dscc2puvfwgcb6cf','htbh','采购合同编号',NULL,1669616400649004992,'2022-07-06 18:09:40','2022-07-07 01:32:57'),('gf4qaxll5apogijed57ubfe22ejqzbw7','rkdbh','合同入库单号',NULL,1669616400649006016,'2022-07-06 18:27:02','2022-07-07 01:11:27'),('gf4w6nry7ql6umrccfzu2xy4jqm3nwyu','ckdbh','零星出库单号',NULL,1669616400649007040,'2022-07-07 01:11:43','2022-07-07 01:12:05');
/*!40000 ALTER TABLE `sys_auto_no_kind` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_data_service`
--

DROP TABLE IF EXISTS `sys_data_service`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_data_service` (
                                    `id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
                                    `table_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `method_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `code_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `timeout_` int DEFAULT NULL,
                                    `source_` text COLLATE utf8mb4_general_ci,
                                    `order_` bigint DEFAULT NULL,
                                    `create_at_` datetime DEFAULT NULL,
                                    `update_at_` datetime DEFAULT NULL,
                                    PRIMARY KEY (`id`),
                                    UNIQUE KEY `sys_data_service_UniqueIndex_table_id_method_code` (`table_id_`,`method_`,`code_`),
                                    CONSTRAINT `sys_data_service_ForeignKey_table_id` FOREIGN KEY (`table_id_`) REFERENCES `sys_table` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_data_service`
--

LOCK TABLES `sys_data_service` WRITE;
/*!40000 ALTER TABLE `sys_data_service` DISABLE KEYS */;
INSERT INTO `sys_data_service` VALUES ('dhp5dmtk9fpt8a82amxca8rkpr6oj9jp','gfahporvfpiepzledb3ur2zgi5a4q46r','GET','query','查询服务(检索)',0,'/* 查询实现：按字段查询(=)、排序、分页、过滤(LIKE) 等常见需求 */\nsql.Select(\"SELECT * \", \"FROM students\");',1676297803526152000,'2023-02-13 22:16:43',NULL),('dhp5dmycbvnkgy5y9g1sm6mj6u97js78','gfahporvfpiepzledb3ur2zgi5a4q46r','POST','save','保存服务(新增 修改 删除 排序)',0,'/* 保存实现：插入、更新、删除、排序 等操作*/\nsql.Save(\"students\");',1676297803526910000,'2023-02-13 22:16:43','2023-02-13 22:21:40'),('dhp5dn1pj2jsby459cwb4b119dxon2rs','gfwuk2jlr4jt4zayd5ouodg2xepfufit','POST','save','保存服务(新增 修改 删除 排序)',0,'/* 保存实现：插入、更新、删除、排序 等操作*/\nsql.Save(\"tests\");',1676297804667675000,'2023-02-13 22:16:44',NULL),('dhp5dn3cfghsypgna11a7w6myk3jrk3j','gfpygrr37uwlrc3kdnjejp3obxgp2u3r','POST','save','保存服务(新增 修改 删除 排序)',0,'/* 保存实现：插入、更新、删除、排序 等操作*/\nsql.Save(\"school\");',1676297804116473000,'2023-02-13 22:16:44',NULL),('dhp5dn3tkgzqcuzd9xjamutvc9gxuwz8','gfpygrr37uwlrc3kdnjejp3obxgp2u3r','GET','query','查询服务(检索)',0,'/* 查询实现：按字段查询(=)、排序、分页、过滤(LIKE) 等常见需求 */\nsql.Select(\"SELECT * \", \"FROM school\");',1676297804115259000,'2023-02-13 22:16:44',NULL),('dhp5dn61v8ny7s1cavysycpd5f6pqpxw','gfwuk2jlr4jt4zayd5ouodg2xepfufit','GET','query','查询服务(检索)',0,'/* 查询实现：按字段查询(=)、排序、分页、过滤(LIKE) 等常见需求 */\nsql.Select(\"SELECT * \", \"FROM tests\");',1676297804666293000,'2023-02-13 22:16:44',NULL),('dhp5z3tt7fd97kfj9gtb49um8cwajvgo','gfwuk2jlr4jt4zayd5ouodg2xepfufit','GET','simple','测试自定义查询服务',0,'sql.Select(\"select id, varchar_256_, parent_id_ \",\"from tests\");',1676300171235299000,'2023-02-13 22:56:11','2023-02-13 23:06:20');
/*!40000 ALTER TABLE `sys_data_service` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_depart`
--

DROP TABLE IF EXISTS `sys_depart`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_depart` (
                              `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                              `code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                              `name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                              `parent_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                              `description_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                              `valid_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                              `order_` bigint DEFAULT NULL,
                              `create_at_` datetime DEFAULT NULL,
                              `update_at_` datetime DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `sys_user_UniqueIndex_code` (`code_`),
                              KEY `sys_depart_ForeignKey_parent_id` (`parent_id_`),
                              CONSTRAINT `sys_depart_ForeignKey_parent_id` FOREIGN KEY (`parent_id_`) REFERENCES `sys_depart` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_depart`
--

LOCK TABLES `sys_depart` WRITE;
/*!40000 ALTER TABLE `sys_depart` DISABLE KEYS */;
INSERT INTO `sys_depart` VALUES ('D000','00','公司办公室',NULL,'','Effective',1654622106198310,'2023-02-15 01:27:16','2023-02-15 01:32:11'),('D001','03','物资部',NULL,'','Effective',1663916515050275000,'2022-06-08 01:12:07','2023-02-15 01:32:05'),('D002','02','生产制造部',NULL,'','Effective',1654622106201638,'2022-06-08 01:12:10','2023-02-15 01:30:13'),('D003','01','财务部',NULL,'','Effective',1654622106199334,'2022-06-08 01:12:14','2023-02-15 01:49:59'),('D004','05','保卫处',NULL,'666','Effective',1675994945236115000,'2022-06-08 01:12:18','2023-02-15 01:29:47'),('D005','04','后勤保障部',NULL,'444','Effective',1675994930715731000,'2022-06-08 01:12:23','2023-02-10 10:09:23'),('D011','0201','管理分部','D002','','Effective',1654622106204454,'2022-06-08 01:15:06','2023-02-15 01:30:57'),('D012','0202','保障分部','D002','','Effective',1654622118669412,'2022-06-08 01:15:18','2023-02-15 01:31:02'),('D013','0203','电装车间','D002','','Effective',1654622125585183,'2022-06-08 01:15:25','2023-02-15 01:31:07'),('D014','0204','精加车间','D002','','Effective',1654622132822165,'2022-06-08 01:15:32','2023-02-15 01:31:11'),('D021','0301','计划组','D001','','Effective',1654622163889488,'2022-06-08 01:16:03','2023-02-15 01:31:18'),('D022','0302','仓库组','D001','','Effective',1654622168520294,'2022-06-08 01:16:08','2023-02-15 01:31:23'),('D023','0303','下料组','D001','','Effective',1654622199535767,'2022-06-08 01:16:39','2023-02-15 01:31:28'),('D031','0103','会计组','D003','','Effective',1654622221718356,'2022-06-08 01:17:01','2023-02-15 01:30:50'),('D032','0102','出纳组','D003','','Effective',1654622221718354,'2022-06-08 01:17:42','2023-02-15 01:30:44'),('D033','0101','审计组','D003','','Effective',1654622221718353,'2022-06-08 01:17:53','2023-02-15 01:30:39');
/*!40000 ALTER TABLE `sys_depart` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dict_item`
--

DROP TABLE IF EXISTS `sys_dict_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dict_item` (
                                 `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                 `kind_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                 `code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                 `name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                 `order_` bigint DEFAULT NULL,
                                 `create_at_` datetime DEFAULT NULL,
                                 `update_at_` datetime DEFAULT NULL,
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `sys_dict_item_UniqueIndex_kind_id_code` (`kind_id_`,`code_`),
                                 CONSTRAINT `sys_dict_item_ForeignKey_kind_id` FOREIGN KEY (`kind_id_`) REFERENCES `sys_dict_kind` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict_item`
--

LOCK TABLES `sys_dict_item` WRITE;
/*!40000 ALTER TABLE `sys_dict_item` DISABLE KEYS */;
INSERT INTO `sys_dict_item` VALUES ('mkzhcpktplkmc2ceisyyecjtj2lqswm1','K000','Unknown','未知',3333,'2022-01-01 01:01:01','2022-07-06 15:58:11'),('mkzhcpktplkmc2ceisyyecjtj2lqswm2','K000','Male','男',1111,'2022-01-01 01:01:01','2022-07-06 15:58:12'),('mkzhcpktplkmc2ceisyyecjtj2lqswm3','K000','Female','女',2222,'2022-01-01 01:01:01','2022-07-06 15:58:11'),('mkzhcpktplkmc2ceisyyecjtj2lqswm4','K002','Effective','生效中',4444,'2022-01-01 01:01:01','2022-06-30 23:11:02'),('mkzhcpktplkmc2ceisyyecjtj2lqswm5','K002','Disable','未启用',5555,'2022-01-01 01:01:01',NULL),('mkzhcpktplkmc2ceisyyecjtj2lqswm6','K002','Locked','已锁定',6666,'2022-01-01 01:01:01',NULL),('mkzhcpktplkmc2ceisyyecjtj2lqswm7','K003','3','机密',1657090065899362,'2022-01-01 01:01:01',NULL),('mkzhcpktplkmc2ceisyyecjtj2lqswm8','K003','1','内部',1657090030884295,'2022-01-01 01:01:01',NULL),('mkzhcpktplkmc2ceisyyecjtj2lqswm9','K003','0','非密',7777,'2022-01-01 01:01:01',NULL),('mkzhcpktplkmc2ceisyyecjtj2lqswmi','K005','Effective','生效中',1655861565702070,'2022-06-22 09:32:46','2022-06-30 23:11:08'),('mkzhcpktplkmc2ceisyyecjtj2lqswn1','K003','2','秘密',1657090065899361,'2022-01-01 01:01:01',NULL),('mkzhcpktplkmc2ceisyyecjtj2lqswn2','K004','1','内部',10002,'2022-01-01 01:01:01',NULL),('mkzhcpktplkmc2ceisyyecjtj2lqswn3','K004','3','机密',1655797887299448,'2022-01-01 01:01:01',NULL),('mkzhcpktplkmc2ceisyyecjtj2lqswn4','K004','0','非密',10001,'2022-01-01 01:01:01',NULL),('mkzhcpktplkmc2ceisyyecjtj2lqswn5','K004','2','秘密',10004,'2022-01-01 01:01:01',NULL),('mkzhdhd6dd7e7rvcjxu37fruc6edzgr2','K005','Disable','已停用',1655861660575229,'2022-06-22 09:34:21',NULL),('mkziywa4x63vzhxpiyczzm5svilzkmu3','K006','Effective','生效中',1655868504736435,'2022-06-22 11:28:25','2022-06-30 23:11:11'),('mkzizkasuzq5tkmhjqs2peiifsv6qxhw','K006','Disable','已停用',1655868584415429,'2022-06-22 11:29:44',NULL);
/*!40000 ALTER TABLE `sys_dict_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dict_kind`
--

DROP TABLE IF EXISTS `sys_dict_kind`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dict_kind` (
                                 `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                 `code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                 `name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                 `description_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                 `order_` bigint DEFAULT NULL,
                                 `create_at_` datetime DEFAULT NULL,
                                 `update_at_` datetime DEFAULT NULL,
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `sys_dict_kind_UniqueIndex_code` (`code_`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict_kind`
--

LOCK TABLES `sys_dict_kind` WRITE;
/*!40000 ALTER TABLE `sys_dict_kind` DISABLE KEYS */;
INSERT INTO `sys_dict_kind` VALUES ('K000','user_sex','性别','',-3747,'2022-01-01 01:01:01','2023-02-13 20:34:14'),('K002','user_valid','用户状态','',-3491,'2022-01-01 01:01:01',NULL),('K003','user_classification','用户密级','',1669618503759354000,'2022-01-01 01:01:01',NULL),('K004','file_classification','文件密级','',1669618501677662000,'2022-01-01 01:01:01','2023-02-13 20:34:31'),('K005','depart_valid','部门状态','',-4003,'2022-06-22 09:32:46',NULL),('K006','menu_valid','菜单状态','',-4515,'2022-06-22 11:28:25',NULL);
/*!40000 ALTER TABLE `sys_dict_kind` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_doc`
--

DROP TABLE IF EXISTS `sys_doc`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_doc` (
                           `id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
                           `name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `size_` int DEFAULT NULL,
                           `mime_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `dir_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `order_` bigint DEFAULT NULL,
                           `create_at_` datetime DEFAULT NULL,
                           `user_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `user_code_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `user_name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `depart_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `depart_code_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `depart_name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_doc`
--

LOCK TABLES `sys_doc` WRITE;
/*!40000 ALTER TABLE `sys_doc` DISABLE KEYS */;
INSERT INTO `sys_doc` VALUES ('dgxvkfzyqd9kcb8f97fuajbgx8g4shtg','JZC-004外协管理系统升级解决方案_20200522.doc',3230720,'application/msword','store/upload/2301/1e/37',1673246903586008000,'2023-01-09 14:48:23','U002','admin','系统管理员','D022',NULL,'仓库组'),('dgxvkn55gc4mx6o2arbcw4m4autpccwu','测试委托单投产结果_20200304.xlsx',11764,'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet','store/upload/2301/af/d9',1673246924870646000,'2023-01-09 14:48:44','U002','admin','系统管理员','D022',NULL,'仓库组'),('dhp93rjrvjyykka7a6hvpja6b2x4gvq6','DSC_0029.JPG',4386668,'image/jpeg','store/upload/2302/39/7d',1676312926745374000,'2023-02-14 02:28:46','U002','admin','系统管理员','D022','2002','仓库组'),('dhp96atn5y6151dxayca3jv97jqwafsx','DSC_0036.JPG',4292333,'image/jpeg','store/upload/2302/88/a5',1676313255089960000,'2023-02-14 02:34:15','U002','admin','系统管理员','D022','2002','仓库组'),('dhp97jjoz7faptpq9amvahnrgt9n9fuk','DSC_0033.JPG',4314435,'image/jpeg','store/upload/2302/54/d3',1676313410221565000,'2023-02-14 02:36:50','U002','admin','系统管理员','D022','2002','仓库组'),('dhp97km8g725ha7k9pqvrgrhnpmnttvo','DSC_0029.JPG',4386668,'image/jpeg','store/upload/2302/09/1a',1676313414977651000,'2023-02-14 02:36:54','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9bc7ba25xt75x9m8ckms5cnhsu579','DSC_0036.JPG',4292333,'image/jpeg','store/upload/2302/60/36',1676313900473260000,'2023-02-14 02:45:00','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9bq9uyc2xpk619k4ubb2fkpb5akpm','DSC_0058.JPG',4345695,'image/jpeg','store/upload/2302/d1/3d',1676313945146196000,'2023-02-14 02:45:45','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9cm96a6mx2prtandtuyaqvm54am9g','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/d3/c9',1676314057842565000,'2023-02-14 02:47:37','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9d699zzzm8s1famjta11wjv4xwd6f','DSC_0036.JPG',4292333,'image/jpeg','store/upload/2302/d1/4c',1676314133541106000,'2023-02-14 02:48:53','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9d6b8nfpwngxtab9ayt2tsjva8843','DSC_0038.JPG',4252686,'image/jpeg','store/upload/2302/01/b2',1676314133556180000,'2023-02-14 02:48:53','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9d6bf67bn4t7aa47cnvmwfctwchm5','DSC_0034.JPG',4794122,'image/jpeg','store/upload/2302/f6/e2',1676314133533998000,'2023-02-14 02:48:53','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9d6dudvrz97y7amuuyb55kxp88djf','DSC_0040.JPG',4369973,'image/jpeg','store/upload/2302/53/e9',1676314133553724000,'2023-02-14 02:48:53','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9dq3281nh7oo3apcu783c6voju2s6','DSC_0038.JPG',4252686,'image/jpeg','store/upload/2302/7e/5b',1676314200737220000,'2023-02-14 02:50:00','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9dttkvk6mf35haaav2fj5akx5qn5g','DSC_0066.JPG',3097217,'image/jpeg','store/upload/2302/0d/73',1676314215167270000,'2023-02-14 02:50:15','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9fku8vtt6c5mqampuoj4ckqg6v3u5','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/65/b3',1676314311650167000,'2023-02-14 02:51:51','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9ggtd19t8n1uj9qdu4jhj59w6wntc','DSC_0040.JPG',4369973,'image/jpeg','store/upload/2302/c1/8d',1676314427916697000,'2023-02-14 02:53:47','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9gqmoufozx5hg9y1boaraqs6o4jah','DSC_0036.JPG',4292333,'image/jpeg','store/upload/2302/ab/7e',1676314458898740000,'2023-02-14 02:54:18','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9h744x1h9bd4g9quuvzwy94645uxf','DSC_0041.JPG',3461628,'image/jpeg','store/upload/2302/7b/b4',1676314520379955000,'2023-02-14 02:55:20','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9j6ppadwq1r4w9x6tbwyo9pfavpcj','DSC_0040.JPG',4369973,'image/jpeg','store/upload/2302/86/94',1676314646282638000,'2023-02-14 02:57:26','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9k5tnq7bum29f9tqcgczcaq8qu445','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/da/70',1676314771963514000,'2023-02-14 02:59:31','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9kx2q2jz7gc149a9csu52m3bxrxrc','DSC_0041.JPG',3461628,'image/jpeg','store/upload/2302/68/7e',1676314868093710000,'2023-02-14 03:01:08','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9o12yuafs2uu39n6bnu94kcvtczan','DSC_0040.JPG',4369973,'image/jpeg','store/upload/2302/be/68',1676315136711121000,'2023-02-14 03:05:36','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9oj3tumjfdz83az9vd8q8mjyszx8y','DSC_0050.JPG',4164901,'image/jpeg','store/upload/2302/a5/b5',1676315200707540000,'2023-02-14 03:06:40','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9or8cu3hkm4mx9d8bgz7y2uokwbqt','DSC_0052.JPG',3842114,'image/jpeg','store/upload/2302/fa/69',1676315228477348000,'2023-02-14 03:07:08','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9pk6dar77orjkaawak16ppwrgqawu','DSC_0050.JPG',4164901,'image/jpeg','store/upload/2302/ad/77',1676315332832677000,'2023-02-14 03:08:52','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9r8j4u4yrgy8ya8ouffuso92qjunf','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/2a/f6',1676315550030545000,'2023-02-14 03:12:30','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9rrsv611hccdnaq7s4hdq31km6m28','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/9d/c9',1676315615751204000,'2023-02-14 03:13:35','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9s2pg4c7mc949a5ybogv9h9cvgsdz','DSC_0042.JPG',3546025,'image/jpeg','store/upload/2302/74/b6',1676315654743721000,'2023-02-14 03:14:14','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9sugcvsx6y6t79xmapc2h64pqzs67','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/d3/46',1676315753878024000,'2023-02-14 03:15:53','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9t2uggathuoh3aa6va7aqd96psyzr','DSC_0050.JPG',4164901,'image/jpeg','store/upload/2302/1e/ab',1676315783136076000,'2023-02-14 03:16:23','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9vshb3mqad5xw9cuuocuskzkgm51d','DSC_0049.JPG',4809263,'image/jpeg','store/upload/2302/87/18',1676316129348604000,'2023-02-14 03:22:09','U002','admin','系统管理员','D022','2002','仓库组'),('dhp9w33x8f7fruyga94u5uyzcmanqxjd','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/bf/c2',1676316168190661000,'2023-02-14 03:22:48','U002','admin','系统管理员','D022','2002','仓库组'),('dhpogtam8qt5vpnkadwvzd2cacb6j5n9','DSC_0040.JPG',4369973,'image/jpeg','store/upload/2302/ab/74',1676363621545248000,'2023-02-14 16:33:41','U002','admin','系统管理员','D022','2002','仓库组'),('dhpoh4zy33nnny2kafxs9vvdpn476h6y','DSC_0040.JPG',4369973,'image/jpeg','store/upload/2302/fc/17',1676363663689049000,'2023-02-14 16:34:23','U002','admin','系统管理员','D022','2002','仓库组'),('dhpohr2j587fmth6a87tswm39ygroy97','DSC_0050.JPG',4164901,'image/jpeg','store/upload/2302/62/72',1676363740984393000,'2023-02-14 16:35:40','U002','admin','系统管理员','D022','2002','仓库组'),('dhpohuqc9rzv81qx9mgsphd9pctak7rm','DSC_0042.JPG',3546025,'image/jpeg','store/upload/2302/d1/1f',1676363754589846000,'2023-02-14 16:35:54','U002','admin','系统管理员','D022','2002','仓库组'),('dhpokkr48444u3gk92zca2r58avhv96h','DSC_0036.JPG',4292333,'image/jpeg','store/upload/2302/25/c0',1676363974619103000,'2023-02-14 16:39:34','U002','admin','系统管理员','D022','2002','仓库组'),('dhpon1o5fy6vjnmz94x9tthzk5cmpomo','DSC_0052.JPG',3842114,'image/jpeg','store/upload/2302/c8/e5',1676364162889699000,'2023-02-14 16:42:42','U002','admin','系统管理员','D022','2002','仓库组'),('dhpoo69ubcu8td2x9tnsgaf44o1q5czo','DSC_0048.JPG',4560846,'image/jpeg','store/upload/2302/a2/4f',1676364309787561000,'2023-02-14 16:45:09','U002','admin','系统管理员','D022','2002','仓库组'),('dhpop4r5msvbhncbaayt681f4cgh5uo3','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/ac/47',1676364430731627000,'2023-02-14 16:47:10','U002','admin','系统管理员','D022','2002','仓库组'),('dhpos55z43n3572kamfags5ur7nxtxkd','DSC_0041.JPG',3461628,'image/jpeg','store/upload/2302/45/fd',1676364816453476000,'2023-02-14 16:53:36','U002','admin','系统管理员','D022','2002','仓库组'),('dhposaj1xo864nkp9oyv4ttb29mbkof6','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/90/59',1676364838447027000,'2023-02-14 16:53:58','U002','admin','系统管理员','D022','2002','仓库组'),('dhpoudfd1p9g918raxosfv1sq72vvvyu','DSC_0041.JPG',3461628,'image/jpeg','store/upload/2302/92/51',1676365105460975000,'2023-02-14 16:58:25','U002','admin','系统管理员','D022','2002','仓库组'),('dhpowhpko6oaqfhua3v9vajcmfzrod1d','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/2a/84',1676365374881311000,'2023-02-14 17:02:54','U002','admin','系统管理员','D022','2002','仓库组'),('dhpowsp3c52uwt9u986chuuqqxodv157','DSC_0048.JPG',4560846,'image/jpeg','store/upload/2302/46/1b',1676365410847149000,'2023-02-14 17:03:30','U002','admin','系统管理员','D022','2002','仓库组'),('dhpoxfogahr3axc5afruf6o3zjws5tdx','DSC_0041.JPG',3461628,'image/jpeg','store/upload/2302/c4/37',1676365494819586000,'2023-02-14 17:04:54','U002','admin','系统管理员','D022','2002','仓库组'),('dhpoxvodz2mzkmwa9whs7p1nmtx7pz9o','DSC_0042.JPG',3546025,'image/jpeg','store/upload/2302/f0/bd',1676365550813576000,'2023-02-14 17:05:50','U002','admin','系统管理员','D022','2002','仓库组'),('dhpoycnpy7zxn6q69z7bvc2jvv6uac34','DSC_0041.JPG',3461628,'image/jpeg','store/upload/2302/8f/de',1676365614872041000,'2023-02-14 17:06:54','U002','admin','系统管理员','D022','2002','仓库组'),('dhpoynxngdk796svagyux3gp7b9cj1x2','DSC_0042.JPG',3546025,'image/jpeg','store/upload/2302/44/a4',1676365647261404000,'2023-02-14 17:07:27','U002','admin','系统管理员','D022','2002','仓库组'),('dhpozoarsjdrotc5915998o5u35zx4p9','DSC_0042.JPG',3546025,'image/jpeg','store/upload/2302/9d/df',1676365777352573000,'2023-02-14 17:09:37','U002','admin','系统管理员','D022','2002','仓库组'),('dhpp1ujj9a5gmhsh9bys3j67v7bjjyc7','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/21/cb',1676365930133197000,'2023-02-14 17:12:10','U002','admin','系统管理员','D022','2002','仓库组'),('dhpp33yjzv6xdqx6a6tbbaohwsygrnrd','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/6e/ec',1676366091871234000,'2023-02-14 17:14:51','U002','admin','系统管理员','D022','2002','仓库组'),('dhpp3pkfwbs3cdn2arxb6zurj7aa62x1','DSC_0036.JPG',4292333,'image/jpeg','store/upload/2302/40/36',1676366166124816000,'2023-02-14 17:16:06','U002','admin','系统管理员','D022','2002','仓库组'),('dhpp4wnh5o8j3jru964v66xbn9c15fg6','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/34/66',1676366322562276000,'2023-02-14 17:18:42','U002','admin','系统管理员','D022','2002','仓库组'),('dhpp66udc12y9fz4axa91uvrq7ybh169','DSC_0050.JPG',4164901,'image/jpeg','store/upload/2302/19/96',1676366487546088000,'2023-02-14 17:21:27','U002','admin','系统管理员','D022','2002','仓库组'),('dhpp6s2yvmdmyucv9dkvkfpk5qazq4s1','DSC_0050.JPG',4164901,'image/jpeg','store/upload/2302/78/07',1676366560870495000,'2023-02-14 17:22:40','U002','admin','系统管理员','D022','2002','仓库组'),('dhpp81tn4tdfauk9991vmrzzmkvp3hfj','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/9e/8d',1676366723092286000,'2023-02-14 17:25:23','U002','admin','系统管理员','D022','2002','仓库组'),('dhpp86uwv638bq17915sckfoao3stfjo','DSC_0048.JPG',4560846,'image/jpeg','store/upload/2302/dc/cf',1676366743618583000,'2023-02-14 17:25:43','U002','admin','系统管理员','D022','2002','仓库组'),('dhpp8nj96n9h6qr2ayhsoa9sbnn8jcw9','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/83/20',1676366798762050000,'2023-02-14 17:26:38','U002','admin','系统管理员','D022','2002','仓库组'),('dhpp9oh1a9k23atc9gmcb7cv1x7jro94','DSC_0040.JPG',4369973,'image/jpeg','store/upload/2302/43/c0',1676366929864794000,'2023-02-14 17:28:49','U002','admin','系统管理员','D022','2002','仓库组'),('dhpp9spgyvyusozz9z2cugnkfpu7vu9q','DSC_0040.JPG',4369973,'image/jpeg','store/upload/2302/73/fe',1676366946192891000,'2023-02-14 17:29:06','U002','admin','系统管理员','D022','2002','仓库组'),('dhpp9ypcc81oo9g49r5b9d61xuxa6vsv','DSC_0042.JPG',3546025,'image/jpeg','store/upload/2302/4b/3e',1676366970523401000,'2023-02-14 17:29:30','U002','admin','系统管理员','D022','2002','仓库组'),('dhppfwqkh837om7n9oabwhbwjgbmqf3s','DSC_0048.JPG',4560846,'image/jpeg','store/upload/2302/ce/46',1676367602775453000,'2023-02-14 17:40:02','U002','admin','系统管理员','D022','2002','仓库组'),('dhppgbdz44qsd99n9zs9jmc63d8f7cvs','DSC_0038.JPG',4252686,'image/jpeg','store/upload/2302/4a/94',1676367657542784000,'2023-02-14 17:40:57','U002','admin','系统管理员','D022','2002','仓库组'),('dhppo7k1cmoo12fxayotsjydoytxcxaq','DSC_0042.JPG',3546025,'image/jpeg','store/upload/2302/db/9a',1676368410844134000,'2023-02-14 17:53:30','U002','admin','系统管理员','D022','2002','仓库组'),('dhpppcmmq9cjtk8raw5s98z71uwwtttr','DSC_0050.JPG',4164901,'image/jpeg','store/upload/2302/28/57',1676368558299886000,'2023-02-14 17:55:58','U002','admin','系统管理员','D022','2002','仓库组'),('dhppr7jhjq3pu7cbajz9wf1mqqpkx2qk','DSC_0042.JPG',3546025,'image/jpeg','store/upload/2302/21/c1',1676368794126304000,'2023-02-14 17:59:54','U002','admin','系统管理员','D022','2002','仓库组'),('dhpprwyyr9ab9s5h9hgsmdg5phvda2vg','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/cb/96',1676368883563114000,'2023-02-14 18:01:23','U002','admin','系统管理员','D022','2002','仓库组'),('dhpps2xypyh2hsq5ad5bad35qf7frujs','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/aa/0b',1676368903848622000,'2023-02-14 18:01:43','U002','admin','系统管理员','D022','2002','仓库组'),('dhppsr2unuwg79xu92dc4ran3pgzdd38','DSC_0041.JPG',3461628,'image/jpeg','store/upload/2302/0f/ed',1676368988271482000,'2023-02-14 18:03:08','U002','admin','系统管理员','D022','2002','仓库组'),('dhpptpp3v7yrxgbf9rguww9qvpoucucr','DSC_0050.JPG',4164901,'image/jpeg','store/upload/2302/6a/bd',1676369110790786000,'2023-02-14 18:05:10','U002','admin','系统管理员','D022','2002','仓库组'),('dhppuo465mu5w6c89hx93f4r8u3ggq9q','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/65/c2',1676369232890247000,'2023-02-14 18:07:12','U002','admin','系统管理员','D022','2002','仓库组'),('dhppvrn2wrhp1rqs91uscyyxzk2mn1v2','DSC_0042.JPG',3546025,'image/jpeg','store/upload/2302/59/e2',1676369374008272000,'2023-02-14 18:09:34','U002','admin','系统管理员','D022','2002','仓库组'),('dhppwar2gcx1dxpv9wgb7yjkbu7bpako','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/c5/f3',1676369446085138000,'2023-02-14 18:10:46','U002','admin','系统管理员','D022','2002','仓库组'),('dhppwxh79bttfqjy93gc5z3cn7gcxfqj','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/c1/28',1676369525149929000,'2023-02-14 18:12:05','U002','admin','系统管理员','D022','2002','仓库组'),('dhppwzhrmtqdfmq6akmtkp83yy53hv6q','DSC_0050.JPG',4164901,'image/jpeg','store/upload/2302/2d/1b',1676369533778933000,'2023-02-14 18:12:13','U002','admin','系统管理员','D022','2002','仓库组'),('dhppxd3tythqx1naaqctt3nbovoujx4m','DSC_0048.JPG',4560846,'image/jpeg','store/upload/2302/9d/e3',1676369584209292000,'2023-02-14 18:13:04','U002','admin','系统管理员','D022','2002','仓库组'),('dhppxkwnpqsfwomman9sdtgpxwy8zyn4','DSC_0049.JPG',4809263,'image/jpeg','store/upload/2302/a1/de',1676369607878293000,'2023-02-14 18:13:27','U002','admin','系统管理员','D022','2002','仓库组'),('dhpq1m9th82a2wx8ark99cdykrhjw7wv','DSC_0049.JPG',4809263,'image/jpeg','store/upload/2302/4f/e0',1676369993475651000,'2023-02-14 18:19:53','U002','admin','系统管理员','D022','2002','仓库组'),('dhpq28yp49hahu3ta1nvhgyf16ufgjs2','DSC_0050.JPG',4164901,'image/jpeg','store/upload/2302/6d/fa',1676370079731987000,'2023-02-14 18:21:19','U002','admin','系统管理员','D022','2002','仓库组'),('dhpq2p3hsmy8uhtbah193aumvm7221rz','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/b5/94',1676370132205450000,'2023-02-14 18:22:12','U002','admin','系统管理员','D022','2002','仓库组'),('dhpq2t6mqnynotsqawvsdkthadysqmrf','DSC_0049.JPG',4809263,'image/jpeg','store/upload/2302/69/b2',1676370148486193000,'2023-02-14 18:22:28','U002','admin','系统管理员','D022','2002','仓库组'),('dhpq2z19gn9jamgmapb9rhrgypsms9jz','DSC_0049.JPG',4809263,'image/jpeg','store/upload/2302/3c/cc',1676370172427551000,'2023-02-14 18:22:52','U002','admin','系统管理员','D022','2002','仓库组'),('dhpq3mrqm4fs4xbx9u6tjap1wm7tyovy','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/68/75',1676370250928176000,'2023-02-14 18:24:10','U002','admin','系统管理员','D022','2002','仓库组'),('dhpq492s87ddv9jq92y9w49p2jvk1gvh','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/54/e0',1676370336728009000,'2023-02-14 18:25:36','U002','admin','系统管理员','D022','2002','仓库组'),('dhpq76g3225fd18wav1b39k87qn8g239','DSC_0042.JPG',3546025,'image/jpeg','store/upload/2302/3f/5c',1676370709337891000,'2023-02-14 18:31:49','U002','admin','系统管理员','D022','2002','仓库组'),('dhpq85amsqkw1vnn9hntkyfwp1c4za3z','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/cd/ea',1676370833276320000,'2023-02-14 18:33:53','U002','admin','系统管理员','D022','2002','仓库组'),('dhpq8kwgftwprz3w919bhscfhav1rhuh','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/88/3d',1676370887027057000,'2023-02-14 18:34:47','U002','admin','系统管理员','D022','2002','仓库组'),('dhpq9f9x1kr3gp8d9sxcncyfq3bgssbk','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/f2/ea',1676370997179503000,'2023-02-14 18:36:37','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqa3bfvzr5s4qsagrvmv4c8pxjtn8d','DSC_0049.JPG',4809263,'image/jpeg','store/upload/2302/79/d0',1676371081579640000,'2023-02-14 18:38:01','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqbkzrauj4usx79f7axgjh2kpxbamn','DSC_0050.JPG',4164901,'image/jpeg','store/upload/2302/09/22',1676371271958234000,'2023-02-14 18:41:11','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqbqhnqr3zcqptaatv2gddmht2hku3','DSC_0048.JPG',4560846,'image/jpeg','store/upload/2302/67/79',1676371289146711000,'2023-02-14 18:41:29','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqc2wxgpkf5qtb9zmcaxtxc3nm31ny','DSC_0056.JPG',3873180,'image/jpeg','store/upload/2302/06/da',1676371335636282000,'2023-02-14 18:42:15','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqccavpmybzdyx9d1twn4zopyq2s9s','DSC_0052.JPG',3842114,'image/jpeg','store/upload/2302/ca/92',1676371373508947000,'2023-02-14 18:42:53','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqct4xppub7yvjaut9ownb78onzvrd','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/8d/4d',1676371428165144000,'2023-02-14 18:43:48','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqdkha93uadsjy9kvvf56z2amw63k4','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/68/01',1676371525561710000,'2023-02-14 18:45:25','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqg6tvz1p2c3n39pbcjxjysknwnrqy','DSC_0040.JPG',4369973,'image/jpeg','store/upload/2302/1a/df',1676371735709729000,'2023-02-14 18:48:55','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqh4yajoytvh2k9pnsxpd94b6hd1n4','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/27/15',1676371855400580000,'2023-02-14 18:50:55','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqhzoyxunb7jhpa8at8m7v8bwsymm6','DSC_0052.JPG',3842114,'image/jpeg','store/upload/2302/ab/37',1676371966268241000,'2023-02-14 18:52:46','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqj78ckqn5fntca1scn9dz3y3o5zqx','DSC_0057.JPG',4105620,'image/jpeg','store/upload/2302/3d/bc',1676371992330044000,'2023-02-14 18:53:12','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqm85yxqj2tm8map2ugfuzo7hcx3ky','DSC_0058.JPG',4345695,'image/jpeg','store/upload/2302/98/cc',1676372252944364000,'2023-02-14 18:57:32','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqn4jndsykn7ysarccoazqbaf5hhv8','DSC_0054.JPG',3766424,'image/jpeg','store/upload/2302/ea/22',1676372366667235000,'2023-02-14 18:59:26','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqnoms9tuod4b9ag49m7tvctvszshb','DSC_0055.JPG',3742602,'image/jpeg','store/upload/2302/6a/aa',1676372434315641000,'2023-02-14 19:00:34','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqnsus5v5bsp959juc25phov8gn7a6','DSC_0062.JPG',4315238,'image/jpeg','store/upload/2302/47/a7',1676372451657662000,'2023-02-14 19:00:51','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqoom2ovn3qfknacy9xadqzncqw9wq','DSC_0229.JPG',3475410,'image/jpeg','store/upload/2302/78/52',1676372562344370000,'2023-02-14 19:02:42','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqqc78v9n767ktanys8wmujsycxtyt','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/2f/dc',1676372780221627000,'2023-02-14 19:06:20','U002','admin','系统管理员','D022','2002','仓库组'),('dhpqqfx91zk5mq7gaoa965udm2cgjpzx','DSC_0087.JPG',4364866,'image/jpeg','store/upload/2302/d7/b9',1676372791867354000,'2023-02-14 19:06:31','U002','admin','系统管理员','D022','2002','仓库组'),('dhpquq2f96jp6u1ga74upwpqzj1rt6q8','DSC_0029.JPG',4386668,'image/jpeg','store/upload/2302/56/eb',1676373336857826000,'2023-02-14 19:15:36','U002','admin','系统管理员','D022','2002','仓库组'),('dhpyjhkdcpfuwf75a8bva5815qkwor7j','DSC_0049.JPG',4809263,'image/jpeg','store/upload/2302/55/45',1676404798466540000,'2023-02-15 03:59:58','U002','admin','系统管理员','D022','2002','仓库组'),('dhpykgjm1ag3xpbaaz5v6pupfxdzzu62','DSC_0029.JPG',4386668,'image/jpeg','store/upload/2302/8a/36',1676404922135172000,'2023-02-15 04:02:02','U002','admin','系统管理员','D022','2002','仓库组'),('dht4rgks7b4b5uyjaupu5gwwznfb8rxw','package.json',1640,'application/json','store/upload/2302/d1/47',1676819386471950000,'2023-02-19 23:09:46','U002','admin','系统管理员','D022','0302','仓库组'),('dhtkvyqfvwk1mqhc9jnbb7fhf8s462ya','29x29.png',842,'image/png','store/upload/2302/3e/10',1676877306493535000,'2023-02-20 15:15:06','U002','admin','系统管理员','D022','0302','仓库组'),('dhtkw7rtcotrjr7ya2kb27tkk2pdo3pd','60x60.png',1414,'image/png','store/upload/2302/e3/02',1676877338518710000,'2023-02-20 15:15:38','U002','admin','系统管理员','D022','0302','仓库组'),('dhtkwzwnwftxu54pa3ya14zh4dpjpo43','60x60.png',1414,'image/png','store/upload/2302/af/07',1676877439915253000,'2023-02-20 15:17:19','U002','admin','系统管理员','D022','0302','仓库组'),('dhtkx3x8u4abnjsqa3vvg4hxym6fzycn','80x80.png',2997,'image/png','store/upload/2302/20/ce',1676877451195552000,'2023-02-20 15:17:31','U002','admin','系统管理员','D022','0302','仓库组'),('dhtkx51vbnn439jgad2tvnghr8h32vso','120x120.png',3174,'image/png','store/upload/2302/2c/9e',1676877456162640000,'2023-02-20 15:17:36','U002','admin','系统管理员','D022','0302','仓库组'),('dhtky5wgzk9j6v6c9xzcqbsxfocmygr4','76x76.png',2675,'image/png','store/upload/2302/07/f0',1676877587380461000,'2023-02-20 15:19:47','U002','admin','系统管理员','D022','0302','仓库组'),('dhtkycz98pyu55f593scw4cpyjrfbtfg','76x76.png',2675,'image/png','store/upload/2302/01/f3',1676877615941392000,'2023-02-20 15:20:15','U002','admin','系统管理员','D022','0302','仓库组'),('dhtkyycf1jgzysgp9b6uyvb9bw2v5r7u','80x80.png',2997,'image/png','store/upload/2302/51/15',1676877689289733000,'2023-02-20 15:21:29','U002','admin','系统管理员','D022','0302','仓库组'),('dhtkza8hf9q4b31s9ak9wahydkgv1jyz','80x80.png',2997,'image/png','store/upload/2302/fd/46',1676877732109590000,'2023-02-20 15:22:12','U002','admin','系统管理员','D022','0302','仓库组'),('dhtm1cfbfacnkun6aydsu8s1b76s2bwb','80x80.png',2997,'image/png','store/upload/2302/9a/d6',1676877869690802000,'2023-02-20 15:24:29','U002','admin','系统管理员','D022','0302','仓库组'),('dhtm1dcywv8u7n57abw91f6by17cmk9h','167x167.png',6707,'image/png','store/upload/2302/f8/61',1676877873370592000,'2023-02-20 15:24:33','U002','admin','系统管理员','D022','0302','仓库组'),('dhtm7oq53r13y2jca8hc5wybyb8ovr43','IMG_7430.MP4',127466957,'video/mp4','store/upload/2302/a7/62',1676878674683736000,'2023-02-20 15:37:54','U002','admin','系统管理员','D022','0302','仓库组'),('dhtmy9ugfsf96kn8a8tat3p7nn1vbx5o','1AFC54EE270D06A0A234DC11D27D19AE.MP4',8907162,'video/mp4','store/upload/2302/78/58',1676881699583580000,'2023-02-20 16:28:19','U002','admin','系统管理员','D022','0302','仓库组'),('dhtmybj5k7a3cr6fakvug481nmnmz2ps','IMG_7098.JPG',1133388,'image/jpeg','store/upload/2302/88/c0',1676881706615084000,'2023-02-20 16:28:26','U002','admin','系统管理员','D022','0302','仓库组'),('dhtn1s3a11f9mobqam3s947g925ouzas','测试委托单投产结果_20200304.xlsx',11764,'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet','store/upload/2302/2d/62',1676882016678522000,'2023-02-20 16:33:36','U002','admin','系统管理员','D022','0302','仓库组'),('dhtn1zocx8q6xrh29jytnka8xntj5m4d','60x60.png',1414,'image/png','store/upload/2302/b1/ab',1676882046657336000,'2023-02-20 16:34:06','U002','admin','系统管理员','D022','0302','仓库组'),('dhtn2chmdmm1pmmca3qtvuunud2mcbvt','8f49e8a80ea906dc1abea937e2dc1a1d.MP4',16303119,'video/mp4','store/upload/2302/83/39',1676882093526044000,'2023-02-20 16:34:53','U002','admin','系统管理员','D022','0302','仓库组'),('dhtn2pvbd1xnaa269s2v1qmmsg5m5pyh','80x80 (3).png',3017,'image/png','store/upload/2302/7e/7b',1676882135285022000,'2023-02-20 16:35:35','U002','admin','系统管理员','D022','0302','仓库组'),('dhtn46kaxu2d585kactuwdsunf7gmkbu','4CD52F11C02933C7E3519EE95EAFC8A0.png',160887,'image/png','store/upload/2302/4d/f7',1676882326927351000,'2023-02-20 16:38:46','U002','admin','系统管理员','D022','0302','仓库组'),('dhtn4hzmf5bqqhwsara9112xrr64s5an','120x120 (1).png',3194,'image/png','store/upload/2302/e7/a1',1676882367650533000,'2023-02-20 16:39:27','U002','admin','系统管理员','D022','0302','仓库组'),('dhtn4qfc3qp82dqmavzc8hfxhhrjm212','123.csv',13117,'text/csv','store/upload/2302/75/6f',1676882393000971000,'2023-02-20 16:39:53','U002','admin','系统管理员','D022','0302','仓库组'),('dhtq8u3z1hjmpambah59mcpo5bffh4qy','40x40.png',1362,'image/png','store/upload/2302/21/31',1676895208089764000,'2023-02-20 20:13:28','U002','admin','系统管理员','D022','0302','仓库组'),('dhtq98jg1dqszbyga83t11ots5tukb5z','40x40.png',1362,'image/png','store/upload/2302/74/e8',1676895262478899000,'2023-02-20 20:14:22','U002','admin','系统管理员','D022','0302','仓库组'),('dhtq9k9znyg39baaa1bc9susbrkvfsay','76x76.png',2675,'image/png','store/upload/2302/8e/5d',1676895301911064000,'2023-02-20 20:15:01','U002','admin','系统管理员','D022','0302','仓库组'),('dhtq9qof5ga991k5ajtsq2h5w1vs7jq7','80x80 (1).png',3017,'image/png','store/upload/2302/6b/76',1676895322071055000,'2023-02-20 20:15:22','U002','admin','系统管理员','D022','0302','仓库组'),('dhtqatdp8fshfgkq9arc8q2yft7j3z7s','60x60.png',1414,'image/png','store/upload/2302/e8/a5',1676895461995253000,'2023-02-20 20:17:41','U002','admin','系统管理员','D022','0302','仓库组'),('dhtqb9g4zjavb5zgadca1ovqd1n4bj5y','60x60.png',1414,'image/png','store/upload/2302/a7/30',1676895521122670000,'2023-02-20 20:18:41','U002','admin','系统管理员','D022','0302','仓库组'),('dhtqm35o247swszf9hp9wr6rrmcj9c68','40x40.png',1362,'image/png','store/upload/2302/a3/73',1676896520318845000,'2023-02-20 20:35:20','U002','admin','系统管理员','D022','0302','仓库组'),('dhtqmx96yz5jd636aduv6dyuq8cpjrw3','1AFC54EE270D06A0A234DC11D27D19AE (1).MP4',8907162,'video/mp4','store/upload/2302/d4/85',1676896629895075000,'2023-02-20 20:37:09','U002','admin','系统管理员','D022','0302','仓库组'),('dhtqwjw5rr3f8u58a2vtgpk4fyog6su6','58x58.png',1818,'image/png','store/upload/2302/5c/e7',1676897859482656000,'2023-02-20 20:57:39','U002','admin','系统管理员','D022','0302','仓库组'),('dhtr67adgwm7j7ksa8cbq4hxhpwgxtov','20x20.png',532,'image/png','store/upload/2302/92/42',1676898969680897000,'2023-02-20 21:16:09','U002','admin','系统管理员','D022','0302','仓库组'),('dhtr6817c55x237j92oa9qupcs2f6r5b','60x60 (1).png',1434,'image/png','store/upload/2302/04/03',1676898972040968000,'2023-02-20 21:16:12','U002','admin','系统管理员','D022','0302','仓库组'),('dhtr6fmqomc92jcfamjvwsoxpdvzcupg','40x40.png',1362,'image/png','store/upload/2302/b4/87',1676898998753852000,'2023-02-20 21:16:38','U002','admin','系统管理员','D022','0302','仓库组'),('dhts7w9a7vm5ygfj96tvbzapcrmy24av','40x40 (1).png',1382,'image/png','store/upload/2302/e1/23',1676903281660073000,'2023-02-20 22:28:01','U002','admin','系统管理员','D022','0302','仓库组'),('dhts7yb6uc6bdmcjarwttos36n9cxhvv','DSC_0029.JPG',4386688,'image/jpeg','store/upload/2302/14/7b',1676903289148201000,'2023-02-20 22:28:09','U002','admin','系统管理员','D022','0302','仓库组'),('dhts7zz4gf55ybgcagncjs55udmsyb3p','IMG_7456.MP4',82054441,'video/mp4','store/upload/2302/f7/a7',1676903295239123000,'2023-02-20 22:28:15','U002','admin','系统管理员','D022','0302','仓库组'),('dhts9a4w591fmxu5acxtcz8s4rbozs2y','DSC_0034.JPG',4794122,'image/jpeg','store/upload/2302/68/55',1676903460886597000,'2023-02-20 22:31:00','U002','admin','系统管理员','D022','0302','仓库组'),('dhts9atvhu5q8snj9g7cgxyjp6r8a8hp','DSC_0044.JPG',4355216,'image/jpeg','store/upload/2302/ef/ae',1676903463892231000,'2023-02-20 22:31:03','U002','admin','系统管理员','D022','0302','仓库组'),('dhts9bq37nkzbz9695wuhf6g7pvsgszn','DSC_0040.JPG',4369973,'image/jpeg','store/upload/2302/34/e5',1676903466845267000,'2023-02-20 22:31:06','U002','admin','系统管理员','D022','0302','仓库组'),('dhts9sv22s6os2qfanrbkdarkr2r37o8','DSC_0041.JPG',3461628,'image/jpeg','store/upload/2302/3c/fc',1676903523831281000,'2023-02-20 22:32:03','U002','admin','系统管理员','D022','0302','仓库组'),('dhts9tq2jkdv9jdj9h4t3mv2r2sg7p6h','DSC_0029.JPG',4386668,'image/jpeg','store/upload/2302/f2/e6',1676903526142507000,'2023-02-20 22:32:06','U002','admin','系统管理员','D022','0302','仓库组'),('dhts9uk88x6ng84w98hs6g2phfogbxwc','DSC_0024.JPG',4340450,'image/jpeg','store/upload/2302/d3/b5',1676903530273436000,'2023-02-20 22:32:10','U002','admin','系统管理员','D022','0302','仓库组'),('dhts9xchkh68vn579zfcj4q3nwa4cbvb','DSC_0050.JPG',4164901,'image/jpeg','store/upload/2302/c0/a3',1676903541299088000,'2023-02-20 22:32:21','U002','admin','系统管理员','D022','0302','仓库组'),('dhtsbybbxvhyumt79xuvfq4jwndu6cxu','DSC_0041.JPG',3461628,'image/jpeg','store/upload/2302/a8/d8',1676903801936055000,'2023-02-20 22:36:41','U002','admin','系统管理员','D022','0302','仓库组'),('dhtsbz83so2qov4j9cg9ohmnqn5qu5kb','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/9b/56',1676903804222269000,'2023-02-20 22:36:44','U002','admin','系统管理员','D022','0302','仓库组'),('dhtsv26abmxmssafaosayuy4ojsadknj','DSC_0045.JPG',4621019,'image/jpeg','store/upload/2302/84/0d',1676905860676386000,'2023-02-20 23:11:00','U002','admin','系统管理员','D022','0302','仓库组'),('dhtsv2j1m61st7km9am9mcurdd99unkt','DSC_0036.JPG',4292333,'image/jpeg','store/upload/2302/1d/8b',1676905862985234000,'2023-02-20 23:11:02','U002','admin','系统管理员','D022','0302','仓库组'),('dhtsv3rc9g9wj3sua2vc9jnvcqu8u5r8','DSC_0029.JPG',4386668,'image/jpeg','store/upload/2302/13/f4',1676905866641463000,'2023-02-20 23:11:06','U002','admin','系统管理员','D022','0302','仓库组'),('dhtta9d4frpj6vd79kq9zc7kzk9q44h6','DSC_0036.JPG',4292333,'image/jpeg','store/upload/2302/e7/29',1676907681332957000,'2023-02-20 23:41:21','U002','admin','系统管理员','D022','0302','仓库组'),('dhuhvrn3b5gs5s7tarpubarnht19a4bg','1AFC54EE270D06A0A234DC11D27D19AE (1) (1).MP4',8907182,'video/mp4','store/upload/2302/81/74',1677000158920815000,'2023-02-22 01:22:38','U002','admin','系统管理员','D022','0302','仓库组'),('dhuhvt645cwo2m9z99jupm7op51hjsyc','IMG_7456 (1).MP4',82054461,'video/mp4','store/upload/2302/1f/dd',1677000164875015000,'2023-02-22 01:22:44','U002','admin','系统管理员','D022','0302','仓库组'),('dhuhvuvdztsb9q2w98pt3ut2yc3jvo5r','IMG_7098.JPG',1133388,'image/jpeg','store/upload/2302/8f/bc',1677000171040457000,'2023-02-22 01:22:51','U002','admin','系统管理员','D022','0302','仓库组'),('dhuhvvtk8hgcojxdanxa65bz3jnk7nyx','IMG_7913.PNG',7640338,'image/png','store/upload/2302/8b/a4',1677000175226312000,'2023-02-22 01:22:55','U002','admin','系统管理员','D022','0302','仓库组'),('dhukvq939g95dpchabks7rgv7gbnsvvx','DSC_0029.JPG',4386688,'image/jpeg','store/upload/2302/c7/be',1677008345907299000,'2023-02-22 03:39:05','U002','admin','系统管理员','D022','0302','仓库组'),('dhuw683x52abknhhan5v4jqn76a75sd8','DSC_0034.JPG',4794142,'image/jpeg','store/upload/2302/98/82',1677050524098020000,'2023-02-22 15:22:04','U002','admin','系统管理员','D022','0302','仓库组'),('dhuwff3ayywz8z5w9j8uf5tho5x1wwy8','DSC_0029.JPG',4386688,'image/jpeg','store/upload/2302/ca/f9',1677051572719934000,'2023-02-22 15:39:32','U002','admin','系统管理员','D022','0302','仓库组'),('dhuwfh4xhzptz4vk9p8bn2ptb69fn8b9','DSC_0036.JPG',4292353,'image/jpeg','store/upload/2302/4e/8b',1677051580497451000,'2023-02-22 15:39:40','U002','admin','系统管理员','D022','0302','仓库组'),('dhuwfjjzwptrv61wag4bqznbd77myjzf','DSC_0044.JPG',4355236,'image/jpeg','store/upload/2302/5c/65',1677051586390321000,'2023-02-22 15:39:46','U002','admin','系统管理员','D022','0302','仓库组'),('dhuyqmy9w2vp9j7s92su6aw7pc3f894n','29x29.png',842,'image/png','store/upload/2302/ba/09',1677060939443708000,'2023-02-22 18:15:39','U002','admin','系统管理员','D022','0302','仓库组'),('dhuyxupx6wptfjwfastc3woypb1jyptt','29x29.png',842,'image/png','store/upload/2302/62/aa',1677061866152666000,'2023-02-22 18:31:06','U002','admin','系统管理员','D022','0302','仓库组'),('dhv3qqrc7f3oxmqb9rmanjp8o8rsh9gf','29x29.png',842,'image/png','store/upload/2302/59/11',1677077338474827000,'2023-02-22 22:48:58','U001','admin','系统管理员3','D005','04','后勤保障部'),('dhv3qz2ysoxzqk6ba1puvux8nzga435y','40x4ssss0.png',1382,'image/png','store/upload/2302/4b/d9',1677077372922147000,'2023-02-22 22:49:32','U001','admin','系统管理员3','D005','04','后勤保障部');
/*!40000 ALTER TABLE `sys_doc` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_menu`
--

DROP TABLE IF EXISTS `sys_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_menu` (
                            `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                            `name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `parent_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `menu_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                            `icon_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `description_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `valid_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `order_` bigint DEFAULT NULL,
                            `create_at_` datetime DEFAULT NULL,
                            `update_at_` datetime DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `sys_menu_ForeignKey_parent_id` (`parent_id_`),
                            CONSTRAINT `sys_menu_ForeignKey_parent_id` FOREIGN KEY (`parent_id_`) REFERENCES `sys_menu` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menu`
--

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;
INSERT INTO `sys_menu` VALUES ('dfdfvq2jwzo44u3n926u481sc6s9mjbw','流程管理',NULL,'NONE','mdi mdi-material-design','','Effective',1668905487363584750,'2022-10-26 15:12:24','2022-10-26 15:13:16'),('dfdfwhjo9vycqv1da5hc5vxkonjkb4fr','流程设计','dfdfvq2jwzo44u3n926u481sc6s9mjbw','diagram_designer','mdi mdi-account-network-outline','','Effective',1666768446457181000,'2022-10-26 15:14:06','2022-10-26 15:16:11'),('dhpwxvxtf5gj5xa194tbtzyt9nwd6tpu','流程发起','dfdfvq2jwzo44u3n926u481sc6s9mjbw','diagram_launch','mdi mdi-clock-start','','Effective',1676398319264347048,'2023-02-15 02:11:59','2023-02-15 02:12:49'),('gfwu7esi4najtccjdmsvlk2an5k5k6zg','代码模版',NULL,'NONE','mdi mdi-protocol','','Effective',1669261660847985375,'2022-08-15 09:34:35','2022-08-15 09:40:05'),('gfwubxwwnieopq5zcgje7i52klvcu27c','DataTable','gfwu7esi4najtccjdmsvlk2an5k5k6zg','system_proto_datatable','mdi mdi-table-large','','Effective',-5117,'2022-08-15 09:40:07','2022-08-15 09:40:34'),('gfxj62zux677of6ycp6fiuvfoqjuvrvz','Tree','gfwu7esi4najtccjdmsvlk2an5k5k6zg','system_proto_tree','mdi mdi-file-tree-outline','','Effective',-4861,'2022-08-16 09:25:23','2022-08-16 09:25:52'),('gfxj6cmkljkolepydgytr5bdjmitnp5y','List','gfwu7esi4najtccjdmsvlk2an5k5k6zg','system_proto_list','mdi mdi-format-list-bulleted','','Effective',-4605,'2022-08-16 09:25:54','2022-08-16 09:26:12'),('gfzn32pzvqtw43zhdlhdqi7g3f6elef6','Form','gfwu7esi4najtccjdmsvlk2an5k5k6zg','system_proto_form','mdi mdi-form-textbox-password','','Effective',-4093,'2022-08-19 14:41:06','2022-08-19 14:42:16'),('mkzt2nkhphcpswaxjrfjhkprgtdjajh0','组织架构',NULL,'NONE','mdi mdi-cogs','','Effective',-3837,'2022-01-01 00:00:00','2022-06-27 11:32:45'),('mkzt2nkhphcpswaxjrfjhkprgtdjajh4','部门管理','mkzt2nkhphcpswaxjrfjhkprgtdjajh0','system_organization_departs','mdi mdi-newspaper-variant-multiple-outline','','Effective',-1789,'2022-01-02 00:00:00','2022-06-28 08:51:42'),('mkzt2nkhphcpswaxjrfjhkprgtdjajh5','用户管理','mkzt2nkhphcpswaxjrfjhkprgtdjajh0','system_organization_users','mdi mdi-account-outline','','Effective',3,'2022-01-03 00:00:00','2023-02-13 23:25:27'),('mkzt2nkhphcpswaxjrfjhkprgtdjajhu','数据管理',NULL,'NONE','mdi mdi-database-cog','','Effective',1655914528114860,'2022-06-23 00:03:02','2022-06-23 00:05:15'),('mkzt2o45th4x7w5ljx5ilephrjnfav63','角色授权','mkztz4ml676ea262jxmzt5ansgqvinqn','system_permission_roles','mdi mdi-account-details','','Effective',1448924619214300,'2022-06-23 00:03:08','2022-06-28 16:59:59'),('mkzt2zl2j6oflbhriji2gizzpkd6xwyq','组织授权','mkztz4ml676ea262jxmzt5ansgqvinqn','system_permission_organization','mdi mdi-chandelier','','Effective',1241935376916825,'2022-06-23 00:03:50','2022-06-28 13:54:11'),('mkzt3bns6n547gm3jbkzxq7bvhgkolga','组织权限查询','mkztz4ml676ea262jxmzt5ansgqvinqn','system_permission_by_organization','mdi mdi-tag-search','','Effective',1655913861511776,'2022-06-23 00:04:22','2022-07-04 15:06:31'),('mkzt3uxlm6jnig4qi66ln2thvi7hqpzo','数据库表','mkzt2nkhphcpswaxjrfjhkprgtdjajhu','system_service_tables','mdi mdi-table-large','','Effective',1655914031095950,'2022-06-23 00:05:38','2022-06-26 23:23:59'),('mkzt5ery42a7dkm3jxk2cnbuk436jneg','数据字典','mkzt2nkhphcpswaxjrfjhkprgtdjajhu','system_service_dicts','mdi mdi-file-sign','','Effective',1655914115888037,'2022-06-23 00:08:51','2022-06-26 23:27:35'),('mkzt5ph445caccwdj542e4wdz2f6xa4k','自动编码','mkzt2nkhphcpswaxjrfjhkprgtdjajhu','system_service_auto_no','mdi mdi-numeric-10-box-outline','','Effective',1655914158284081,'2022-06-23 00:09:33','2022-07-06 15:29:08'),('mkzt5wbpa2gwkdmricojln7bg6naukii','数据服务','mkzt2nkhphcpswaxjrfjhkprgtdjajhu','system_service_data','mdi mdi-cable-data','','Effective',1655914200680125,'2022-06-23 00:10:01','2022-07-14 08:44:02'),('mkzt7f7ydjraqtzli4xjdts6vmxcabsi','系统管理',NULL,'NONE','mdi mdi-remote-desktop','','Effective',1669617834332386000,'2022-06-23 00:13:11','2022-06-27 11:24:29'),('mkztz4ml676ea262jxmzt5ansgqvinqn','权限管理',NULL,'NONE','mdi mdi-timer-cog','','Effective',-3581,'2022-06-23 00:01:53','2022-06-23 00:05:35'),('mkzuc45sz7h5nowqjjkjsi24ifk7dqvq','角色管理','mkzt7f7ydjraqtzli4xjdts6vmxcabsi','system_admin_roles','mdi mdi-account-multiple-outline','','Effective',1655914873281018,'2022-06-23 00:21:08','2022-06-23 00:21:44'),('mkzucq7mis655rf7izqzfv5gii3jr3af','系统设置','mkzt7f7ydjraqtzli4xjdts6vmxcabsi','system_admin_setting','mdi mdi-cog-outline','','Effective',1655914851820353,'2022-06-23 00:20:19','2022-06-23 00:21:06'),('mkzudgkfzs4dzotmjatkhn5mmaqo7ayn','菜单管理','mkzt7f7ydjraqtzli4xjdts6vmxcabsi','system_admin_menus','mdi mdi-menu','','Effective',1655914862550685,'2022-06-23 00:21:46','2022-06-23 00:22:16'),('mkzudoj77ovek6asjvgywl6dbbqef4qm','操作日志','mkzt7f7ydjraqtzli4xjdts6vmxcabsi','system_admin_operate_logs','mdi mdi-math-log','','Effective',208864089066376185,'2022-06-23 00:22:18','2022-06-26 23:24:02');
/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_operate_log`
--

DROP TABLE IF EXISTS `sys_operate_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_operate_log` (
                                   `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                   `ip_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `size_` int DEFAULT NULL,
                                   `agent_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                                   `method_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `menu_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `path_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `params_` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
                                   `values_` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
                                   `user_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `user_code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `user_name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `depart_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `depart_code_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `depart_name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `start_` datetime DEFAULT NULL,
                                   `end_` datetime DEFAULT NULL,
                                   `duration_` int DEFAULT NULL,
                                   `status_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `message_` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_operate_log`
--

LOCK TABLES `sys_operate_log` WRITE;
/*!40000 ALTER TABLE `sys_operate_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_operate_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_organization_role`
--

DROP TABLE IF EXISTS `sys_organization_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_organization_role` (
                                         `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                         `organization_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                         `role_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                         `create_at_` datetime DEFAULT NULL,
                                         PRIMARY KEY (`id`),
                                         UNIQUE KEY `sys_organization_role_UniqueIndex_organization_id_role_id` (`organization_id_`,`role_id_`),
                                         KEY `sys_organization_role_ForeignKey_role_id` (`role_id_`),
                                         CONSTRAINT `sys_organization_role_ForeignKey_role_id` FOREIGN KEY (`role_id_`) REFERENCES `sys_role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_organization_role`
--

LOCK TABLES `sys_organization_role` WRITE;
/*!40000 ALTER TABLE `sys_organization_role` DISABLE KEYS */;
INSERT INTO `sys_organization_role` VALUES ('dhnwz8s43bzqjbro9atvkkuqb1sq9rka','U081','R006','2023-02-12 01:25:51'),('dhnwzfqrt7zb6o179anav18u1qbb4dac','U021','R006','2023-02-12 01:26:14'),('dhnwzkt69f4guc63a8csr9qbmrcmvbak','U017','R003','2023-02-12 01:26:31'),('dhnwzkwsbyya8qsoawtudryyn7owrvnm','D022','R003','2023-02-12 01:26:31'),('dhnxmpa5n795jodjahtuwsy9kgsj1wqt','D004','R009','2023-02-12 02:07:17'),('dhnxmpbnk1puff56a1bvstd9cw1b6skr','D004','R010','2023-02-12 02:07:17'),('dhnxmpd55xhbfu259cmsyrtjt1yutk19','D004','R002','2023-02-12 02:07:17'),('dhnxmpgw9m4883qvagzs6sjtucd8rqpa','D004','R006','2023-02-12 02:07:17'),('dhnyjcafwtwzkm67awqaqn4654418ukv','D003','R007','2023-02-12 03:10:37'),('dhnyjcd6yvogk9d8attcko8a4dxdko47','D003','R002','2023-02-12 03:10:37'),('dhnyjcdhwjc1ktwja8hao934jssnxc9p','D003','R001','2023-02-12 03:10:37'),('dhv31v9dqf2ozfm6asw9ykdfn3d54fzf','U003','R010','2023-02-22 22:02:21'),('dhv31vh1bym3gfzd9z7u1ot7zd19sopp','U001','R010','2023-02-22 22:02:21'),('OR001','U002','R010','2023-01-09 00:00:00');
/*!40000 ALTER TABLE `sys_organization_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role`
--

DROP TABLE IF EXISTS `sys_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_role` (
                            `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                            `code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `description_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `order_` bigint DEFAULT NULL,
                            `create_at_` datetime DEFAULT NULL,
                            `update_at_` datetime DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `sys_role_UniqueIndex_code` (`code_`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role`
--

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;
INSERT INTO `sys_role` VALUES ('R001','008','生产计划员','',1659114559779873429,'2022-06-07 08:52:28','2023-02-13 23:56:50'),('R002','006','采购员','',1659230818887538574,'2022-06-07 08:52:30','2023-02-13 23:56:06'),('R003','009','车间计划员','',1659095183261929238,'2022-06-07 08:52:32','2023-02-13 23:56:08'),('R004','003','项目主管','',1659799196747234839,'2022-06-07 08:52:34','2023-02-13 23:47:08'),('R005','004','项目负责人','',1659695855318199155,'2022-06-07 08:52:36','2023-02-13 23:58:09'),('R006','010','工序检验员','',1659075806743985048,'2022-06-07 08:54:30','2023-02-13 23:57:58'),('R007','007','仓库保管员','',1659153312815761811,'2022-06-07 08:55:22','2023-02-13 23:56:19'),('R008','005','设计师','',1659385831031092101,'2022-06-07 08:56:41','2023-02-13 23:56:16'),('R009','002','系统运维员','',1660264233177895419,'2022-06-07 09:02:29','2023-02-13 23:47:10'),('R010','001','系统管理员','',1660729269608556000,'2022-06-07 09:03:45','2023-02-13 23:58:16');
/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_menu`
--

DROP TABLE IF EXISTS `sys_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_role_menu` (
                                 `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                 `role_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                 `menu_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                 `create_at_` datetime DEFAULT NULL,
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `sys_role_menu_UniqueIndex_role_id_menu_id` (`role_id_`,`menu_id_`),
                                 KEY `sys_role_menu_ForeignKey_menu_id` (`menu_id_`),
                                 CONSTRAINT `sys_role_menu_ForeignKey_menu_id` FOREIGN KEY (`menu_id_`) REFERENCES `sys_menu` (`id`),
                                 CONSTRAINT `sys_role_menu_ForeignKey_role_id` FOREIGN KEY (`role_id_`) REFERENCES `sys_role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_menu`
--

LOCK TABLES `sys_role_menu` WRITE;
/*!40000 ALTER TABLE `sys_role_menu` DISABLE KEYS */;
INSERT INTO `sys_role_menu` VALUES ('dhp7hs4c4nqu33549h698dumhdxw3v2a','R006','mkzt2o45th4x7w5ljx5ilephrjnfav63','2023-02-14 00:40:00'),('dhp7hs6cabaom3zjaf8u6xxk69s4zt4b','R006','mkzt3bns6n547gm3jbkzxq7bvhgkolga','2023-02-14 00:40:00'),('dhp7hs7j511koj8yau8bgzmybkp4a3bb','R006','mkzt2zl2j6oflbhriji2gizzpkd6xwyq','2023-02-14 00:40:00'),('dhpwyhs4cpy5ph7m9tbtr756vs8fnn4f','R010','gfzn32pzvqtw43zhdlhdqi7g3f6elef6','2023-02-15 02:13:19'),('dhpwyhtppf64h1m19wd9fn2z8nufynwq','R010','mkzudgkfzs4dzotmjatkhn5mmaqo7ayn','2023-02-15 02:13:19'),('dhpwyhub6na8sgjta72brnoyd8tfqhbn','R010','mkzt5wbpa2gwkdmricojln7bg6naukii','2023-02-15 02:13:19'),('dhpwyhupcujmbusk9uub9v7q993zp4pm','R010','dfdfwhjo9vycqv1da5hc5vxkonjkb4fr','2023-02-15 02:13:19'),('dhpwyhuq3wtgw1s29j5ugxtbg2pnzygq','R010','mkzt5ery42a7dkm3jxk2cnbuk436jneg','2023-02-15 02:13:19'),('dhpwyhvtbrympkdp9f5tgzr2n97xaa5a','R010','mkzuc45sz7h5nowqjjkjsi24ifk7dqvq','2023-02-15 02:13:19'),('dhpwyhvxfzutaq1ba7uv7a84tx7rdo6b','R010','gfxj6cmkljkolepydgytr5bdjmitnp5y','2023-02-15 02:13:19'),('dhpwyhw4qtut6rnv9159uhd7pr71226m','R010','dhpwxvxtf5gj5xa194tbtzyt9nwd6tpu','2023-02-15 02:13:19'),('dhpwyhwg5bgzgdpb9c8tsbhbdqmpv13n','R010','mkzudoj77ovek6asjvgywl6dbbqef4qm','2023-02-15 02:13:19'),('dhpwyhwmax3vfy7g95cuafbuacdqw7gs','R010','mkzt2nkhphcpswaxjrfjhkprgtdjajh4','2023-02-15 02:13:19'),('dhpwyhwojvj6z7qua78at5gq8r6oyyu4','R010','mkzt5ph445caccwdj542e4wdz2f6xa4k','2023-02-15 02:13:19'),('dhpwyhwqhr9gtxapa8n9nknxqk79j3cn','R010','mkzt3uxlm6jnig4qi66ln2thvi7hqpzo','2023-02-15 02:13:19'),('dhpwyhwtwtdzc6wk97aa6rcod214wvyo','R010','mkzucq7mis655rf7izqzfv5gii3jr3af','2023-02-15 02:13:19'),('dhpwyhxs7jmd8g4qajjbv1urbjgnbtv5','R010','gfxj62zux677of6ycp6fiuvfoqjuvrvz','2023-02-15 02:13:19'),('dhpwyhyjrwh3j4w1azdaunkbmkkf5wu5','R010','mkzt2zl2j6oflbhriji2gizzpkd6xwyq','2023-02-15 02:13:19'),('dhpwyhykhkgfaa4s9wvvrdh4rzukdhjf','R010','gfwubxwwnieopq5zcgje7i52klvcu27c','2023-02-15 02:13:19'),('dhpwyhynrznwopr49xwu1ukau42j8x3y','R010','mkzt2o45th4x7w5ljx5ilephrjnfav63','2023-02-15 02:13:19'),('dhpwyhyorm6bx6spaxguzhgqc1fb1t64','R010','mkzt2nkhphcpswaxjrfjhkprgtdjajh5','2023-02-15 02:13:19'),('dhpwyhyp9xn9cqzfas7ajr6gmgncqufm','R010','mkzt3bns6n547gm3jbkzxq7bvhgkolga','2023-02-15 02:13:19');
/*!40000 ALTER TABLE `sys_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_setting`
--

DROP TABLE IF EXISTS `sys_setting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_setting` (
                               `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                               `field_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `value_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `description_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `order_` bigint DEFAULT NULL,
                               `create_at_` datetime DEFAULT NULL,
                               `update_at_` datetime DEFAULT NULL,
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_setting`
--

LOCK TABLES `sys_setting` WRITE;
/*!40000 ALTER TABLE `sys_setting` DISABLE KEYS */;
INSERT INTO `sys_setting` VALUES ('10','name','Phoenix','系统名称',10,'2022-01-01 01:01:01','2023-02-22 21:31:47'),('20','version','1.0.0','版本号',20,'2022-01-01 01:01:01','2023-02-22 21:31:47'),('30','copyright','2023@Phoenix','系统底部版权信息',30,'2022-01-01 01:01:01','2023-02-22 21:31:47'),('50','password_default','12345678','初始默认密码',50,'2022-01-01 01:01:01','2023-02-22 21:31:47'),('60','password_min_length','6','密码最小长度',60,'2022-01-01 01:01:01','2023-02-22 21:31:47'),('70','password_max_length','24','密码最大长度',70,'2022-01-01 01:01:01','2023-02-22 21:31:47'),('80','classification_enable','No','启用密级管理',80,'2022-01-01 01:01:01','2023-02-22 21:31:47'),('90','token_expire','86400','Token时效(秒)',90,'2022-01-01 01:01:01','2023-02-22 21:31:47');
/*!40000 ALTER TABLE `sys_setting` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_table`
--

DROP TABLE IF EXISTS `sys_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_table` (
                             `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                             `code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                             `name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                             `sync_status_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                             `description_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                             `order_` bigint DEFAULT NULL,
                             `create_at_` datetime DEFAULT NULL,
                             `update_at_` datetime DEFAULT NULL,
                             PRIMARY KEY (`id`),
                             UNIQUE KEY `sys_table_UniqueIndex_code` (`code_`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_table`
--

LOCK TABLES `sys_table` WRITE;
/*!40000 ALTER TABLE `sys_table` DISABLE KEYS */;
INSERT INTO `sys_table` VALUES ('gfahporvfpiepzledb3ur2zgi5a4q46r','students','学生','Done','',1669608933289343000,'2022-07-12 10:22:11','2023-02-13 01:36:44'),('gfpygrr37uwlrc3kdnjejp3obxgp2u3r','school','学校','Done','',1669609070769693000,'2022-08-04 23:31:42','2023-02-13 01:36:41'),('gfwuk2jlr4jt4zayd5ouodg2xepfufit','tests','测试','Done','',1676223417607055000,'2022-08-15 09:57:21','2023-02-13 01:36:47');
/*!40000 ALTER TABLE `sys_table` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_table_column`
--

DROP TABLE IF EXISTS `sys_table_column`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_table_column` (
                                    `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                    `table_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `is_sys_` tinyint DEFAULT NULL,
                                    `code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `type_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `description_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `order_` bigint DEFAULT NULL,
                                    `create_at_` datetime DEFAULT NULL,
                                    `update_at_` datetime DEFAULT NULL,
                                    PRIMARY KEY (`id`),
                                    UNIQUE KEY `sys_table_column_UniqueIndex_table_id_code` (`table_id_`,`code_`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_table_column`
--

LOCK TABLES `sys_table_column` WRITE;
/*!40000 ALTER TABLE `sys_table_column` DISABLE KEYS */;
INSERT INTO `sys_table_column` VALUES ('gfahpokswfq76c6xcanvykhraafv4uys','gfahporvfpiepzledb3ur2zgi5a4q46r',1,'create_at_','创建时间','DATETIME','【系统自动创建】该条记录的创建时间',1658197330945805000,'2022-07-12 10:22:11',NULL),('gfahpol7k7rj2qh4c3bfer5sl5xwd5sk','gfahporvfpiepzledb3ur2zgi5a4q46r',1,'update_at_','更新时间','DATETIME','【系统自动创建】该条记录的最新更新时间',1658197330946093000,'2022-07-12 10:22:11',NULL),('gfahponkmef4in3jcmbtramdbjlic2nd','gfahporvfpiepzledb3ur2zgi5a4q46r',1,'create_depart_id_','部门ID','VARCHAR(256)','【系统自动创建】该条记录的创建者的部门ID',1658197330944064000,'2022-07-12 10:22:11',NULL),('gfahpooeiqng6llwdxlfrmdqdiau3sqw','gfahporvfpiepzledb3ur2zgi5a4q46r',1,'create_user_id_','用户ID','VARCHAR(256)','【系统自动创建】该条记录的创建者ID',1658197330944796000,'2022-07-12 10:22:11',NULL),('gfahpopalawqhchldg4suizheutd7wie','gfahporvfpiepzledb3ur2zgi5a4q46r',1,'create_depart_name_','部门名称','VARCHAR(256)','【系统自动创建】该条记录的创建者的部门名称',1658197330944407000,'2022-07-12 10:22:11',NULL),('gfahpopyzqcdmayhd6wvn2jvzsps6ppj','gfahporvfpiepzledb3ur2zgi5a4q46r',1,'create_user_code_','工号','VARCHAR(256)','【系统自动创建】该条记录的创建者工号',1658197330945150000,'2022-07-12 10:22:11',NULL),('gfahpoqevjq75kkjdufftoqpttqa4tpp','gfahporvfpiepzledb3ur2zgi5a4q46r',1,'order_','排序号','BIGINT','【系统自动创建】拖拽排序服务的排序号',1658197330943371001,'2022-07-12 10:22:11',NULL),('gfahpoqo4nsjtmzqcfcfx2iu5e7p2jbz','gfahporvfpiepzledb3ur2zgi5a4q46r',1,'create_user_name_','用户名','VARCHAR(256)','【系统自动创建】该条记录的创建者用户名',1658197330945458000,'2022-07-12 10:22:11',NULL),('gfahpor2m24jvpgbcb2fr6qjrqpc3bod','gfahporvfpiepzledb3ur2zgi5a4q46r',1,'parent_id_','父级ID','VARCHAR(256)','【系统自动创建】自动建树服务的父级ID',1658197330943714000,'2022-07-12 10:22:11',NULL),('gfahporsvtd24n3ldvoed5cobtn52krn','gfahporvfpiepzledb3ur2zgi5a4q46r',1,'id','ID','VARCHAR(256)','【系统自动创建】该条记录的全局唯一ID',1657592530942502000,'2022-07-12 10:22:11',NULL),('gfahqzp6kklqgbyhdqaunrhkfwp4x6yd','gfahporvfpiepzledb3ur2zgi5a4q46r',0,'code_','学号','VARCHAR(256)','',1657592702722928000,'2022-07-12 10:25:03','2022-07-12 10:25:15'),('gfahr6v6cpgfg5dbdireobrvhwkxhbbj','gfahporvfpiepzledb3ur2zgi5a4q46r',0,'name_','姓名','VARCHAR(256)','',1657592734268036000,'2022-07-12 10:25:23',NULL),('gfahrbppyecqsgowdqjfpizsmfbrqyn3','gfahporvfpiepzledb3ur2zgi5a4q46r',0,'age_','年龄','TINYINT','',1657592734268038000,'2022-07-12 10:25:34','2022-07-29 15:25:06'),('gfahreg2xpvbjke6du4c445ngmtadpql','gfahporvfpiepzledb3ur2zgi5a4q46r',0,'score_','学分','NUMERIC(13,2)','',1657592786101048000,'2022-07-12 10:25:46','2022-07-12 10:25:47'),('gfahrifiav3evzzncxxt2jgzmoqkibci','gfahporvfpiepzledb3ur2zgi5a4q46r',0,'birth_','生日','DATE','',1657592734268037998,'2022-07-12 10:26:02','2022-07-29 14:21:31'),('gfahropeq3zfckffdc3ew5eemdrr36xd','gfahporvfpiepzledb3ur2zgi5a4q46r',0,'description_','简介','TEXT','',1658197330943542500,'2022-07-12 10:26:26','2022-07-12 10:26:31'),('gfch7fmuaobxvycidahfc2ndsx4g56fz','gfahporvfpiepzledb3ur2zgi5a4q46r',0,'pay_','学费','NUMERIC(13,2)','',1657592786101047000,'2022-07-15 10:36:30','2022-07-15 10:37:21'),('gfcue3x36vf3cctrcizsradi23vaj4tz','gfahporvfpiepzledb3ur2zgi5a4q46r',0,'province_','省份','VARCHAR(256)','',1657592786101045000,'2022-07-16 01:34:00',NULL),('gfcue7h7r5c7c5yqdgmfdydnkqjyx7gm','gfahporvfpiepzledb3ur2zgi5a4q46r',0,'sex_','性别','VARCHAR(256)','',1657592734268037000,'2022-07-16 01:34:13',NULL),('gfcueediknf7pby5cbudoqrfdjm62o6f','gfahporvfpiepzledb3ur2zgi5a4q46r',0,'is_full_','是否全日制','VARCHAR(256)','',1657592786101048999,'2022-07-16 01:34:34',NULL),('gfnede4bhzcn4fivcpwf4b6tijrqivw6','gfahporvfpiepzledb3ur2zgi5a4q46r',0,'course_','课程数','INT','',1657592786101048499,'2022-07-31 23:50:00','2022-07-31 23:52:20'),('gfpygrkaqgwg4nytcquusqykjo7avwip','gfpygrr37uwlrc3kdnjejp3obxgp2u3r',1,'parent_id_','父级ID','VARCHAR(256)','【系统自动创建】自动建树服务的父级ID',1660231902863508000,'2022-08-04 23:31:42',NULL),('gfpygrkrkmtez65ddzodgupnzoizzetk','gfpygrr37uwlrc3kdnjejp3obxgp2u3r',1,'id','ID','VARCHAR(256)','【系统自动创建】该条记录的全局唯一ID',1659627102862793000,'2022-08-04 23:31:42',NULL),('gfpygrlhk3xmot42dsesuuglbokqtqzx','gfpygrr37uwlrc3kdnjejp3obxgp2u3r',1,'order_','排序号','BIGINT','【系统自动创建】拖拽排序服务的排序号',1660231902864893000,'2022-08-04 23:31:42',NULL),('gfpygro4uzhndxszdgdcpqu7fczrco72','gfpygrr37uwlrc3kdnjejp3obxgp2u3r',1,'update_at_','更新时间','DATETIME','【系统自动创建】该条记录的最新更新时间',1660231902867671000,'2022-08-04 23:31:42',NULL),('gfpygrpapuszpjmacyqfmemw4zf4yqtd','gfpygrr37uwlrc3kdnjejp3obxgp2u3r',1,'create_depart_name_','部门名称','VARCHAR(256)','【系统自动创建】该条记录的创建者的部门名称',1660231902865819000,'2022-08-04 23:31:42',NULL),('gfpygrpesbtywzzsd7zf4dzucw2d3j6z','gfpygrr37uwlrc3kdnjejp3obxgp2u3r',1,'create_user_id_','用户ID','VARCHAR(256)','【系统自动创建】该条记录的创建者ID',1660231902866221000,'2022-08-04 23:31:42',NULL),('gfpygrpk2hirpbpmd2duug2poczxkta6','gfpygrr37uwlrc3kdnjejp3obxgp2u3r',1,'create_user_name_','用户名','VARCHAR(256)','【系统自动创建】该条记录的创建者用户名',1660231902866996000,'2022-08-04 23:31:42',NULL),('gfpygrq44ir2lamodzycf7lntjyxiwtm','gfpygrr37uwlrc3kdnjejp3obxgp2u3r',1,'create_at_','创建时间','DATETIME','【系统自动创建】该条记录的创建时间',1660231902867298000,'2022-08-04 23:31:42',NULL),('gfpygrqwnyz4i5nlcbofqpkeojkfe7sf','gfpygrr37uwlrc3kdnjejp3obxgp2u3r',1,'create_user_code_','工号','VARCHAR(256)','【系统自动创建】该条记录的创建者工号',1660231902866660000,'2022-08-04 23:31:42',NULL),('gfpygrriogf2d463dx5fporkobq7272h','gfpygrr37uwlrc3kdnjejp3obxgp2u3r',1,'create_depart_id_','部门ID','VARCHAR(256)','【系统自动创建】该条记录的创建者的部门ID',1660231902865424000,'2022-08-04 23:31:42',NULL),('gfpyhlsog7mij7wmclhfufewjj3taubw','gfpygrr37uwlrc3kdnjejp3obxgp2u3r',0,'name_','部门名称','VARCHAR(256)','',1659627207203481000,'2022-08-04 23:33:27',NULL),('gfwuk2cmk6bvalzddzjv7efketlfi5pg','gfwuk2jlr4jt4zayd5ouodg2xepfufit',1,'create_user_id_','用户ID','VARCHAR(256)','【系统自动创建】该条记录的创建者ID',1661133441033177000,'2022-08-15 09:57:21',NULL),('gfwuk2co7xjr5wcgcxathgzgrfbm5sc5','gfwuk2jlr4jt4zayd5ouodg2xepfufit',1,'create_user_code_','工号','VARCHAR(256)','【系统自动创建】该条记录的创建者工号',1661133441033714000,'2022-08-15 09:57:21',NULL),('gfwuk2ddc2zwqdncd2vej3lxcbkcku4s','gfwuk2jlr4jt4zayd5ouodg2xepfufit',1,'update_at_','更新时间','DATETIME','【系统自动创建】该条记录的最新更新时间',1661133441034934000,'2022-08-15 09:57:21',NULL),('gfwuk2dzi4ip5vafda3uwy7nhnqdz5ch','gfwuk2jlr4jt4zayd5ouodg2xepfufit',1,'create_depart_name_','部门名称','VARCHAR(256)','【系统自动创建】该条记录的创建者的部门名称',1661133441032894000,'2022-08-15 09:57:21',NULL),('gfwuk2f34zjzijduc4uuyrn5joqji4si','gfwuk2jlr4jt4zayd5ouodg2xepfufit',1,'create_user_name_','用户名','VARCHAR(256)','【系统自动创建】该条记录的创建者用户名',1661133441034126000,'2022-08-15 09:57:21',NULL),('gfwuk2fgcjp6zc3pdxvtvxhxa5mo2hk2','gfwuk2jlr4jt4zayd5ouodg2xepfufit',1,'id','ID','VARCHAR(256)','【系统自动创建】该条记录的全局唯一ID',1660528641031131000,'2022-08-15 09:57:21',NULL),('gfwuk2gttn5tqt4edjruzxcel26cysca','gfwuk2jlr4jt4zayd5ouodg2xepfufit',1,'order_','排序号','BIGINT','【系统自动创建】拖拽排序服务的排序号',1661133441032362000,'2022-08-15 09:57:21',NULL),('gfwuk2ibxs3v42pqdudfrokxzybkl2xo','gfwuk2jlr4jt4zayd5ouodg2xepfufit',1,'create_at_','创建时间','DATETIME','【系统自动创建】该条记录的创建时间',1661133441034506000,'2022-08-15 09:57:21',NULL),('gfwuk2jq3nudd7g4ct6chtwkftcyl3q1','gfwuk2jlr4jt4zayd5ouodg2xepfufit',1,'create_depart_code_','部门名称','VARCHAR(256)','【系统自动创建】该条记录的创建者的部门名称',1661133441032736000,'2022-08-15 09:57:21',NULL),('gfwuk2jq3nudd7g4ct6chtwkftcyl3qm','gfwuk2jlr4jt4zayd5ouodg2xepfufit',1,'create_depart_id_','部门ID','VARCHAR(256)','【系统自动创建】该条记录的创建者的部门ID',1661133441032636000,'2022-08-15 09:57:21',NULL),('gfwuk2jqwtifl2nod3jufruwhi6kzqa5','gfwuk2jlr4jt4zayd5ouodg2xepfufit',1,'parent_id_','父级ID','VARCHAR(256)','【系统自动创建】自动建树服务的父级ID',1661133441031474000,'2022-08-15 09:57:21',NULL),('gfwukc2bd7qy4wurcxvcbc6jp7jx3qpn','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'varchar_256_','VARCHAR 256','VARCHAR(256)','',1660528672879045000,'2022-08-15 09:57:52','2022-08-15 10:03:50'),('gfwukcf5lvwlijbrcfltiopbgv32z2tp','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'varchar_1024_','VARCHAR 1024','VARCHAR(1024)','',1660528673976990000,'2022-08-15 09:57:53','2022-08-15 10:03:54'),('gfwukd5pjemx4oo5dvgvmcotb4rnuaor','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'varchar_4096_','VARCHAR 4096','VARCHAR(4096)','',1660528676393987000,'2022-08-15 09:57:56','2022-08-15 10:03:58'),('gfwukhhflgxsqbk7cgyubm27oyurlvoq','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'tinyint_','TINYINT','TINYINT','',1660528693493042000,'2022-08-15 09:58:13','2022-08-15 10:04:39'),('gfwukkqgqch5ypclctoubjfpllgpsabi','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'int_','INT','INT','',1660528706080265000,'2022-08-15 09:58:26','2022-08-15 10:04:42'),('gfwukl57lpox77eddqmdzaeyshix2mzx','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'bigint_','BIGINT','BIGINT','',1660528708116104000,'2022-08-15 09:58:28','2022-08-15 10:04:44'),('gfwunuyy47kpecsidgrdx4kkdhjbgbto','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'numeric_13_2_','NUMERIC 13(2)','NUMERIC(13,2)','',1660529131161935000,'2022-08-15 10:05:31','2022-08-15 10:11:12'),('gfwuo2iznjtiss5lcaytzbehkhdxedpg','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'numeric_18_4_','NUMERIC 18(4)','NUMERIC(18,4)','',1660529153300668000,'2022-08-15 10:05:53','2022-08-15 10:11:14'),('gfwuo6fhylajbtnddqftji4mzqz4lw3n','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'numeric_23_6_','NUMERIC 23(6)','NUMERIC(23,6)','',1660529169199826000,'2022-08-15 10:06:09','2022-08-15 10:11:17'),('gfwuoex3bikk37kacbktiisddf2sykuh','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'date_1_','DATE 1','DATE','',1660831497960443000,'2022-08-15 10:06:35','2022-08-15 11:34:35'),('gfwuohkdlrw7ht5bdpuvblzyfjfnckdh','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'date_2_','DATE 2','DATE','',1660982469495958500,'2022-08-15 10:06:46','2022-08-15 11:34:37'),('gfwuow3lq7l73oa3cnccku3vmatcuz26','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'date_3_','DATE 3','DATE','',1661057955263716250,'2022-08-15 10:07:44','2022-08-15 11:34:40'),('gfwup5hrw5maguuyc7pfm5o3fiafhxxj','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'varchar_1_','VARCHAR 1','VARCHAR(256)','',1660529188960827500,'2022-08-15 10:08:13','2022-08-15 11:34:24'),('gfwupaoodimwnf63dksdzvxfvzispguf','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'varchar_2_','VARCHAR 2','VARCHAR(256)','',1660529192254327750,'2022-08-15 10:08:26','2022-08-15 11:34:27'),('gfwupdltqsubst46c7iclqt3czh2cswm','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'varchar_3_','VARCHAR 3','VARCHAR(256)','',1660529193901077875,'2022-08-15 10:08:38','2022-08-15 11:34:29'),('gfwuqi76avr7xvq2d6lvumbx74v4rcxk','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'date_','DATE','DATE','',1660529182373827000,'2022-08-15 10:11:04','2022-08-15 10:14:09'),('gfwuqwyn4cny5bixcl5vwnyjaroegrra','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'int_1_','INT 1','INT','',1660529523414216000,'2022-08-15 10:12:03','2022-08-15 10:12:43'),('gfwuqxl2ps5txfgadqesirf6betp3mr3','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'int_2_','INT 2','INT','',1660529526916467000,'2022-08-15 10:12:06','2022-08-15 10:12:44'),('gfwuqy6fphslsrebcrmvcgkqiq74yqw4','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'int_3_','INT 3','INT','',1660529528413998000,'2022-08-15 10:12:08','2022-08-15 10:12:46'),('gfwur3wtv5oswhgjdgkuxozqz4ir6tah','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'numeric_1_','NUMERIC 1','NUMERIC(13,2)','',1660529543962922000,'2022-08-15 10:12:23','2022-08-15 10:12:37'),('gfwur6h4fswp47nidmzfaxqbryonnmin','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'numeric_2_','NUMERIC 2','NUMERIC(13,2)','',1660529553432567000,'2022-08-15 10:12:33','2022-08-15 10:12:38'),('gfwur6k4sulhr6ftdj4sncslgfloby3m','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'numeric_3_','NUMERIC 3','NUMERIC(13,2)','',1660529554889412000,'2022-08-15 10:12:34','2022-08-15 10:12:40'),('gfwz7bfyiafa5sslcn6smfsa2vjdh2lr','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'varchar_4_','VARCHAR 4','VARCHAR(256)','',1660529358657646937,'2022-08-15 15:15:41',NULL),('gfwz7c4qrel4c2kbcd2td3jyithmocch','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'varchar_5_','VARCHAR 5','VARCHAR(256)','',1660529441035931468,'2022-08-15 15:15:44',NULL),('gfwz7d3wo6gq5nccc3ocgwloocjujrle','gfwuk2jlr4jt4zayd5ouodg2xepfufit',0,'varchar_6_','VARCHAR 6','VARCHAR(256)','',1660529482225073734,'2022-08-15 15:15:48',NULL);
/*!40000 ALTER TABLE `sys_table_column` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user`
--

DROP TABLE IF EXISTS `sys_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user` (
                            `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                            `user_code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `user_name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `account_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `password_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `depart_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `sex_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `is_depart_leader_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `valid_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `classification_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `telephone_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `email_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `birth_` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `description_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                            `order_` bigint DEFAULT NULL,
                            `create_at_` datetime DEFAULT NULL,
                            `update_at_` datetime DEFAULT NULL,
                            `login_at_` datetime DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `sys_user_UniqueIndex_depart_id_account_id` (`depart_id_`,`account_id_`),
                            CONSTRAINT `sys_user_ForeignKey_depart_id` FOREIGN KEY (`depart_id_`) REFERENCES `sys_depart` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user`
--

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;
INSERT INTO `sys_user` VALUES ('dhpwodsg5hf9gha39qtsjsryxcf94phs','U000000','孙红雷','U000000','b0a23ec09ca282d3','D000','Male','Yes','Effective','0','','','','',1677072759952908000,'2023-02-15 01:51:47','2023-02-15 01:51:57','2023-02-22 22:13:38'),('U001','admin','系统管理员3','admin','b0a23ec09ca282d3','D005','Male','Yes','Effective','0','','','','',1655102498331312,'2022-06-13 14:41:38','2023-02-22 22:46:37','2023-02-22 22:48:03'),('U002','admin','系统管理员2','admin','b0a23ec09ca282d3','D022','Male','Yes','Effective','0','','','','',1655102498331312,'2022-06-13 14:41:38','2023-02-22 22:46:48','2023-02-22 22:47:13'),('U003','admin','系统管理员1','admin','b0a23ec09ca282d3','D012','Male','Yes','Effective','0','','','','',1655102498331312,'2022-06-13 14:41:38','2023-02-22 22:47:04','2023-02-22 22:40:18'),('U011','ZZB1015','刘青青','ZZB1015','b0a23ec09ca282d3','D011','Female','No','Effective','0','','','','',1654699223948413,'2022-06-08 22:40:23',NULL,NULL),('U012','ZZB1016','徐文娟','ZZB1016','b0a23ec09ca282d3','D011','Female','No','Locked','0','','','','',1654699248413238,'2022-06-08 22:40:48',NULL,NULL),('U013','ZZB1017','韦雪晴','ZZB1017','b0a23ec09ca282d3','D011','Female','No','Effective','0','','','','',1654699269160939,'2022-06-08 22:41:09',NULL,NULL),('U014','ZZB2011','付雪','ZZB2011','b0a23ec09ca282d3','D011','Female','Yes','Disable','0','','','','',1654699333006215,'2022-06-08 22:42:13',NULL,NULL),('U015','ZZB2012','韦心怡','ZZB2012','b0a23ec09ca282d3','D012','Female','No','Effective','0','','','','',1654699361356931,'2022-06-08 22:42:41',NULL,NULL),('U016','ZZB2013','史晓东','ZZB2013','b0a23ec09ca282d3','D012','Male','No','Effective','1','','','','',1654699382109389,'2022-06-08 22:43:02',NULL,NULL),('U017','WZB0001','万文杰','WZB0001','b0a23ec09ca282d3','D001','Male','Yes','Disable','0','','','1978-06-05 00:00:00','',1654699488685696,'2022-06-08 22:44:48','2022-07-31 14:25:35',NULL),('U018','WZB0002','陈世','WZB0002','b0a23ec09ca282d3','D022','Male','Yes','Effective','1','','','','',1654699512539108,'2022-06-08 22:45:12','2023-02-10 10:57:36',NULL),('U019','WZB0003','周亮亮','WZB0003','b0a23ec09ca282d3','D001','Male','No','Effective','0','','','','',1654699488685695,'2022-06-08 22:45:32','2022-06-08 22:47:09',NULL),('U020','WZB1011','钟成成','WZB1011','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699594267998,'2022-06-08 22:46:34','2022-07-04 15:30:43',NULL),('U021','WZB1012','尹子龙','WZB1012','b0a23ec09ca282d3','D001','Male','No','Effective','0','','','','',1654699651885570,'2022-06-08 22:47:31',NULL,NULL),('U022','WZB1013','卢一凡','WZB1013','b0a23ec09ca282d3','D022','Male','No','Locked','0','','','','',1654699670953892,'2022-06-08 22:47:50','2022-06-08 23:26:38',NULL),('U023','WZB1014','许兰','WZB1014','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699694167801,'2022-06-08 22:48:14',NULL,NULL),('U024','WZB1015','万思琪','WZB1015','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699718108239,'2022-06-08 22:48:38',NULL,NULL),('U025','WZB1016','郭婷婷','WZB1016','b0a23ec09ca282d3','D022','Female','No','Locked','0','','','','',1654699736625552,'2022-06-08 22:48:56',NULL,NULL),('U026','WZB2011','夏婷婷','WZB2011','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','1990-09-03 00:00:00','',1654699763282821,'2022-06-08 22:49:23','2022-06-09 10:08:37',NULL),('U027','WZB2012','于慧敏','WZB2012','b0a23ec09ca282d3','D022','Female','No','Locked','0','','','','',1654699967029607,'2022-06-08 22:49:43','2022-06-08 23:23:05',NULL),('U028','WZB2013','杨晴','WZB2013','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699822167101,'2022-06-08 22:50:22',NULL,NULL),('U029','CWB1011','薛安琪','CWB1011','b0a23ec09ca282d3','D003','Female','No','Effective','0','','','','',1654699861748633,'2022-06-08 22:51:01',NULL,NULL),('U030','CWB1012','顾文博','CWB1012','b0a23ec09ca282d3','D003','Male','No','Effective','0','','','','',1654699882782207,'2022-06-08 22:51:22',NULL,NULL),('U031','CWB1013','吴姗姗','CWB1013','b0a23ec09ca282d3','D003','Female','No','Effective','0','','','','',1654699901735822,'2022-06-08 22:51:41',NULL,NULL),('U032','CWB1014','毛慧','CWB1014','b0a23ec09ca282d3','D031','Female','No','Locked','0','','','','',1654699920116870,'2022-06-08 22:52:00',NULL,NULL),('U033','CWB1015','王雪婷','CWB1015','b0a23ec09ca282d3','D031','Female','No','Disable','0','','','','',1654699941711322,'2022-06-08 22:52:21',NULL,NULL),('U034','WZB2032','韦兰兰','WZB2032','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029621,'2022-06-08 22:52:47',NULL,NULL),('U035','WZB2053','胡梦雨','WZB2053','b0a23ec09ca282d3','D022','Female','No','Locked','0','','','','',1654699763282820,'2022-06-08 22:52:47','2022-06-08 23:23:09',NULL),('U036','WZB2054','曹婉莹','WZB2054','b0a23ec09ca282d3','D022','Female','No','Effective','1','','','','',1654699763282822,'2022-06-08 22:52:47','2022-06-09 01:26:04',NULL),('U037','WZB2055','段雅倩','WZB2055','b0a23ec09ca282d3','D022','Female','No','Effective','1','','','1989-03-07 00:00:00','',1654699763282824,'2022-06-08 22:52:47','2022-06-10 02:40:59',NULL),('U038','WZB2056','唐珊珊','WZB2056','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699763282825,'2022-06-08 22:52:47',NULL,NULL),('U039','CWB2011','万亚玲','CWB2011','b0a23ec09ca282d3','D031','Female','No','Effective','0','','','','',1654699967029590,'2022-06-08 22:52:47',NULL,NULL),('U040','0012','杨幂','0012','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1668066511910885000,'2022-11-10 15:48:31',NULL,NULL),('U041','WZB2057','冯甜甜','WZB2057','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029603,'2022-06-08 22:52:47',NULL,NULL),('U042','WZB2061','吕一鸣','WZB2061','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029606,'2022-06-08 22:52:47','2022-06-08 23:22:45',NULL),('U043','WZB2039','郑佳','WZB2039','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029628,'2022-06-08 22:52:47',NULL,NULL),('U044','WZB2040','王翔宇','WZB2040','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029629,'2022-06-08 22:52:47','2022-06-08 23:22:01',NULL),('U045','WZB2038','田志刚','WZB2038','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029627,'2022-06-08 22:52:47','2022-06-08 23:21:54',NULL),('U046','WZB2052','何玲','WZB2052','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029641,'2022-06-08 22:52:47',NULL,NULL),('U047','WZB2051','许海军','WZB2051','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029640,'2022-06-08 22:52:47','2022-06-08 23:22:30',NULL),('U048','WZB2037','崔彬彬','WZB2037','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029626,'2022-06-08 22:52:47','2022-06-08 23:21:50',NULL),('U049','WZB2050','罗明珠','WZB2050','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029639,'2022-06-08 22:52:47',NULL,NULL),('U050','WZB2027','吴志远','WZB2027','b0a23ec09ca282d3','D022','Male','No','Disable','0','','','','',1654699967029616,'2022-06-08 22:52:47','2022-06-08 23:18:49',NULL),('U051','ZZB0002','吴文丽','ZZB0002','b0a23ec09ca282d3','D022','Female','Yes','Locked','0','','','1989-01-26 00:00:00','',1654697293585078,'2022-06-08 22:08:13','2023-02-10 10:58:49',NULL),('U052','WZB2036','曹曼曼','WZB2036','b0a23ec09ca282d3','D022','Female','No','Locked','0','','','','',1654699967029625,'2022-06-08 22:52:47','2022-06-08 23:21:46',NULL),('U053','WZB2047','沈亚男','WZB2047','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029636,'2022-06-08 22:52:47',NULL,NULL),('U054','WZB2026','林家乐','WZB2026','b0a23ec09ca282d3','D022','Male','No','Locked','0','','','','',1654699967029615,'2022-06-08 22:52:47','2022-06-08 23:18:43',NULL),('U055','WZB2035','罗明','WZB2035','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029624,'2022-06-08 22:52:47','2022-06-08 23:21:42',NULL),('U056','WZB2046','吕晶','WZB2046','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029635,'2022-06-08 22:52:47',NULL,NULL),('U057','WZB2049','苏欣悦','WZB2049','b0a23ec09ca282d3','D005','Female','No','Effective','0','','','','',1654699967029638,'2022-06-08 22:52:47',NULL,NULL),('U058','WZB2025','黄海波','WZB2025','b0a23ec09ca282d3','D022','Male','No','Locked','0','','','','',1654699967029614,'2022-06-08 22:52:47','2022-06-08 23:18:33',NULL),('U059','WZB2065','潘杰','WZB2065','b0a23ec09ca282d3','D005','Female','No','Effective','0','','','','',1654699967029605,'2022-06-08 22:52:47',NULL,NULL),('U060','WZB2019','唐强','WZB2019','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029608,'2022-06-08 22:52:47','2022-06-08 23:26:11',NULL),('U061','WZB2020','闫鹏飞','WZB2020','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029609,'2022-06-08 22:52:47','2022-06-08 23:18:04',NULL),('U062','ZZB1011','赵杰','ZZB1011','b0a23ec09ca282d3','D002','Male','Yes','Effective','1','','','','',1654699152535984,'2022-06-08 22:38:07','2022-06-09 01:26:36',NULL),('U063','WZB2045','毛燕','WZB2045','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029634,'2022-06-08 22:52:47',NULL,NULL),('U064','WZB2034','夏文','WZB2034','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029623,'2022-06-08 22:52:47','2022-06-08 23:19:20',NULL),('U065','WZB2048','董兰','WZB2048','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029637,'2022-06-08 22:52:47',NULL,NULL),('U066','WZB2024','章文静','WZB2024','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029613,'2022-06-08 22:52:47',NULL,NULL),('U067','WZB2033','尹玉龙','WZB2033','b0a23ec09ca282d3','D022','Male','No','Locked','0','','','','',1654699967029622,'2022-06-08 22:52:47','2022-06-08 23:19:14',NULL),('U068','WZB2044','陈静静','WZB2044','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029633,'2022-06-08 22:52:47',NULL,NULL),('U069','WZB2043','施永康','WZB2043','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029632,'2022-06-08 22:52:47','2022-06-08 23:22:19',NULL),('U070','WZB2023','曾阳','WZB2023','b0a23ec09ca282d3','D022','Male','No','Disable','0','','','','',1654699967029612,'2022-06-08 22:52:47','2022-06-08 23:18:27',NULL),('U071','WZB2031','黄宏','WZB2031','b0a23ec09ca282d3','D022','Male','No','Disable','0','','','','',1654699967029620,'2022-06-08 22:52:47','2022-06-08 23:19:08',NULL),('U072','WZB2022','洪明明','WZB2022','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029611,'2022-06-08 22:52:47','2022-06-08 23:18:21',NULL),('U073','ZZB1012','姜子豪','ZZB1012','b0a23ec09ca282d3','D022','Male','No','Disable','0','','','','',1654699131578140,'2022-06-08 22:38:51','2022-06-08 22:41:37',NULL),('U074','WZB2042','姜娟','WZB2042','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029631,'2022-06-08 22:52:47',NULL,NULL),('U075','WZB2030','丁文强','WZB2030','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029619,'2022-06-08 22:52:47','2022-06-08 23:19:00',NULL),('U076','WZB2021','余豪','WZB2021','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029610,'2022-06-08 22:52:47','2022-06-08 23:18:17',NULL),('U077','WZB2041','韩志','WZB2041','b0a23ec09ca282d3','D022','Male','No','Effective','0','','','','',1654699967029630,'2022-06-08 22:52:47','2022-06-08 23:22:10',NULL),('U078','WZB2029','马婉莹','WZB2029','b0a23ec09ca282d3','D022','Female','No','Effective','0','','','','',1654699967029618,'2022-06-08 22:52:47',NULL,NULL),('U079','WZB2028','林兵','WZB2028','b0a23ec09ca282d3','D023','Male','No','Effective','0','','','','',1654699967029617,'2022-06-08 22:52:47','2022-06-08 23:18:53',NULL),('U080','ZZB3011','孔志强','ZZB3011','b0a23ec09ca282d3','D002','Male','No','Locked','0','','','','',1654702048487260,'2022-06-08 23:27:28','2022-06-10 02:41:02',NULL),('U081','ZZB0001','万凯旋','ZZB0001','b0a23ec09ca282d3','D001','Male','Yes','Disable','1','','','1980-11-20 00:00:00','',1654697293585077,'2022-06-08 22:07:45','2022-07-11 21:57:53',NULL),('U082','ZZB1013','马龙','ZZB1013','b0a23ec09ca282d3','D002','Male','No','Locked','0','','','','',1654699173493829,'2022-06-08 22:39:33','2022-07-29 16:01:01',NULL),('U083','ZZB1014','蔡兵兵','ZZB1014','b0a23ec09ca282d3','D003','Male','No','Effective','1','','','','',1654699201510832,'2022-06-08 22:40:01','2022-06-08 22:41:24',NULL);
/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tests`
--

DROP TABLE IF EXISTS `tests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tests` (
                         `id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `varchar_256_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `varchar_1024_` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `varchar_4096_` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `tinyint_` tinyint DEFAULT NULL,
                         `int_` int DEFAULT NULL,
                         `bigint_` bigint DEFAULT NULL,
                         `numeric_13_2_` decimal(13,2) DEFAULT NULL,
                         `numeric_18_4_` decimal(18,4) DEFAULT NULL,
                         `numeric_23_6_` decimal(23,6) DEFAULT NULL,
                         `date_` date DEFAULT NULL,
                         `varchar_1_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `varchar_2_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `varchar_3_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `varchar_4_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `varchar_5_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `varchar_6_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `int_1_` int DEFAULT NULL,
                         `int_2_` int DEFAULT NULL,
                         `int_3_` int DEFAULT NULL,
                         `numeric_1_` decimal(13,2) DEFAULT NULL,
                         `numeric_2_` decimal(13,2) DEFAULT NULL,
                         `numeric_3_` decimal(13,2) DEFAULT NULL,
                         `date_1_` date DEFAULT NULL,
                         `date_2_` date DEFAULT NULL,
                         `date_3_` date DEFAULT NULL,
                         `parent_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `order_` bigint DEFAULT NULL,
                         `create_depart_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `create_depart_code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `create_depart_name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `create_user_id_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `create_user_code_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `create_user_name_` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `create_at_` datetime DEFAULT NULL,
                         `update_at_` datetime DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tests`
--

LOCK TABLES `tests` WRITE;
/*!40000 ALTER TABLE `tests` DISABLE KEYS */;
INSERT INTO `tests` VALUES ('1','111',NULL,NULL,NULL,67890123,NULL,NULL,12345.6790,NULL,'2023-02-28','Female','Unknown','Yes',NULL,NULL,NULL,NULL,NULL,NULL,6789.75,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,'2023-02-13 23:09:25'),('2','222',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,'2023-03-06',NULL,NULL,'Yes',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,'2023-02-13 23:09:26'),('3','333',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,'2023-02-13 23:02:59'),('dhp63b8t1hjm4kgyayyu8bhpwppa6zda','111-111',NULL,'111-100',NULL,12345,NULL,4343.43,12345.1235,13579.246869,'2023-02-20',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,1234.50,NULL,NULL,NULL,NULL,NULL,'1',1676300584038709000,'D022','2002','仓库组','U002','admin','系统管理员','2023-02-13 23:03:04','2023-02-13 23:09:23');
/*!40000 ALTER TABLE `tests` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wf_diagram`
--

DROP TABLE IF EXISTS `wf_diagram`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wf_diagram` (
                              `id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
                              `code_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                              `name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                              `icon_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                              `description_` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
                              `model_` text COLLATE utf8mb4_general_ci,
                              `options_` text COLLATE utf8mb4_general_ci,
                              `order_` bigint DEFAULT NULL,
                              `create_at_` datetime DEFAULT NULL,
                              `update_at_` datetime DEFAULT NULL,
                              `publish_at_` datetime DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `wf_diagram_UniqueIndex_code` (`code_`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wf_diagram`
--

LOCK TABLES `wf_diagram` WRITE;
/*!40000 ALTER TABLE `wf_diagram` DISABLE KEYS */;
INSERT INTO `wf_diagram` VALUES ('dhpusz1tvjjj9ovp9129qwgv4t5h5pou','A01','员工请假单','mdi mdi-alarm-snooze','填写请假单，经部门领导和总经理审批后生效','{ \"class\": \"GraphLinksModel\",\n  \"linkFromPortIdProperty\": \"fromPort\",\n  \"linkToPortIdProperty\": \"toPort\",\n  \"nodeDataArray\": [\n{\"category\":\"Start\",\"text\":\"\\u5f00\\u59cb\",\"key\":-1,\"loc\":\"230 120\"},\n{\"category\":\"Execute\",\"text\":\"\\u90e8\\u95e8\\u9886\\u5bfc\",\"key\":-2,\"loc\":\"230 220\"},\n{\"category\":\"Branch\",\"text\":\"\\u8d85\\u8fc73\\u5929\",\"key\":-3,\"loc\":\"230 330\"},\n{\"category\":\"Execute\",\"text\":\"\\u603b\\u7ecf\\u7406\",\"key\":-4,\"loc\":\"440 330\"},\n{\"category\":\"End\",\"text\":\"\\u7ed3\\u675f\",\"key\":-5,\"loc\":\"230 460\"}\n],\n  \"linkDataArray\": [\n{\"from\":-1,\"to\":-2,\"fromPort\":\"B\",\"toPort\":\"T\",\"points\":[230,153,230,163,230,174,230,174,230,185,230,195],\"category\":\"Link\"},\n{\"from\":-2,\"to\":-3,\"fromPort\":\"B\",\"toPort\":\"T\",\"points\":[230,245,230,255,230,269,230,269,230,283,230,293],\"category\":\"Link\"},\n{\"from\":-3,\"to\":-4,\"fromPort\":\"R\",\"toPort\":\"L\",\"visible\":true,\"points\":[291,330,301,330,335,330,335,330,369,330,379,330],\"category\":\"Link\",\"text\":\"\\u662f\"},\n{\"from\":-3,\"to\":-5,\"fromPort\":\"B\",\"toPort\":\"T\",\"visible\":true,\"points\":[230,367,230,377,230,397,230,397,230,417,230,427],\"category\":\"Link\",\"text\":\"\\u5426\"},\n{\"from\":-4,\"to\":-5,\"fromPort\":\"B\",\"toPort\":\"R\",\"points\":[440,355,440,365,440,460,356.5,460,273,460,263,460],\"category\":\"Link\"}\n]}','{\"diagram\":{\"category\":\"Diagram\",\"code_\":\"A01\",\"name_\":\"员工请假单\",\"icon_\":\"mdi mdi-alarm-snooze\",\"description_\":\"填写请假单，经部门领导和总经理审批后生效\",\"keyword_\":\"#type_#: 请假#days_#天\",\"exceed_days_\":\"365\"},\"nodes\":[{\"code_\":\"Start\",\"key\":-1,\"category\":\"Start\",\"revocable_\":true},{\"rejectable_\":true,\"require_reject_comment_\":true,\"executor_custom_num_\":\"3\",\"executor_selectable_num_\":\"3\",\"executor_savable_\":true,\"executor_policy_\":\"StartDepartLeader\",\"executor_script_\":\"[]\",\"key\":-2,\"category\":\"Execute\"},{\"key\":-3,\"category\":\"Branch\"},{\"rejectable_\":true,\"require_reject_comment_\":true,\"executor_custom_num_\":\"0\",\"executor_selectable_num_\":\"0\",\"executor_savable_\":true,\"executor_policy_\":\"None\",\"executor_script_\":\"[]\",\"key\":-4,\"category\":\"Execute\",\"executor_roles_\":\"\",\"executor_name_roles_\":\"\",\"executor_name_departs_\":\"公司办公室\",\"executor_departs_\":\"D000\",\"executor_users_\":\"U003,U002,U001\",\"executor_name_users_\":\"系统管理员1,系统管理员2,系统管理员3\"},{\"code_\":\"End\",\"key\":-5,\"category\":\"End\"}],\"links\":[{\"on_script_\":\"true\",\"from\":-1,\"to\":-2,\"category\":\"Link\"},{\"on_script_\":\"true\",\"from\":-2,\"to\":-3,\"category\":\"Link\"},{\"on_script_\":\"$values[\\\"days_\\\"] > 3;\",\"from\":-3,\"to\":-4,\"category\":\"Link\"},{\"on_script_\":\"$values[\\\"days_\\\"] <=3;\",\"from\":-3,\"to\":-5,\"category\":\"Link\"},{\"on_script_\":\"true\",\"from\":-4,\"to\":-5,\"category\":\"Link\"}]}',1676393796157209000,'2023-02-14 23:45:00','2023-02-22 22:48:40','2023-02-22 22:48:41'),('dhuyswx6vyt2ghk4auvsyyhmtsx445r3','A02','新员工入职单','mdi mdi-account-clock','为新员工办理入职申请','{ \"class\": \"GraphLinksModel\",\n  \"linkFromPortIdProperty\": \"fromPort\",\n  \"linkToPortIdProperty\": \"toPort\",\n  \"nodeDataArray\": [],\n  \"linkDataArray\": []}','{\"diagram\":{\"category\":\"Diagram\",\"code_\":\"A02\",\"name_\":\"新员工入职单\",\"icon_\":\"mdi mdi-account-clock\",\"description_\":\"为新员工办理入职申请\",\"keyword_\":\"#user_name_# #education_#  #work_years_#年\",\"exceed_days_\":\"0\"},\"nodes\":[],\"links\":[]}',1677061235312607000,'2023-02-22 18:20:35','2023-02-22 18:24:18',NULL);
/*!40000 ALTER TABLE `wf_diagram` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wf_flow`
--

DROP TABLE IF EXISTS `wf_flow`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wf_flow` (
                           `id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
                           `values_` text COLLATE utf8mb4_general_ci,
                           `keyword_` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `diagram_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `start_key_` int DEFAULT NULL,
                           `executed_keys_` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `activated_keys_` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `status_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `status_text_` varchar(4096) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `create_at_` datetime DEFAULT NULL,
                           `start_at_` datetime DEFAULT NULL,
                           `active_at_` datetime DEFAULT NULL,
                           `end_at_` datetime DEFAULT NULL,
                           `order_` bigint DEFAULT NULL,
                           `create_depart_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `create_depart_code_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `create_depart_name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `create_user_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `create_user_code_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `create_user_name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           PRIMARY KEY (`id`),
                           KEY `wf_flow_ForeignKey_diagram_id` (`diagram_id_`),
                           CONSTRAINT `wf_flow_ForeignKey_diagram_id` FOREIGN KEY (`diagram_id_`) REFERENCES `wf_options_diagram` (`diagram_id_`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wf_flow`
--

LOCK TABLES `wf_flow` WRITE;
/*!40000 ALTER TABLE `wf_flow` DISABLE KEYS */;
INSERT INTO `wf_flow` VALUES ('dhts924oa63cm11j9gnua1fy7h69skwo','{\"doc_\":\"[]\",\"type_\":\"事假\",\"start_\":\"2023-02-20 00:00:00\",\"end_\":\"2023-03-10 00:00:00\",\"days_\":\"18\",\"reason_\":\"444444444ertertretertret\"}','事假: 请假18天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'[-1,-2,-3,-4,-5]','[]','Finished','流程实例已结束','2023-02-20 22:30:28','2023-02-22 17:00:06','2023-02-22 17:01:11','2023-02-22 17:01:11',1676903428318932000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhts9cfzuu38vrnqafjbsxcvzuk8horo','{\"doc_\":\"[{\\\"id\\\":\\\"dhts9a4w591fmxu5acxtcz8s4rbozs2y\\\",\\\"value\\\":\\\"dhts9a4w591fmxu5acxtcz8s4rbozs2y\\\",\\\"name\\\":\\\"DSC_0034.JPG\\\",\\\"sizetext\\\":\\\"4.57 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhts9atvhu5q8snj9g7cgxyjp6r8a8hp\\\",\\\"value\\\":\\\"dhts9atvhu5q8snj9g7cgxyjp6r8a8hp\\\",\\\"name\\\":\\\"DSC_0044.JPG\\\",\\\"sizetext\\\":\\\"4.15 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhts9bq37nkzbz9695wuhf6g7pvsgszn\\\",\\\"value\\\":\\\"dhts9bq37nkzbz9695wuhf6g7pvsgszn\\\",\\\"name\\\":\\\"DSC_0040.JPG\\\",\\\"sizetext\\\":\\\"4.17 Mb\\\",\\\"status\\\":\\\"server\\\"}]\",\"type_\":\"事假\",\"start_\":\"2023-02-20 00:00:00\",\"end_\":\"2023-02-21 00:00:00\",\"days_\":\"1\",\"reason_\":\"111111\"}','事假: 请假1天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'[-1]','[-2]','Executing','等待 【部门领导】陈世,系统管理员 执行中','2023-02-20 22:31:09','2023-02-20 22:31:17','2023-02-22 02:16:17',NULL,1676903469424771000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhts9mr6q72h47k29pna7v6vgov4csfa','{\"type_\":\"事假\",\"start_\":\"2023-02-20 00:00:00\",\"end_\":\"2023-02-22 00:00:00\",\"days_\":\"2\",\"reason_\":\"22222222\",\"doc_\":\"[]\"}','事假: 请假2天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,NULL,NULL,'Draft','等待流程实例启动','2023-02-20 22:31:38',NULL,NULL,NULL,1676903498367334000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhts9qgz5wnxg9z394cv77pgvymd6xdo','{\"doc_\":\"[{\\\"id\\\":\\\"dhts9sv22s6os2qfanrbkdarkr2r37o8\\\",\\\"value\\\":\\\"dhts9sv22s6os2qfanrbkdarkr2r37o8\\\",\\\"name\\\":\\\"DSC_0041.JPG\\\",\\\"sizetext\\\":\\\"3.3 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhts9tq2jkdv9jdj9h4t3mv2r2sg7p6h\\\",\\\"value\\\":\\\"dhts9tq2jkdv9jdj9h4t3mv2r2sg7p6h\\\",\\\"name\\\":\\\"DSC_0029.JPG\\\",\\\"sizetext\\\":\\\"4.18 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhts9uk88x6ng84w98hs6g2phfogbxwc\\\",\\\"value\\\":\\\"dhts9uk88x6ng84w98hs6g2phfogbxwc\\\",\\\"name\\\":\\\"DSC_0024.JPG\\\",\\\"sizetext\\\":\\\"4.14 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhts9xchkh68vn579zfcj4q3nwa4cbvb\\\",\\\"value\\\":\\\"dhts9xchkh68vn579zfcj4q3nwa4cbvb\\\",\\\"name\\\":\\\"DSC_0050.JPG\\\",\\\"sizetext\\\":\\\"3.97 Mb\\\",\\\"status\\\":\\\"server\\\"}]\",\"type_\":\"陪产假\",\"start_\":\"2023-02-20 00:00:00\",\"end_\":\"2023-02-25 00:00:00\",\"days_\":\"5\",\"reason_\":\"55555\"}','陪产假: 请假5天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'[-1,-2,-3,-4,-5]','[]','Finished','流程实例已结束','2023-02-20 22:31:53','2023-02-20 23:07:12','2023-02-22 16:38:23','2023-02-22 16:38:23',1676903513067336000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhtsbzdyhpyzt31v9hjsk415238xvtqh','{\"doc_\":\"[{\\\"id\\\":\\\"dhtsbybbxvhyumt79xuvfq4jwndu6cxu\\\",\\\"value\\\":\\\"dhtsbybbxvhyumt79xuvfq4jwndu6cxu\\\",\\\"name\\\":\\\"DSC_0041.JPG\\\",\\\"sizetext\\\":\\\"3.3 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhtsbz83so2qov4j9cg9ohmnqn5qu5kb\\\",\\\"value\\\":\\\"dhtsbz83so2qov4j9cg9ohmnqn5qu5kb\\\",\\\"name\\\":\\\"DSC_0045.JPG\\\",\\\"sizetext\\\":\\\"4.41 Mb\\\",\\\"status\\\":\\\"server\\\"}]\",\"type_\":\"事假\",\"start_\":\"2023-02-20 00:00:00\",\"end_\":\"2023-03-31 00:00:00\",\"days_\":\"39\",\"reason_\":\"qqqqqqq555555555555rrrrrrrrr\"}','事假: 请假39天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'[-1,-2,-3,-4,-5]','[]','Finished','流程实例已结束','2023-02-20 22:36:45','2023-02-20 22:36:50','2023-02-22 17:01:16','2023-02-22 17:01:16',1676903805273642000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhtsv5njawd1cuz6aq9su7soqcuzggog','{\"doc_\":\"[{\\\"id\\\":\\\"dhtsv26abmxmssafaosayuy4ojsadknj\\\",\\\"value\\\":\\\"dhtsv26abmxmssafaosayuy4ojsadknj\\\",\\\"name\\\":\\\"DSC_0045.JPG\\\",\\\"sizetext\\\":\\\"4.41 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhtsv2j1m61st7km9am9mcurdd99unkt\\\",\\\"value\\\":\\\"dhtsv2j1m61st7km9am9mcurdd99unkt\\\",\\\"name\\\":\\\"DSC_0036.JPG\\\",\\\"sizetext\\\":\\\"4.09 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhtsv3rc9g9wj3sua2vc9jnvcqu8u5r8\\\",\\\"value\\\":\\\"dhtsv3rc9g9wj3sua2vc9jnvcqu8u5r8\\\",\\\"name\\\":\\\"DSC_0029.JPG\\\",\\\"sizetext\\\":\\\"4.18 Mb\\\",\\\"status\\\":\\\"server\\\"}]\",\"type_\":\"事假\",\"start_\":\"2023-02-20 00:00:00\",\"end_\":\"2023-02-22 00:00:00\",\"days_\":\"2\",\"reason_\":\"sssssss\"}','事假: 请假2天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'[-1,-2,-3,-5]','[]','Finished','流程实例已结束','2023-02-20 23:11:14','2023-02-22 15:16:43','2023-02-22 17:02:02','2023-02-22 17:02:02',1676905874483801000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhtta9t2qvy8qhoqaqmt15tnfs3jgj92','{\"doc_\":\"[{\\\"id\\\":\\\"dhtta9d4frpj6vd79kq9zc7kzk9q44h6\\\",\\\"value\\\":\\\"dhtta9d4frpj6vd79kq9zc7kzk9q44h6\\\",\\\"name\\\":\\\"DSC_0036.JPG\\\",\\\"sizetext\\\":\\\"4.09 Mb\\\",\\\"status\\\":\\\"server\\\"}]\",\"type_\":\"事假\",\"start_\":\"2023-02-20 00:00:00\",\"end_\":\"2023-02-22 00:00:00\",\"days_\":\"2\",\"reason_\":\"呜呜呜呜\"}','事假: 请假2天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'[-1]','[-2]','Executing','等待 【部门领导】陈世,系统管理员 执行中','2023-02-20 23:41:23','2023-02-22 03:37:48','2023-02-22 03:37:48',NULL,1676907683047621000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuhvoaz38fkmry2af79zoccfgadxjds','{\"doc_\":\"[{\\\"id\\\":\\\"dhuhvrn3b5gs5s7tarpubarnht19a4bg\\\",\\\"value\\\":\\\"dhuhvrn3b5gs5s7tarpubarnht19a4bg\\\",\\\"name\\\":\\\"1AFC54EE270D06A0A234DC11D27D19AE (1) (1).MP4\\\",\\\"sizetext\\\":\\\"8.49 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhuhvt645cwo2m9z99jupm7op51hjsyc\\\",\\\"value\\\":\\\"dhuhvt645cwo2m9z99jupm7op51hjsyc\\\",\\\"name\\\":\\\"IMG_7456 (1).MP4\\\",\\\"sizetext\\\":\\\"78.25 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhuhvuvdztsb9q2w98pt3ut2yc3jvo5r\\\",\\\"value\\\":\\\"dhuhvuvdztsb9q2w98pt3ut2yc3jvo5r\\\",\\\"name\\\":\\\"IMG_7098.JPG\\\",\\\"sizetext\\\":\\\"1.08 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhuhvvtk8hgcojxdanxa65bz3jnk7nyx\\\",\\\"value\\\":\\\"dhuhvvtk8hgcojxdanxa65bz3jnk7nyx\\\",\\\"name\\\":\\\"IMG_7913.PNG\\\",\\\"sizetext\\\":\\\"7.29 Mb\\\",\\\"status\\\":\\\"server\\\"}]\",\"type_\":\"事假\",\"start_\":\"2023-02-22 00:00:00\",\"end_\":\"2023-02-25 00:00:00\",\"days_\":\"3\",\"reason_\":\"sssssssssssss\"}','事假: 请假3天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'[-1]','[-2]','Executing','等待 【部门领导】系统管理员 执行中','2023-02-22 01:22:25','2023-02-22 01:23:22','2023-02-22 01:23:48',NULL,1677000145133821000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhukvr5oqnbbrp6n92o9wbath7y9y5an','{\"doc_\":\"[{\\\"id\\\":\\\"dhukvq939g95dpchabks7rgv7gbnsvvx\\\",\\\"value\\\":\\\"dhukvq939g95dpchabks7rgv7gbnsvvx\\\",\\\"name\\\":\\\"DSC_0029.JPG\\\",\\\"sizetext\\\":\\\"4.18 Mb\\\",\\\"status\\\":\\\"server\\\"}]\",\"type_\":\"事假\",\"start_\":\"2023-02-22 00:00:00\",\"end_\":\"2023-02-24 00:00:00\",\"days_\":\"2\",\"reason_\":\"ssssss\"}','事假: 请假2天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'[-1,-2,-3,-5]','[]','Finished','流程实例已结束','2023-02-22 03:39:08','2023-02-22 03:39:12','2023-02-22 17:02:09','2023-02-22 17:02:09',1677008348369237000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuw68gzv82gayak9v8t6c42qz89nb35','{\"type_\":\"事假\",\"start_\":\"2023-02-22 00:00:00\",\"end_\":\"2023-02-25 00:00:00\",\"days_\":\"3\",\"reason_\":\"2233232\",\"doc_\":\"[{\\\"id\\\":\\\"dhuw683x52abknhhan5v4jqn76a75sd8\\\",\\\"value\\\":\\\"dhuw683x52abknhhan5v4jqn76a75sd8\\\",\\\"name\\\":\\\"DSC_0034.JPG\\\",\\\"sizetext\\\":\\\"4.57 Mb\\\",\\\"status\\\":\\\"server\\\"}]\"}','事假: 请假3天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'[-1,-2,-3,-5]','[]','Finished','流程实例已结束','2023-02-22 15:22:05','2023-02-22 15:22:08','2023-02-22 17:01:55','2023-02-22 17:01:55',1677050525523807000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuwfk31b6fkj1yoaruunhwokxs7yvys','{\"doc_\":\"[{\\\"id\\\":\\\"dhuwff3ayywz8z5w9j8uf5tho5x1wwy8\\\",\\\"value\\\":\\\"dhuwff3ayywz8z5w9j8uf5tho5x1wwy8\\\",\\\"name\\\":\\\"DSC_0029.JPG\\\",\\\"sizetext\\\":\\\"4.18 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhuwfh4xhzptz4vk9p8bn2ptb69fn8b9\\\",\\\"value\\\":\\\"dhuwfh4xhzptz4vk9p8bn2ptb69fn8b9\\\",\\\"name\\\":\\\"DSC_0036.JPG\\\",\\\"sizetext\\\":\\\"4.09 Mb\\\",\\\"status\\\":\\\"server\\\"},{\\\"id\\\":\\\"dhuwfjjzwptrv61wag4bqznbd77myjzf\\\",\\\"value\\\":\\\"dhuwfjjzwptrv61wag4bqznbd77myjzf\\\",\\\"name\\\":\\\"DSC_0044.JPG\\\",\\\"sizetext\\\":\\\"4.15 Mb\\\",\\\"status\\\":\\\"server\\\"}]\",\"type_\":\"路途假\",\"start_\":\"2023-02-22 00:00:00\",\"end_\":\"2023-03-11 00:00:00\",\"days_\":\"17\",\"reason_\":\"3232234324354353543534543543534543cccvvvvvvbbbbbbccxxxxx\"}','路途假: 请假17天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'[-1,-2,-3,-4,-5]','[]','Finished','流程实例已结束','2023-02-22 15:39:48','2023-02-22 15:40:04','2023-02-22 21:33:30','2023-02-22 21:33:30',1677051588710242000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhv3qqsjo3gjdqas9xgbnnqj9wostqd8','{\"type_\":\"事假\",\"start_\":\"2023-02-22 00:00:00\",\"end_\":\"2023-02-24 00:00:00\",\"days_\":\"2\",\"reason_\":\"wwww\",\"doc_\":\"[{\\\"id\\\":\\\"dhv3qqrc7f3oxmqb9rmanjp8o8rsh9gf\\\",\\\"value\\\":\\\"dhv3qqrc7f3oxmqb9rmanjp8o8rsh9gf\\\",\\\"name\\\":\\\"29x29.png\\\",\\\"sizetext\\\":\\\"842 b\\\",\\\"status\\\":\\\"server\\\"}]\"}','事假: 请假2天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'[-1]','[-2]','Executing','等待 【部门领导】系统管理员3 执行中','2023-02-22 22:48:59','2023-02-22 22:49:04','2023-02-22 22:49:04',NULL,1677077339496171000,'D005','04','后勤保障部','U001','admin','系统管理员3'),('dhv3qz9b4dp24y5taurujpv6xrsympzx','{\"type_\":\"事假\",\"start_\":\"2023-02-22 00:00:00\",\"end_\":\"2023-03-08 00:00:00\",\"days_\":\"14\",\"reason_\":\"我千千万万期望\",\"doc_\":\"[{\\\"id\\\":\\\"dhv3qz2ysoxzqk6ba1puvux8nzga435y\\\",\\\"value\\\":\\\"dhv3qz2ysoxzqk6ba1puvux8nzga435y\\\",\\\"name\\\":\\\"40x4ssss0.png\\\",\\\"sizetext\\\":\\\"1.35 Kb\\\",\\\"status\\\":\\\"server\\\"}]\"}','事假: 请假14天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'[-1]','[-2]','Executing','等待 【部门领导】系统管理员3 执行中','2023-02-22 22:49:33','2023-02-22 22:49:36','2023-02-22 22:49:36',NULL,1677077373763127000,'D005','04','后勤保障部','U001','admin','系统管理员3'),('dhv3r4851wmqsxn2aqptf67fptxy8hc4','{\"type_\":\"事假\",\"start_\":\"2023-02-22 00:00:00\",\"end_\":\"2023-03-10 00:00:00\",\"days_\":\"16\",\"reason_\":\"2121312321\",\"doc_\":\"[]\"}','事假: 请假16天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,NULL,NULL,'Draft','等待流程实例启动','2023-02-22 22:49:48',NULL,NULL,NULL,1677077388409439000,'D005','04','后勤保障部','U001','admin','系统管理员3'),('dhv3r4vq2m2j2518amyavnoxy1anp561','{\"type_\":\"事假\",\"start_\":\"2023-02-22 00:00:00\",\"end_\":\"2023-03-10 00:00:00\",\"days_\":\"16\",\"reason_\":\"2121312321\",\"doc_\":\"[]\"}','事假: 请假16天','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,NULL,NULL,'Draft','等待流程实例启动','2023-02-22 22:49:51',NULL,NULL,NULL,1677077391546712000,'D005','04','后勤保障部','U001','admin','系统管理员3');
/*!40000 ALTER TABLE `wf_flow` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wf_flow_executors`
--

DROP TABLE IF EXISTS `wf_flow_executors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wf_flow_executors` (
                                     `id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
                                     `diagram_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                     `key_` int DEFAULT NULL,
                                     `order_` bigint DEFAULT NULL,
                                     `executor_user_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                     `executor_user_name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                     `create_at_` datetime DEFAULT NULL,
                                     `create_user_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                     `create_user_name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                     PRIMARY KEY (`id`),
                                     KEY `wf_flow_executors_ForeignKey_diagram_id` (`diagram_id_`),
                                     CONSTRAINT `wf_flow_executors_ForeignKey_diagram_id` FOREIGN KEY (`diagram_id_`) REFERENCES `wf_options_diagram` (`diagram_id_`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wf_flow_executors`
--

LOCK TABLES `wf_flow_executors` WRITE;
/*!40000 ALTER TABLE `wf_flow_executors` DISABLE KEYS */;
/*!40000 ALTER TABLE `wf_flow_executors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wf_flow_task`
--

DROP TABLE IF EXISTS `wf_flow_task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wf_flow_task` (
                                `id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
                                `flow_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `executed_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `diagram_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `key_` int DEFAULT NULL,
                                `category_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `code_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `executor_user_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `executor_user_name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `activated_at_` datetime DEFAULT NULL,
                                `canceled_at_` datetime DEFAULT NULL,
                                `executed_at_` datetime DEFAULT NULL,
                                `comment_` text COLLATE utf8mb4_general_ci,
                                `status_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `order_` bigint DEFAULT NULL,
                                `executed_depart_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `executed_depart_code_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `executed_depart_name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `executed_user_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `executed_user_code_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                `executed_user_name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                PRIMARY KEY (`id`),
                                KEY `wf_flow_task_ForeignKey_diagram_id` (`diagram_id_`),
                                CONSTRAINT `wf_flow_task_ForeignKey_diagram_id` FOREIGN KEY (`diagram_id_`) REFERENCES `wf_options_diagram` (`diagram_id_`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wf_flow_task`
--

LOCK TABLES `wf_flow_task` WRITE;
/*!40000 ALTER TABLE `wf_flow_task` DISABLE KEYS */;
INSERT INTO `wf_flow_task` VALUES ('dhts9f9rdwjqvhw79oztqhq4nyr76xn3','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhts9fcoou1j52kw9tjv4nfrd8vxry82','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-20 22:31:17','2023-02-20 22:40:18',NULL,NULL,'Canceled',1676903477243937000,NULL,NULL,NULL,NULL,NULL,NULL),('dhts9fd2c59bv62n9josyw39gc96899a','dhts9cfzuu38vrnqafjbsxcvzuk8horo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-20 22:31:17',NULL,'2023-02-20 22:31:17','1111','Executed Auto',1676903477236776000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhts9ffajkzuzkfdascvbkx7o8fswtrh','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhts9fcoou1j52kw9tjv4nfrd8vxry82','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-20 22:31:17','2023-02-20 22:40:18',NULL,NULL,'Canceled',1676903477242573000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtsc1jcgxjxjx2xaq6cbk8vk9p5aasz','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhtsc1qjx81brdyrasythzv3caccmsft','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-20 22:36:50','2023-02-21 00:00:29',NULL,NULL,'Canceled',1676903810890598000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtsc1jzuqvv9m3gasdcdy2wc9nxgbpx','dhtsbzdyhpyzt31v9hjsk415238xvtqh',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-20 22:36:50',NULL,'2023-02-20 22:36:50','','Executed Auto',1676903810889166000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhtsc1nggoqhcpy3ad2s7b1jyyh825uo','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhtsc1qjx81brdyrasythzv3caccmsft','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-20 22:36:50','2023-02-21 00:00:29',NULL,NULL,'Canceled',1676903810891721000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtsc1nkgtnh1zmyasq9mb9o2gtc981k','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhtsc1qjx81brdyrasythzv3caccmsft','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-20 22:36:50','2023-02-21 00:00:29',NULL,NULL,'Canceled',1676903810891238000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtst921ry1c6bg7aymt36b1sy146ov6','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhtst92fjkkyn8bvah1uyz4v3ct3azyq','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-20 23:07:12','2023-02-20 23:37:47',NULL,NULL,'Canceled',1676905632044733000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtst958r8k1f8b59oyu82ucr34c58qa','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhtst92fjkkyn8bvah1uyz4v3ct3azyq','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-20 23:07:12','2023-02-20 23:37:47',NULL,NULL,'Canceled',1676905632045402000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtst962a58gw3mma8vvxqrrdb9mx9m8','dhts9qgz5wnxg9z394cv77pgvymd6xdo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-20 23:07:12',NULL,'2023-02-20 23:07:12','1111','Executed Auto',1676905632042439000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhtst97kvn99m81s9djc7xj732mn4k5r','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhtst92fjkkyn8bvah1uyz4v3ct3azyq','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-20 23:07:12','2023-02-20 23:37:47',NULL,NULL,'Canceled',1676905632046190000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtt8oa415xmuwqaaftankc23g2mcnss','dhts9qgz5wnxg9z394cv77pgvymd6xdo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-20 23:37:53',NULL,'2023-02-20 23:37:53','','Executed Auto',1676907473636235000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhtt8odag49jma6w9hscxff8xcntt23n','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhtt8odwbcyzn75t92dtocksytbm8qr9','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-20 23:37:53','2023-02-20 23:37:58',NULL,NULL,'Canceled',1676907473637631000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtt8oh4s2wfp6zga6waqw6a3njwa4m7','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhtt8odwbcyzn75t92dtocksytbm8qr9','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-20 23:37:53','2023-02-20 23:37:58',NULL,NULL,'Canceled',1676907473638556000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtt8oh6sjcufd699jysuhmgx5f7uc5x','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhtt8odwbcyzn75t92dtocksytbm8qr9','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-20 23:37:53','2023-02-20 23:37:58',NULL,NULL,'Canceled',1676907473638163000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtt8q9yxvfrn7nya8w9sy7j3hvgdgt9','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhtt8qhgs93rqmya9p599jp396py22kg','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-20 23:38:01','2023-02-20 23:38:26',NULL,NULL,'Canceled',1676907481603643000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtt8qa7a7655d93ah6bqpgoxss8uosk','dhts9qgz5wnxg9z394cv77pgvymd6xdo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-20 23:38:01',NULL,'2023-02-20 23:38:01','','Executed Auto',1676907481602490000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhtt8qaop2z2gvo8a7gsfmc6rz45wroy','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhtt8qhgs93rqmya9p599jp396py22kg','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-20 23:38:01','2023-02-20 23:38:26',NULL,NULL,'Canceled',1676907481604738000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtt8qbmox92mojxah9sgrsx4xnjkp8g','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhtt8qhgs93rqmya9p599jp396py22kg','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-20 23:38:01','2023-02-20 23:38:26',NULL,NULL,'Canceled',1676907481604279000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtt92nwhk3brsvs94atvckr8b2kdbhh','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhtt92qrr7ofrgkpa7gsf3p769xwvn1k','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-20 23:38:46','2023-02-21 00:00:27',NULL,NULL,'Canceled',1676907526069358000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtt92of53z15vsm9z7926ovtj28pq4g','dhts9qgz5wnxg9z394cv77pgvymd6xdo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-20 23:38:46',NULL,'2023-02-20 23:38:46','珊珊珊珊珊珊珊珊珊珊珊珊珊珊','Executed Auto',1676907526067917000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhtt92r45r6tdkp99dst47cc3b55o9k1','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhtt92qrr7ofrgkpa7gsf3p769xwvn1k','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-20 23:38:46','2023-02-21 00:00:27',NULL,NULL,'Canceled',1676907526070008000,NULL,NULL,NULL,NULL,NULL,NULL),('dhtt92radyfw2gt6akdvfhsmc2bpo3wq','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhtt92qrr7ofrgkpa7gsf3p769xwvn1k','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-20 23:38:46','2023-02-21 00:00:27',NULL,NULL,'Canceled',1676907526068823000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttksauhnsocrg1ar7btocoaygj4uxw','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhttksdf47b6nuw7acbbqdaapdhjqxn6','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-20 23:59:29','2023-02-21 00:00:25',NULL,NULL,'Canceled',1676908769953888000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttksb5zhrdxchmar1tfrwpjz9avvv5','dhts9cfzuu38vrnqafjbsxcvzuk8horo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-20 23:59:29',NULL,'2023-02-20 23:59:29','','Executed Auto',1676908769951349000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhttksg9gts7k94ta63tumnuhuo2pzca','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhttksdf47b6nuw7acbbqdaapdhjqxn6','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-20 23:59:29','2023-02-21 00:00:25',NULL,NULL,'Canceled',1676908769954412000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttm9kbxjnfsk899mm9wuwqmpaqgt98','dhtsbzdyhpyzt31v9hjsk415238xvtqh',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-21 00:00:34',NULL,'2023-02-21 00:00:34','','Executed Auto',1676908834501787000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhttm9oyws7yvx1h98tvdvxr7yu86gng','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhttm9jwrq72y7wm9a1sxofmxbq64pa6','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-21 00:00:34','2023-02-21 00:01:19',NULL,NULL,'Canceled',1676908834503582000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttm9pnmvuzg3vxafgskjy2gz8npxp5','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhttm9jwrq72y7wm9a1sxofmxbq64pa6','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-21 00:00:34','2023-02-21 00:01:19',NULL,NULL,'Canceled',1676908834503226000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttm9prcfnwrucjagot63yquubv48so','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhttm9jwrq72y7wm9a1sxofmxbq64pa6','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-21 00:00:34','2023-02-21 00:01:19',NULL,NULL,'Canceled',1676908834502778000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttmaawdvtz1cdj9r7cyn28xg2631zw','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhttmadunnndahwhauva157kkvnkujpx','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-21 00:00:37','2023-02-21 00:01:17',NULL,NULL,'Canceled',1676908837938063000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttmabj5acay6jz9wmvd89kofp693q4','dhts9qgz5wnxg9z394cv77pgvymd6xdo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-21 00:00:37',NULL,'2023-02-21 00:00:37','','Executed Auto',1676908837935582000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhttmabjax8xh33oapwsbqupostbtqsw','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhttmadunnndahwhauva157kkvnkujpx','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-21 00:00:37','2023-02-21 00:01:17',NULL,NULL,'Canceled',1676908837937485000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttmabu996byh5o925s8q14p2n75mbp','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhttmadunnndahwhauva157kkvnkujpx','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-21 00:00:37','2023-02-21 00:01:17',NULL,NULL,'Canceled',1676908837937034000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttmbm7g9ngsjrh92s9na6kutzbvxw6','dhts9cfzuu38vrnqafjbsxcvzuk8horo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-21 00:00:42',NULL,'2023-02-21 00:00:42','','Executed Auto',1676908842801374000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhttmbpojowqhgxfakxbbgj1tjahxsgd','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhttmbrygfj35z1t9pxucj1uq1cd4y1p','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-21 00:00:42','2023-02-21 00:01:14',NULL,NULL,'Canceled',1676908842802738000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttmpjkmzm7obmfa45u6a2ymds2kw1r','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhttmpr843osaz81a8rv8wcqbjvmobbu','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-21 00:01:26','2023-02-22 15:56:32',NULL,NULL,'Canceled',1676908886153492000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttmpk2ny5wtt259ohsdjqbn9wob13z','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhttmpr843osaz81a8rv8wcqbjvmobbu','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-21 00:01:26','2023-02-22 15:56:32',NULL,NULL,'Canceled',1676908886153139000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttmpqvaw9c5k31a6avknk4646w9gy9','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhttmpr843osaz81a8rv8wcqbjvmobbu','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-21 00:01:26',NULL,'2023-02-22 15:56:32','','Executed Accepted',1676908886152732000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhttmprbqvnsgubf9x7vd67vw8krkhfp','dhts9qgz5wnxg9z394cv77pgvymd6xdo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-21 00:01:26',NULL,'2023-02-21 00:01:26','','Executed Auto',1676908886151474000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhttmqn1vx2ddnynac1uyvpkt78skyqg','dhtsbzdyhpyzt31v9hjsk415238xvtqh',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-21 00:01:30',NULL,'2023-02-21 00:01:30','','Executed Auto',1676908890546971000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhttmqr3d32u1854ag3a3ashczb7bnb1','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhttmqqh9uvzmf5zaz4bw14rkybajxjq','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-21 00:01:30','2023-02-22 16:10:33',NULL,NULL,'Canceled',1676908890547918000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttmqrsj74ccbq9a5vu4377t5q11xsy','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhttmqqh9uvzmf5zaz4bw14rkybajxjq','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-21 00:01:30',NULL,'2023-02-22 16:10:33','werewrw','Executed Accepted',1676908890548274000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhttms43qb6b21sx96yszdksx63jtadm','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhttms6gnaqw2a5y9z89g9tfuqy6yjzr','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-21 00:01:36','2023-02-22 01:05:31',NULL,'流程发起者已撤回','Canceled',1676908896835687000,NULL,NULL,NULL,NULL,NULL,NULL),('dhttms7zbn1a8pkvabzuaqqkrb4x9c4f','dhts9cfzuu38vrnqafjbsxcvzuk8horo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-21 00:01:36',NULL,'2023-02-21 00:01:36','','Executed Auto',1676908896834526000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuhjx9572nxqsvvab6b5b23baxcp95r','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhuhjx99h3dg6wog9twvccg9pfc86dvh','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-22 00:59:33','2023-02-22 01:05:31',NULL,'流程发起者已撤回','Canceled',1676998773436733000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuhjxbfshyhb8459cyvup5s9g3u3qoy','dhts9cfzuu38vrnqafjbsxcvzuk8horo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 00:59:33',NULL,'2023-02-22 00:59:33','','Executed Auto',1676998773435100000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuhjxdhqwp657mga6193byhcovkat65','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhuhjx99h3dg6wog9twvccg9pfc86dvh','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 00:59:33','2023-02-22 01:05:31',NULL,'流程发起者已撤回','Canceled',1676998773435990000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuhjxgrbjpq8oxnazgatwyvtfpsr3wn','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhuhjx99h3dg6wog9twvccg9pfc86dvh','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 00:59:33','2023-02-22 01:05:31',NULL,'流程发起者已撤回','Canceled',1676998773436376000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuhnpu93ru3k9qx9kjchozp1scdxx5f','dhts9cfzuu38vrnqafjbsxcvzuk8horo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 01:05:27',NULL,'2023-02-22 01:05:27','','Executed Auto',1676999127203276000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuhnpvbwcgn28k3am4arou3hot4ygvp','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhuhnpx7gw6h9c9xaz2c72oy2p7hjpog','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-22 01:05:27','2023-02-22 01:05:31',NULL,'流程发起者已撤回','Canceled',1676999127205143000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuhnpw1z3xftpcdajucvyt48mprsdsa','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhuhnpx7gw6h9c9xaz2c72oy2p7hjpog','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 01:05:27','2023-02-22 01:05:31',NULL,'流程发起者已撤回','Canceled',1676999127204289000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuhnpwzz44q9q29aq3uudhdzt2g5oxy','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhuhnpx7gw6h9c9xaz2c72oy2p7hjpog','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 01:05:27','2023-02-22 01:05:31',NULL,'流程发起者已撤回','Canceled',1676999127205942000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuhw3ku8o5fsyp5a54tcmbsdgn8kwvn','dhuhvoaz38fkmry2af79zoccfgadxjds','dhuhw3nkacqpfc419tda13r7vn2shcwc','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 01:23:22','2023-02-22 01:23:40',NULL,'流程发起者已撤回','Canceled',1677000202150113000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuhw3nhp513ghb19gycqqn1bm9ojz5d','dhuhvoaz38fkmry2af79zoccfgadxjds','dhuhw3nkacqpfc419tda13r7vn2shcwc','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 01:23:22','2023-02-22 01:23:40',NULL,'流程发起者已撤回','Canceled',1677000202150473000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuhw3nzdj58y1akayatasxhpf4x8hf1','dhuhvoaz38fkmry2af79zoccfgadxjds',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 01:23:22',NULL,'2023-02-22 01:23:22','','Executed Auto',1677000202149174000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuhwa4tm9cmkpvk9zwt35738fpf9924','dhuhvoaz38fkmry2af79zoccfgadxjds',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 01:23:48',NULL,'2023-02-22 01:23:48','','Executed Auto',1677000228282164000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuhwa6y8f1tuj4n9joa3428ankbf4ap','dhuhvoaz38fkmry2af79zoccfgadxjds','dhuhwa8hxuqx3bpx977b8jyg5majmpad','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 01:23:48',NULL,NULL,NULL,'Executing',1677000228283355000,NULL,NULL,NULL,NULL,NULL,NULL),('dhujow9dvgxsy9n59kta8b1rdo5oqrum','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhujowaa7z6ngpqfaj2a2wrtnv7s35a9','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 02:16:17',NULL,NULL,NULL,'Executing',1677003377854564000,NULL,NULL,NULL,NULL,NULL,NULL),('dhujowftcghsf4wxatxtxjwsjjw3otqx','dhts9cfzuu38vrnqafjbsxcvzuk8horo','dhujowaa7z6ngpqfaj2a2wrtnv7s35a9','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 02:16:17',NULL,NULL,NULL,'Executing',1677003377856095000,NULL,NULL,NULL,NULL,NULL,NULL),('dhujowg7zdsr64xn9dzvct3hmqxgy7q7','dhts9cfzuu38vrnqafjbsxcvzuk8horo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 02:16:17',NULL,'2023-02-22 02:16:17','xxxxxxxxxx','Executed Auto',1677003377852870000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhukv454s25t5tkb9agu5x7jh7djnn7y','dhtta9t2qvy8qhoqaqmt15tnfs3jgj92','dhukv44p3rykzgt39rut1awzbgqff46n','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 03:37:48',NULL,NULL,NULL,'Executing',1677008268053922000,NULL,NULL,NULL,NULL,NULL,NULL),('dhukv45yuarsr3u59bkts6tworpcq49r','dhtta9t2qvy8qhoqaqmt15tnfs3jgj92','dhukv44p3rykzgt39rut1awzbgqff46n','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 03:37:48',NULL,NULL,NULL,'Executing',1677008268053031000,NULL,NULL,NULL,NULL,NULL,NULL),('dhukv487ayaqsj5u9auttchx4oatxpnx','dhtta9t2qvy8qhoqaqmt15tnfs3jgj92',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 03:37:48',NULL,'2023-02-22 03:37:48','xzzzzzzzzzzzzzzz星星星星星星','Executed Auto',1677008268050811000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhukvs1vtm4v1j8f94cupzkts1yktft5','dhukvr5oqnbbrp6n92o9wbath7y9y5an',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 03:39:12',NULL,'2023-02-22 03:39:12','ssss','Executed Auto',1677008352584525000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhukvs35u2vymsv69vybnrcz7h6fc3r4','dhukvr5oqnbbrp6n92o9wbath7y9y5an','dhukvs1xcn4tjj6f954t29y9374at416','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 03:39:12','2023-02-22 17:02:08',NULL,NULL,'Canceled',1677008352585552000,NULL,NULL,NULL,NULL,NULL,NULL),('dhukvs4scryz57oj9d3sodr99mua3ycd','dhukvr5oqnbbrp6n92o9wbath7y9y5an','dhukvs1xcn4tjj6f954t29y9374at416','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 03:39:12',NULL,'2023-02-22 17:02:08','3sssss','Executed Accepted',1677008352587087000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhukvs8q27k1cgmj9j6ubmqwcvsztrbd','dhukvr5oqnbbrp6n92o9wbath7y9y5an','dhukvs1xcn4tjj6f954t29y9374at416','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-22 03:39:12','2023-02-22 17:02:08',NULL,NULL,'Canceled',1677008352586718000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuw3qvzkmnophy69szsr3oz8hhmt72x','dhtsv5njawd1cuz6aq9su7soqcuzggog','dhuw3qz7k92x18m6a5nsd3kuz35m6r55','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-22 15:16:43','2023-02-22 17:02:02',NULL,NULL,'Canceled',1677050203221734000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuw3qw53gz8chbcadpahkqpcahvd6vc','dhtsv5njawd1cuz6aq9su7soqcuzggog','dhuw3qz7k92x18m6a5nsd3kuz35m6r55','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 15:16:43','2023-02-22 17:02:02',NULL,NULL,'Canceled',1677050203222876000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuw3qw7xnhbx9u2968cm26nq6o8wqch','dhtsv5njawd1cuz6aq9su7soqcuzggog',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 15:16:43',NULL,'2023-02-22 15:16:43','','Executed Auto',1677050203220215000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuw3qxwad6q7at8aawbcgana12pmr26','dhtsv5njawd1cuz6aq9su7soqcuzggog','dhuw3qz7k92x18m6a5nsd3kuz35m6r55','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 15:16:43',NULL,'2023-02-22 17:02:02','4354336547567','Executed Accepted',1677050203222299000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuw695yt37mav119vhbd3zc2dt8jc1t','dhuw68gzv82gayak9v8t6c42qz89nb35','dhuw691mdnou2vkn9x4swx8b491m9kor','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 15:22:08','2023-02-22 17:01:55',NULL,NULL,'Canceled',1677050528454635000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuw6964t25cg6fc9yvbgoskrrj3pqu7','dhuw68gzv82gayak9v8t6c42qz89nb35',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 15:22:08',NULL,'2023-02-22 15:22:08','','Executed Auto',1677050528453654000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuw696bm25u6nrg9uxcmcanqj7z8ssp','dhuw68gzv82gayak9v8t6c42qz89nb35','dhuw691mdnou2vkn9x4swx8b491m9kor','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-22 15:22:08','2023-02-22 17:01:55',NULL,NULL,'Canceled',1677050528455972000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuw698ggn7b8vxc9g4v4zg7y6a1c3hn','dhuw68gzv82gayak9v8t6c42qz89nb35','dhuw691mdnou2vkn9x4swx8b491m9kor','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 15:22:08',NULL,'2023-02-22 17:01:55','345325','Executed Accepted',1677050528455331000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuwfp2yao81n86q9gacntq1waws5xn6','dhuwfk31b6fkj1yoaruunhwokxs7yvys',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 15:40:04',NULL,'2023-02-22 15:40:04','ccccc','Executed Auto',1677051604469711000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuwfp3pj6cxx3jsa5bs3ng2a9bfurpn','dhuwfk31b6fkj1yoaruunhwokxs7yvys','dhuwfp7mut1r86vv91cavfn1xj3rtctj','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-22 15:40:04','2023-02-22 15:40:39',NULL,'流程发起者已撤回','Canceled',1677051604473196000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuwfp4ug7hmz4xtaa2sf4pnjr7g7oxb','dhuwfk31b6fkj1yoaruunhwokxs7yvys','dhuwfp7mut1r86vv91cavfn1xj3rtctj','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 15:40:04','2023-02-22 15:40:39',NULL,'流程发起者已撤回','Canceled',1677051604472189000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuwfp76nch8npap95kufp5msgkz29sw','dhuwfk31b6fkj1yoaruunhwokxs7yvys','dhuwfp7mut1r86vv91cavfn1xj3rtctj','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 15:40:04','2023-02-22 15:40:39',NULL,'流程发起者已撤回','Canceled',1677051604474651000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuwg152uuhug2xnacdsm2jnc5h7utnb','dhuwfk31b6fkj1yoaruunhwokxs7yvys','dhuwg18vyha84v9yaf1b44qap3jaq12r','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-22 15:40:48','2023-02-22 15:48:18',NULL,NULL,'Canceled',1677051648690722000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuwg1837yjyr3s4a55b8xko3f4nz3ns','dhuwfk31b6fkj1yoaruunhwokxs7yvys','dhuwg18vyha84v9yaf1b44qap3jaq12r','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 15:40:48',NULL,'2023-02-22 15:48:18','5555555','Executed Accepted',1677051648691901000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuwg18fv7gr5rxpawubbk5b6cu6avot','dhuwfk31b6fkj1yoaruunhwokxs7yvys',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 15:40:48',NULL,'2023-02-22 15:40:48','','Executed Auto',1677051648689532000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuwg18obxp8oogw923cy56wdgy8hvyv','dhuwfk31b6fkj1yoaruunhwokxs7yvys','dhuwg18vyha84v9yaf1b44qap3jaq12r','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 15:40:48','2023-02-22 15:48:18',NULL,NULL,'Canceled',1677051648691307000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuwkjjgbgraprypak69q2tn816n12fw','dhuwfk31b6fkj1yoaruunhwokxs7yvys','dhuwkjpmjq81931n94z98rt7q4sbnawf','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-4,'Execute','','总经理','dhpwodsg5hf9gha39qtsjsryxcf94phs','孙红雷','2023-02-22 15:48:18',NULL,'2023-02-22 21:33:30','66666','Executed Accepted',1677052098973096000,'D000','00','公司办公室','dhpwodsg5hf9gha39qtsjsryxcf94phs','U000000','孙红雷'),('dhuwkjrw3mnd1f4o9ds9jn6nttg293fw','dhuwfk31b6fkj1yoaruunhwokxs7yvys',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-3,'Branch','','超过3天','U002','系统管理员','2023-02-22 15:48:18',NULL,'2023-02-22 15:48:18',NULL,'Executed Auto',1677052098972103000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuwpd2zs7gwsv2hatt9p2tqv3o4bfot','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhuwpd6667dmw7mh9h2sf4swnhqqo14w','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-4,'Execute','','总经理','dhpwodsg5hf9gha39qtsjsryxcf94phs','孙红雷','2023-02-22 15:56:32','2023-02-22 16:38:23',NULL,NULL,'Canceled',1677052592994841000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuwpd38o63r9k779j9ajxdmr5ux8zjw','dhts9qgz5wnxg9z394cv77pgvymd6xdo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-3,'Branch','','超过3天','U002','系统管理员','2023-02-22 15:56:32',NULL,'2023-02-22 15:56:32',NULL,'Executed Auto',1677052592993782000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuwpd7gjyhzk5r3ajbb89cdjtbj5wop','dhts9qgz5wnxg9z394cv77pgvymd6xdo','dhuwpd6667dmw7mh9h2sf4swnhqqo14w','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-4,'Execute','','总经理','U002','系统管理员','2023-02-22 15:56:32',NULL,'2023-02-22 16:38:23','xxxxxxx','Executed Accepted',1677052592995227000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuwvy9com1cp35ca3fsdf5xnwndsk78','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhuwvyfk1jtsskf8azjcsjgh25jqhrr9','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-4,'Execute','','总经理','dhpwodsg5hf9gha39qtsjsryxcf94phs','孙红雷','2023-02-22 16:10:33','2023-02-22 16:52:10',NULL,NULL,'Canceled',1677053433315558000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuwvybwanf8n6ms9s2cryhq946udmgb','dhtsbzdyhpyzt31v9hjsk415238xvtqh',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-3,'Branch','','超过3天','U002','系统管理员','2023-02-22 16:10:33',NULL,'2023-02-22 16:10:33',NULL,'Executed Auto',1677053433313056000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuwvydfuuq7afr795gbpzvun9j3hg8p','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhuwvyfk1jtsskf8azjcsjgh25jqhrr9','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-4,'Execute','','总经理','U002','系统管理员','2023-02-22 16:10:33',NULL,'2023-02-22 16:52:10','55566666666666666','Executed Rejected',1677053433314822000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhux9zwpxa3arhcf97w9c4vx4wcs1c1v','dhts9qgz5wnxg9z394cv77pgvymd6xdo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-5,'End','End','结束','U002','系统管理员','2023-02-22 16:38:23',NULL,'2023-02-22 16:38:23',NULL,'Executed Auto',1677055103474029000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxhrtu7drwczod9f4vjj7p1jm9fpg6','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhuxhrub91nu8sjya1jadpykt9s7cq69','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-22 16:52:47','2023-02-22 16:53:20',NULL,'流程发起者已撤回','Canceled',1677055967180462000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuxhrutcj5zgucc9k4vjndsbxot9ayu','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhuxhrub91nu8sjya1jadpykt9s7cq69','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 16:52:47','2023-02-22 16:53:20',NULL,'流程发起者已撤回','Canceled',1677055967180979000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuxhry7y94qk6sra4w9gf768f4xqymm','dhtsbzdyhpyzt31v9hjsk415238xvtqh',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 16:52:47',NULL,'2023-02-22 16:52:47','','Executed Auto',1677055967179305000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxj5t29f9m2hmy9b99o8yfsga7q4on','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhuxj5tx5gza2u7q9ouaxw4btzb8dnf9','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-22 16:53:39','2023-02-22 16:53:46',NULL,'流程发起者已撤回','Canceled',1677056019579160000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuxj5vdc2dq7nqx9vqu9adhm27xb9qg','dhtsbzdyhpyzt31v9hjsk415238xvtqh',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 16:53:39',NULL,'2023-02-22 16:53:39','','Executed Auto',1677056019577747000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxj5ypgfwfrdcway69h2taq4gv9n6q','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhuxj5tx5gza2u7q9ouaxw4btzb8dnf9','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 16:53:39','2023-02-22 16:53:46',NULL,'流程发起者已撤回','Canceled',1677056019580799000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuxj5zdgcp7r7gf96otug9oqb6c5vxm','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhuxj5tx5gza2u7q9ouaxw4btzb8dnf9','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 16:53:39','2023-02-22 16:53:46',NULL,'流程发起者已撤回','Canceled',1677056019580351000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuxk71uacruc19z9tbut72ztbbattpo','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhuxk72puycxgkwn9rhv3h3vornc7vby','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-22 16:55:52','2023-02-22 16:56:10',NULL,NULL,'Canceled',1677056152378911000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuxk72dsrpf7k8u9kdvrs31ormnazj5','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhuxk72puycxgkwn9rhv3h3vornc7vby','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 16:55:52',NULL,'2023-02-22 16:56:10','555555','Executed Accepted',1677056152379562000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxk76qaqkayj2hau8txr2a74xk91nn','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhuxk72puycxgkwn9rhv3h3vornc7vby','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 16:55:52','2023-02-22 16:56:10',NULL,NULL,'Canceled',1677056152379258000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuxk78u4ju7w4rva59s39o4dkozkfpf','dhtsbzdyhpyzt31v9hjsk415238xvtqh',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 16:55:52',NULL,'2023-02-22 16:55:52','','Executed Auto',1677056152377957000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxkbmaohs81rdfah39xm9pu9f49jtp','dhtsbzdyhpyzt31v9hjsk415238xvtqh',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-3,'Branch','','超过3天','U002','系统管理员','2023-02-22 16:56:10',NULL,'2023-02-22 16:56:10',NULL,'Executed Auto',1677056170429974000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxkbnccytfvumz95hsummjxdtqqvyp','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhuxkbo73oakfstf92fbwmjroktn38vy','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-4,'Execute','','总经理','dhpwodsg5hf9gha39qtsjsryxcf94phs','孙红雷','2023-02-22 16:56:10','2023-02-22 17:01:16',NULL,NULL,'Canceled',1677056170432611000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuxkbrf2ahctb149op9ftbmpfxukr93','dhtsbzdyhpyzt31v9hjsk415238xvtqh','dhuxkbo73oakfstf92fbwmjroktn38vy','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-4,'Execute','','总经理','U002','系统管理员','2023-02-22 16:56:10',NULL,'2023-02-22 17:01:16','ewrtrewtre','Executed Accepted',1677056170434223000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxn6kj48z97a9a9sb9ov687da7ro1t','dhts924oa63cm11j9gnua1fy7h69skwo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 17:00:06',NULL,'2023-02-22 17:00:06','55555','Executed Auto',1677056406894554000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxn6kym12bbzvq9omvpxs6ks27rhg4','dhts924oa63cm11j9gnua1fy7h69skwo','dhuxn6kka5fna7gh95bvgbb52cc57xkc','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 17:00:06',NULL,'2023-02-22 17:00:16','ertreterwt','Executed Rejected',1677056406897025000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxn6p76v2p96pya21syu7rnjp3tof2','dhts924oa63cm11j9gnua1fy7h69skwo','dhuxn6kka5fna7gh95bvgbb52cc57xkc','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 17:00:06','2023-02-22 17:00:16',NULL,NULL,'Canceled',1677056406898918000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuxn6rmaxgv9wxz9b39z9c6rfb9chph','dhts924oa63cm11j9gnua1fy7h69skwo','dhuxn6kka5fna7gh95bvgbb52cc57xkc','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-22 17:00:06','2023-02-22 17:00:16',NULL,NULL,'Canceled',1677056406897998000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuxncam8k6r8nzzawovoapank2zppvs','dhts924oa63cm11j9gnua1fy7h69skwo','dhuxncbqfbq13jbf94har6rgukyhf6r1','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U002','系统管理员','2023-02-22 17:00:29',NULL,'2023-02-22 17:00:54','ertertewt','Executed Accepted',1677056429847291000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxncb5dj6ft72p9rss4yobhac1w1nh','dhts924oa63cm11j9gnua1fy7h69skwo','dhuxncbqfbq13jbf94har6rgukyhf6r1','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U018','陈世','2023-02-22 17:00:29','2023-02-22 17:00:54',NULL,NULL,'Canceled',1677056429846042000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuxncg32ko52dgu9w3tcbz5fp2c8nty','dhts924oa63cm11j9gnua1fy7h69skwo','dhuxncbqfbq13jbf94har6rgukyhf6r1','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U051','吴文丽','2023-02-22 17:00:29','2023-02-22 17:00:54',NULL,NULL,'Canceled',1677056429846693000,NULL,NULL,NULL,NULL,NULL,NULL),('dhuxncg93d5n4d7n94mvpzmyr646y6uc','dhts924oa63cm11j9gnua1fy7h69skwo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U002','系统管理员','2023-02-22 17:00:29',NULL,'2023-02-22 17:00:29','ertertre','Executed Auto',1677056429844927000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxnkk32xoju15say89a43ubrnnpu7a','dhts924oa63cm11j9gnua1fy7h69skwo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-3,'Branch','','超过3天','U002','系统管理员','2023-02-22 17:00:54',NULL,'2023-02-22 17:00:54',NULL,'Executed Auto',1677056454565641000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxnkkc7bsyomq99uvv7pr1drcpgsvu','dhts924oa63cm11j9gnua1fy7h69skwo','dhuxnkp99avayuj3aozcn2ppowxfomvu','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-4,'Execute','','总经理','U002','系统管理员','2023-02-22 17:00:54',NULL,'2023-02-22 17:01:11','ertwrew','Executed Accepted',1677056454566413000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxnpwcr324sro19qncj48dkx2nf1hp','dhts924oa63cm11j9gnua1fy7h69skwo',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-5,'End','End','结束','U002','系统管理员','2023-02-22 17:01:11',NULL,'2023-02-22 17:01:11',NULL,'Executed Auto',1677056471267091000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxnr3s13f3mjsza2pu2na6brj51kcf','dhtsbzdyhpyzt31v9hjsk415238xvtqh',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-5,'End','End','结束','U002','系统管理员','2023-02-22 17:01:16',NULL,'2023-02-22 17:01:16',NULL,'Executed Auto',1677056476566920000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxo1uf3qrdt5ntad2uoov63hqghk5g','dhuw68gzv82gayak9v8t6c42qz89nb35',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-3,'Branch','','超过3天','U002','系统管理员','2023-02-22 17:01:55',NULL,'2023-02-22 17:01:55',NULL,'Executed Auto',1677056515731944000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxo1vazy8jsjuya4dvu8osgoc6o5yr','dhuw68gzv82gayak9v8t6c42qz89nb35',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-5,'End','End','结束','U002','系统管理员','2023-02-22 17:01:55',NULL,'2023-02-22 17:01:55',NULL,'Executed Auto',1677056515733047000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxo3kagdk2tons9s5vjt9zrztgwg6u','dhtsv5njawd1cuz6aq9su7soqcuzggog',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-5,'End','End','结束','U002','系统管理员','2023-02-22 17:02:02',NULL,'2023-02-22 17:02:02',NULL,'Executed Auto',1677056522038073000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxo3p1nvy9adgx9379vp9fg5n76umq','dhtsv5njawd1cuz6aq9su7soqcuzggog',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-3,'Branch','','超过3天','U002','系统管理员','2023-02-22 17:02:02',NULL,'2023-02-22 17:02:02',NULL,'Executed Auto',1677056522037247000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxo56x72c8a48n9ubcu445qju5bqv5','dhukvr5oqnbbrp6n92o9wbath7y9y5an',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-3,'Branch','','超过3天','U002','系统管理员','2023-02-22 17:02:08',NULL,'2023-02-22 17:02:08',NULL,'Executed Auto',1677056528998523000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhuxo5g4kxgcyfbh9hp9c11fanoo1v44','dhukvr5oqnbbrp6n92o9wbath7y9y5an',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-5,'End','End','结束','U002','系统管理员','2023-02-22 17:02:09',NULL,'2023-02-22 17:02:09',NULL,'Executed Auto',1677056529002691000,'D022','0302','仓库组','U002','admin','系统管理员'),('dhv2nbpsspz365mjadutynssa6woj5sn','dhuwfk31b6fkj1yoaruunhwokxs7yvys',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-5,'End','End','结束','dhpwodsg5hf9gha39qtsjsryxcf94phs','孙红雷','2023-02-22 21:33:30',NULL,'2023-02-22 21:33:30',NULL,'Executed Auto',1677072810915774000,'D000','00','公司办公室','dhpwodsg5hf9gha39qtsjsryxcf94phs','U000000','孙红雷'),('dhv3qs3q1hhns1zna6osp351mvos9ff8','dhv3qqsjo3gjdqas9xgbnnqj9wostqd8',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U001','系统管理员3','2023-02-22 22:49:04',NULL,'2023-02-22 22:49:04','呜呜呜呜','Executed Auto',1677077344691011000,'D005','04','后勤保障部','U001','admin','系统管理员3'),('dhv3qs89jcfhhcj99j3sokzghszycdsg','dhv3qqsjo3gjdqas9xgbnnqj9wostqd8','dhv3qs5zcuawhb1y9dhbj1a31xnzss72','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U001','系统管理员3','2023-02-22 22:49:04',NULL,NULL,NULL,'Executing',1677077344691926000,NULL,NULL,NULL,NULL,NULL,NULL),('dhv3r11ma2j5o6y79sy9ysyzu93p4c3h','dhv3qz9b4dp24y5taurujpv6xrsympzx','dhv3r18s25g5jm4jao5bups38fdbsxta','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导','U001','系统管理员3','2023-02-22 22:49:36',NULL,NULL,NULL,'Executing',1677077376460235000,NULL,NULL,NULL,NULL,NULL,NULL),('dhv3r12q5ubk27m898va1zws7q2shum9','dhv3qz9b4dp24y5taurujpv6xrsympzx',NULL,'dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始','U001','系统管理员3','2023-02-22 22:49:36',NULL,'2023-02-22 22:49:36','其望闻问切','Executed Auto',1677077376459085000,'D005','04','后勤保障部','U001','admin','系统管理员3');
/*!40000 ALTER TABLE `wf_flow_task` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wf_options_diagram`
--

DROP TABLE IF EXISTS `wf_options_diagram`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wf_options_diagram` (
                                      `id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
                                      `diagram_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                      `diagram_code_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                      `diagram_name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                      `model_` text COLLATE utf8mb4_general_ci,
                                      `keyword_` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                      `icon_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                      `description_` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                      `exceed_days_` int DEFAULT NULL,
                                      `start_key_` int DEFAULT NULL,
                                      `order_` bigint DEFAULT NULL,
                                      PRIMARY KEY (`id`),
                                      UNIQUE KEY `wf_options_diagram_UniqueIndex_diagram_id` (`diagram_id_`),
                                      CONSTRAINT `wf_options_diagram_ForeignKey_diagram_id` FOREIGN KEY (`diagram_id_`) REFERENCES `wf_diagram` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wf_options_diagram`
--

LOCK TABLES `wf_options_diagram` WRITE;
/*!40000 ALTER TABLE `wf_options_diagram` DISABLE KEYS */;
INSERT INTO `wf_options_diagram` VALUES ('dhttjgtbfm2t3uzf9vpc2hsoqfhs6r8h','dhpusz1tvjjj9ovp9129qwgv4t5h5pou','A01','员工请假单','{ \"class\": \"GraphLinksModel\",\n  \"linkFromPortIdProperty\": \"fromPort\",\n  \"linkToPortIdProperty\": \"toPort\",\n  \"nodeDataArray\": [\n{\"category\":\"Start\",\"text\":\"\\u5f00\\u59cb\",\"key\":-1,\"loc\":\"230 120\"},\n{\"category\":\"Execute\",\"text\":\"\\u90e8\\u95e8\\u9886\\u5bfc\",\"key\":-2,\"loc\":\"230 220\"},\n{\"category\":\"Branch\",\"text\":\"\\u8d85\\u8fc73\\u5929\",\"key\":-3,\"loc\":\"230 330\"},\n{\"category\":\"Execute\",\"text\":\"\\u603b\\u7ecf\\u7406\",\"key\":-4,\"loc\":\"440 330\"},\n{\"category\":\"End\",\"text\":\"\\u7ed3\\u675f\",\"key\":-5,\"loc\":\"230 460\"}\n],\n  \"linkDataArray\": [\n{\"from\":-1,\"to\":-2,\"fromPort\":\"B\",\"toPort\":\"T\",\"points\":[230,153,230,163,230,174,230,174,230,185,230,195],\"category\":\"Link\"},\n{\"from\":-2,\"to\":-3,\"fromPort\":\"B\",\"toPort\":\"T\",\"points\":[230,245,230,255,230,269,230,269,230,283,230,293],\"category\":\"Link\"},\n{\"from\":-3,\"to\":-4,\"fromPort\":\"R\",\"toPort\":\"L\",\"visible\":true,\"points\":[291,330,301,330,335,330,335,330,369,330,379,330],\"category\":\"Link\",\"text\":\"\\u662f\"},\n{\"from\":-3,\"to\":-5,\"fromPort\":\"B\",\"toPort\":\"T\",\"visible\":true,\"points\":[230,367,230,377,230,397,230,397,230,417,230,427],\"category\":\"Link\",\"text\":\"\\u5426\"},\n{\"from\":-4,\"to\":-5,\"fromPort\":\"B\",\"toPort\":\"R\",\"points\":[440,355,440,365,440,460,356.5,460,273,460,263,460],\"category\":\"Link\"}\n]}','#type_#: 请假#days_#天','mdi mdi-alarm-snooze','填写请假单，经部门领导和总经理审批后生效',365,-1,1676393796157209000);
/*!40000 ALTER TABLE `wf_options_diagram` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wf_options_link`
--

DROP TABLE IF EXISTS `wf_options_link`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wf_options_link` (
                                   `id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
                                   `diagram_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `from_key_` int DEFAULT NULL,
                                   `to_key_` int DEFAULT NULL,
                                   `on_script_` text COLLATE utf8mb4_general_ci,
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `wf_options_link_UniqueIndex_diagram_id_from_key_to_key` (`diagram_id_`,`from_key_`,`to_key_`),
                                   CONSTRAINT `wf_options_link_ForeignKey_diagram_id` FOREIGN KEY (`diagram_id_`) REFERENCES `wf_diagram` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wf_options_link`
--

LOCK TABLES `wf_options_link` WRITE;
/*!40000 ALTER TABLE `wf_options_link` DISABLE KEYS */;
INSERT INTO `wf_options_link` VALUES ('dhv3qmbhxroj1ngh9z2cdpfqnn5swpwp','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,-3,'true'),('dhv3qmf58xrdz7jm9jccaqrwcgr17trh','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,-2,'true'),('dhv3qmfuncdvgzb39z7urpgmp2uvkcw6','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-4,-5,'true'),('dhv3qmgjrpo76s7dapav15zoxowur21f','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-3,-5,'$values[\"days_\"] <=3;'),('dhv3qmhbdo4325s1abovkx6f4xq14w5d','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-3,-4,'$values[\"days_\"] > 3;');
/*!40000 ALTER TABLE `wf_options_link` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wf_options_node`
--

DROP TABLE IF EXISTS `wf_options_node`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wf_options_node` (
                                   `id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
                                   `diagram_id_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `key_` int DEFAULT NULL,
                                   `category_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `code_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `name_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `rejectable_` tinyint DEFAULT NULL,
                                   `require_reject_comment_` tinyint DEFAULT NULL,
                                   `on_reject_script_` text COLLATE utf8mb4_general_ci,
                                   `revocable_` tinyint DEFAULT NULL,
                                   `on_revoke_script_` text COLLATE utf8mb4_general_ci,
                                   `on_before_script_` text COLLATE utf8mb4_general_ci,
                                   `on_after_script_` text COLLATE utf8mb4_general_ci,
                                   `executor_custom_num_` int DEFAULT NULL,
                                   `executor_selectable_num_` int DEFAULT NULL,
                                   `executor_savable_` tinyint DEFAULT NULL,
                                   `executor_users_` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `executor_name_users_` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `executor_departs_` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `executor_name_departs_` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `executor_roles_` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `executor_name_roles_` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `executor_policy_` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `executor_script_` text COLLATE utf8mb4_general_ci,
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `wf_options_node_UniqueIndex_diagram_id_key` (`diagram_id_`,`key_`),
                                   CONSTRAINT `wf_options_node_ForeignKey_diagram_id` FOREIGN KEY (`diagram_id_`) REFERENCES `wf_diagram` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wf_options_node`
--

LOCK TABLES `wf_options_node` WRITE;
/*!40000 ALTER TABLE `wf_options_node` DISABLE KEYS */;
INSERT INTO `wf_options_node` VALUES ('dhv3qm9vtzo1kxmma8uavfvmycwmk52c','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-1,'Start','Start','开始',0,0,'',1,'','','',0,0,0,'','','','','','','',''),('dhv3qmaq3sor2ru69rza7po31ot9prkk','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-4,'Execute','','总经理',1,1,'',0,'','','',0,0,1,'U003,U002,U001','系统管理员1,系统管理员2,系统管理员3','D000','公司办公室','','','None','[]'),('dhv3qmbnmg6896n1ayas14fxqymx6pn3','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-5,'End','End','结束',0,0,'',0,'','','',0,0,0,'','','','','','','',''),('dhv3qmdmsosg1jz8azg9mmcf51abtwxv','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-2,'Execute','','部门领导',1,1,'',0,'','','',3,3,1,'','','','','','','StartDepartLeader','[]'),('dhv3qmg5rdyp1qyka8d957b6jdgtzzmu','dhpusz1tvjjj9ovp9129qwgv4t5h5pou',-3,'Branch','','超过3天',0,0,'',0,'','','',0,0,0,'','','','','','','','');
/*!40000 ALTER TABLE `wf_options_node` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-02-22 23:02:20
