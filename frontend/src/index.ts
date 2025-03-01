// index.ts
window.addEventListener('DOMContentLoaded', () => {
    // Add event listener to the Load Data button
    const loadDataButton = document.getElementById('loadDataButton') as HTMLButtonElement;

    loadDataButton.addEventListener('click', () => {
        // Fetch user data from the Go backend when the button is clicked
        fetch('http://localhost:8080/LoadData')
            .then(response => response.json())  // Parse the JSON response
            .then(users => {
                const userList = document.getElementById('userList') as HTMLUListElement;
                if (Array.isArray(users)) {
                    // Clear any previous data
                    userList.innerHTML = '';

                    // Populate the list with user data
                    users.forEach((user: { id: string, name: string, email: string }) => {
                        const li = document.createElement('li');
                        li.textContent = `ID: ${user.id}, Name: ${user.name}, Email: ${user.email}`;
                        userList.appendChild(li);
                    });
                }
            })
            .catch(error => {
                console.error('Error fetching user data:', error);
            });
    });
});
