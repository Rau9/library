# Library
Library is an application to manage an online book catalog. 
The first thing we will find will be a login screen, in which we will have to add the username and password. 

In this first version, to make it longer, the default user will be "admin" and the password "admin". 

This user will be checked by the front to allow us access. In future versions we should have a table with the users, 
a registration screen and save the registered users in that table, in addition, in the first login it would persist 
in the localStorage to maintain the session. 

Once past the login screen, we will find a screen with a menu where we can go between the different options (Consult, Create, Modify and DELETE):

- Modify screen: Here we will find the entire catalog with the mandatory data of each book, we do not have an option to search for a specific one by time, but for future versions I would add a button to search for a single book, or an author to see all his books ...
- Create screen: Here we can create a new book, we will have to add the required fields marked with "*" and the optional ones if we have that information. By clicking on the create button, it would be created, it would persist in the table and the query appears on the screen.
- Modify screen: Here we can modify the book based on its ID, we can modify all the fields except the author and its associated fields since if a book is written by an author we do not see the sense to modify it.
- DELETE screen: Here we can delete any book created, simply by copying its ID from the consult screen, pasting it in the ID field and clicking on delete.

It is a simple application but in it I wanted to touch all the parts of the frotn and the back so that they know how my way of programming is and the knowledge that I have. I would have liked to finish all the tests on the back, to have made a more complete and good login and to have created the categories in a related table as I do with author. But it has been quite long and I did not want to spend much longer and I think that with this I will be able to explain well my knowledge in development.

## Design Considerations
- The id is the unique identifier for each book.
- The application allows you to modify the isbn because the unique identifier is ID.
- The application logs in a structured way in JSON format for better diagnosis and troubleshooting in a centralized logging system.
- The application has a healthcheck endpoint that validates the status of the backend dependencies. If it returns 500, the application is not ready to receive traffic.
- For simplicity we consider a book can only have 1 author
- For simplicity we consider a book can only have 1 category
- For simplicity, the categories is a combobox but only one can be selected and it is done from the front.
- It not allows a search a one register only in the consult view because I did not want to lengthen the delivery time but I should have.

## API layout

| Resource | Rest Action | HTTP verb | Endpoint     | Status code | Req body/params/headers    | Res body/headers      | Comments                                                              |
|----------|-------------|-----------|--------------|-------------|----------------------------|-----------------------|-----------------------------------------------------------------------|
| Books    |             |           |              |             |                            |                       |                                                                       |
|          | Create      | POST      | /books       | 204         | {libros, autor, categoría} | Location /books/:uuid |                                                                       |
|          |             |           |              | 400         | {libros, categoría}        | {error}               | i.e:Missing mandatory information                                     |
|          |             |           |              | 422         | {libros, autor, categoría} | {error}               | i.e:Invalid isbn                                                      |
|          |             |           |              | 409         | {libros, autor, categoría} | {error}               | i.e:Duplicated ISBN                                                   |
|          |             |           |              | 500         | {libros, autor, categoría} | {error}               |                                                                       |
|          | Read        | GET       | /books/:uuid | 200         |                            | {book}                |                                                                       |
|          |             |           |              | 404         |                            | {error}               | i.e:uuid does not exist                                               |
|          |             |           |              | 422         |                            | {error}               | i.e: Invalid uuid                                                     |
|          |             |           |              | 500         |                            | {error}               |                                                                       |
|          | Update      | PUT       | /books/:uuid | 204         | {libros, autor, categoría} | Location /books/:uuid |                                                                       |
|          |             |           |              | 404         |                            | {error}               | i.e:uuid does not exist                                               |
|          |             |           |              | 400         | {libros, categoría}        | {error}               | i.e:Missing mandatory information                                     |
|          |             |           |              | 422         | {libros, autor, categoría} | {error}               | i.e:Invalid isbn or existing isbn                                     |
|          |             |           |              | 500         | {libros, autor, categoría} | {error}               |                                                                       |
|          | Delete      | DELETE    | /books/:uuid | 200         |                            | {success}             |                                                                       |
|          |             |           |              | 404         |                            | {error}               | i.e:uuid does not exist                                               |
|          |             |           |              | 422         |                            | {error}               | i.e:Invalid uuid                                                      |
|          |             |           |              | 500         |                            | {error}               |                                                                       |
|          | List        | GET       | /books       | 200         | {isbn}                     | [{book}]              | ISBN is optional                                                      |
|          |             |           |              | 500         |                            | {error}               |                                                                       |
|          | HealthCheck | GET       | /healthz     | 200         |                            | {healthCheck}         |                                                                       |
|          |             |           |              | 500         |                            | {healthCheck}         |                                                                       |
|          |             |           |              |             |                            |                       |                                                                       |


## Local development
To run the project locally, first create your own secrets using `.env.sample` as a template:

```shell
cp .env.sample .env
```

And run the application:

```shell
source .env
make run
```

To run the tests:

```shell
make test
```

## Deployment to production considerations
- Authors and books relationship should be one to many, a book can have multiple authors
- Categories should be added to the list of models and handled by the backend. It should no longer be hardcoded in the frontend
- Categories and books relationship should be one to many, a book can have multiple categories
- List endpoint should allow searches by fields, i.e, by isbn, name, author, etc
- List endpoint should be paginated and return a pagination object in the response
- Migration task should seed the database with default users
- Users should no longer be hardcoded in the fronted
  - Option 1: Let the backend handle it and add an users model and expose its endpoints
  - Option 2: Use an external users management backend API, like AWS cognito
- Application must validate format of ISBN
- Application must validate format of dates 
- Add metrics and traces for real-time monitoring.
- The frontend should be deployable independently on a cdn like AWS Cloudfront or Vercel
- Add integration tests suite to the books API endpoints
- Add CI/CD pipeline to build, package and release the app
- Package backend app in an ECS task definition or a Kubernetes helm chart
