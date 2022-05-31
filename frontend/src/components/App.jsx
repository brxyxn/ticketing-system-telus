import "./App.css";
import { Route, Routes } from "react-router-dom";
import Login from "./Authentication/Login";
import { Home } from "./Home/Home";
import { Register } from "./Registration/Register";
import PrivateRoute from "./PrivateRoute";
import GetTickets from "./Tickets/Ticket";

import Container from "react-bootstrap/Container";

import NavbarComponent from "./helpers/Navbar";
import { history } from "./helpers/History";

function App() {
    return (
        <>
            <NavbarComponent />
            <main>
                <Container fluid="md">
                    <Routes history={history}>
                        <Route path="/" element={<Home />} />
                        <Route path="register" element={<Register />} />
                        <Route path="login" element={<Login />} />
                        <Route
                            path="tickets"
                            element={
                                <PrivateRoute>
                                    <GetTickets test="Hello Tickets" />
                                </PrivateRoute>
                            }
                        />
                    </Routes>
                </Container>
            </main>
        </>
    );
}

export default App;
