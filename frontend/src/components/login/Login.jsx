import React from "react";

export default class Login extends React.Component {
    render() {
        return (
            <div>
                <h1>Login</h1>
                <form>
                    <label>
                        Username:
                        <input type="text" name="username" />
                    </label>
                    <label>
                        Password:
                        <input type="password" name="password" />
                    </label>
                    <input type="submit" value="Submit" />
                </form>
            </div>
        );
    }
}
