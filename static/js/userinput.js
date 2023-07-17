function buttonSubmit() {
    const titleInput = document.getElementById('linkYoutube');		//* GET input
	const inputValue = titleInput.value;	

    const urlData = {
        url: inputValue
    };
    alert(urlData);

    fetch('/api/link', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(urlData)		
		})
    
	.catch(error => {
		console.error('Error:', error);
	});
}