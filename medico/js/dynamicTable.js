$(document).ready(function(){
    $.each(data.objects, function(index, item) {
        console.log(item)
        var result = `
            <tr>
            <td>${item.provider_name}</td>
            <td>${item.provider_street + " " + item.provider_city}</td>
            <td>${item.average_total_payments}</td>
            </tr>
            `
            $("#resultsTable tbody").append(result)
    })
})

