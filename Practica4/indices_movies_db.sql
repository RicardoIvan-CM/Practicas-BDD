use movies_db;
describe movies;

insert into movies values (null,null,null,"Super Mario Bros",7.0,2,"2023-03-21",90,null);
insert into genres values(null,null,null,"Videojuegos",13,1);
update movies set genre_id = 13 where title = "Super Mario Bros";
update actors set favorite_movie_id = 22 where first_name = "Adam" and last_name = "Sandler"; 

create temporary table my_movies as select * from movies;
delete from my_movies where awards < 5;
select name from genres g inner join movies m on m.genre_id = g.id group by g.id;
select a.*,awards from actors a inner join movies m on a.favorite_movie_id = m.id where awards > 3;
create index name_idx on my_movies(title);
show indexes from my_movies;
/*No habría una mejora significativa ya que la tabla no realiza busquedas por nombre de la pelicula con tanta frecuencia*/
/*Considero que no deberían crearse indices, ya que las tablas aparentan ser actualizadas con frecuencia*/