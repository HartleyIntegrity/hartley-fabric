import { Routes, Route } from 'react-router-dom';
import { CreateTenancyAgreement } from '../CreateTenancyAgreement/CreateTenancyAgreement';
import { EditTenancyAgreement } from '../EditTenancyAgreement/EditTenancyAgreement';
import { TenancyAgreementList } from '../TenancyAgreementList/TenancyAgreementList';
import Header from '../Header/Header';

function App() {
  return (
    <div className="App">
      <Header />
      <Routes>
        <Route path="/" element={<TenancyAgreementList />} />
        <Route path="/create" element={<CreateTenancyAgreement />} />
        <Route path="/edit/:id" element={<EditTenancyAgreement />} />
      </Routes>
    </div>
  );
}

export default App;
