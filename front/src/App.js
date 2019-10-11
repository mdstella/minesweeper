import React, { Component } from 'react';
import Button from '@material-ui/core/Button';
import Board from './Board'


class App extends Component {
  newGame = () => {
    fetch("http://localhost:5000/minesweeper/v1/game", {
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
        this.clearBoard()
        this.setState({ gameBoard: data.board, gameId: data.gameId })
      })
      .catch((error) => {
        console.log(error, "catch the hoop")
      })
  }

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

    return (
      <div>
        <Button variant="contained" color="primary" onClick={this.newGame}>
          New Game
        </Button>

        <div style={divStyle} />



        {gameBoard.length > 0 && gameId !== '' &&
          <Board appCallback={callback} board={gameBoard} gameId={gameId} />
        }
      </div >
    );
  }
}
export default App;