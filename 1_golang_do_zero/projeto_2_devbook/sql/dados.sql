insert into usuarios (nome, nick, email, senha) values
("Teste 1", "teste1", "teste1@gmail.com", "$2a$10$9yuX74rEmuVyc5O0f2v2qefjGof8Gop1y9xFSnETyKVwTfltclpsa"),
("Teste 2", "teste2", "teste2@gmail.com", "$2a$10$9yuX74rEmuVyc5O0f2v2qefjGof8Gop1y9xFSnETyKVwTfltclpsa"),
("Teste 3", "teste3", "teste3@gmail.com", "$2a$10$9yuX74rEmuVyc5O0f2v2qefjGof8Gop1y9xFSnETyKVwTfltclpsa");

insert into seguidores (seguido_id, seguidor_id) values
(1, 2),
(3, 1),
(1, 3);

 insert into publicacoes (titulo, conteudo, autor_id) values
 ("Publicação do usuario 1", "Conteúdo da publicação do usuario 1", 1),
 ("Publicação do usuario 2", "Conteúdo da publicação do usuario 2", 2),
 ("Publicação do usuario 3", "Conteúdo da publicação do usuario 3", 3);