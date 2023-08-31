function buttonSubmit() {
    const titleInput = document.getElementById('linkYoutube');		//* GET input
	const inputValue = titleInput.value;	

    const urlData = {
        url: inputValue
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
	alert("Downloaded");
}
