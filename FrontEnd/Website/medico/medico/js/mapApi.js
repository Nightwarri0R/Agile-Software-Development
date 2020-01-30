/**
 * Funvtion that gets the users ive location from the browser
 * 
 */
function get_location() {

    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(position => {
            //console.log(position.coords);
            alert(position, coords);
            var live_locaion = {};
            position = live_locaion[0];
            coords = live_locaion[1];
            return live_locaion;
        });
    } else {
        console.log("Error occured");



    }

}



function return_injury() {
    var injury = document.getElementById("procedure_type").value;

    return injury;
}

function return_price() {
    var price = document.getElementById("price_range").value;

    return price;
}

function return_distance() {
    var distance = document.getElementById("distance_range").value;

    return distance;
}

/**
 * Method that is used for returning the value of the zip code entered by the user
 */
function return_input() {
    var zip = document.getElementById("zip_code").value;


    return zip;
}


function user_input() {
    return {
        injury: return_injury(),
        price: return_price(),
        distance: return_distance(),
        zip: return_input()
    }
}





function zip_code_info(zip) {

    console.log("hi there")

    console.log(zip);
    $.ajax({
        url: 'https://geocoder.ls.hereapi.com/6.2/geocode.json',
        type: 'GET',
        dataType: 'jsonp',
        jsonp: 'jsoncallback',
        data: {
            PostalCode: zip,
            country: 'usa',
            gen: '9',
            apiKey: 'wS3zBaE7wLZ5Im9u7TfCFwbDmPABPKlSCCg_7s4JM-U'
        },
        success: function (data) {
            JSON.stringify(data);


            var lat = (data.Response.View[0].Result[0].Location.DisplayPosition.Latitude);
            var lon = (data.Response.View[0].Result[0].Location.DisplayPosition.Longitude);

            send_backend(lat,lon);
        }
    });
}

function send_backend(lat,lon){

    console.log(lat,lon);
}

function submision() {
    var input = user_input();
    zip_code_info(input.zip.value);
    console.log(input.injury, input.price, input.distance, input.zip);

}

