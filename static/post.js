export default class PostPage {
    constructor() {
      // super();
      //this.setTitle("Login");
    }

    renderHTML() {
        return `
        <h2>Create a New Post</h2>
        <form id="create-post-form">
            <label for="post-title">Title:</label>
            <input type="text" id="post-title" name="post-title" required>

            <label for="post-content">Content:</label>
            <textarea id="post-content" name="post-content" required></textarea>

            <button type="submit">Create Post</button>
        </form>
        `;
      }
}