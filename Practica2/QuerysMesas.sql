use empresa_internet;

select idclientes, nombre, dni
from clientes 
where idclientes >= 10;

select precio,velocidad_ofrecida 
from planes_internet 
where velocidad_ofrecida 
between 100 and 1000;

SELECT AVG(precio) as 'Precio promedio' 
FROM planes_internet;

select min(velocidad_ofrecida) 
from planes_internet;

select * 
from clientes
where provincia like "bu%" 
Order By dni desc;

select year(fecha_nacimiento), nombre, apellido 
from clientes;

SELECT COUNT(idclientes) as 'Total clientes' 
FROM clientes;

UPDATE planes_internet SET precio = 1500 
WHERE velocidad_ofrecida = 500;

DELETE FROM planes_internet 
WHERE idclientes = (select idclientes from clientes where nombre = 'Maury');

Delete From clientes
where nombre = 'Maury';

select * from planes_internet;

select * from clientes;