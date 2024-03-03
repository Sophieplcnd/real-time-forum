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

            <label for="firstname">First Name:</label>
            <input type="text" id="firstname" name="firstname" required><br>

            <label for="lastname">Last Name:</label>
            <input type="text" id="lastname" name="lastname"><br>
            
            <label for="age">Age:</label>
            <input type="text" id="age" name="age" pattern="[0-9]+"><br>

            <label for="gender">Gender:</label>
            <select id="gender">
              <option value="female">Female</option>
              <option value="male">Male</option>
              <option value="other">Other</option>
              <option value="N/A">Prefer Not to Say</option>
            </select><br> 

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
  const userData = {
    username: document.getElementById("username").value,
    password: document.getElementById("password").value,
    email: document.getElementById("email").value,
    firstname: document.getElementById("firstname").value,
    lastname: document.getElementById("lastname").value,
    age: document.getElementById("age").value,
    gender: document.getElementById("gender").value,
  }

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
<<<<<<< HEAD
        console.log("registration attempted")
=======
>>>>>>> 4c48961abef3d431cccca8c9f7f8a6c0ae67f07d
        if (response.ok) {
          alert("registration successful")
				} else {
          alert("registration unsuccessful, please check your details")        
        }
			} catch (error) {
        console.error(error);
			}
}