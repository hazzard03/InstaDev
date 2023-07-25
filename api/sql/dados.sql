insert into usuarios (nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$mbw5HuLPEvSX.PRhfY1ba.JlWef08x6MSo/u.nlHJZbD4MlM2Zu4C"),  
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$2Agfjyr5gWRIu/FJD07hZOSYGHXrSp2LgIOONxLuyw/sEVc0BPBXK"), 
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$Wy02EXT71oZ0zH.n4bKAJ.VQtw/NMhv1y7CEA03dCXeFmVLAMc4HK"); 

insert into seguidores(usuario_id, seguidor_id)
values
(36, 37),
(38, 36),
(36, 38);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do Usuário 1", "Essa é a publicação do usuário 1! Oba!", 36),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2! Oba!", 37),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3! Oba!", 38);