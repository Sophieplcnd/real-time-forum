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
  console.log("JS loginhandler called");

  // get user login details
  const loginDetails = {
    emailUsername: document.getElementById("email-username").value,
    password: document.getElementById("password").value,
  }

  try {
    const response = await fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(loginDetails),
    });
    if (response.ok) {
      alert("login successful")
    } else {
      alert("login unsuccessful, please check your username/email and/or password")        
    }
  } catch (error) {
    console.error(err);
  }
}
