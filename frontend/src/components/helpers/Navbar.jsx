import React from "react";
import Container from "react-bootstrap/Container";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import NavDropdown from "react-bootstrap/NavDropdown";
import Button from "react-bootstrap/esm/Button";
import { NavLink as Link } from "react-router-dom";

const NavbarComponent = () => {
    return (
        <Navbar bg="dark" variant="dark" expand="lg">
            <Container>
                <Navbar.Brand href="#home">React-Bootstrap</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="me-auto">
                        <Nav.Item>
                            <Link to="/" className="nav-link">
                                Home
                            </Link>
                        </Nav.Item>
                        <Nav.Item>
                            <Link to="/register" className="nav-link">
                                Register
                            </Link>
                        </Nav.Item>
                        <Nav.Item>
                            <Link to="/login" className="nav-link">
                                Login
                            </Link>
                        </Nav.Item>
                        <Nav.Item>
                            <Link to="/tickets" className="nav-link">
                                Tickets
                            </Link>
                        </Nav.Item>

                        <NavDropdown title="Profile" id="basic-nav-dropdown">
                            <NavDropdown.Item href="#action/3.2">
                                Settings
                            </NavDropdown.Item>
                            <NavDropdown.Item href="#action/3.3">
                                Manage
                            </NavDropdown.Item>
                            <NavDropdown.Divider />
                            <NavDropdown.Item href="#action/3.4">
                                Logout
                            </NavDropdown.Item>
                        </NavDropdown>
                    </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
};

export default NavbarComponent;
