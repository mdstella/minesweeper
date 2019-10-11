import React, { Component } from 'react';

class Cell extends Component {
    pickCell = () => {
        console.log("Row: " + this.props.rowId)
        console.log("Column: " + this.props.colId)
    }
    render() {
        let color = "white"
        if (this.props.value === "") {
            color = "grey"
        }
        return (
            <td onClick={this.pickCell} bgcolor={color} height="20px" width="20px" align="center" id={this.props.colId}> {this.props.value} </td>
        );
    }
}
export default Cell;