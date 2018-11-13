CREATE TABLE `membership` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) DEFAULT NULL,
  `user_id` int(10) NOT NULL,
  `expiry_date` date DEFAULT NULL,
  `earn_pts` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `outlet` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `postal_code` int(6) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `payment` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `payment_ref_no` int(20) NOT NULL,
  `total_amount` decimal(18,2) DEFAULT NULL,
  `payment_datetime` datetime DEFAULT NULL,
  `outlet_id` int(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `payment_items` (
  `id` int(10) NOT NULL,
  `payment_id` int(10) DEFAULT NULL,
  `sub_total` decimal(18,2) DEFAULT NULL,
  `discount` double DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `points` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `product_id` int(10) NOT NULL,
  `pts` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `product` (
  `id` int(10) NOT NULL,
  `name` varchar(200) DEFAULT NULL,
  `price` decimal(18,2) DEFAULT NULL,
  `status` enum('Y','N') DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `firstname` varchar(100) NOT NULL,
  `lastname` varchar(100) NOT NULL,
  `password` varchar(200) NOT NULL,
  `email` varchar(250) NOT NULL,
  `isAdmin` tinyint(4) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
