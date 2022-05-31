import React from "react";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import axios from "axios";

export const setAuthToken = (token) => {
    if (token) {
        axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
    } else {
        delete axios.defaults.headers.common["Authorization"];
    }
};

export default class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isLoading: true,
            email: "",
            password: "",
        };
    }

    async handleSubmit(e) {
        e.preventDefault();
        const login = {
            email: this.state.email,
            password: this.state.password,
        };

        await axios
            .post("http://localhost:5000/api/customer/login", login)
            .then((response) => {
                const token = response.data.token;
                localStorage.setItem("token", token);
                setAuthToken(token);
                window.location.href = "/";
            })
            .catch((err) => {
                console.log(err);
            });
    }

    componentDidMount() {
        const token = localStorage.getItem("token");
        if (token) {
            window.location.href = "/";
        } else {
            this.setState({ isLoading: false });
        }
    }

    render() {
        if (this.state.isLoading) {
            return <div>Loading...</div>;
        }
        if (!this.state.isLoading)
            return (
                <Row>
                    <Col>
                        <Form className="col-md-4 mx-auto my-5">
                            <Form.Group
                                className="mb-3"
                                controlId="formBasicEmail"
                            >
                                <Form.Label>Email address</Form.Label>
                                <Form.Control
                                    type="email"
                                    placeholder="Enter email"
                                    onChange={(e) =>
                                        this.setState({ email: e.target.value })
                                    }
                                    value={this.state.email}
                                />
                                <Form.Text className="text-muted">
                                    We'll never share your email with anyone
                                    else.
                                </Form.Text>
                            </Form.Group>

                            <Form.Group
                                className="mb-3"
                                controlId="formBasicPassword"
                            >
                                <Form.Label>Password</Form.Label>
                                <Form.Control
                                    type="password"
                                    placeholder="Password"
                                    onChange={(e) =>
                                        this.setState({
                                            password: e.target.value,
                                        })
                                    }
                                    value={this.state.password}
                                />
                            </Form.Group>
                            <Form.Group
                                className="mb-3"
                                controlId="formBasicCheckbox"
                            >
                                <Form.Check
                                    type="checkbox"
                                    label="Check me out"
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
