import React, { Component } from 'react';
import Button from '@material-ui/core/Button';
import Board from './Board'


class App extends Component {
  // this functions is invoked when the user clicks on the "NEW GAME" button
  // we invoke the endpoint /game to obtain a new board an gameId
  newGame = () => {
    fetch("http://localhost:8000/minesweeper/v1/game", {
      method: 'post',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      },
      body: {}
    })
      .then((resp) => {
        return resp.json()
      })
      .then((data) => {
        // the API response. First we clear the board (in case is not a the first game, to remove the old board), then we set the state
        // of the App component with the board and gameId
        this.clearBoard()
        this.setState({ gameBoard: data.board, gameId: data.gameId })
      })
      .catch((error) => {
        console.log(error, "catch the hoop")
      })
  }

  // this clears the board (set the initial state of App component, empty board and gameId)
  clearBoard = () => {
    this.setState({
      gameBoard: [],
      gameId: ''
    })
  }

  render() {
    const divStyle = {
      marginBottom: '40px',
    };

    let gameBoard = []
    let gameId = ""
    if (this.state !== null) {
      if (this.state.gameBoard !== null) {
        gameBoard = this.state.gameBoard
      }
      if (this.state.gameId !== null) {
        gameId = this.state.gameId
      }
    }

    let callback = this.clearBoard.bind(this)
    let newGame = this.newGame.bind(this)

    return (
      <div>
        <Button variant="contained" color="primary" onClick={this.newGame}>
          New Game
        </Button>

        <div style={divStyle} />



        {gameBoard.length > 0 && gameId !== '' &&
          <Board newGameCallback={newGame} appCallback={callback} board={gameBoard} gameId={gameId} />
        }
      </div >
    );
  }
}
export default App;