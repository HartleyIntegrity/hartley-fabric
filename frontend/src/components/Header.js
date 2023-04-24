import React from "react";
import { Link } from "react-router-dom";

function Header() {
  return (
    <header>
      <nav>
        <Link to="/">Tenancy Agreements</Link>
        <Link to="/create">Create Tenancy Agreement</Link>
      </nav>
    </header>
  );
}

export default Header;
