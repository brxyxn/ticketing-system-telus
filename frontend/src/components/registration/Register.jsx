import React from "react";

export class Register extends React.Component {
    render() {
        return (
            <div>
                <h1>Register</h1>
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
