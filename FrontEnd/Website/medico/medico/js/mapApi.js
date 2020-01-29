


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
function zip_code_info(){
    var zip_code = document.getElementById('Submit');
    var zip = document.getElementsByName("zip_code").values(); 
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
        alert(JSON.stringify(data));
        }
    });

}


function submision(){
    zip_code_info();
}

