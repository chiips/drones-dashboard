# Drone Dashboard

This project is a live dashboard for tracking the activity of drones. Sample drones are created on start, and visitors can add new drones, specifying their starting coordinates, which will then began to travel on their own. A map presents the location of the drones and the table present each drone's details: ID, latitude, longitude, speed, and stationary status. The rows of stationary drones will be highlighted red.

## Backend

The backend is built on Go, uses websockets for regular updates to the frontend, and uses in-memory data storage.

## Frontend

The frontend is built on HTML5, CSS3, and Javascript, and uses websockets to receive updates from the backend.

## Licence

This code is licensed under the MIT License.