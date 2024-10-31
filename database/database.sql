create database e_learning;

-- create table Role (
--   id serial primary key not null,
--   name varchar(100)
-- );

create table Users (
  id SERIAL PRIMARY KEY NOT NULL,
  username varchar(100),
  password varchar(100),
  role varchar(30)
  );

  create table Admin (
    id serial primary key not null,
    name varchar(100),
    user_id int references Users(id)
  );

  create table Student(
    id serial primary key not null,
    name varchar(100),
    user_id int references Users(id)
  );


  create table Mentor(
    id serial primary key not null,
    name varchar(100),
    user_id int references Users(id)
  )