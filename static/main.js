import LoginPage, { loginHandler } from "./login.js";
import HomePage from "./home.js";
import RegisterPage, { registerHandler } from "./registration.js";
import PostPage from "./post.js";

// function to update and redirect browser url
const navigateTo = (page) => {
  history.pushState(null, null, page);
  pageLoader();
};


const pageLoader = async () => {
  // paths and imported JS pages
  const pages = [
    { path: "/", view: HomePage },
    { path: "/login", view: LoginPage },
    { path: "/register", view: RegisterPage },
    { path: "/posts", view: PostPage },
  ];

  // goes through pages array, returns true if current URL matches page in array
  let page = pages.find((page) => location.pathname === page.path);
  // if page cannot be found, redirect to homepage
  if (page === undefined) page = pages[0];

  // reassigning view depending on desired page
  const view = new page.view();
  // declaring html variable using renderHTML function in JS pages
  const html = view.renderHTML();

  // find and replace div container element with html of chosen page
  document.querySelector("#container").innerHTML = html;

  // find current page, get elements needed for that page
  switch (page.view) {
    case LoginPage:
      console.log("login page loaded");
      document
        .getElementById("login-button")
        .addEventListener("click", loginHandler);
      break;
    case RegisterPage:
      console.log("registration page loaded");
      document
      .getElementById("register-button")
      .addEventListener("click", registerHandler);
      break;
    case HomePage:
      console.log("home page loaded");
      break;
    case PostPage:
      console.log("create post page loaded");
      break;
    default:
      console.log("default page loaded");
  }
};

// replacing default forwards and backwards in browser with our pageLoader function
window.addEventListener("popstate", pageLoader);

// add event listener so that all links will have the same behaviour as our pageLoader function
document.addEventListener("click", (e) => {
  if (e.target.tagName === "A") {
    e.preventDefault();
    navigateTo(e.target.href);
  }
});
