INSERT INTO ParentRoute (parent_route_id, parent_route_name) VALUES
    ('parent-route-1', '鹿島号（上り）'),
    ('parent-route-2', '鹿島号（下り）');

INSERT INTO Stop (stop_id, stop_name) VALUES
    ('stop-1', '東京駅'),
    ('stop-2', '水郷潮来BT'),
    ('stop-3', 'アートホテル鹿島セントラル'),
    ('stop-4', '鹿島神宮駅');

INSERT INTO Route (route_id, parent_route_id, route_name, origin_stop_id, destination_stop_id) VALUES
    ('route-1', 'parent-route-1', '東京駅行き', 'stop-4', 'stop-1'),
    ('route-2', 'parent-route-2', '鹿島神宮駅行き', 'stop-1', 'stop-4'),
    ('route-3', 'parent-route-2', 'アートホテル鹿島セントラル行き', 'stop-3', 'stop-1');

INSERT INTO Calendar (service_id, monday, tuesday, wednesday, thursday, friday, saturday, sunday, description) VALUES
    ('weekday-id', TRUE, TRUE, TRUE, TRUE, TRUE, FALSE, FALSE, 'weekday service'),
    ('holiday-id', FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, TRUE, 'holiday service');

INSERT INTO Trip (trip_id, route_id, service_id) VALUES
    ('trip-1', 'route-1', 'weekday-id'),
    ('trip-2', 'route-1', 'weekday-id'),
    ('trip-3', 'route-1', 'holiday-id'),
    ('trip-4', 'route-2', 'weekday-id'),
    ('trip-5', 'route-2', 'holiday-id'),
    ('trip-6', 'route-3', 'weekday-id'),
    ('trip-7', 'route-3', 'holiday-id');

INSERT INTO StopTime (trip_id, stop_id, arrival_time, departure_time, stop_sequence, pickup_only_flag, drop_off_only_flag) VALUES
    ('trip-1', 'stop-2', '05:05', '05:05', 1, TRUE, FALSE),
    ('trip-1', 'stop-1', '06:40', '06:40', 2, FALSE, TRUE),
    ('trip-2', 'stop-2', '05:25', '05:25', 1, FALSE, TRUE),
    ('trip-2', 'stop-1', '06:55', '06:55', 1, TRUE, FALSE);
