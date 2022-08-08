BEGIN;
CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password varchar(255) not null
);





-- insert into category(genre)
-- values
-- ('Fantasy'),
-- ('Drama'),
-- ('Comedy'),
-- ('Triller');

-- insert into book (author,title)
-- values
-- ('Stephene Meyer','Twilight'),
-- ('Cristopher Nolan','Interstellar'),
-- ('Nurlan Saburov','StandUp'),
-- ('Timur Bekmambetov','Find');



-- insert into bookcategory(book_id,category_id)
-- values
-- (1,2),
-- (2,1),
-- (3,2),
-- (4,3),
-- (2,4),
-- (4,1);

-- SELECT * FROM book 
-- left  JOIN bookcategory  ON bookcategory.book_id = book.id
-- left join category on bookcategory.category_id =category.id;

COMMIT;