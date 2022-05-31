import "./App.css";
import { NavLink as Link, Route, Routes } from "react-router-dom";
import Login from "./login/Login";
import { Home } from "./home/Home";
import { Register } from "./registration/Register";
import PrivateRoute from "./PrivateRoute";
import GetTickets from "./tickets/Ticket";

function App() {
    return (
        <>
            <nav>
                <ul>
                    <li>
                        <Link to="/">Home</Link>
                    </li>
                    <li>
                        <Link to="/register">Register</Link>
                    </li>
                    <li>
                        <Link to="/login">Login</Link>
                    </li>
                    <li>
                        <Link to="/tickets">Tickets</Link>
                    </li>
                </ul>
            </nav>
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
