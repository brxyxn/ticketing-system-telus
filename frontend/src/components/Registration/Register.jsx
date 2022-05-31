import React from "react";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import axios from "axios";

export class Register extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            firstname: "",
            lastname: "",
            company: "",
            email: "",
            password: "",
            repeatPassword: "",
        };
    }

    async handleSubmit(event) {
        event.preventDefault();

        if (this.state.password === this.state.repeatPassword) {
            var _headers = {
                headers: {
                    Accept: "application/json",
                    "Content-Type": "application/json",
                },
            };

            await axios
                .post(
                    "http://localhost:5000/api/customer/register",
                    JSON.stringify({
                        user: {
                            email: this.state.email,
                            password: this.state.password,
                        },
                        profile: {
                            first_name: this.state.firstname,
                            last_name: this.state.lastname,
                        },
                        company: {
                            name: this.state.company,
                        },
                    }),
                    _headers
                )
                .then((response) => {
                    window.location.href = "/login";
                })
                .catch((error) => {
                    console.log(error);
                });
        }
    }

    componentDidMount() {}

    render() {
        return (
            <Row>
                <Col>
                    <Form className="col-md-4 mx-auto my-5">
                        <Form.Group className="mb-3" controlId="formFirstName">
                            <Form.Label>Firstname</Form.Label>
                            <Form.Control
                                type="text"
                                placeholder="Enter firstname"
                                onChange={(e) =>
                                    this.setState({ firstname: e.target.value })
                                }
                                value={this.state.firstname}
                            />
                        </Form.Group>

                        <Form.Group className="mb-3" controlId="formLastName">
                            <Form.Label>Lastname</Form.Label>
                            <Form.Control
                                type="text"
                                placeholder="Enter lastname"
                                onChange={(e) =>
                                    this.setState({ lastname: e.target.value })
                                }
                                value={this.state.lastname}
                            />
                        </Form.Group>

                        <Form.Group className="mb-3" controlId="formCompany">
                            <Form.Label>Company Name</Form.Label>
                            <Form.Control
                                type="text"
                                placeholder="Enter company name"
                                onChange={(e) =>
                                    this.setState({ company: e.target.value })
                                }
                                value={this.state.company}
                            />
                        </Form.Group>

                        <Form.Group className="mb-3" controlId="formEmail">
                            <Form.Label>Email</Form.Label>
                            <Form.Control
                                type="email"
                                placeholder="Enter email"
                                onChange={(e) =>
                                    this.setState({ email: e.target.value })
                                }
                                value={this.state.email}
                            />
                        </Form.Group>

                        <Form.Group className="mb-3" controlId="formPassword">
                            <Form.Label>Password</Form.Label>
                            <Form.Control
                                type="password"
                                placeholder="Password"
                                onChange={(e) =>
                                    this.setState({ password: e.target.value })
                                }
                                value={this.state.password}
                            />
                        </Form.Group>
                        <Form.Group
                            className="mb-3"
                            controlId="formRepeatPassword"
                        >
                            <Form.Label>Repeat Password</Form.Label>
                            <Form.Control
                                type="password"
                                placeholder="Repeat Password"
                                onChange={(e) =>
                                    this.setState({
                                        repeatPassword: e.target.value,
                                    })
                                }
                                value={this.state.repeatPassword}
                            />
                        </Form.Group>
                        <Button
                            variant="primary"
                            onClick={(e) => this.handleSubmit(e)}
                        >
                            Submit
                        </Button>
                    </Form>
                </Col>
            </Row>
        );
    }
}
