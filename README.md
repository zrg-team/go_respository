# About (In progress)

Project for building RESTful APIs and microservices using Golang, [mux](https://github.com/gorilla/mux) and Mariadb

Inspiring by [node_express_respository](https://github.com/zrg-team/node_express_respository)

We focus on:
 + Easy to start
 + Reduce copied code
 + Readable code
 + Simple to implement new api for a model like: search, create, update
 + etc

## Features

+ Mux + Mariadb
+ Uses [Gorm](https://github.com/jinzhu/gorm) as ORM
+ [Docker](https://www.docker.com/) support
+ Load environment variables from env file
+ Authentication and Authorization with [JWT](https://jwt.io/)
+ Support repositories
+ Default search feature

## Requirements

+ [Golang](https://golang.org/)
+ [Docker]((https://www.docker.com/))

## Structure motivation

#### 1. We separate routes into many files based on main model. ex:

  user.js

    - POST / => /user/
    - POST /login => /user/login
    - POST /verify => /user/verify
    - GET /  => /user/
    - GET /version => /user/version

    ...

#### 3. Repositories is used to abstract the data layer, making our application more flexible to maintain.

Methods:

- create
- update
- findById
- find
- paginate

#### 4. Controllers handles requests using repository layer 

#### 5. Criteria: Default criteria supported fast search on a repository, push criteria to your request using below code


http://api.local/users?search=name:John%20Doe

or 

http://api.local/users?search=name:John&searchFields=name:like

or

http://api.local/users?search=email:john@gmail.com;name=john&searchFields=email:=;name:like

By default DefaultCriteria makes its queries using the AND comparison operator for each query parameter.

http://api.local/users?search=email:john@gmail.com;age=18&searchFields=email:=;age:>=

The above example will execute the following query:

```
SELECT * FROM users WHERE age >= 18 AND email = 'john@gmail.com';
```

In order for it to query using the OR, pass the searchJoin parameter as shown below:

http://api.local/users?search=email:john@gmail.com;age=18&searchFields=email:=;age:>=&searchJoin=or

Filtering fields

http://api.local/users?fields=id;name

Sorting the results

http://api.local/users?order=name:desc

#### 6.Domain logic should put to services folders
