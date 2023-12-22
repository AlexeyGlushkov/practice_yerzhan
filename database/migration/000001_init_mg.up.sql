CREATE TABLE employees (
    employee_id BIGINT PRIMARY KEY,
    firstname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    email VARCHAR(254) NOT NULL,
    manager_id BIGINT NOT NULL
);

CREATE TABLE managers (
    manager_id BIGINT PRIMARY KEY,
    firstname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    email VARCHAR(254) NOT NULL
);

ALTER TABLE employees
    ADD FOREIGN KEY (manager_id) REFERENCES managers (manager_id);