select title, g.name from series s inner join genres g on s.genre_id = g.id;
select title, first_name, last_name from episodes e inner join actor_episode ae on e.id = ae.episode_id inner join actors a on ae.actor_id = a.id;
select s.title, count(x.id) from series s inner join seasons x on x.serie_id = s.id group by s.id;
select name, count(m.id) as cuenta from genres g inner join movies m on g.id = m.genre_id group by g.id having cuenta >= 3;
select first_name, last_name from actors a where a.id in (select actor_id from actor_movie am inner join movies m on am.movie_id = m.id where title like "%galaxias%") group by a.id;
select first_name, last_name from actors a inner join actor_movie am on a.id = am.actor_id inner join movies m on am.movie_id = m.id where title like "%galaxias%" group by a.id;