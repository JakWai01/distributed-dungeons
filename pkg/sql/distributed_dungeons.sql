-- Drop old tables
drop table if exists characters;
drop table if exists sessions;
-- Create new tables
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
    session_id integer
);
create table sessions (
    id integer not null primary key,
    password text not null
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