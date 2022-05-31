import React from "react";

export default class GetTickets extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            test: "",
            tickets: [],
            isLoading: true,
            error: null,
        };
    }

    render() {
        return <h1>{this.props.test}</h1>;
    }
}
