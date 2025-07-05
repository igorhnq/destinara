CREATE TABLE travel_package (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    paragraph1 TEXT,
    paragraph2 TEXT
);

CREATE TABLE app_user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    cpf VARCHAR(11) NOT NULL,
    birth_date DATE NOT NULL,
    sex CHAR(1) NOT NULL,
    email VARCHAR(100),
    password VARCHAR(25) NOT NULL
);

CREATE TABLE purchase (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    travel_package_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES app_user(id),
    FOREIGN KEY (travel_package_id) REFERENCES travel_package(id)
);

INSERT INTO travel_package (name, type, price, paragraph1, paragraph2) VALUES
('Manaus, AM', 'Nacional', 1500, 'Manaus, no coracao da Floresta Amazonica, e conhecida pela biodiversidade e importancia historica no ciclo da borracha. Principais atracoes: o Teatro Amazonas e o Encontro das Aguas, onde os rios Negro e Solimoes se encontram sem se misturar.', 'Pacotes turisticos para Manaus incluem visitas guiadas pela floresta, cruzeiros e excursoes para ver fauna local como botos e macacos. Pacotes de 3 a 7 dias custam de R$ 1.500 a R$ 4.500, com transporte, hospedagem e algumas refeicoes.');

INSERT INTO travel_package (name, type, price, paragraph1, paragraph2) VALUES
('Foz do Iguacu, PR', 'Nacional', 1000, 'Foz do Iguacu, localizada na fronteira entre Brasil, Argentina e Paraguai, e famosa pelas Cataratas do Iguacu e pela Usina de Itaipu. A cidade tambem abriga o Parque das Aves e o Marco das Tres Fronteiras, onde os rios Iguacu e Parana se encontram.', 'Os pacotes turisticos costumam incluir visitas as cataratas (tanto pelo lado brasileiro quanto pelo argentino), passeios de barco, e ate sobrevoos de helicoptero. Um pacote de 4 dias custa entre R$ 1.500 e R$ 3.000, incluindo hospedagem e transporte.');

INSERT INTO travel_package (name, type, price, paragraph1, paragraph2) VALUES
('Salvador, BA', 'Nacional', 1500, 'A capital baiana e um dos principais centros culturais do Brasil, com o Pelourinho listado como Patrimonio Mundial pela UNESCO. A cidade possui atrações como a Igreja de Sao Francisco, o Elevador Lacerda, que liga a cidade alta à cidade baixa.', 'Os pacotes turisticos variam de R$ 1500 a R$ 3000 para estadias de 5 a 7 dias, com hospedagem, visitas guiadas e acesso as praias. Durante o Carnaval, os pacotes incluem camarotes e ingressos para trios eletricos, com precos mais elevados.');

INSERT INTO travel_package (name, type, price, paragraph1, paragraph2) VALUES
('Gramado, RS', 'Nacional', 1500, 'Gramado, na Serra Gaucha, e famosa pelo clima europeu, arquitetura alpina e eventos como Natal Luz e o Festival de Cinema. Atrações incluem Lago Negro, Mini Mundo, Rua Coberta e o parque Snowland com neve artificial no Brasil.', 'Os pacotes turisticos para Gramado sao populares no inverno e Natal Luz, com precos de R$ 2.000 a R$ 4.000 para pacotes de 4 a 6 dias, incluindo transporte, hospedagem em hoteis charmosos, ingressos para parques e eventos especiais.');

INSERT INTO travel_package (name, type, price, paragraph1, paragraph2) VALUES
('BALI, Indonesia', 'Internacional', 4000, 'Bali e rica em diversidade cultural e paisagens, com praias, montanhas e campos de arroz. As principais cidades sao Ubud, Canggu e Seminyak. Destaques incluem os templos Tanah Lot e Uluwatu, praias de Nusa Dua e Sanur, e o terraço de arroz de Tegalalang.', 'Pacotes turisticos incluem visitas a templos, aulas de surf em Kuta e passeios para as ilhas Gili. Aluguel de moto custa de R$ 114 a R$ 171 por semana. Motorista privado custa cerca de R$ 285. Algumas atracoes cobram taxas de entrada, cerca de 15.000 IDR.');

INSERT INTO travel_package (name, type, price, paragraph1, paragraph2) VALUES
('Toquio, Japao', 'Internacional', 3800, 'Toquio e o coracao do Japao, com pontos turisticos como Shibuya, Shinjuku e Akihabara. Atraçoes incluem o Templo Senso-ji, Torre de Toquio, Jardim Imperial e Tokyo Disneyland. Visite Roppongi Hills e os bares de Golden Gai para uma experiencia unica.', 'Pacotes turisticos para Toquio começam em R$ 8.534,10, incluindo passagens, hospedagem e passeios opcionais como barco yakata-bune e ingressos para Tokyo Disneyland. O metrô e o Japan Rail Pass facilitam viagens pela cidade e para Kyoto e Osaka.');

INSERT INTO travel_package (name, type, price, paragraph1, paragraph2) VALUES
('Vik, Islandia', 'Internacional', 10000, 'Reykjavik e um destino para quem busca aventura e natureza. Atrações incluem a Blue Lagoon, o Circulo Dourado (Thingvellir, Geysir, Gullfoss), a Harpa Concert Hall e a Igreja Hallgrimskirkja, que oferece vistas panoramicas da cidade.', 'Pacotes turisticos incluem passeios para ver a aurora boreal (setembro a abril), cavernas de gelo e trilhas por vulcoes e geleiras');

INSERT INTO travel_package (name, type, price, paragraph1, paragraph2) VALUES
('Santorini, Grecia', 'Internacional', 6900, 'Santorini e famosa pelas paisagens pitorescas e destino popular para lua de mel. As principais cidades sao Fira, centro economico, e Oia, famosa pelos pores do sol.', 'Os pacotes turisticos para Santorini incluem hospedagem em hoteis de luxo com vista para o mar Egeu e passeios de barco. Os precos variam de R$6.827 a R$14.223 por semana, podendo chegar a mais de R$17.068 para hoteis boutique com spa e jantares privados.');
