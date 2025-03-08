this would be my note of my learning
- **dont forget**, exported component need to be start with capital letter!!
- use `go get -u <module_name>` to get and update module in `go.mod`
- most code come from [mikelopster's tutorial](https://docs.mikelopster.dev/c/goapi-essential/chapter-3/crud). I'll try to extend it for my own variant(for learning)
- to connect dockerized postgres to desktop pgAdmin, use hostname as `127.0.0.1`
- to connect it, use port `5430` since the port `5432` is taken by postgres installed by my `home-manager`(at this time)

# Script to create DB
## Buildings
```sql
INSERT INTO public.buildings(
	id, location, num_floor)
    values (1,'north',5),
    (2,'west',6),
    (3,'south',4);
```

## People
### Create Table
```sql
-- Table: public.people

-- DROP TABLE IF EXISTS public.people;

CREATE TABLE IF NOT EXISTS public.people
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    firstname text COLLATE pg_catalog."default",
    lastname text COLLATE pg_catalog."default",
    age integer,
    role text COLLATE pg_catalog."default",
    CONSTRAINT people_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.people
    OWNER to admin;
```
### Insert
```sql
INSERT INTO public.people (firstname, lastname, age, role) VALUES
('Alice', 'Smith', 30, 'Tenant'),
('Bob', 'Johnson', 25, 'Tenant'),
('Charlie', 'Williams', 35, 'Manager'),
('David', 'Brown', 28, 'Tenant'),
('Eve', 'Davis', 40, 'Manager'),
('Frank', 'Miller', 32, 'Tenant'),
('Grace', 'Wilson', 27, 'Tenant'),
('Henry', 'Moore', 38, 'Manager'),
('Ivy', 'Taylor', 29, 'Tenant'),
('Jack', 'Anderson', 33, 'Tenant'),
('Karen', 'Thomas', 42, 'Manager'),
('Liam', 'Jackson', 26, 'Tenant'),
('Mia', 'White', 31, 'Tenant'),
('Noah', 'Harris', 36, 'Manager'),
('Olivia', 'Martin', 29, 'Tenant');
```

## Rooms
```sql
INSERT INTO public.rooms (id, room_number, building_id, type, floor, tenant_id) VALUES
(1, 101, 1, 'single', 1, 1),  -- Alice Smith
(2, 102, 1, 'double', 1, 2),  -- Bob Johnson
(3, 201, 1, 'single', 2, 3),  -- Charlie Williams
(4, 202, 1, 'double', 2, 4),  -- David Brown
(5, 103, 2, 'single', 1, 5),  -- Eve Davis
(6, 203, 2, 'double', 2, 6),  -- Frank Miller
(7, 301, 2, 'single', 3, 7),  -- Grace Wilson
(8, 104, 3, 'double', 1, 8),  -- Henry Moore
(9, 204, 3, 'single', 2, 9),  -- Ivy Taylor
(10, 302, 3, 'double', 3, 10); -- Jack Anderson
```


# ORM Approach
using ORM is to represent relation(table) as a object in Go. Then the library will help convert it to be relation and deal with db for us.

now we don't need to create table by our own. We will be using **Migration** to create table from object
