DROP TABLE IF EXISTS Payments;
DROP TABLE IF EXISTS Trips;
DROP TABLE IF EXISTS Drivers;
DROP TABLE IF EXISTS Clients;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Vehicles;

CREATE TABLE Users (
    id INT AUTO_INCREMENT NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE Vehicles (
    id INT AUTO_INCREMENT NOT NULL,
    vehicle_number VARCHAR(255) NOT NULL UNIQUE,
    max_weight VARCHAR(255) NOT NULL,

    PRIMARY KEY(id)
);

CREATE TABLE Clients (
    id INT AUTO_INCREMENT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    PRIMARY KEY(id)
);

CREATE TABLE Drivers (
    id INT AUTO_INCREMENT NOT NULL,
    user_id INT NOT NULL,
    vehicle_id INT NOT NULL,
    unit_cost DECIMAL(5, 2) NOT NULL,
    FOREIGN KEY (User_Id) REFERENCES Users(id),
    FOREIGN KEY (Vehicle_Id) REFERENCES Vehicles(id),
    PRIMARY KEY(id)
);

CREATE TABLE Trips (
    id INT AUTO_INCREMENT NOT NULL,
    client_id INT NOT NULL,
    driver_id INT NOT NULL,
    start DECIMAL(5, 2) NOT NULL,
    end DECIMAL(5, 2) NOT NULL,
    FOREIGN KEY (Client_Id) REFERENCES Clients(id),
    FOREIGN KEY (Driver_Id) REFERENCES Drivers(id),
    PRIMARY KEY(id)
);

CREATE TABLE Payments (
    id INT AUTO_INCREMENT NOT NULL,
    trip_id INT NOT NULL,
    amount DECIMAL(5, 2) NOT NULL,
    FOREIGN KEY (Trip_Id) REFERENCES Trips(id),
    PRIMARY KEY(id)
);
