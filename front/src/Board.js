import React, { Component } from 'react';

import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import FormGroup from '@material-ui/core/FormGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Switch from '@material-ui/core/Switch';

import Row from './Row';

class Board extends Component {
    constructor(props) {
        super(props);
        this.state = { board: [], addFlag: false, endedGame: false };
    }

    componentDidMount() {
        this.setState({
            board: this.props.board
        });
    }

    handleClose = () => {
        this.setState({
            endedGame: false
        });
    };

    // this is callback that will be invoked when a cell is picked by the user. It will be invoked after
    // calling the /game/{gameId} endpoint in the BE, and will override the properties with the new board to
    // re render it
    cellPickedCallback = (cellPicked) => {
        if (cellPicked.error !== undefined) {
            this.setState({ board: [] })
            this.props.appCallback()
        }
        else {
            this.setState({ board: cellPicked.board })
            if (cellPicked.endedGame && !cellPicked.won) {
                this.setState({ endedGame: true, won: false, board: cellPicked.board })
            } else if (cellPicked.endedGame && cellPicked.won) {
                this.setState({ endedGame: true, won: true, board: cellPicked.board })
            }
        }
    }

    handleActiveFlagChange = (event) => {
        let addFlag = this.state.addFlag
        addFlag = !addFlag
        this.setState({
            addFlag: addFlag
        });
    }

    render() {
        let gameId = this.props.gameId
        let callback = this.cellPickedCallback.bind(this)
        let addFlag = this.state.addFlag

        if (this.state.board !== [] && gameId !== '') {
            // iterates the rows and return the Row component that will render each row
            let rows = this.state.board.map(function (columns, i) {
                return (
                    <Row boardCallback={callback} key={i} rowId={i} columns={columns} gameId={gameId} addFlag={addFlag} />
                );
            });
            return (
                <div>
                    <FormGroup row>
                        <FormControlLabel
                            control={
                                <Switch
                                    checked={this.state.addFlag}
                                    onChange={this.handleActiveFlagChange}
                                    value="checkedB"
                                    color="primary"
                                />
                            }
                            label="Add flag"
                        />
                    </FormGroup>

                    <table border="1">
                        <tbody>
                            {rows}
                        </tbody>
                    </table>

                    <Dialog
                        open={this.state.endedGame}
                        onClose={this.handleClose}
                        aria-labelledby="alert-dialog-title"
                        aria-describedby="alert-dialog-description"
                    >
                        <DialogTitle id="alert-dialog-title">{"GAME OVER!!!"}</DialogTitle>
                        <DialogContent>
                            <DialogContentText id="alert-dialog-description">
                                {this.state.won ? "Cograts!!! You won!!! Want to play again?" : "You lost!!!! Want to play again?"}
                            </DialogContentText>
                        </DialogContent>
                        <DialogActions>
                            <Button onClick={this.handleClose} color="primary">
                                No
          </Button>
                            <Button onClick={this.props.newGameCallback} color="primary" autoFocus>
                                Yes
          </Button>
                        </DialogActions>
                    </Dialog>
                </div >
            );
        } else {
            return null
        }
    }
}
export default Board;