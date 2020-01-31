let pos;
let cord;
var url = "results.html?"; 

/**
 * Function that gets the users ive location from the browser
 * 
 */
function get_location() {

    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(position => {
            //console.log(position.coords);
            alert(position, coords);
            pos = position; 
            cord = coords; 
        });
    } else {
        console.log("Error occured");

    }

}



/**
 * Method that is used for returning the value of the current  code entered by the user
 */
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
    };
}





function zip_code_info(zip) {

    console.log("hi there");

    console.log(zip);
    return new Promise((resolve, reject) => {
        $.ajax({
            url: 'https://geocoder.ls.hereapi.com/6.2/geocode.json',
            type: 'GET',
            dataType: 'jsonp',
            jsonp: 'jsoncallback',
            data: {
                PostalCode: zip,
                country: 'usa',
                gen: '9',
                apiKey: '-eGJp8gx7bCBX_ZnC56TIvGCssDMgCXdPl9ecE87g8Q'
            },
            success: function (data) {
                JSON.stringify(data);


                var lat = (data.Response.View[0].Result[0].Location.DisplayPosition.Latitude);
                var lon = (data.Response.View[0].Result[0].Location.DisplayPosition.Longitude);
                resolve({lat, lon})


            }
        });
    });
}





function submision() {
var input = user_input();


{

    zip_code_info(input.zip)
    .then((data)=>{
        console.log("submission")
        console.log(data.lat, data.lon);
        window.location = url +'query=' + input.injury & 'max_price='+input.price & 'proximity=' + input.distance & 'lat=' + data.lat & 'lon=' + data.lon; 

    })


    console.log(input.zip, input.price, input.distance, input.injury);

    
    
    
    }
}