# Express Bus Time Table app
Ths is a simple WearOS app that provides a bus time table for a given route. 

# Features

## App
The app consists of the following three screens.

| Screen | Description | Screenshot |
|--------|-------------|------------|
| Route List Screen | Displays a list of available express bus routes. | ![Route List Screen](images/route_list.png) |
| Bus Stop List Screen | Displays a list of bus stops associated with the selected route. | ![Bus Stop List Screen](images/bus_stop_list_1.png) |
| Timetable Screen | Displays the timetable for the selected bus stop. You can also register bus stops to be shown on a tile. | ![Timetable Screen](images/timetable_1.png) | 

## Tile

| Screen | Description | Screenshot |
|--------|-------------|------------|
| Tile | The tile displays the upcoming timetable for the registered bus stop. | ![Tile](images/tile.png) | 

# Architecure
```mermaid
graph LR
    MobileApp[Mobile App] -->|Request| API
    API -->|Response| MobileApp

    subgraph koyeb[Cloud: Koyeb]
        API[API Server] -->|Query| DB[(Database)]
        DB -->|Data| API
    end
```

# Tech Stack
- Mobile
    - Kotlin
    - Jetpack Compose
- Backend
    - Go
    - Echo
- Database
    - PostgreSQL
- Hosting
    - [Koyeb](https://www.koyeb.com/)