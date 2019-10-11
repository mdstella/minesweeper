import React, { Component } from 'react';
import Cell from './Cell'

class Row extends Component {

    render() {
        let rowId = this.props.rowId

        let cells = this.props.columns.map(function (cellVal, colId) {
            return (
                <Cell key={colId} colId={colId} rowId={rowId} value={cellVal} />
            );
        });
        return (
            <tr>{cells}</tr>
        );
    }
}
export default Row;