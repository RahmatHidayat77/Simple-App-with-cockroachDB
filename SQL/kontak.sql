DROP TABLE IF EXISTS kontak_db;

create database kontak_db;

set database=kontak_db;

CREATE TABLE kontak (
id SERIAL NOT NULL,
nama STRING(25),
alamat STRING(50),
telp STRING(20)
);

INSERT INTO kontak (nama, alamat, telp) VALUES ('Andi','Sleman','085123789');
INSERT INTO kontak (nama, alamat, telp) VALUES ('Indah','Klaten','08513513');
INSERT INTO kontak (nama, alamat, telp) VALUES ('Bagas','Bantul','08651314');



