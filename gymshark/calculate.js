// calculate.js

// Wait for the document to be ready before executing the JavaScript code
$(document).ready(function() {
    // Attach a submit event handler to the form
    $("form").submit(function(e) {
        e.preventDefault(); // Prevent the default form submission

        // Get the number of items from the input field
        var items = $("#items").val();

        // Send an AJAX POST request to the server
        $.ajax({
            type: "POST",
            url: "/calculate",
            data: { items: items },
            success: function(response) {
                // Update the result in the HTML based on the server response
                $("#orderedItems").text(items);
                $("#packsNeeded").text(response.PacksNeeded);
            },
            error: function() {
                // Display an alert if an error occurs during the AJAX request
                alert("An error occurred while processing the request.");
            }
        });
    });
});
