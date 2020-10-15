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
        'testname 2',
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