import React from "react";
import { Link } from "react-router-dom";

function Header({ isLoggedIn }) {
  return (
    <header className="App-header">
      <nav>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/login">{isLoggedIn ? "Logout" : "Login"}</Link>
          </li>
          <li>
            <Link to="/signup">Signup</Link>
          </li>
          {isLoggedIn && (
            <li>
              <Link to="/admin">Admin Dashboard</Link>
            </li>
          )}
        </ul>
      </nav>
    </header>
  );
}

export default Header;
