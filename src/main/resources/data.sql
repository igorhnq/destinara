CREATE TABLE travel_package (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    description TEXT
);

INSERT INTO travel_package (name, type, price, description) VALUES
('Manaus, AM', 'Nacional', 1500, 'Pacote Manaus show de bola');

INSERT INTO travel_package (name, type, price, description) VALUES
('Foz do Iguacu, PR', 'Nacional', 1000, 'Pacote Foz do Iguaçu show de bola');

INSERT INTO travel_package (name, type, price, description) VALUES
('Salvador, BA', 'Nacional', 1500, 'Pacote Salvador show de bola');

INSERT INTO travel_package (name, type, price, description) VALUES
('Gramado, RS', 'Nacional', 1500, 'Pacote Gramado show de bola');

INSERT INTO travel_package (name, type, price, description) VALUES
('BALI, Indonesia', 'Internacional', 4000, 'Pacote BALI show de bola');

INSERT INTO travel_package (name, type, price, description) VALUES
('Toquio, Japao', 'Internacional', 3800, 'Pacote Japao show de bola');

INSERT INTO travel_package (name, type, price, description) VALUES
('Vik, Islandia', 'Internacional', 10000, 'Pacote Vik show de bola');

INSERT INTO travel_package (name, type, price, description) VALUES
('Santorini, Grecia', 'Internacional', 6900, 'Pacote Santorini show de bola');
