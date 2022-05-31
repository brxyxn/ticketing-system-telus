import React from "react";
import Container from "react-bootstrap/Container";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import { NavLink as Link } from "react-router-dom";

function setAuthenticated() {
    const token = localStorage.getItem("token");
    if (token !== null) {
        return true;
    } else {
        return false;
    }
}

function handleLogout() {
    localStorage.removeItem("token");
    window.location.reload();
}

const NavbarComponent = () => {
    let isAuthenticated = setAuthenticated();

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
                            <Link to="/tickets" className="nav-link">
                                Tickets
                            </Link>
                        </Nav.Item>

                        {!isAuthenticated ? (
                            <>
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
                            </>
                        ) : (
                            <Nav.Item>
                                <Link
                                    to="/logout"
                                    className="nav-link"
                                    onClick={handleLogout}
                                >
                                    Logout
                                </Link>
                            </Nav.Item>
                        )}
                    </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
};

export default NavbarComponent;
