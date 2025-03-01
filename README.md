this would be my note of my learning
- **dont forget**, exported component need to be start with capital letter!!
- use `go get -u <module_name>` to get and update module in `go.mod`
- most code come from [mikelopster's tutorial](https://docs.mikelopster.dev/c/goapi-essential/chapter-3/crud). I'll try to extend it for my own variant(for learning)
- to connect dockerized postgres to desktop pgAdmin, use hostname as `127.0.0.1`

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
```sql
INSERT INTO public.people (id, firstname, lastname, role) VALUES
(1, 'Alice', 'Smith', 'Tenant'),
(2, 'Bob', 'Johnson', 'Tenant'),
(3, 'Charlie', 'Williams', 'Tenant'),
(4, 'David', 'Brown', 'Tenant'),
(5, 'Eve', 'Davis', 'Tenant'),
(6, 'Frank', 'Miller', 'Tenant'),
(7, 'Grace', 'Wilson', 'Tenant'),
(8, 'Henry', 'Moore', 'Tenant'),
(9, 'Ivy', 'Taylor', 'Tenant'),
(10, 'Jack', 'Anderson', 'Tenant'),
(11, 'Karen', 'Thomas', 'Manager'),
(12, 'Liam', 'Jackson', 'Manager'),
(13, 'Mia', 'White', 'Manager'),
(14, 'Noah', 'Harris', 'Manager'),
(15, 'Olivia', 'Martin', 'Manager');
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


