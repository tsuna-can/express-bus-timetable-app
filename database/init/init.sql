CREATE TABLE parent_route (
    parent_route_id TEXT PRIMARY KEY,
    parent_route_name TEXT NOT NULL
);

CREATE TABLE route (
    route_id TEXT PRIMARY KEY,
    parent_route_id TEXT NOT NULL,
    route_name TEXT NOT NULL,
    origin_stop_id TEXT NOT NULL,
    destination_stop_id TEXT NOT NULL,
    FOREIGN KEY (parent_route_id) REFERENCES parent_route (parent_route_id)
);

CREATE TABLE stop (
    stop_id TEXT PRIMARY KEY,
    stop_name TEXT NOT NULL
);

CREATE TABLE calendar (
    service_id TEXT PRIMARY KEY,
    monday BOOLEAN NOT NULL,
    tuesday BOOLEAN NOT NULL,
    wednesday BOOLEAN NOT NULL,
    thursday BOOLEAN NOT NULL,
    friday BOOLEAN NOT NULL,
    saturday BOOLEAN NOT NULL,
    sunday BOOLEAN NOT NULL,
    description TEXT
);

CREATE TABLE trip (
    trip_id TEXT PRIMARY KEY,
    route_id TEXT NOT NULL,
    service_id TEXT NOT NULL,
    FOREIGN KEY (route_id) REFERENCES route (route_id),
    FOREIGN KEY (service_id) REFERENCES calendar (service_id)
);

CREATE TABLE stop_time (
    trip_id TEXT NOT NULL,
    stop_id TEXT NOT NULL,
    arrival_time TIME NOT NULL,
    departure_time TIME NOT NULL,
    stop_sequence INTEGER NOT NULL,
    pickup_only_flag BOOLEAN NOT NULL,
    drop_off_only_flag BOOLEAN NOT NULL,
    PRIMARY KEY (trip_id, stop_id, stop_sequence),
    FOREIGN KEY (trip_id) REFERENCES trip (trip_id),
    FOREIGN KEY (stop_id) REFERENCES stop (stop_id)
);
