# minesweeper-API
API test

## Changelog/decisions
1. Breakdown and task creation for the complete project
    - Created githb issues from T1 to T12 based on the priority (this might change during the development process, but at least will provide some guide)
    - Also as a branching strategy will create, if possible, 1 branch from master for each issue under the name T{X}. In a regular process should go under code reviwe before merging it into master, here I will merge it as soon as I have some working code to show progress
2. Doing T1 issue on github tasks definition. Adding the main structure of the BE. Creating a test skeleton POST endpoint, adding service layer (with support for generating mocks - important for generating test cases). 
    - NOTE: this version was generated with golang 1.12 (supporting go modules). To obtain the needed dependencies you can run `go build` under *core* folder, and it should download the dependencies defined in `go.mod` file. 
    - NOTE 2: To run it just execute `go run main.go` under *core* folder and it will startup a server running on port 8080. To test it's workig just send the following CURL:

    ```
    curl -X POST \
        http://localhost:8080/skeleton \
        -H 'Accept: */*' \
        -H 'Content-Type: application/json' \
        -d '{
	            "param":"this is the parameter you sent"
        }'
    ```