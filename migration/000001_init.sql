CREATE TABLE IF NOT EXISTS `roles` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(30) NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `departments` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(45) NOT NULL,
    `address` VARCHAR(100) NOT NULL,
    `status` TINYINT NOT NULL COMMENT '0: inactive\n1: active',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `countries` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(45) NOT NULL,
    `status` TINYINT NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `users` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `role_id` INT NOT NULL,
    `department_id` INT DEFAULT NULL,
    `email` VARCHAR(45) NOT NULL,
    `password` TEXT NOT NULL,
    `name` VARCHAR(45) NOT NULL,
    `surname` VARCHAR(45) NOT NULL,
    `gender` VARCHAR(20) NOT NULL,
    `dob` DATE NOT NULL,
    `mobile` VARCHAR(15) NOT NULL,
    `country_id` INT NOT NULL,
    `resident_country_id` INT NOT NULL,
    `avatar` VARCHAR(100) DEFAULT NULL,
    `verification_status` TINYINT DEFAULT 0 COMMENT '0: unverified\n1: verified',
    `status` TINYINT NOT NULL COMMENT '0: inactive\n1: active',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY `fk_users_roles_idx` (`role_id`),
    KEY `fk_users_depts_idx` (`department_id`),
    KEY `fk_users_countries_idx` (`country_id`),
    KEY `fk_users_resident_countries_idx` (`resident_country_id`),
    CONSTRAINT `fk_users_roles` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
    CONSTRAINT `fk_users_depts` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`),
    CONSTRAINT `fk_users_countries` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`),
    CONSTRAINT `fk_users_resident_countries` FOREIGN KEY (`resident_country_id`) REFERENCES `countries` (`id`)
);

CREATE TABLE IF NOT EXISTS `volunteer_details` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `department_id` INT NOT NULL,
    `status` TINYINT NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY `fk_volunteer_details_depts_idx` (`department_id`),
    KEY `fk_volunteer_details_users_idx` (`user_id`),
    CONSTRAINT `fk_volunteer_details_depts` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`),
    CONSTRAINT `fk_volunteer_details_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE IF NOT EXISTS `requests` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `type` VARCHAR(45) NOT NULL,
    `status` TINYINT NOT NULL,
    `reject_notes` VARCHAR(255) DEFAULT NULL,
    `verifier_id` INT DEFAULT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY `fk_requests_users_idx` (`user_id`),
    KEY `fk_requests_verifiers_idx` (`verifier_id`),
    CONSTRAINT `fk_requests_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    CONSTRAINT `fk_requests_verifiers` FOREIGN KEY (`verifier_id`) REFERENCES `users` (`id`)
);

CREATE TABLE IF NOT EXISTS `user_identities` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `number` VARCHAR(30) NOT NULL,
    `type` VARCHAR(45) NOT NULL,
    `status` TINYINT NOT NULL,
    `expiry_date` DATE NOT NULL,
    `place_issued` VARCHAR(100) NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY `fk_user_identities_users_idx` (`user_id`),
    CONSTRAINT `fk_user_identities_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
