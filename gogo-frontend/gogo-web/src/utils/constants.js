export const NAV_LINKS = {
    admin: [
        { name: "Dashboard", icon: "📊", path: "/admin/dashboard" },
        { name: "All Rides", icon: "🚗", path: "/admin/rides" },
        { name: "Drivers List", icon: "👥", path: "/admin/drivers" },
        { name: "Settings", icon: "⚙️", path: "/admin/settings" },
    ],
    rider: [
        { name: "Book a Ride", icon: "📍", path: "/dashboard" },
        { name: "My History", icon: "📜", path: "/history" },
        { name: "Profile", icon: "👤", path: "/profile" },
    ],
    driver: [
        { name: "Available Jobs", icon: "🛣️", path: "/jobs" },
        { name: "Earnings", icon: "💰", path: "/earnings" },
        { name: "Vehicle Docs", icon: "📄", path: "/vehicle" },
    ],
};