import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "@/pages/Login";
import Onboarding from "@/pages/Onboarding";
import Dashboard from "@/pages/Dashboard";
import ProtectedRoute from "@/components/ProtectedRoute";
import AdminRides from "@/pages/admin/Rides.jsx";
import AdminDrivers from "@/pages/admin/Drivers.jsx";
import AdminSettings from "@/pages/admin/AdminSettings.jsx";
import RiderHistory from "@/pages/rider/RiderHistory.jsx";

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Onboarding />} />
                <Route path="/login" element={<Login />} />
                <Route path="/admin" element={<Login />} />

                <Route
                    path="/dashboard"
                    element={
                        <ProtectedRoute>
                            <Dashboard />
                        </ProtectedRoute>
                    }
                />

                <Route
                    path="/admin/dashboard"
                    element={
                        <ProtectedRoute role="admin">
                            <Dashboard />
                        </ProtectedRoute>
                    }
                />

                <Route
                    path="/admin/rides"
                    element={
                        <ProtectedRoute role="admin">
                            <AdminRides />
                        </ProtectedRoute>
                    }
                />

                <Route
                    path="/admin/drivers"
                    element={
                        <ProtectedRoute role="admin">
                            <AdminDrivers />
                        </ProtectedRoute>
                    }
                />

                <Route
                    path="/admin/settings"
                    element={
                        <ProtectedRoute role="admin">
                            <AdminSettings />
                        </ProtectedRoute>
                    }
                />

                <Route
                    path="/history"
                    element={
                        <ProtectedRoute>
                            <RiderHistory />
                        </ProtectedRoute>
                    }
                />
            </Routes>
        </Router>
    );
}

export default App;
