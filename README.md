# minesweeper-API
API test

## Changelog/decisions 
**NOTE:** after having both front and back on a web version will start adding test cases
1. Breakdown and task creation for the complete project
    - Created githb issues from T1 to T12 based on the priority (this might change during the development process, but at least will provide some guide)
    - Also as a branching strategy will create, if possible, 1 branch from master for each issue under the name T{X}. In a regular process should go under code reviwe before merging it into master, here I will merge it as soon as I have some working code to show progress
2. Doing T1 issue on github tasks definition. Adding the main structure of the BE. Creating a test skeleton POST endpoint, adding service layer (with support for generating mocks - important for generating test cases). **COMMIT: 2711ebe0ad7c0c527051dcbea47d03aa7addadad**
    - NOTE: this version was generated with golang 1.12 (supporting go modules). To obtain the needed dependencies you can run `go build` under **core** folder, and it should download the dependencies defined in `go.mod` file. 
    - NOTE 2: To run it just execute `go run main.go` under **core** folder and it will startup a server running on port 8080. To test it's workig just send the following CURL:

    ```
    curl -X POST \
        http://localhost:8080/skeleton \
        -H 'Accept: */*' \
        -H 'Content-Type: application/json' \
        -d '{
	            "param":"this is the parameter you sent"
        }'
    ```
3. Creating new game endpoint (**COMMIT: 6003e20863c2ec5ed195e78a5448a0249efde145**). We need to do the following CURL to create the new board:
    ```
    curl -X POST \
        http://localhost:8080/minesweeper/v1/game \
        -H 'Accept: */*' \
        -H 'Content-Type: application/json' \
        -d '{}'
    ```

    Response

    ```
       {
           "gameId":"Q9ZnJqGwcGVJ6ZeWYPx49A",
           "board":[
               ["0","0","0","0","0","0","1","2","2"],
               ["0","0","0","0","0","0","1","*","*"],
               ["0","0","1","1","1","0","1","2","2"],
               ["0","0","2","*","2","0","0","0","0"],
               ["0","1","3","*","3","1","0","0","0"],
               ["0","1","*","3","*","2","1","0","0"],
               ["0","1","1","3","3","*","1","1","1"],
               ["0","0","0","1","*","3","2","1","*"],
               ["0","0","0","1","2","*","1","1","1"]
            ]
       }
 
    ```

    - It will retrieve 2 fields:
        - **gameId**: String that will be used on the next stage to identify which game we are playing
        - **board**: The board showing where are the bombs. This step is just for showing how the board is generated. In next stages we will keep this generated board in memory and retrieve an empty board with all the values hidden that will be the one that the Client should render.
4. Creating endpoint to pick and reveal a cell (**COMMIT: fc8b37b**). 
    - Adding **gorilla/mux** library to be able to route and dispatch the endpoints with the ability to extracts parameters from the URL path. 
    - Adding memory cache LRU to store the games in memory to be able to play N different games at the same time (hardcoded as 10). In the cache we are keeping 2 boards by game. One has the complete solution of the game, the other has the same solution that the user is seeing.
    - Adding error handling (so far only HTTP status code 400 and 500)
    - To play the game using the API:
        1. Create a new game (The response is different than the previous point, now we hidden all the values from the board)
            REQUEST
            ```
            curl -X POST \
                http://localhost:8080/minesweeper/v1/game \
                -H 'Accept: */*' \
                -H 'Content-Type: application/json' \
                -d '{}'
            ```
            RESPONSE
            ```
            {
                "gameId": "zF8JeVqn3tj4Q3KBYP2SMR",
                "board": [
                    [ "", "", "", "", "", "", "", "", "" ], 
                    [ "", "", "", "", "", "", "", "", "" ], 
                    [ "", "", "", "", "", "", "", "", "" ], 
                    [ "", "", "", "", "", "", "", "", "" ], 
                    [ "", "", "", "", "", "", "", "", "" ],
                    [ "", "", "", "", "", "", "", "", "" ], 
                    [ "", "", "", "", "", "", "", "", "" ], 
                    [ "", "", "", "", "", "", "", "", "" ], 
                    [ "", "", "", "", "", "", "", "", "" ]
                ]
            }
            ```
        2. Invoke the endpoint to pick the cell. For this you have to add the gameId obtained when creating the new game into the path. Consider that is a 9X9 board, and the allowed indexes values are from 0 to 8.
            REQUEST
            ```
            curl -X POST \
                http://localhost:8080/minesweeper/v1/game/zF8JeVqn3tj4Q3KBYP2SMR \
                -H 'Accept: */*' \
                -H 'Content-Type: application/json' \
                -d '{
	                "row":1,
	                "column": 1
                }'    
            ```
            RESPONSE (will retrieve the board with the cell revealed)
            ```
            {
                "gameId":"zF8JeVqn3tj4Q3KBYP2SMR",
                "endedGame":true,
                "won":false,
                "board":[
                    ["","","","","","","","",""],
                    ["","*","","","","","","",""],
                    ["","","","","","","","",""],
                    ["","","","","","","","",""],
                    ["","","","","","","","",""],
                    ["","","","","","","","",""],
                    ["","","","","","","","",""],
                    ["","","","","","","","",""],
                    ["","","","","","","","",""]
                ]
            }
            ```
