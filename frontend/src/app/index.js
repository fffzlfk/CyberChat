import React from "react";
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
} from "react-router-dom";
import { Chat } from "./pages/Chat";
import About from "./pages/About";
import "./styles.css"

export default function App() {
    return (
        <div>
            <Router>
                <nav>
                    <div>
                        <ul>
                            <li>
                                <Link to="/">Chat</Link>
                            </li>
                            <li>
                                <Link to="/about">About</Link>
                            </li>
                        </ul>
                    </div>
                </nav>
                <hr />
                <Switch>
                    <Route exact path="/">
                        <Chat />
                    </Route>
                    <Route exact path="/about">
                        <About />
                    </Route>
                </Switch>
            </Router>
        </div>
    );
}
