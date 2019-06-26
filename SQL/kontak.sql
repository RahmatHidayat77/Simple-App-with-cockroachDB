create database kontak_db;

set database=kontak_db;

CREATE TABLE kontak (
id SERIAL NOT NULL,
nama STRING(25),
alamat STRING(50),
telp STRING(20)
);