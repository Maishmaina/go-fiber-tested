# My Awesome Project

This is a description of my Go with Fiber Framework.

## Getting Started

Follow these steps to get started with the project:
- Step 1: Clone the repository
   - (git clone https://github.com/Maishmaina/go-fiber-tested.git)
- Step 2: Install dependencies
  - (cd go-fiber-tested)
  - go mod tidy
- Step 3: Run the project
  - (go run .)

## Usage

below show some screen short of use scenario 

User this code to create mysql db and table 

       CREATE DATABASE IF NOT EXISTS chama_soft;
       USE chama_soft;

    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL
    );

User register request API.
<img src="/assets/Postman register.PNG" />

User Login request API.
<img src="/assets/loginsuccessful.PNG" />

Logger Middleware.
<img src="/assets/loggermiddleware.PNG" />

List of Users in mysqldb.
<img src="/assets/userindb.PNG" />

User List fetch by API.
<img src="/assets/userlistig.PNG.PNG" />

Auth Middleware.
<img src="/assets/accessing protectedroutes.PNG" />




## License

This project is licensed under the MIT License - see the LICENSE.md file for details.
