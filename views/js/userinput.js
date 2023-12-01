function buttonSubmit() {
    const titleInput = document.getElementById('linkYoutube');		//* GET input
	if (titleInput.value === "") {
		console.log(titleInput.value)
		alert("Fill the link!")
		return false
	}
    const urlData = {
        url: titleInput.value
    };

    fetch('/download/', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(urlData)		
	})
	.then(response => response.json())
	.then(data => {
		// 'data' will contain the JSON data received from the server
		if (data.status === "OK") {
			popup();
		}
		else alert("Error");
	})
	.catch(error => {
		console.error('Error:', error);
	})
}

function popup() {
	alert("Downloaded Successfully!");
}
