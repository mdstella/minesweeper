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

    cellPickedCallback = (cellPicked) => {
        if (cellPicked.error !== undefined) {
            alert("GAME OVER, PLAY AGAIN?")
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