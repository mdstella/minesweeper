import React, { Component } from 'react';

class Cell extends Component {
    constructor(props) {
        super(props);
        this.state = { value: '' };
    }

    componentDidMount() {
        this.setState({
            value: this.props.value,
        });
    }

    pickCell = () => {
        fetch("http://localhost:5000/minesweeper/v1/game/" + this.props.gameId, {
            method: 'post',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            },
            body: JSON.stringify({
                "row": this.props.rowId,
                "column": this.props.colId,
            })
        })
            .then((resp) => {
                return resp.json()
            })
            .then((data) => {
                this.props.boardCallback(data)
            })
            .catch((error) => {
                console.log(error, "catch the hoop")
            })
    }

    render() {
        let color = "white"
        if (this.state.value === "") {
            color = "grey"
        }
        return (
            <td onClick={this.pickCell} bgcolor={color} height="20px" width="20px" align="center" id={this.props.colId}> {this.state.value} </td>
        );
    }
}
export default Cell;