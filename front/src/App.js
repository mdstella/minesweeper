import React, { Component } from 'react';
import Button from '@material-ui/core/Button';
import Board from './Board'


class App extends Component {
  render() {
    return (
      <div>
        <Button variant="contained" color="primary">
          New Game
        </Button>

        <Board />
      </div>
    );
  }
}
export default App;