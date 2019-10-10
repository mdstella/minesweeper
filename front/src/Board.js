import React, { Component } from 'react';
import { Table, TableBody, TableCell, TableRow } from '@material-ui/core';

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
            var row = columns.map(function (cell, j) {
                var color = "white"
                if (cell === "") {
                    color = "grey"
                }
                return (
                    <td bgcolor={color} width="20px" align="center" id={j}> {cell} </td>
                );
            });
            return (
                <tr height="20px" align="center" id={i}> {row} </tr>
            );
        });
        return (
            <div>
                <table border="1">
                    <tbody>
                        {rows}
                    </tbody>
                </table>
                <div />
                <div />
                <div />
                <Table>
                    <TableBody>
                        {board.map((rows, i) => (
                            < TableRow key={i} >
                                {rows.map((col, j) => (
                                    < TableCell key={j} component="th" scope="row">
                                        {col}
                                    </TableCell>

                                ))}
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </div >
        );
    }
}
export default Board;