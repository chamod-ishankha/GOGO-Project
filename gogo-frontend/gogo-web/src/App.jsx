import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "@/pages/Login";
import Onboarding from "@/pages/Onboarding.jsx";

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Onboarding />} />
                <Route path="/login" element={<Login />} />
                <Route path="/admin" element={<Login />} />
            </Routes>
        </Router>
    );
}

export default App;
