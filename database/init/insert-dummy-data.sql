INSERT INTO ParentRoute (parent_route_id, parent_route_name) VALUES
    ('route-1', 'Route 1'),
    ('route-2', 'Route 2'),
    ('route-3', 'Route 3');

INSERT INTO Stop (stop_id, stop_name) VALUES
    ('stop-1', 'Station A'),
    ('stop-2', 'Station B'),
    ('stop-3', 'Station C');

INSERT INTO Route (route_id, parent_route_id, route_name, origin_stop_id, destination_stop_id) VALUES
    ('route-1', 'route-1', 'Route 1', 'stop-1', 'stop-2'),
    ('route-2', 'route-2', 'Route 2', 'stop-2', 'stop-3'),
    ('route-3', 'route-3', 'Route 3', 'stop-3', 'stop-1');

INSERT INTO Calendar (service_id, monday, tuesday, wednesday, thursday, friday, saturday, sunday, description) VALUES
    ('weekday', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, 'Weekday service'),
    ('daily', TRUE, TRUE, TRUE, TRUE, TRUE, TRUE, TRUE, 'Daily service');

INSERT INTO Trip (trip_id, route_id, service_id) VALUES
    ('trip-1', 'route-1', 'weekday'),
    ('trip-2', 'route-1', 'weekday'),
    ('trip-3', 'route-1', 'weekday'),
    ('trip-4', 'route-1', 'weekday'),
    ('trip-5', 'route-1', 'weekday'),
    ('trip-6', 'route-1', 'weekday'),
    ('trip-7', 'route-1', 'weekday'),
    ('trip-8', 'route-1', 'weekday'),
    ('trip-9', 'route-1', 'weekday'),
    ('trip-10', 'route-1', 'weekday'),
    ('trip-11', 'route-1', 'weekday'),
    ('trip-12', 'route-1', 'weekday'),
    ('trip-13', 'route-1', 'weekday'),
    ('trip-14', 'route-1', 'daily');

INSERT INTO StopTime (trip_id, stop_id, arrival_time, departure_time, stop_sequence, pickup_only_flag, drop_off_only_flag) VALUES
    ('trip-1', 'stop-1', '08:00', '08:00', 1, TRUE, FALSE),
    ('trip-2', 'stop-1', '10:00', '10:00', 1, TRUE, FALSE),
    ('trip-3', 'stop-1', '12:00', '12:00', 1, TRUE, FALSE),
    ('trip-4', 'stop-1', '14:00', '14:00', 1, TRUE, FALSE),
    ('trip-5', 'stop-1', '16:00', '16:00', 1, TRUE, FALSE),
    ('trip-6', 'stop-1', '18:00', '18:00', 1, TRUE, FALSE),
    ('trip-7', 'stop-1', '19:00', '19:00', 1, TRUE, FALSE),
    ('trip-8', 'stop-1', '20:00', '20:00', 1, TRUE, FALSE),
    ('trip-9', 'stop-1', '20:15', '20:15', 1, TRUE, FALSE),
    ('trip-10', 'stop-1', '20:17', '20:17', 1, TRUE, FALSE),
    ('trip-11', 'stop-1', '21:00', '21:00', 1, TRUE, FALSE),
    ('trip-12', 'stop-1', '22:00', '22:00', 1, TRUE, FALSE),
    ('trip-13', 'stop-1', '23:00', '23:00', 1, TRUE, FALSE),
    ('trip-14', 'stop-1', '23:30', '23:30', 1, TRUE, FALSE);

