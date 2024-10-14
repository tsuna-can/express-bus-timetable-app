# Databse

## ER diagram

```mermaid
---
title: Bus Time Table information ER diagram
---
erDiagram
    ParentRoute {
        string parent_route_id PK "Primary key"
        string parent_route_name
    }

    Route {
        string route_id PK "Primary key"
        string parent_route_id FK "Foreign key to ParentRoute"
        string route_name
        string origin_stop_id
        string destination_stop_id
    }

    Stop {
        string stop_id PK "Primary key"
        string stop_name
    }

    Trip {
        string trip_id PK "Primary key"
        string route_id FK "Foreign key to Route"
        string service_id FK "Foreign key to Calender"
    }

    StopTime {
        string trip_id FK "Foreign key to Trip"
        string stop_id FK "Foreign key to Stop"
        time arrival_time
        time departure_time
        int stop_sequence
        boolean pickup_only_flag
        boolean drop_off_only_flag
    }

    Calendar {
        string service_id PK "Primary key"
        boolean monday
        boolean tuesday
        boolean wednesday
        boolean thursday
        boolean friday
        boolean saturday
        boolean sunday
        string description
    }

    ParentRoute ||--|{ Route: "A ParentRoute includes multiple Routes"
    Route ||--o{ Trip: "A Route includes multiple Trips"
    Trip ||--o{ StopTime: "A Trip has multiple StopTimes"
    Stop ||--o{ StopTime: "A Stop is visited by multiple StopTimes"
    Trip ||--|{ Calendar: "A Trip operates on a Calendar"

```
