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

```psql -U $DBNAME```

You can access your exisiting databases with ```\l``` 
 
At this point, the database and tables should be initialized. This can be done using basic SQL.
```
CREATE DATABASE distributed_dungeons;
CREATE TABLE session (
  sessionkey VARCHAR(255),
  password VARCHAR(255)
  );
CREATE TABLE characters ( 
  name VARCHAR(255),
  player VARCHAR(255),
  occupation VARCHAR(255),
  age INT,
  sex VARCHAR(255),
  residence VARCHAR(255),
  birthplace VARCHAR(255),
  str INT,
  dex INT,
  pow INT,
  con INT,
  app INT,
  edu INT,
  siz INT,
  int INT,
  hitpoints INT,
  sanity INT,
  luck INT,
  magicpoints INT,
  diceroll INT
  );
```

Note: The character table is currently designed for the "Call of Cthulu" game.

A value in a table can be updated like this: 

```UPDATE characters SET diceroll=1 WHERE name='Oliver Smith';```
  
  
