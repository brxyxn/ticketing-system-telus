import "./App.css";
import { NavLink as Link, Route, Routes } from "react-router-dom";
import Login from "./login/Login";
import { Home } from "./home/Home";
import { Register } from "./registration/Register";
import PrivateRoute from "./PrivateRoute";
import GetTickets from "./tickets/Ticket";

import NavbarComponent from "./helpers/Navbar";

function App() {
    return (
        <>
            <NavbarComponent />
            <main>
                <Routes>
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
            </main>
        </>
    );
}

export default App;
