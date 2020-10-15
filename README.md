# Distributed Dungeons

## About

The idea behind Distributed Dungeons was to make online pen & paper sessions easier to manage. Distributed Dungeons is open source web application. Players can connect to the session via a sessionkey and password they got from the DM or whoever started the session. Said session allows the DM to see the stats of their players in realtime and even their dicerolls (they have to roll the dice in the webapp though :D). 

## Database structure

The database consists of two seperate tables. First, there is the "session" table. This table contains two columns: sessionkey and password. 
The second table "characters" constits of 20 columns: Name, Player, Occupation, Age, Sex, Residence, Birthplace, STR, DEX, POW, CON, APP, EDU, SIZ, INT, Hitpoints, Sanity, Luck, Magicpoints, Diceroll. 
One of the "session" entries is in relation to many "character" entries (depends on how many characters are playing).

## Setting up the Database

Creating a persistent postgres database in a docker container: 

```docker run -p 5432:5432 -d -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=stripe-example -v pgdata:/var/lib/postgresql/data postgres```

After that, execute to get into the containers bash: 

```docker exec -it $CONTAINER_ID bash```

You can access your container ID with ```docker ps```

At last, execute: 

```psql -U postgres```

You can access your exisiting databases with ```\l``` 
 
At this point, the database and tables should be initialized. This can be done using basic SQL.
```
CREATE DATABASE distributed_dungeons;
CREATE TABLE sessions (
    id serial primary key,
    password text not null
);
CREATE TABLE characters (
    id serial not null,
    name text not null,
    player text not null,
    occupation text not null,
    age integer not null,
    sex text not null,
    residence text not null,
    birthplace text not null,
    strength integer not null,
    dexterity integer not null,
    power integer not null,
    constitution integer not null,
    appearance integer not null,
    education integer not null,
    size integer not null,
    intelligence integer not null,
    hitpoints integer not null,
    sanity integer not null,
    luck integer not null,
    magic_points integer not null,
    diceroll integer not null,
    session_id integer not null,
    foreign key (session_id) references sessions(id)
);
```

Note: The character table is currently designed for the "Call of Cthulu" game.

A value in a table can be updated like this: 

```UPDATE characters SET diceroll=1 WHERE name='Oliver Smith';```
  
To delete the persistent data you have to remove the docker volume: 

```docker volume prune```
  
 ## database relations
 
 In the distributed_dungeons database there is a one-to-many relation from a key of the "session" table to n rows of the "character" table.
 
 ## Sql File
 
 The sql file looks like this: 
 
 ```
 -- Drop old tables
drop table if exists characters;
drop table if exists sessions;
-- Create new tables
create table sessions (
    id serial primary key,
    password text not null
);
create table characters (
    id serial not null,
    name text not null,
    player text not null,
    occupation text not null,
    age integer not null,
    sex text not null,
    residence text not null,
    birthplace text not null,
    strength integer not null,
    dexterity integer not null,
    power integer not null,
    constitution integer not null,
    appearance integer not null,
    education integer not null,
    size integer not null,
    intelligence integer not null,
    hitpoints integer not null,
    sanity integer not null,
    luck integer not null,
    magic_points integer not null,
    diceroll integer not null,
    session_id integer not null,
    foreign key (session_id) references sessions(id)
);
-- Create test data
insert into sessions (password)
values ('password123');
insert into characters (
        name,
        player,
        occupation,
        age,
        sex,
        residence,
        birthplace,
        strength,
        dexterity,
        power,
        constitution,
        appearance,
        education,
        size,
        intelligence,
        hitpoints,
        sanity,
        luck,
        magic_points,
        diceroll,
        session_id
    )
values (
        'testname',
        'testplayer',
        'testoccupation',
        1,
        'female',
        'testresidence',
        'stuttgart',
        1,
        1,
        1,
        1,
        1,
        1,
        1,
        1,
        1,
        1,
        1,
        1,
        1,
        1
    );
    ```
