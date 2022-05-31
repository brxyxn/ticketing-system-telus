import React from "react";
import { Navigate, useLocation } from "react-router-dom";

function hasJWT() {
    let flag = false;
    localStorage.getItem("token") ? (flag = true) : (flag = false);
    return flag;
}

const PrivateRoute = ({ children }) => {
    let location = useLocation();
    return hasJWT() ? (
        children
    ) : (
        <Navigate to="/login" state={{ from: location }} />
    );
};

export default PrivateRoute;
