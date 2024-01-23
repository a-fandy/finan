/*
SQLyog Community v13.2.1 (64 bit)
MySQL - 8.0.30 : Database - finan
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
/*Table structure for table `ref_format` */

DROP TABLE IF EXISTS `ref_format`;

CREATE TABLE `ref_format` (
  `format_id` varchar(255) NOT NULL COMMENT 'The id of current format',
  `format` varchar(255) NOT NULL COMMENT 'The format of current string',
  `counter` int NOT NULL COMMENT 'Counter',
  `counter_pad_length` int DEFAULT NULL COMMENT 'Counter''s padding length',
  `counter_pad_char` char(1) DEFAULT NULL COMMENT 'Counter''s padding char',
  `description` text COMMENT 'The description of current format',
  `extra_field` varchar(255) DEFAULT NULL COMMENT 'Extra field for common purposes',
  `created_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Datetime created',
  `updated_date` timestamp NULL DEFAULT NULL COMMENT 'Updated date',
  PRIMARY KEY (`format_id`,`format`,`counter`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `ref_format` */

insert  into `ref_format`(`format_id`,`format`,`counter`,`counter_pad_length`,`counter_pad_char`,`description`,`extra_field`,`created_date`,`updated_date`) values 
('referral_code','{prefix}{seq}',0,5,'0',NULL,'FIN','2024-01-23 02:21:38','2024-01-23 02:21:52');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
