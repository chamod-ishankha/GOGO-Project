import { useNavigate } from "react-router-dom";
import gogoLogoLight from "@/assets/gogo_wo_bg.png";
import gogoLogoDark from "@/assets/gogo_wo_bg_dark.png";
import ThemeToggle from "@/components/ThemeToggle";
import {useTheme} from "@/context/ThemeContext.js";

export default function Onboarding() {
    const { isDarkMode } = useTheme();
    const navigate = useNavigate();

    return (
        <div className="relative min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-blue-50 to-purple-50 dark:from-slate-900 dark:to-slate-800 p-6 transition-colors duration-300">

            {/* Position the toggle in the top-right corner */}
            <div className="absolute top-6 right-6">
                <ThemeToggle />
            </div>

            <img
                src={isDarkMode ? gogoLogoDark : gogoLogoLight}
                alt="GoGo Logo"
                className="h-32 mb-8 drop-shadow-xl transition-all duration-500"
            />

            <h1 className="text-4xl font-black text-gray-800 dark:text-white mb-2 transition-colors">
                Ready to GoGo?
            </h1>

            <p className="text-gray-500 dark:text-gray-400 mb-12 text-center max-w-sm transition-colors">
                Choose how you want to use the platform today.
            </p>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-6 w-full max-w-2xl">
                {/* Rider Option */}
                <button
                    onClick={() => navigate("/login?role=rider")}
                    className="group bg-white dark:bg-slate-800 p-8 rounded-3xl shadow-lg dark:shadow-2xl transition-all border-2 border-transparent hover:border-blue-500 text-left"
                >
                    <div className="text-4xl mb-4">🏠</div>
                    <h3 className="text-2xl font-bold text-blue-600 dark:text-blue-400 group-hover:translate-x-1 transition-transform">I want a Ride</h3>
                    <p className="text-gray-400 dark:text-gray-500 text-sm">Find a driver nearby and get to your destination quickly.</p>
                </button>

                {/* Driver Option */}
                <button
                    onClick={() => navigate("/login?role=driver")}
                    className="group bg-white dark:bg-slate-800 p-8 rounded-3xl shadow-lg dark:shadow-2xl transition-all border-2 border-transparent hover:border-blue-500 text-left"
                >
                    <div className="text-4xl mb-4">🚗</div>
                    <h3 className="text-2xl font-bold text-blue-600 dark:text-blue-400 group-hover:translate-x-1 transition-transform">I want to Drive</h3>
                    <p className="text-gray-400 dark:text-gray-500 text-sm">Earn money by picking up passengers in your area.</p>
                </button>
            </div>
        </div>
    );
}