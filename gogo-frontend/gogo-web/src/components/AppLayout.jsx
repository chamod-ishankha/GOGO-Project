import { useAuth } from "@/context/AuthContext";
import { useTheme } from "@/context/ThemeContext";
import ThemeToggle from "@/components/ThemeToggle";
import gogoLogoLight from "@/assets/gogo_wo_bg.png";
import gogoLogoDark from "@/assets/gogo_wo_bg_dark.png";
import { Link, useLocation } from "react-router-dom";
import { NAV_LINKS } from "@/utils/constants";
import logoutIco from "@/assets/icons/logout.png";

export default function AppLayout({ children }) {
    const { logout, user } = useAuth();
    const { isDarkMode } = useTheme();
    const location = useLocation();

    // Determine role (defaulting to rider if not found)
    const role = user?.role || "rider";
    const links = NAV_LINKS[role] || NAV_LINKS.rider;

    // Role-based theme colors
    const brandColors = {
        admin: "text-red-500",
        driver: "text-purple-600",
        rider: "text-blue-600"
    };

    return (
        <div className="h-screen flex bg-slate-50 dark:bg-slate-950 transition-colors duration-500 overflow-hidden">
            {/* --- SIDEBAR --- */}
            {/* Change 2: Ensure h-full is set so the sidebar spans the viewport */}
            <aside className="w-72 h-full bg-white dark:bg-slate-900 border-r border-gray-200 dark:border-slate-800 flex flex-col transition-colors duration-500 hidden md:flex shrink-0">
                {/* --- BRAND SECTION --- */}
                <div className="px-6 py-8 overflow-y-auto flex-1">
                    <div className="flex items-center gap-3 mb-8 px-2">
                        <img
                            src={isDarkMode ? gogoLogoDark : gogoLogoLight}
                            className="h-10 w-auto object-contain"
                            alt="GoGo Logo"
                        />
                        <div className="flex flex-col">
                            <span className={`text-2xl font-black tracking-tighter leading-none ${isDarkMode ? 'text-white' : 'text-slate-900'}`}>
                                GOGO
                            </span>
                            <span className="text-[10px] font-bold uppercase tracking-[0.2em] text-blue-600 dark:text-blue-400 leading-none mt-1">
                                Platform
                            </span>
                        </div>
                    </div>

                    {/* --- NAVIGATION --- */}
                    <nav className="w-full space-y-1">
                        {links.map((link) => (
                            <Link
                                key={link.name}
                                to={link.path}
                                className={`flex items-center gap-3 px-4 py-3 rounded-xl font-medium transition-all group ${
                                    location.pathname === link.path
                                        ? `bg-blue-50 dark:bg-blue-900/20 ${brandColors[role]}`
                                        : "text-gray-500 dark:text-gray-400 hover:bg-gray-50 dark:hover:bg-slate-800/50"
                                }`}
                            >
                                <span className={`flex items-center justify-center w-6 h-6 text-xl transition-transform group-hover:scale-110`}>
                                    {link.icon}
                                </span>
                                <span className="tracking-tight">{link.name}</span>
                            </Link>
                        ))}
                    </nav>
                </div>

                {/* User Profile Summary at bottom */}
                <div className="mt-auto p-6 border-t border-gray-100 dark:border-slate-800 bg-gray-50/50 dark:bg-slate-900/50">
                    <div className="flex items-center gap-3 mb-4">
                        <div className="shrink-0 w-10 h-10 rounded-full bg-gradient-to-tr from-blue-500 to-purple-500 flex items-center justify-center text-white font-bold shadow-sm">
                            {user?.email?.[0].toUpperCase()}
                        </div>
                        <div className="min-w-0"> {/* Fixed truncate by adding min-w-0 */}
                            <p className="text-sm font-bold dark:text-white truncate">{user?.email}</p>
                            <p className="text-xs text-gray-400 capitalize">{role} Account</p>
                        </div>
                    </div>
                    <button
                        onClick={logout}
                        className="w-full flex items-center gap-3 text-sm font-medium text-red-500 hover:bg-red-100/50 dark:hover:bg-red-900/20 p-2.5 rounded-xl transition-all border border-transparent hover:border-red-200 dark:hover:border-red-900/50"
                    >
                        <img
                            src={logoutIco}
                            alt="Logout"
                            className="w-5 h-5 object-contain"
                        />
                        Sign Out
                    </button>
                </div>
            </aside>

            {/* --- MAIN CONTENT AREA --- */}
            <div className="flex-1 flex flex-col h-full overflow-hidden">
                {/* HEADER */}
                <header className="h-20 shrink-0 bg-white/80 dark:bg-slate-900/80 backdrop-blur-md border-b border-gray-200 dark:border-slate-800 px-8 flex items-center justify-between sticky top-0 z-10 transition-colors duration-500">
                    <div>
                        <h2 className="text-sm font-medium text-gray-400 uppercase tracking-widest">Platform</h2>
                        <h1 className="text-xl font-bold dark:text-white capitalize">
                            {location.pathname.split('/').pop() || 'Dashboard'}
                        </h1>
                    </div>

                    <div className="flex items-center gap-4">
                        <ThemeToggle />
                        <div className="h-8 w-[1px] bg-gray-200 dark:bg-slate-700 mx-2"></div>
                        <span className={`px-3 py-1 rounded-full text-xs font-bold uppercase ${
                            role === 'admin' ? 'bg-red-100 text-red-600' : 'bg-green-100 text-green-600'
                        }`}>
                            {role}
                        </span>
                    </div>
                </header>

                {/* MAIN PAGE BODY */}
                <main className="flex-1 overflow-y-auto p-8 custom-scrollbar">
                    <div className="max-w-7xl mx-auto">
                        {children}
                    </div>

                    {/*/!* FOOTER (Moved inside scroll area so it appears at bottom of content) *!/*/}
                    {/*<footer className="mt-20 py-8 text-center text-gray-400 text-sm border-t border-gray-100 dark:border-slate-800">*/}
                    {/*    &copy; 2026 GoGo Project &bull; {role.charAt(0).toUpperCase() + role.slice(1)} Dashboard*/}
                    {/*</footer>*/}
                </main>

                {/* FOOTER */}
                <footer className="mt-auto p-8 text-center text-gray-400 text-sm border-t border-gray-100 dark:border-slate-800">
                    &copy; 2026 GoGo Project &bull; {role.charAt(0).toUpperCase() + role.slice(1)} Dashboard
                </footer>
            </div>
        </div>
    );
}