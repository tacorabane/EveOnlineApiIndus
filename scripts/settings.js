document.getElementById('configForm').addEventListener('submit', function(event) {
    event.preventDefault();

    var screenSize = document.getElementById('screenSize').value;
    var theme = document.getElementById('theme').value;

    var data = {
        screenSize: screenSize,
        theme: theme
    }

    fetch('/update-config', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    })
    .then(function(response) {
        return response.json();
    })
    .then(function(data) {
        console.log('Reponse reçue: ', data);
        alert('Configuration sauvegardée avec succès');
    })
    .catch(function(error) {
        console.error('Erreur: ', error);
        alert('Erreur lors de la sauvegarde de la configuration');
    })
});