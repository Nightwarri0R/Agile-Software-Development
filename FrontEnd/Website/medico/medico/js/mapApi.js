


var curr_location = document.getElementById('curr_location');

function get_location(){

    if(navigator.geolocation){
    navigator.geolocation.getCurrentPosition(position => {
        console.log(position.coords);
        print(position,coords);
        
    });
    } else {
        console.log("Error occured");
        
}
}
function return_input(){
    var zip = document.getElementById("zip_code").nodeValue;
    return zip;
}


function zip_code_info(zip){
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
        JSON.parse(JSON.stringify(data));
        console.log(data);
        }
    });
    
    console.log("Hello there ")
}


function submision(){
    zip_code_info();
}

