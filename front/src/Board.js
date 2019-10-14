import React, { Component } from 'react';

import FormGroup from '@material-ui/core/FormGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Switch from '@material-ui/core/Switch';

import Row from './Row';

class Board extends Component {
    constructor(props) {
        super(props);
        this.state = { board: [], addFlag: false };
    }

    componentDidMount() {
        this.setState({
            board: this.props.board
        });
    }

    // this is callback that will be invoked when a cell is picked by the user. It will be invoked after
    // calling the /game/{gameId} endpoint in the BE, and will override the properties with the new board to
    // re render it
    cellPickedCallback = (cellPicked) => {
        if (cellPicked.error !== undefined) {
            alert("GAME OVER, PLAY AGAIN?")
            this.setState({ board: [] })
            this.props.appCallback()
        } else if (cellPicked.endedGame && cellPicked.won) {
            alert("CONGRATS!!! YOU WON!!! PLAY AGAIN?")
            this.setState({ board: [] })
            this.props.appCallback()
        } else {
            this.setState({ board: cellPicked.board })
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
                </div >
            );
        } else {
            return null
        }
    }
}
export default Board;