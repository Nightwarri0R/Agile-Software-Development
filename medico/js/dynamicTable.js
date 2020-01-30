// Disable searching on table
$.extend( true, $.fn.dataTable.defaults, {
    "searching": false
} );

// Init table on document load
$(document).ready(function() {
    $('#resultsTable').DataTable({
        "scrollY": "80vh",
        "scrollCollaps": true,
        info: false,
        select: true,
        paging: false,
        data: data.objects,
        columns: [
            {
                class_name: "full-span-col",
                data: null,
                title: "Name",
                render: function(data, type, row, meta) {
                    return `
                        <div class="card">
                            <div class="card-body">
                                <h5 class="card-title">${row.provider_name}</h5>
                                <h6 class="card-subtitle mb-2 text-muted">${row.hrr_description}</h6>
                                <div class="row">
                                    <div class="col">
                                        <i class="fas fa-road"></i> Street - ${row.provider_street}</i>
                                    </div>
                                    <div class="col">
                                        <i class="fas fa-city"></i> City - ${row.provider_city}</i>
                                    </div>
                                    <div class="w-100"></div>
                                    <div class="col">
                                        <i class="fas fa-location-arrow"></i> Zip Code - ${row.provider_zip_code}</i>
                                    </div>
                                    <div class="col">
                                        <i class="fas fa-dollar-sign"> Price - $${row.average_total_payments}</i>
                                    </div>
                                </div>
                            </div>
                        </div>
                    `
                }
            },
            {
                data: "average_total_payments",
                title: "Price",
            },
            {
                data: "distance",
                title: "Distance",
            }
        ],
        createdRow: function(row, data, dataIndex){
            // Add COLSPAN attribute
            $('td:eq(0)', row).attr('colspan', 3);

            // Hide required number of columns
            // next to the cell with COLSPAN attribute
            $('td:eq(1)', row).css('display', 'none');
            $('td:eq(2)', row).css('display', 'none');
        }
    });
} );