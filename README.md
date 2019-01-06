# Introduction to Go

This is a project for simulating the data flow between school and the database of education ministry. It is just a sample project for learning `Go` language connected by `MySQL` database.

First you can send the student info such image, name, and address to your school server (via **Postman**). Then your school will passing these info to education ministry server. Then education ministry will confirm this student status to school, and the school will inform you your status.

## Prerequisite
I assume you already have `mysql-server` installed on your PC. Then you need to make new tables for the project like this command.

School database :
```
create table school(
    id int not null auto_increment,
    ministry_id int,
    name varchar(50),
    address varchar(100),
    status int,
    created timestamp default current_timestamp,
    updated timestamp default current_timestamp on update current_timestamp,
    primary key (id)
);
```
Education ministry database :
```
create table edu_ministry(
    id int not null auto_increment,
    name varchar(50),
    address varchar(100),
    status int,
    created timestamp default current_timestamp,
    updated timestamp default current_timestamp on update current_timestamp,
    primary key (id)
);
```


## Implementation
Run the school server :

```
cd school
go build && ./school
```

Run the education ministry server (from another terminal):
```
cd ministry
go build && ministry
```

Open your browser, and open url at `http://localhost:9000`. You can do the document uploading from here.

Or you can test it using **Postman**. Send data image, name and address to school server like this format via **Postman**.  

```json
URL: http:localhost:9000/upload  
Method: POST  
Content-Type multipart/form-data  
  
Request:  
image_file: multipart.File
student_name: string
student_address: string

Response:
{
    "code": 1000,
    "data": {
        "student_name": "Student Name",
        "id": 77
    },
    "message": "Success"
}
```
You will get your school ID and your status in education ministry status. Then the school server will hit the education ministry server using this API  
```json
URL: http:localhost:8000/student  
Method: POST  
Content-Type multipart/form-data  
  
Request:  
imageFile: multipart.File
studentName: string
studentAddress: string

Response:
{
    "code": 1000,
    "data": {
        "student_name": "Student Name",
        "id": 1457
    },
    "message": "Success"
}
```
You will get ministry ID. This ID is use to communicate between school and education ministry.