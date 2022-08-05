BEGIN;
CREATE table IF NOT EXISTS category (
    id serial primary key,
    genre text not NULL
);

CREATE TABLE IF NOT EXISTS book(
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    author TEXT NOT NULL

  
);

CREATE TABLE IF NOT EXISTS bookcategory(
    id serial PRIMARY KEY,
    book_id int NOT NULL REFERENCES book,
	category_id int NOT NULL REFERENCES category
  
);


insert into category(genre)
values
('Fantasy'),
('Drama'),
('Comedy'),
('Triller');

insert into book (author,title)
values
('Stephene Meyer','Twilight'),
('Cristopher Nolan','Interstellar'),
('Nurlan Saburov','StandUp'),
('Timur Bekmambetov','Find');



insert into bookcategory(book_id,category_id)
values
(1,2),
(2,1),
(3,2),
(4,3),
(2,4),
(4,1);

SELECT * FROM book 
left  JOIN bookcategory  ON bookcategory.book_id = book.id
left join category on bookcategory.category_id =category.id;

COMMIT;