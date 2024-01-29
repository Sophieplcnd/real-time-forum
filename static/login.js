export default class LoginPage {
  constructor() {
    // super();
    //this.setTitle("Login");
  }

  renderHTML() {
    return `
        <h2>Login</h2>
        <form id="login-form">
            <label for="username">Username:</label>
            <input type="text" id="username" name="username" required><br>

            <label for="password">Password:</label>
            <input type="password" id="password" name="password" required><br>

            <button type="button" id="login-button">Login</button>
        </form>
    `;
  }
}

export function loginHandler() {
  alert("Button working");
}

// document.addEventListener("DOMContentLoaded", function () {
//   const loginButton = document.getElementById("login-button");
//   if (loginButton) {
//     console.log("Button found!");
//     loginButton.addEventListener("click", loginHandler);
//   } else {
//     console.error("Button not found!");
//   }
// });
