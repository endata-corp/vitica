# ************************************************************
# Sequel Pro SQL dump
# Version 4096
#
# http://www.sequelpro.com/
# http://code.google.com/p/sequel-pro/
#
# Host: 127.0.0.1 (MySQL 5.6.23)
# Database: vitica
# Generation Time: 2015-06-19 04:04:48 +0000
# ************************************************************

# Dump of table products
# ------------------------------------------------------------

INSERT INTO `products` (`id`, `code`, `slug`, `name`, `description`, `design_id`, `garment_id_m`, `garment_id_w`, `garment_color`, `garment_hex`, `garment_description`, `available_sizes_m`, `available_sizes_w`, `price_suggested`, `price_cost`, `price`, `is_new_release`, `is_on_sale`, `is_top_seller`, `is_designer_pick`, `image_url`, `image_url_m`, `image_url_w`, `image_side_url`, `image_back_url`, `image_design_url`, `is_active`, `is_approved`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'example1','my-example','My Example Product','The desriptions goes here',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,28,NULL,18,0,1,0,1,NULL,NULL,NULL,NULL,NULL,NULL,1,0,NULL,NULL,NULL),
	(2,'example2','my-example2','My Example2 Product','The desriptions goes here',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,26,1,0,0,1,NULL,NULL,NULL,NULL,NULL,NULL,1,0,NULL,NULL,NULL),
	(3,'example3','my-example3','Popular Product','The desriptions goes here',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,26,0,0,1,0,NULL,NULL,NULL,NULL,NULL,NULL,1,0,NULL,NULL,NULL),
	(4,'example4','my-example4','Product On Sale','The desriptions goes here',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,28,NULL,22,0,1,0,0,NULL,NULL,NULL,NULL,NULL,NULL,1,0,NULL,NULL,NULL);


# Dump of table tags
# ------------------------------------------------------------

INSERT INTO `tags` (`id`, `code`, `name`)
VALUES
	(1,'technology','Technology'),
	(2,'programming','Programming'),
	(3,'funny','Funny');


# Dump of table product_tags
# ------------------------------------------------------------

INSERT INTO `product_tags` (`id`, `product_id`, `tag_code`)
VALUES
	(1,1,'technology'),
	(2,2,'programming'),
	(3,1,'funny'),
	(5,3,'programming');


