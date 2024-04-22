DROP DATABASE IF EXISTS hw_gallery;

CREATE DATABASE hw_gallery CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE hw_gallery;

CREATE TABLE images (
  name        VARCHAR (50)  NOT NULL,
  created_at  DATETIME      NOT NULL,

  PRIMARY KEY (name)
);