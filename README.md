# Golang Forum

Golang Forum is a web application that allows users to engage in discussions, create posts, associate categories with posts, and interact with the community through likes and dislikes. The project utilizes the SQLite database for data storage and Docker for containerization.

## Features

- User Registration: Users can register on the forum by providing their email, username, and password. The system checks for duplicate emails and securely stores the user's credentials.

- Authentication: Registered users can log in to the forum to access additional features such as creating posts and comments. Sessions are managed using cookies to ensure a seamless user experience.

- Communication: Users can create posts and comments to engage in discussions with other users. The forum supports associating categories with posts, allowing users to categorize their content effectively.

- Likes and Dislikes: Registered users can express their preferences by liking or disliking posts and comments. The number of likes and dislikes is visible to all users, contributing to community engagement.

- Filtering: Users can filter posts based on categories, created posts, and liked posts. This feature allows users to find specific content and personalize their browsing experience.

## Database - SQLite

The SQLite database is used to store the necessary data for the Golang Forum. It provides a reliable and efficient storage solution, ensuring the persistence of user information, posts, comments, and other forum-related data.

## Docker

The project is containerized using Docker, which simplifies deployment and ensures consistency across different environments. Docker allows for easy setup and configuration of the Golang Forum, making it portable and scalable.

## Prerequisites

To run the Golang Forum locally, you need to have the following dependencies installed:

- Go 
- SQLite 
- Docker 

## Installation and Setup

Follow the steps below to install and set up the project:

1. Clone the repository: `git clone https://github.com/buraksyagmur/golang-forum.git`
2. Change to the project directory: `cd golang-forum`
3. Install the project dependencies: `go get -d ./...`
4. Build the project: `go build`
5. Start the application: `./golang-forum`

## Usage

Once the application is up and running, you can access the Golang Forum by opening your web browser and navigating to `http://localhost:8080`. From there, you can register as a new user, log in, and start participating in discussions.

