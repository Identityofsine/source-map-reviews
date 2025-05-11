CREATE DATABASE app;
CREATE USER 'docker'@'%' IDENTIFIED BY 'docker';
GRANT ALL PRIVILEGES ON your_database_name.* TO 'docker'@'%';
FLUSH PRIVILEGES;
