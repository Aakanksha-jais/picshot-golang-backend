-- name: drop
DROP Database IF EXISTS picshot;
-- name: create
CREATE DATABASE picshot;
-- name: use
USE picshot;

-- name: create-table
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

-- password: hello123
-- name: insert-aakanksha
INSERT INTO `accounts` (`user_name`, `password`, `email` , `f_name`, `l_name`, `phone_no` , `status`)
VALUES ("aakanksha_jais", "$2a$10$.HUjOWXbMuVBXkpRLX9fuOg623yZP0/UTF4EAGHCJu1fXNWP4M7eS", "jaiswal14aakanksha@gmail.com" , "Aakanksha", "Jaiswal", "7807052049" , "ACTIVE");

-- password: demo123
-- name: insert-mainak
INSERT INTO `accounts` (`user_name`, `password`, `email` , `f_name`, `l_name`, `phone_no` , `status`)
VALUES ("mainak_pandit", "$2a$10$PFaxwsavIJuQEqga4tFMQ.oofBlx6qE/RebQVHdofxXKcvJbAc0xW", "mainakpandit@gmail.com" , "Mainak", "Pandit", "9149137433" , "ACTIVE");

-- password: demo123
-- name: insert-divij
INSERT INTO `accounts` (`user_name`, `password`, `email` , `f_name`, `l_name`, `phone_no` , `status`)
VALUES ("divij_gupta", "$2a$10$PFaxwsavIJuQEqga4tFMQ.oofBlx6qE/RebQVHdofxXKcvJbAc0xW", "divijgupta@gmail.com" , "Divij", "Gupta", "9682622125" , "ACTIVE");
