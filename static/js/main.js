//declare global variables
let success = document.getElementById("success"); //for successful API responses
let error = document.getElementById("error"); //for errors from API responses
let drones = []; //array to store drone data
let markers = []; //array to hold map markers
let map; //to init google map

//create websocket connection
const websocket = new WebSocket(`ws://${window.location.host}/drones`);

websocket.onopen = function(event) {
  console.log("Successfully connected to websocket server");
};

websocket.onerror = function(error) {
  console.log("Error connecting to websocket server");
  console.log(error);
};

websocket.onmessage = function(event) {

  drones = JSON.parse(event.data);

  let table = document.getElementById("tbody");

  //on every update clear all existing markers for full refresh
  for (let i = 0; i < markers.length; i++) {
    markers[i].setMap(null);
  }
  markers = [];

  drones.forEach(function(drone) {
    //if the drone is not already listed then create a new row for it
    let exists = document.getElementById(drone.id);
    if (exists === null) {
      let row = table.insertRow(table.rows.length);
      row.id = drone.id;

      //highlight stationary drones
      if (drone.stationary == true) {
        row.style.color='red';
      }

      let cell1 = row.insertCell(0);
      let cell2 = row.insertCell(1);
      let cell3 = row.insertCell(2);
      let cell4 = row.insertCell(3);
      let cell5 = row.insertCell(4);

      cell1.innerHTML = drone.id;
      cell2.innerHTML = drone.lat;
      cell3.innerHTML = drone.lon;
      cell4.innerHTML = drone.speed;
      cell5.innerHTML = `<button class="btn btn-danger" onclick="removeDrone(this)">Remove</button>`

    } else {
      //otherwise update existing rows
      let row = document.getElementById(drone.id)
      row.cells[0].innerHTML = drone.id;
      row.cells[1].innerHTML = drone.lat;
      row.cells[2].innerHTML = drone.lon;
      row.cells[3].innerHTML = drone.speed;

      //highlight stationary drones
      if (drone.stationary == true) {
          row.style.color='red';
      } else {
          row.style.color='black';
      }
    }

      //add a marker for each drone
      let droneLatlng = new google.maps.LatLng(drone.lat, drone.lon);
      let marker = new google.maps.Marker({
          position: droneLatlng,
          title: drone.id,
      });

      //add marker to array and map
      markers.push(marker);
      marker.setMap(map);

  });

  //hide waiting notice
  success.style.display = "none"

  //clear rows of since deleted drones
  for (let i = 0, row; row = table.rows[i]; i++) {
    if (!drones.some(drone => drone.id === row.id)) {
      row.parentNode.removeChild(row);
    }

  }

};

//add drone
const droneForm = document.getElementById("droneForm");

droneForm.addEventListener("submit", function(e) {
  e.preventDefault();

  //reset error message if was visible from previous attempt
  error.style.display="none";

  let formLat = document.getElementById("formLat");
  let formLon = document.getElementById("formLon");

  let lat = parseFloat(formLat.value);
  let lon = parseFloat(formLon.value);

  //post values to API
  fetch("/add", {
    method: "POST",
    body: JSON.stringify({"lat": lat, "lon": lon})
  })
  .then(() => {
    //reset form values and notify "sucess, waiting for update"
    this.formLat.value = "";
    this.formLon.value = "";
    success.style.display = "block";
  })
  .catch(() => {
    error.style.display="block";
  })

})

//remove drone
function removeDrone(e) {

  //get drone row and id
  let p = e.parentNode.parentNode;
  let id = p.id;

  //delete drone request to API
  fetch(`/delete?id=${p.id}`, {
    method: "DELETE",
  })
  .then(() => {
    //remove row
    p.parentNode.removeChild(p);
    //remove marker from array and map
    let marker = markers.filter(m => m.title === id)[0]
    let index = markers.indexOf(marker);
    if (index > -1) {
      markers.splice(index, 1);
    }
    marker.setMap(null);
  })
  .catch(() => {
    error.style.display="block";
  })


}

//init google map function
function initMap() {
  map = new google.maps.Map(document.getElementById('map'), {
    center: {lat: 43.65, lng: -79.35},
    zoom: 8
  });
}



