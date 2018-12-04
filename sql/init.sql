CREATE TABLE `tax_db`.`products` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `tax_code` int(1) unsigned NOT NULL,
  `price` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
);