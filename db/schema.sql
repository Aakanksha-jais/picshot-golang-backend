DROP Database IF EXISTS picshot;
CREATE DATABASE picshot;
USE picshot;

CREATE TABLE `accounts` (
                                  `id` int(11) NOT NULL AUTO_INCREMENT,
                                  `user_name` VARCHAR(255) NOT NULL,
                                  `password` VARCHAR(255) NOT NULL,
                                  `email` VARCHAR(255),
                                  `f_name` VARCHAR(255) NOT NULL,
                                  `l_name` VARCHAR(255) NOT NULL,
                                  `phone_no` VARCHAR(255),
                                  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `pwd_update` timestamp DEFAULT NULL,
                                  `del_req` timestamp DEFAULT NULL,
                                  `status` enum('ACTIVE','INACTIVE') NOT NULL DEFAULT 'ACTIVE',
                                  PRIMARY KEY (`id`)
);

-- run the following commands from project root:
-- sudo mysql -u root < db/schema.sql
-- docker exec -it mysql-container mysql -u admin -padmin123