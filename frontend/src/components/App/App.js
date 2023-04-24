import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Header from "../Header/Header";
import Home from "../../pages/Home/Home";
import CreateTenancyAgreement from "../../pages/CreateTenancyAgreement/CreateTenancyAgreement";
import EditTenancyAgreement from "../../pages/EditTenancyAgreement/EditTenancyAgreement";
import TenancyAgreementList from "../../pages/TenancyAgreementList/TenancyAgreementList";

function App() {
  return (
    <Router>
      <div className="App">
        <Header />
        <Routes>
          <Route exact path="/" component={Home} />
          <Route
            exact
            path="/create-tenancy-agreement"
            component={CreateTenancyAgreement}
          />
          <Route
            exact
            path="/edit-tenancy-agreement/:id"
            component={EditTenancyAgreement}
          />
          <Route
            exact
            path="/tenancy-agreements"
            component={TenancyAgreementList}
          />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
