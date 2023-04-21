# Project Overview:


The Hartley-Fabric platform is a full-stack property rental and leasing platform based on blockchain technology. The platform allows users to create, view, and update tenancy agreements automatically. The platform is designed to have a beautiful and intuitive user interface, with a mining status display that shows transaction IDs like a digital clock. The platform will also have user authentication and authorization features, including the ability to register and login and sign out.



# Folder and File Structure:


We will use a simple and organized file structure to keep our codebase organized and maintainable. Here is an example of the proposed file structure for the Hartley-Fabric project:

├── backend
│   ├── cmd
│   │   └── main.go
│   ├── config
│   │   └── config.go
│   ├── controller
│   │   ├── auth_controller.go
│   │   └── tenancy_controller.go
│   ├── database
│   │   ├── migrations
│   │   └── models
│   ├── middleware
│   │   ├── auth_middleware.go
│   │   └── logging_middleware.go
│   ├── repository
│   ├── router
│   │   └── router.go
│   ├── service
│   └── main.go
├── frontend
│   ├── public
│   │   ├── index.html
│   │   ├── css
│   │   │   ├── bootstrap.min.css
│   │   │   └── style.css
│   │   ├── js
│   │   │   ├── bootstrap.bundle.min.js
│   │   │   ├── jquery-3.6.0.min.js
│   │   │   └── main.js
│   └── src
│       ├── components
│       │   ├── auth
│       │   ├── dashboard
│       │   ├── home
│       │   └── tenancy
│       ├── config
│       ├── pages
│       ├── services
│       └── utils
├── mining
│   ├── mining.go
│   ├── block.go
│   └── transaction.go
├── README.md
├── LICENSE
├── go.mod
├── go.sum
├── .gitignore
├── .env
├── .env.example
├── docker-compose.yml
├── Dockerfile
├── .travis.yml
└── .editorconfig



backend: contains the backend code for the Hartley-Fabric platform, including the main application file, configuration files, controller logic, database migrations, middleware functions, repository code, service code, and router code.
frontend: contains the frontend code for the Hartley-Fabric platform, including public assets, component code, page code, service code, and utility code.


mining: contains the mining code for the Hartley-Fabric platform, including block code, transaction code, and mining code.


README.md: contains information about the Hartley-Fabric platform and how to use it.


LICENSE: contains the license for the Hartley-Fabric platform.


go.mod and go.sum: contain Go module information for the Hartley-Fabric platform.


.gitignore: specifies which files and directories to ignore when pushing to Git.

.env: contains environment variables used by the Hartley-Fabric platform.


.env.example: contains an example environment file with placeholder values.


docker-compose.yml: contains Docker Compose configuration for the Hartley-Fabric platform.


Dockerfile: contains Docker configuration for the Hartley-Fabric platform.


.travis.yml: contains Travis CI configuration for the Hartley-Fabric platform.


.editorconfig: contains EditorConfig configuration for the Hartley-Fabric platform.


# Backend Overview:


The backend of the Hartley-Fabric platform will be responsible for handling user requests, authenticating and authorizing users, creating, viewing, and updating tenancy agreements, and interacting with the blockchain system to store and retrieve data. The backend will use Go programming language and will be divided into different packages to make the codebase more modular and maintainable.



# Frontend Overview:


The frontend of the Hartley-Fabric platform will be responsible for providing a beautiful and intuitive user interface for users to interact with the platform. The frontend will use Bootstrap to create a responsive and mobile-friendly design and will be built using JavaScript, HTML, and CSS. The frontend will also use AJAX to interact with the backend and retrieve data from the blockchain system.



# Blockchain Overview:


The blockchain system of the Hartley-Fabric platform will be implemented using Go programming language and will be responsible for storing and retrieving tenancy agreement data in a decentralized and secure manner. The blockchain system will use a proof-of-work consensus mechanism and will be able to generate new blocks and transactions using mining algorithms. The blockchain system will also provide transaction validation and verification to ensure the integrity of data stored on the blockchain.