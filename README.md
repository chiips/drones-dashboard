# Drone Dashboard

This project is a live dashboard for tracking the activity of drones. Sample drones are created on start, and visitors can add new drones, specifying their starting coordinates, which will then begin to travel on their own. A map presents the location of the drones, and a table presents each drone's details: ID, latitude, longitude, speed, and stationary status. The rows of stationary drones will be highlighted red.

## Backend

The backend is built on Go, uses websockets for regular updates to the frontend, and uses in-memory data storage.

## Frontend

The frontend is built on HTML5, the Bootstrap CSS Framework, Javascript, and the Google Maps Javascript API, and uses websockets to receive updates from the backend.

## Licence

This code is licensed under the MIT License.