5. Hosting the backend on the web (**COMMIT: deec7737b26a33a02a50d71c4c6e46adc08f1fc6**). The host will be sent by email. Now the API is on the web, you can use it by runninng the following CURL's    
    1. Start a new game
        REQUEST
        ```
        curl -X POST \
            {{host}}/minesweeper/v1/game \
            -H 'Accept: */*' \
            -H 'Content-Type: application/json' \
            -d '{}'
        ```
        RESPONSE
        ```
        {
            "gameId": "zF8JeVqn3tj4Q3KBYP2SMR",
            "board": [
                [ "", "", "", "", "", "", "", "", "" ], 
                [ "", "", "", "", "", "", "", "", "" ], 
                [ "", "", "", "", "", "", "", "", "" ], 
                [ "", "", "", "", "", "", "", "", "" ], 
                [ "", "", "", "", "", "", "", "", "" ],
                [ "", "", "", "", "", "", "", "", "" ], 
                [ "", "", "", "", "", "", "", "", "" ], 
                [ "", "", "", "", "", "", "", "", "" ], 
                [ "", "", "", "", "", "", "", "", "" ]
            ]
        }
        ```
    2. Invoke the endpoint to pick the cell. For this you have to add the gameId obtained when creating the new game into the path. Consider that is a 9X9 board, and the allowed indexes values are from 0 to 8
        REQUEST
        ```
        curl -X POST \
            {{host}}/minesweeper/v1/game/zF8JeVqn3tj4Q3KBYP2SMR \
            -H 'Accept: */*' \
            -H 'Content-Type: application/json' \
            -d '{
                "row":1,
                "column": 1
            }'    
        ```
        RESPONSE (will retrieve the board with the cell revealed)
        ```
        {
            "gameId":"zF8JeVqn3tj4Q3KBYP2SMR",
            "endedGame":true,
            "won":false,
            "board":[
                ["","2","1","0","","","","",""],
                ["","*","","","","","","",""],
                ["","","","","","","","",""],
                ["","","","","","","","",""],
                ["","","","","","","","",""],
                ["","","","","","","","",""],
                ["","","","","","","","",""],
                ["","","","","","","","",""],
                ["","","","","","","","",""]
            ]
        }
        ```
6. Development of a React application that invokes the core API to render the game and allows the user to play it.
    - So far is only available locally. 
    - The game was splited on 4 different components:
        - App -> it renders the NEW GAME button and the Board component
        - Board -> it iterates the rows from the board rendering each Row component
        - Row -> it iterates every column in the row rendering each Cell component
        - Cell -> shows the value of each cell
    - Node: 8.9.4, npm: 5.6.0
    - To run it run `npm install` under **front** folder. Then run `npm start`
    - Change the localhost pointing to the hosted core API, or start also the API locally. To do this follow the following steps:
        - Use Golang 1.12.10
        - Under core folder run `go buil`
        - Run `go run main.go` and server should start
