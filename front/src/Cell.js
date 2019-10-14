import React, { Component } from 'react';

class Cell extends Component {
    constructor(props) {
        super(props);
        this.state = { value: '', addFlag: false };
    }

    componentDidMount() {
        this.setState({
            value: this.props.value,
            addFlag: this.props.addFlag
        });
    }

    // this will update the state if there are new props on the component
    static getDerivedStateFromProps(props, state) {
        if (props.value !== state.value || props.addFlag !== state.addFlag) {
            return {
                value: props.value,
                addFlag: props.addFlag
            };
        }
        // Return null to indicate no change to state.
        return null;
    }

    // invokes the API /game/{gameId} when a cell is clicked. With the response it will re render
    // the board. In the future maybe the cell could be the only thing we re render, but now we are getting the 
    // complete board from the API
    pickCell = () => {
        let addFlag = this.state.addFlag
        let url = "http://localhost:8000/minesweeper/v1/game/" + this.props.gameId
        if (addFlag) {
            url = "http://localhost:8000/minesweeper/v1/flag/" + this.props.gameId
        }

        if ((addFlag && (this.state.value === "" || this.state.value === "?")) ||
            (!addFlag && this.state.value === "")) {

            fetch(url, {
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
    }

    render() {
        let color = "white"
        let hasImage = false
        let img
        if (this.state.value === "") {
            color = "grey"
        } else if (this.state.value === "*") {
            hasImage = true
            img = <img alt="mine" width="20" height="20" align="center" src={require('./images/mine.png')} />
            color = "red"
        } else if (this.state.value === "?") {
            hasImage = true
            img = <img alt="mine" width="20" height="20" align="center" src={require('./images/flag.png')} />
        }

        return (
            <td onClick={this.pickCell} bgcolor={color} height="20px" width="20px" align="center" id={this.props.colId}>
                {hasImage ? img : this.state.value}
            </td>
        );
    }
}
export default Cell;