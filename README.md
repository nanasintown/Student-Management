# A simple student management application

This is a terminal-based application, where it asks for a file name contains information about a specific student.
After reading the information from the file, the app will sere several options for user to interact.

For example:

> 1): Get name
> 2): Get ID
> 3): Get status
> 4): Get job
> 5): Get Salary
> 6): Print statistic
> 7):Exit program

- In option 5, the salary will be calculated based on the input number of hours and the base salary from the file.

---

## Run the application

To run the app, make sure you have installed Go in your local computer. The following commands can be applied assume you have Go run fine.

```
cd student
go run main.go management.go
```

### This app is not fully implemented to handle all possible error

For example, an invalid file will not be handled at the moment, the application assume the input file is valid.
