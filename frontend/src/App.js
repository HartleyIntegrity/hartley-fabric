import React from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import TenancyAgreementList from "./components/TenancyAgreementList";
import Property from "./components/Property";
import TenancyAgreementForm from "./components/TenancyAgreementForm";
import { AppBar, Toolbar, Typography } from "@mui/material";

function App() {
  return (
    <Router>
      <div>
        <AppBar position="static">
          <Toolbar>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              Hartley Fabric
            </Typography>
            <Link to="/" style={{ color: "white", textDecoration: "none" }}>
              Tenancy Agreements
            </Link>
          </Toolbar>
        </AppBar>

        <Switch>
          <Route exact path="/">
            <TenancyAgreementList />
          </Route>
          <Route path="/property/:id">
            <Property />
          </Route>
          <Route exact path="/create">
            <TenancyAgreementForm />
          </Route>
          <Route path="/update/:id">
            <TenancyAgreementForm />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
