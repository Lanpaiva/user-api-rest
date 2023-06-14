create table users(
    id serial primary key,
    nome varchar,
    senha varchar
);

INSERT INTO users(nome, senha) VALUES
('username1', '70567254aa'),
('username2', '70567254aa');