export default class RegisterPage {
  constructor() {
    // super();
    //this.setTitle("Home");
  }

  renderHTML() {
    return `
        <h2>Register</h2>
        <form id="register-form">
            <label for="username">Username:</label>
            <input type="text" id="username" name="username" required><br>
            
            <label for="age">Age:</label>
            <input type="text" id="age" name="age" required><br>

            <label for="firstname">First Name:</label>
            <input type="text" id="firstname" name="firstname" required><br>

            <label for="lastname">Last Name:</label>
            <input type="text" id="lastname" name="lastname" required><br>

            <label for="email">Email:</label>
            <input type="text" id="email" name="email" required><br>

            <label for="password">Password:</label>
            <input type="password" id="password" name="password" required><br>

            <button type="button" id="register-button">Register</button>
        </form>
        <div id="login-link">
          <p>Already have an account? <a href="/login" id="link">Login here!</a></p>
        </div>
    `;
  }
}

export async function registerHandler() {
  console.log("JS registerHandler triggered");

  // get user registration details
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;
  const email = document.getElementById("email").value;
  const firstname = document.getElementById("firstname").value;

  const userData = {
    username: username,
    password: password,
    email: email,
    firstname: firstname,
  }

  // // testing functionality with pageLoader on main.js
  // if (email === "email" && password === "password") {
  //   alert("registration successful");
  // } else {
  //   alert("registration unsuccessful");
  // }

  // attempt registration
  try {
    const response = await fetch("/register", {
					method: "POST",
					headers: {
						Accept: "application/json",
						"Content-Type": "application/json",
					},
					body: JSON.stringify(userData),
					// include cookies
				})
        if (response.ok) {
          console.log("registration successful")
				}
			} catch (error) {
				console.error(error)
			}
}