document.addEventListener("DOMContentLoaded", function () {
    navigateTo(getCurrentRoute());

    window.addEventListener("hashchange", function () {
        navigateTo(getCurrentRoute());
    });
});

function getCurrentRoute() {
    return window.location.hash.substring(2);
}

function navigateTo(route) {
    const contentDiv = document.getElementById("content");
    contentDiv.innerHTML = "";  // Clear previous content

    switch (route) {
        case "":
        case "/":
            loadHomePage();
            break;
        case "/create-post":
            loadCreatePostPage();
            break;
        // default:
        //     loadNotFoundPage();
        //     break;
    }
}

function loadHomePage() {
    const contentDiv = document.getElementById("content");
    contentDiv.innerHTML = "<h2>Welcome to the Forum!</h2>";
}

function loadCreatePostPage() {
    const contentDiv = document.getElementById("content");
    contentDiv.innerHTML = `
        <h2>Create a New Post</h2>
        <form id="create-post-form">
            <label for="post-title">Title:</label>
            <input type="text" id="post-title" name="post-title" required>

            <label for="post-content">Content:</label>
            <textarea id="post-content" name="post-content" required></textarea>

            <button type="submit">Create Post</button>
        </form>
    `;
    document.getElementById("create-post-form").addEventListener("submit", function (event) {
        event.preventDefault();
        handleCreatePost();
    });
    
}

// function loadNotFoundPage() {
//     const contentDiv = document.getElementById("content");
//     contentDiv.innerHTML = "<h2>404 - Not Found</h2>";
// }

function handleCreatePost() {
    // Handle post creation logic here
    // ...

    // Display a success message for now
    const contentDiv = document.getElementById("content");
    contentDiv.innerHTML = "<h2>Post created successfully!</h2>";
}
