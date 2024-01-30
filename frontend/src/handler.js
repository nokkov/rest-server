function submitForm() {
    let longUrl = document.getElementById('longUrlField').value; 

    fetch('localhost:8080/save-url', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },

        body: JSON.stringify({"LONG_URL":longUrl})
    })

    .then(response => response.json())
    .then(response => console.log(JSON.stringify(response)))
}

export default submitForm;