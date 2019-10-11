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

    // this will update the state if there are new props on the component
    static getDerivedStateFromProps(props, state) {
        if (props.value !== state.value) {
            return {
                value: props.value
            };
        }
        // Return null to indicate no change to state.
        return null;
    }

    // invokes the API /game/{gameId} when a cell is clicked. With the response it will re render
    // the board. In the future maybe the cell could be the only thing we re render, but now we are getting the 
    // complete board from the API
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
                // invoking the board callback to update the board status and re render
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
        } else if (this.state.value === "*") {
            color = "red"
        }
        return (
            <td onClick={this.pickCell} bgcolor={color} height="20px" width="20px" align="center" id={this.props.colId}> {this.state.value} </td>
        );
    }
}
export default Cell;