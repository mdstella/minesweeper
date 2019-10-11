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

    // this will update the state if there are new props on the component
    static getDerivedStateFromProps(props, state) {
        if (props.columns !== state.columns) {
            return {
                columns: props.columns
            };
        }
        // Return null to indicate no change to state.
        return null;
    }

    render() {
        let gameId = this.props.gameId
        let rowId = this.props.rowId
        let callback = this.props.boardCallback

        if (gameId !== '') {
            // Iterates all the columns and with the value it will render each Cell
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