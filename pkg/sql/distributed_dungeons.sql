-- Drop old tables
drop table if exists characters;
drop table if exists sessions;
-- Create new tables
create table characters (
    id serial,
    name text,
    player text,
    occupation text,
    age integer,
    sex text,
    residence text,
    birthplace text,
    strength integer,
    dexterity integer,
    power integer,
    constitution integer,
    appearance integer,
    education integer,
    size integer,
    intelligence integer,
    hitpoints integer,
    sanity integer,
    luck integer,
    magic_points integer,
    diceroll integer,
    session_id integer
);
create table sessions (
    id integer not null primary key,
    password text
);
-- Create test data
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
        null
    );