SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

DROP SCHEMA IF EXISTS `blog` ;
CREATE SCHEMA IF NOT EXISTS `blog` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ;
USE `blog` ;

-- -----------------------------------------------------
-- Table `blog`.`section`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `blog`.`section` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `number` INT NULL,
  `name` VARCHAR(45) NULL,
  `text` TEXT NULL,
  `image` VARCHAR(45) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `blog`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `blog`.`user` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(45) NULL,
  `pseudonym` VARCHAR(45) NULL,
  `password` VARCHAR(45) NULL,
  `it_is_administrator` TINYINT(1) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `blog`.`article`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `blog`.`article` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(45) NULL,
  `text` TEXT NULL,
  `image` VARCHAR(45) NULL,
  `date` DATETIME NULL,
  `section_id` INT UNSIGNED NULL,
  `user_id` INT UNSIGNED NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_article_section_idx` (`section_id` ASC),
  INDEX `fk_article_user1_idx` (`user_id` ASC),
  CONSTRAINT `fk_article_section`
    FOREIGN KEY (`section_id`)
    REFERENCES `blog`.`section` (`id`)
    ON DELETE SET NULL
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_article_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `blog`.`user` (`id`)
    ON DELETE SET NULL
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `blog`.`comment`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `blog`.`comment` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `text` TEXT NULL,
  `date` DATETIME NULL,
  `article_id` INT UNSIGNED NULL,
  `user_id` INT UNSIGNED NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_comment_article1_idx` (`article_id` ASC),
  INDEX `fk_comment_user1_idx` (`user_id` ASC),
  CONSTRAINT `fk_comment_article1`
    FOREIGN KEY (`article_id`)
    REFERENCES `blog`.`article` (`id`)
    ON DELETE SET NULL
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_comment_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `blog`.`user` (`id`)
    ON DELETE SET NULL
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `blog`.`subscriber`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `blog`.`subscriber` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `email` VARCHAR(45) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
