import React, { Component } from 'react';
import Cell from './Cell'

class Row extends Component {
    constructor(props) {
        super(props);
        this.state = { columns: [] };
    }

    componentDidMount() {
        this.setState({
            columns: this.props.columns,
        });
    }

    render() {
        let gameId = this.props.gameId
        let rowId = this.props.rowId
        let callback = this.props.boardCallback

        if (gameId !== '') {
            let cells = this.state.columns.map(function (cellVal, colId) {
                return (
                    <Cell boardCallback={callback} key={colId} colId={colId} rowId={rowId} value={cellVal} gameId={gameId} />
                );
            });
            return (
                <tr>{cells}</tr>
            );
        }
    }
}
export default Row;