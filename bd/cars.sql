CREATE DATABASE IF NOT EXISTS carsdb;
USE carsdb;
DROP TABLE IF EXISTS cars;
CREATE TABLE cars(id INT PRIMARY KEY AUTO_INCREMENT, manufacturer VARCHAR(255), name VARCHAR(255), fuel VARCHAR(255));
INSERT INTO cars(manufacturer, name, fuel) VALUES('Mersedes','Viano', 'diesel');
INSERT INTO cars(manufacturer, name, fuel) VALUES('Volvo','v40', 'diesel');
INSERT INTO cars(manufacturer, name, fuel) VALUES('Opel','Rekord', 'benzin');
INSERT INTO cars(manufacturer, name, fuel) VALUES('Skoda','Kodiaq', 'diesel');
INSERT INTO cars(manufacturer, name, fuel) VALUES('Mersedes','c180', 'benzin');
INSERT INTO cars(manufacturer, name, fuel) VALUES('VAZ','2106', 'gas');
INSERT INTO cars(manufacturer, name, fuel) VALUES('tesla','1', 'electro');
INSERT INTO cars(manufacturer, name, fuel) VALUES('tesla','2', 'electro');