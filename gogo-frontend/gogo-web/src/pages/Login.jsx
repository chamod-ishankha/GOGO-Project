import { useState } from "react";
import { useLocation } from "react-router-dom";
import api from "@/api/axios";
import { useAuth } from "@/context/AuthContext";
import gogoLogoLight from "@/assets/gogo_wo_bg.png";
import gogoLogoDark from "@/assets/gogo_wo_bg_dark.png";
import {useTheme} from "@/context/ThemeContext.js";

export default function Login() {
    const {isDarkMode} = useTheme();

    const { login } = useAuth();
    const location = useLocation();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    // Determine role based on URL path or query string
    const queryParams = new URLSearchParams(location.search);
    const roleParam = queryParams.get("role"); // 'rider' or 'driver'
    const isAdminPath = location.pathname === "/admin";

    const config = {
        title: isAdminPath ? "Admin Portal" : roleParam === "driver" ? "Driver Login" : "Rider Login",
        color: isAdminPath ? "from-red-600 to-orange-600" : roleParam === "driver" ? "from-purple-600 to-indigo-600" : "from-blue-600 to-cyan-600",
        welcomeText: isAdminPath ? "Management System" : roleParam === "driver" ? "Welcome back, Partner!" : "Where are we going today?"
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError("");
        setLoading(true);
        try {
            const res = await api.post("/user/login", { email, password });
            login(res.data.user, res.data.token);
            window.location.href = isAdminPath ? "/admin/dashboard" : "/dashboard";
        } catch (err) {
            setError(err.response?.data?.message || "Login failed");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-slate-50 dark:bg-slate-900 transition-colors duration-300">
            <div className="max-w-md w-full">
                <div className="flex flex-col items-center mb-8">
                    <img src={isDarkMode ? gogoLogoDark : gogoLogoLight} alt="GoGo" className="h-20 mb-4 drop-shadow-md" />
                    <h2 className={`text-3xl font-black text-transparent bg-clip-text bg-gradient-to-r ${config.color}`}>
                        {config.title}
                    </h2>
                    <p className="text-gray-500 font-medium mt-1">{config.welcomeText}</p>
                </div>

                <form onSubmit={handleSubmit} className="bg-white dark:bg-slate-800 p-8 rounded-3xl shadow-2xl border border-white/20 dark:border-slate-700">
                    {error && <div className="bg-red-50 text-red-500 p-3 rounded-xl mb-4 text-center text-sm">{error}</div>}

                    <div className="space-y-4">
                        <input
                            type="email"
                            placeholder="Email Address"
                            className="w-full px-5 py-4 bg-gray-50 dark:bg-slate-700 text-gray-900 dark:text-white border-none rounded-2xl focus:ring-2 focus:ring-blue-400 outline-none transition-all placeholder:text-gray-400"
                            value={email} onChange={(e) => setEmail(e.target.value)} required
                        />
                        <input
                            type="password"
                            placeholder="Password"
                            className="w-full px-5 py-4 bg-gray-50 dark:bg-slate-700 text-gray-900 dark:text-white border-none rounded-2xl focus:ring-2 focus:ring-blue-400 outline-none transition-all placeholder:text-gray-400"
                            value={password} onChange={(e) => setPassword(e.target.value)} required
                        />
                    </div>

                    <button
                        type="submit"
                        disabled={loading}
                        className={`w-full mt-8 bg-gradient-to-r ${config.color} text-white font-bold py-4 rounded-2xl shadow-lg hover:scale-[1.02] active:scale-[0.98] transition-all`}
                    >
                        {loading ? "Verifying..." : "Login"}
                    </button>
                </form>

                {!isAdminPath && (
                    <button onClick={() => {window.location.href = '/'}} className="w-full mt-4 text-gray-400 hover:text-gray-600 text-sm">
                        ← Change account type
                    </button>
                )}
            </div>
        </div>
    );
}