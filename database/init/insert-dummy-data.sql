INSERT INTO parent_route (parent_route_id, parent_route_name) VALUES
    ('parent-route-1', '千葉号（上り）'),
    ('parent-route-2', '千葉号（下り）');

INSERT INTO stop (stop_id, stop_name) VALUES
    ('stop-1', '東京駅'),
    ('stop-2', '船橋駅'),
    ('stop-3', '幕張'),
    ('stop-4', '千葉駅');

INSERT INTO route (route_id, parent_route_id, route_name, origin_stop_id, destination_stop_id) VALUES
    ('route-1', 'parent-route-1', '東京駅行き', 'stop-4', 'stop-1'),
    ('route-2', 'parent-route-2', '千葉駅行き', 'stop-1', 'stop-4');

INSERT INTO calendar (service_id, monday, tuesday, wednesday, thursday, friday, saturday, sunday, description) VALUES
    ('weekday-id', TRUE, TRUE, TRUE, TRUE, TRUE, FALSE, FALSE, 'weekday service'),
    ('holiday-id', FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, TRUE, 'holiday service');

INSERT INTO trip (trip_id, route_id, service_id) VALUES
    ('trip-1', 'route-1', 'weekday-id'),
    ('trip-2', 'route-1', 'weekday-id'),
    ('trip-3', 'route-1', 'weekday-id'),
    ('trip-4', 'route-1', 'weekday-id');

INSERT INTO stop_time (trip_id, stop_id, arrival_time, departure_time, stop_sequence, pickup_only_flag, drop_off_only_flag) VALUES
    ('trip-1', 'stop-4', '19:00', '19:00', 1, TRUE, FALSE),
    ('trip-1', 'stop-3', '19:10', '19:10', 1, TRUE, FALSE),
    ('trip-1', 'stop-2', '19:35', '19:35', 1, TRUE, FALSE),
    ('trip-1', 'stop-1', '20:00', '20:00', 1, TRUE, FALSE),
    ('trip-2', 'stop-4', '19:15', '19:15', 1, TRUE, FALSE),
    ('trip-3', 'stop-4', '19:30', '19:30', 1, TRUE, FALSE),
    ('trip-4', 'stop-4', '19:45', '19:45', 1, TRUE, FALSE);
