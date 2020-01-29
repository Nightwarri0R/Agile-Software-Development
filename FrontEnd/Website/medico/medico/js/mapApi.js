

var curr_location = document.getElementById('curr_location');

function get_location(){

    if(navigator.geolocation){
    navigator.geolocation.getCurrentPosition(position => {
        console.log(position.coords);
        print(position,coords);
    });
    } else {
        console.log("Error occured");
        console.log("Hi there")
}


}
