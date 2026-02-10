import { useAuth } from "@/context/AuthContext";
import AdminDashboard from "./dashboards/AdminDashboard";
import RiderDashboard from "./dashboards/RiderDashboard";
import DriverDashboard from "./dashboards/DriverDashboard";

export default function Dashboard() {
    const { user } = useAuth();

    // Switch based on user role
    if (user?.role === 'admin') return <AdminDashboard />;
    if (user?.role === 'driver') return <DriverDashboard />;
    return <RiderDashboard />; // Default to Rider
}