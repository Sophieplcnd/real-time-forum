export default class LoginPage {
  constructor() {
    // super();
    //this.setTitle("Login");
  }

  renderHTML() {
    return `
        <h2>Login</h2>
        <form id="login-form">
            <label for="email-username">Email or Username:</label>
            <input type="text" id="email-username" name="email-username" required><br>

            <label for="password">Password:</label>
            <input type="password" id="password" name="password" required><br>

            <button type="button" id="login-button">Login</button>
        </form>
    `;
  }
}

export async function loginHandler() {
  console.log("Button working");

  // get user login details
  const emailUsername = document.getElementById("email-username").value;
  const password = document.getElementById("password").value;

  try {
    const response = await fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: emailUsername,
        password: password,
      }),
    });
    // handle login, set cookies
    const data = await response.json();
  } catch (error) {
    console.error(err);
  }
}

// if (emailUsername === "email" && password === "password") {
//   alert("You have successfully logged in.");
// } else {
//   alert("Incorrect Email/Username and/or Password");
// }

// document.addEventListener("DOMContentLoaded", function () {
//   const loginButton = document.getElementById("login-button");
//   if (loginButton) {
//     console.log("Button found!");
//     loginButton.addEventListener("click", loginHandler);
//   } else {
//     console.error("Button not found!");
//   }
// });
