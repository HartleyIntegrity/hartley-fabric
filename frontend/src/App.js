import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import TenancyAgreementForm from './components/TenancyAgreementForm';
import TenancyAgreementList from './components/TenancyAgreementList';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<TenancyAgreementList />} />
        <Route path="/create" element={<TenancyAgreementForm />} />
      </Routes>
    </Router>
  );
}

export default App;
