## Introduction
There are two kind of testing:<br/>
1. touchy testing<br/>
2. mock testing<br/>
let me explain by an example imagine you have a project that communicates with a database, and you want to test the <br/>
part of the code where you write to and read from the database in touchy testing you need to run a database and literally<br/>
write to and read from that database to see if it works well, so you do need the database and all the operations are real,>br/>
but some people say that whether the database is up and running, or it's down is non of their business, and they want to<br/>
test their code not the database, so they use mocking even when testing the database.<br/>
Now I want to explain more precisely about storages.<br/>
1. redis: To mock redis and test it [miniredis](https://github.com/alicebob/miniredis) is used