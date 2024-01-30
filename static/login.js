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
        <div id="register-link">
          <p>Don't have an account? <a href="/register" id=link>Register here!</a></p>
        </div>
    `;
  }
}

export async function loginHandler() {
  console.log("Button working");

  // get user login details
  const emailUsername = document.getElementById("email-username").value;
  const password = document.getElementById("password").value;

  // testing functionality with pageLoader on main.js
  if (emailUsername === "username" && password === "password") {
    alert("login successful");
  } else {
    alert("login unsuccessful");
  }

  // try {
  //   const response = await fetch("/login", {
  //     method: "POST",
  //     headers: {
  //       "Content-Type": "application/json",
  //     },
  //     body: JSON.stringify({
  //       username: emailUsername,
  //       password: password,
  //     }),
  //   });
  //   // handle login, set cookies
  //   const data = await response.json();
  // } catch (error) {
  //   console.error(err);
  // }
}
