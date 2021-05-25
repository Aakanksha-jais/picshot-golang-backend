-- name: use
USE picshot;

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

-- run the following commands from project root:
-- sudo mysql -u root < db/schema.sql
-- sudo mysql -u root < db/test_data.sql