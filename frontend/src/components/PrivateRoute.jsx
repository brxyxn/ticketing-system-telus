import React from "react";
import { Navigate, useLocation } from "react-router-dom";

let isAuthenticated = false; // temp implementation

const PrivateRoute = ({ children }) => {
    let location = useLocation();
    return isAuthenticated ? (
        children
    ) : (
        <Navigate to="/login" state={{ from: location }} />
    );
};

export default PrivateRoute;
