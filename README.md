gonote
======

A simple note app written in goLang with revel.

Here is a postman link to the APIs - https://www.getpostman.com/collections/3a62e2584909e7e3f9b5

After cloning the repo into your local machine, database needs to be configured. The database configurations can be seen in the file "gonote/app/database/connection.go" in the function NewConnection.

To view the web UI -<br/>
Open a browser and ping the url (localhost:9000/ if you have not changed the port).

Features -<br/>
CRUD API for notes<br/>
Web user interface for managing notes<br/>

Why did I do this?<br/>
I wanted to check if I could use Go, Revel, Gorp and make some serious backend with it.
