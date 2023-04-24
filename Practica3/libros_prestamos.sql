
drop database if exists libros_prestamos;
create database libros_prestamos;
use libros_prestamos; 

create table autor(idAutor int primary key auto_increment, nombre varchar(100), nacionalidad varchar(50));
create table estudiante(idLector int primary key auto_increment, nombre varchar(50), apellido varchar(50), direccion varchar(100), carrera varchar(100), edad int);
create table libro(idLibro int primary key auto_increment, titulo varchar(100), editorial varchar(100), area varchar(50));
create table libro_autor(idAutor int, idLibro int, primary key(idAutor, idLibro), foreign key (idAutor) references autor(idAutor), foreign key (idLibro) references libro(idLibro));
create table prestamo(idLector int, idLibro int, FechaPrestamo date default (current_date), FechaDevolucion date, devuelto boolean default false, primary key(idLector,idLibro), foreign key (idLector) references estudiante (idLector), foreign key (idLibro) references  libro (idLibro));

insert into libro values (null, "Harry Potter 1", "Salamandra", "fantasia");
insert into libro values (null, "Harry Potter 2", "Salamandra", "fantasia");
insert into libro values (null, "El Universo: Guía de Viaje", "Planeta", "astronomía");
insert into libro values (null, "La Historia de la Web", "Planeta", "internet");
insert into libro values (null, "Geometría Analítica", "Editorial ABC", "matemáticas");

insert into autor values (null,"J.K. Rowling","inglesa");
insert into autor values (null,"Pierre Dupont","francesa");
insert into autor values (null,"Marie Roux","francesa");
insert into autor values (null,"Giuseppe Reggiani","italiana");
insert into autor values (null,"Ronald McDonald","estadounidense");

insert into estudiante values (null, "Roberto", "González", "Calle Falsa 123", "informática",20);
insert into estudiante values (null, "Pedro", "Fuentes", "Calle Inventada 321", "informática",17);
insert into estudiante values (null, "Isabel", "Pérez", "Calle Dirección 648", "psicología",21);
insert into estudiante values (null, "Federico", "Sánchez", "Calle Fake 44", "química",24);
insert into estudiante values (null, "Filippo", "Galli", "Calle Mock 85", "literatura",20);

insert into libro_autor values(1,1),(2,1),(3,2),(4,3),(4,4),(5,3);

insert into prestamo values(1,1,null,"2024-04-15",false);
insert into prestamo values(2,1,"2023-04-01","2024-05-18",false);
insert into prestamo values(1,4,"2019-12-06","2020-03-05",false);
insert into prestamo values(3,5,null,"2023-09-11",false);
insert into prestamo values(4,2,"2020-11-07","2021-07-16",false);


select * from autor;
select nombre, edad from estudiante;
select nombre, apellido from estudiante where carrera = "informática";
select nombre from autor where nacionalidad in ("francesa","italiana");
select titulo from libro where area != "internet";
select titulo from libro where editorial = "Salamandra";
select * from estudiante where edad > (select avg(edad) from estudiante);
select nombre, apellido from estudiante where apellido like "G%";

/*Con doble inner join*/
select nombre from autor a inner join libro_autor la on a.idAutor = la.idAutor inner join libro l on l.idLibro = la.idLibro where titulo = "El Universo: Guía de Viaje";
/*Con inner join y subconsulta*/
select nombre from autor a inner join libro_autor la on a.idAutor = la.idAutor where idLibro in (select idLibro from libro where titulo = "El Universo: Guía de Viaje");

select titulo from libro l inner join prestamo p on l.idLibro = p.idLibro inner join estudiante e on e.idLector = p.idLector where nombre = "Filippo Galli";
select nombre,apellido from estudiante where edad = (select min(edad) from estudiante);
select nombre, apellido from estudiante e inner join prestamo p on e.idLector = p.idLector;
select titulo from autor a inner join libro_autor la on a.idAutor = la.idAutor inner join libro l on l.idLibro = la.idLibro where nombre = "J.K. Rowling";
select titulo from libro l inner join prestamo p on l.idLibro = p.idLibro where FechaDevolucion = "2021-07-16";

