import React, { Component } from 'react';
import Row from './Row'

class Board extends Component {
    render() {
        var board = [
            ["", "2", "1", "0", "", "", "", "", ""],
            ["", "*", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", ""]
        ]
        var rows = board.map(function (columns, i) {
            return (
                <Row key={i} rowId={i} columns={columns} />
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
    }
}
export default Board;