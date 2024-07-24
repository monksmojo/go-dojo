<div align=center> LEAGUE MATRIX ASSIGNMENT </div>

##### developer: Udit Makkar

### documentation content content 
1. [Set up the Assignment on Your Machine](#set-up-the-assignment-on-your-machine)
2. [swagger-url](#swagger-url)
3. [curl-request](#swagger-url)
4. [Objective of the Assignment](#objective-of-the-assignment)
5. [external Packages the App Is Using](#packages-the-app-is-using)
6. [external Packages the App Is Using](#project-structures)

#### how to set up the assignment on your machine
1. download / clone the file
    - [download link from the drive](https://drive.google.com/drive/folders/1gcLidTJR_GBWx5a0VwL2zmAYOhBU_u65?usp=sharing)
    - [clone link from github]()
2. open the project in command prompt/terminal
3. the app is using `.makefile` as the automation tool that you can use to run and build the application
4. to run it

    > linux-os `make spin-up-linux`

    > windows-os type `make spin-up-windows`

    > mac-os type `make spin-up-macos`
5. what will the above commands do

        - make sure you are at the root of the project
        - it will remove any previous build in /bin directory and run go mod tdy command
        - it will run all the test cases
        - only if the test cases pass it will build a new executable 
        - run the executable
        - check -makefile for more info
6. if step 4 fails / or you are using os not mentioned above you can use alternative command `go run ./cmd/app/main.go`
5. the application is running on port 8080 but you can change it in .env file


#### swagger-url
http://localhost:8080/swagger/index.html


### curl-request
`curl -X POST http://localhost:8080/<endpoint> -H "Content-Type: multipart/form-data" -F "file=@/database/matrix.csv"`

#### objective of the assignment
read a csv file as a request and perform the following operation where each operation is bind to an endpoint
ps: .csv file can be found in the database directory
1. 	POST request ("/echo", handler.MatrixHandler.Echo)
2.	POST request ("/invert", handler.MatrixHandler.Invert)
3.	POST request ("/flatten", handler.MatrixHandler.Flatten)
5.	POST request ("/sum", handler.MatrixHandler.Sum)
6.	POST request ("/multiply", handler.MatrixHandler.Multiply) 


1. Echo (given)
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

#### external packages the app is using
1. github.com/gin-gonic/gin v1.10.0
2. github.com/stretchr/testify v1.9.0
3. github.com/swaggo/files v1.0.1
4. github.com/swaggo/gin-swagger v1.6.0
5. github.com/swaggo/swag v1.16.3


### project structure
![image](https://gist.github.com/user-attachments/assets/093eb0c2-1098-4181-9228-94fd7b4f9e0d)

- api # contains api response related functions and swagger api
- bin  # contains app executable run by makefile
- main.go # set environment variables and spin the server by calling NewServer() in server package
- database # for now contain matrix.csv file used in curl || later can contain multiple data store
- docs # swagger-docs
- handler # contain all the handlers of the domains
- internal # use partial domain-driven-design
- domain -matrix matrix_handler.go # matrix specific handlers
- domain -matrix matrix_repository.go # empty for now
- domain -matrix matrix_service.go # all the matrix related logic called by handler
- domain -matrix matrix_service_test.go # test cases verifying the logics 
- Makefile - automated tool to run the 
- models # contain models and entity used by the add
- errors.go # all the custom error used by apis
- router # contain router the app is using with call to fetch all the handlers from the handler package
- server # spin the new net/http package with specific router 