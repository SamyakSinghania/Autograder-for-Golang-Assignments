# Autograder
This `autograder` evaluates all the `go assignments` on the basis of running `go test` commands on the bash for all `_test.go` files.
## Setup
To use this Autograder, clone the repository and all the submissions in the Assignment folder.
Then run the following command:
```bash
go run main.go
```
## Results
The scores of each individual will be stored in the `results.txt` file.
## Requirements
The assignment must have a `_test.go` file.

## Approach
The `autograder` uses the `go test` command to run the tests on the bash for all the `_test.go` files.\
Using `os/exec package`, go test commands are run on the bash for each submission. 

A `map` is created with its `key` as `username` and `value` as the `number of tasks` successfully completed.

If the output received on `stdout` starts with `PASS`, the values for that username is `incremented` on the map. All the data of this map is then copied to a user defined struct array so that it can be sorted in a decreasing order.

Then this key value pair is written on a file named `results.txt`

If any error happens while running the `go test` command, the error is printed on `stdout`

This `autograder` can be effectively used to evaluate the `PCLUB` summer project submissions. 

Some demo submissions are provided in the `Assignments` folder along with the results in the `results.txt` file.