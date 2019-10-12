import React, { Component } from 'react';
import Row from './Row';

class Board extends Component {
    constructor(props) {
        super(props);
        this.state = { board: [] };
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

    render() {
        let gameId = this.props.gameId
        let callback = this.cellPickedCallback.bind(this)

        if (this.state.board !== [] && gameId !== '') {
            // iterates the rows and return the Row component that will render each row
            let rows = this.state.board.map(function (columns, i) {
                return (
                    <Row boardCallback={callback} key={i} rowId={i} columns={columns} gameId={gameId} />
                );
            });
            return (
                <div>
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