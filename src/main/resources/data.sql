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
('Foz do Iguaçu, PR', 'Nacional', 1000, 'Pacote Foz do Iguaçu show de bola');

INSERT INTO travel_package (name, type, price, description) VALUES
('Salvador, BA', 'Nacional', 1500, 'Pacote Salvador show de bola');

INSERT INTO travel_package (name, type, price, description) VALUES
('Gramado, RS', 'Nacional', 1500, 'Pacote Gramado show de bola');

