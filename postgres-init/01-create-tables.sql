BEGIN;

CREATE TABLE IF NOT EXISTS devices (
    device_id SERIAL PRIMARY KEY,
    mac VARCHAR(17), -- e.g. AA:BB:CC:DD:EE:FF
    UNIQUE(mac),
    device_type TEXT,
    device_name TEXT
);

CREATE TABLE IF NOT EXISTS measurements (
    id SERIAL PRIMARY KEY,
    device_id INT,
    FOREIGN KEY (device_id) REFERENCES devices(device_id),
    temperature INT,
    humidity INT,
    battery INT
);
COMMIT;