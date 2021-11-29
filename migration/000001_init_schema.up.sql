CREATE TABLE `publishers` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `slogan` TEXT NOT NULL,
  `founded` DATE NOT NULL,

  PRIMARY KEY (`id`)
);

CREATE TABLE `authors` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(100) NOT NULL, 
  `last_name` VARCHAR(100) NOT NULL,
  `biography` TEXT NOT NULL,
  `birthday` DATE NOT NULL,
  `death` DATE,

  PRIMARY KEY (`id`)
);

CREATE TABLE `books` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL,
  `original_release` DATE NOT NULL,
  `description` TEXT NOT NULL,
  `short_description` TEXT NOT NULL,

  `published` DATE NOT NULL,
  `publisher_id` INT UNSIGNED NOT NULL,
  `pages`  INT UNSIGNED NOT NULL,

  `seller_id` INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`),
  CONSTRAINT `books_constr_publishers`
    FOREIGN KEY (`publisher_id`) REFERENCES `publishers`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE `authorship` (
  `book_id` INT UNSIGNED NOT NULL,
  `author_id` INT UNSIGNED NOT NULL,

  PRIMARY KEY (`book_id`, `author_id`),

  CONSTRAINT `authorship_constr_book`
    FOREIGN KEY (`book_id`) REFERENCES `books`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,

  CONSTRAINT `authorship_constr_author`
    FOREIGN KEY (`author_id`) REFERENCES `authors`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE `published` (
  `author_id` INT UNSIGNED NOT NULL,
  `publisher_id` INT UNSIGNED NOT NULL, 

  PRIMARY KEY (`author_id`, `publisher_id`),

  CONSTRAINT `published_constr_publisher`
    FOREIGN KEY (`publisher_id`) REFERENCES `publishers`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,

  CONSTRAINT `published_constr_author`
    FOREIGN KEY (`author_id`) REFERENCES `authors`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
)