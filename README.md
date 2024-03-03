# Real-time Forum
## Overview

Real-time Forum is a basic online forum application built using JavaScript for the frontend, Go for the backend, and SQLite for the database. The application is designed as a single-page application with a focus on simplicity, utilizing vanilla JavaScript and websockets for real-time communication.


### Features

#### User Registration and Login:
- Users can register an account with a unique username and password.
- Existing users can log in to access the forum.

#### Text Posts:
- Registered users can create text posts to share their thoughts or ask questions.

#### Comments:
- Users can comment on posts, providing a platform for discussions.

#### Real-time Chat:
- Users can engage in real-time chat with other online users.

#### Websockets:
- Real-time communication is implemented using websockets, allowing instant updates on posts, comments, and chat.

#### SQLite Database:
- User data, posts, and comments are stored in an SQLite database.

#### Single Page Application (SPA):
- The application is designed as a single-page application, enhancing user experience by reducing page reloads.


### Technologies Used

Frontend:
HTML, CSS, Vanilla JavaScript

Backend:
Go

Database:
SQLite

Also:
- sqlite3
- Bcrypt
- Websockets
- UUID


## Setup Instructions

1. Clone the Repository:

_git clone [https://github.com/your-username/simple-forum.git](https://github.com/Sophieplcnd/real-time-forum)_

2. Navigate to the Project Directory:

_cd simple-forum_

3. Run the Go Backend:

_go run main.go_

4. Open the Application in a Browser:

Open your preferred web browser and navigate to http://localhost:8080.

