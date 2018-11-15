CREATE TABLE `membership` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) DEFAULT NULL,
  `pts_given` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

CREATE TABLE `membership_product` (
  `id` int(10) NOT NULL,
  `product_id` int(10) DEFAULT NULL,
  `membership_id` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `outlet` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `postal_code` int(6) NOT NULL,
  `contactNo` varchar(45) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

CREATE TABLE `payment` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `payment_ref_no` int(20) NOT NULL,
  `total_amount` decimal(18,2) DEFAULT NULL,
  `payment_datetime` datetime DEFAULT NULL,
  `outlet_id` int(10) NOT NULL,
  `discount` decimal(10,0) DEFAULT NULL,
  `payment_mode` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `payment_items` (
  `id` int(10) NOT NULL,
  `payment_id` int(10) DEFAULT NULL,
  `sub_total` decimal(18,2) DEFAULT NULL,
  `product_id` int(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `points` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `product_id` int(10) NOT NULL,
  `pts` int(11) NOT NULL,
  `price` decimal(18,2) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

CREATE TABLE `product` (
  `id` int(10) NOT NULL,
  `name` varchar(200) DEFAULT NULL,
  `price` decimal(18,2) DEFAULT NULL,
  `status` enum('Y','N') DEFAULT NULL,
  `remarks` varchar(200) DEFAULT NULL,
  `gender` enum('M','F') DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `firstname` varchar(100) NOT NULL,
  `lastname` varchar(100) NOT NULL,
  `password` varchar(200) NOT NULL,
  `email` varchar(250) NOT NULL,
  `isAdmin` tinyint(4) DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `pts_balance` bigint(20) DEFAULT NULL,
  `expiry_date` timestamp NULL DEFAULT NULL,
  `membership_id` int(10) DEFAULT NULL,
  `outlet_id` int(10) DEFAULT NULL,
  `cut_cnt` int(11) DEFAULT NULL,
  `treatment_cnt` int(11) DEFAULT NULL,
  `hairloss_treatment_cnt` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
