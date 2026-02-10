import AppLayout from "@/components/AppLayout";
import { useState, useEffect } from "react";

export default function RiderDashboard() {
    // 1. UI States
    const [view, setView] = useState("search"); // search | select | finding | tracking | completed
    const [pickup, setPickup] = useState("");
    const [destination, setDestination] = useState("");

    // New State for Mock Driver
    const [driverInfo, setDriverInfo] = useState(null);

    // Add this to your useEffect or create a button to simulate ride completion
    const simulateRideCompletion = () => {
        setView("completed");
    };

    // Simulation: Automatically "find" a driver after 4 seconds of searching
    useEffect(() => {
        if (view === "finding") {
            const timer = setTimeout(() => {
                setDriverInfo({
                    name: "Arjun Ratnayake",
                    rating: "4.9",
                    vehicle: "White Toyota Prius",
                    plate: "WP CAS-8821",
                    eta: "3 mins"
                });
                setView("tracking");
            }, 4000); // 4 second delay
            return () => clearTimeout(timer);
        }

        // if (view === "tracking") {
        //     const timer = setTimeout(() => {
        //         simulateRideCompletion();
        //     }, 3000)
        //     return () => clearTimeout(timer);
        // }
    }, [view]);

    // Mock Data for Vehicle Types
    const vehicles = [
        { id: 'eco', name: 'GoGo Economy', price: 450, time: '3 mins', icon: '🚗' },
        { id: 'lux', name: 'GoGo Premium', price: 850, time: '5 mins', icon: '✨' },
        { id: 'bike', name: 'GoGo Bike', price: 200, time: '2 mins', icon: '🏍️' },
    ];

    return (
        <AppLayout>
            <div className="max-w-5xl mx-auto space-y-8">

                {/* --- STEP 1: INITIAL SEARCH VIEW --- */}
                {view === "search" && (
                    <>
                        <div className="bg-gradient-to-r from-blue-600 to-indigo-600 rounded-[2rem] p-8 text-white shadow-xl">
                            <h2 className="text-3xl font-black mb-2">Where to, Chamod?</h2>
                            <p className="text-blue-100 mb-8">Get a reliable ride in minutes.</p>

                            <div className="bg-white dark:bg-slate-800 p-2 rounded-2xl flex flex-col md:flex-row gap-2">
                                <input
                                    type="text"
                                    placeholder="Pickup location"
                                    value={pickup}
                                    onChange={(e) => setPickup(e.target.value)}
                                    className="flex-1 p-4 text-gray-800 dark:text-white dark:bg-slate-700 rounded-xl outline-none"
                                />
                                <input
                                    type="text"
                                    placeholder="Where to?"
                                    value={destination}
                                    onChange={(e) => setDestination(e.target.value)}
                                    className="flex-1 p-4 text-gray-800 dark:text-white dark:bg-slate-700 rounded-xl outline-none"
                                />
                                <button
                                    onClick={() => setView("select")}
                                    disabled={!pickup || !destination}
                                    className="bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-bold px-8 py-4 rounded-xl transition-all"
                                >
                                    Search GoGo
                                </button>
                            </div>
                        </div>

                        {/* Quick Shortcuts */}
                        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                            {['Home', 'Work', 'Gym', 'Recent'].map((label) => (
                                <button key={label} className="bg-white dark:bg-slate-900 p-4 rounded-2xl border border-gray-100 dark:border-slate-800 flex flex-col items-center hover:border-blue-500 transition-all group">
                                    <span className="text-2xl mb-2 group-hover:scale-110 transition-transform">📍</span>
                                    <span className="text-sm font-bold dark:text-gray-300">{label}</span>
                                </button>
                            ))}
                        </div>

                        {/* --- NEW: RECENT ACTIVITY SECTION --- */}
                        <div className="space-y-4">
                            <div className="flex justify-between items-center px-2">
                                <h3 className="font-black text-lg dark:text-white">Recent Activity</h3>
                                <button className="text-sm font-bold text-blue-600">View All</button>
                            </div>

                            <div className="space-y-3">
                                {[
                                    { from: "Panadura", to: "Bambalapitiya", date: "Yesterday", fare: "LKR 1,200", icon: "🚗" },
                                    { from: "Colombo Fort", to: "Mount Lavinia", date: "2 days ago", fare: "LKR 850", icon: "🏍️" }
                                ].map((ride, index) => (
                                    <div
                                        key={index}
                                        className="bg-white dark:bg-slate-900 p-5 rounded-[2rem] border border-gray-100 dark:border-slate-800 flex items-center justify-between hover:shadow-md transition-all cursor-pointer group"
                                    >
                                        <div className="flex items-center gap-4">
                                            <div className="w-12 h-12 bg-gray-50 dark:bg-slate-800 rounded-2xl flex items-center justify-center text-xl group-hover:bg-blue-50 dark:group-hover:bg-blue-900/20 transition-colors">
                                                {ride.icon}
                                            </div>
                                            <div>
                                                <p className="font-bold text-sm dark:text-white">{ride.from} → {ride.to}</p>
                                                <p className="text-xs text-gray-400 font-medium">{ride.date}</p>
                                            </div>
                                        </div>
                                        <div className="text-right">
                                            <p className="font-black text-sm dark:text-white">{ride.fare}</p>
                                            <p className="text-[10px] text-green-500 font-bold uppercase tracking-tighter">Completed</p>
                                        </div>
                                    </div>
                                ))}
                            </div>
                        </div>

                        {/* Available Near You Simulation (Moved to bottom) */}
                        <div className="bg-white dark:bg-slate-900 rounded-3xl p-6 border border-gray-100 dark:border-slate-800">
                            <h3 className="font-bold text-lg mb-4 dark:text-white">Available Near You</h3>
                            {/* ... (Keep your vehicles map code here) ... */}
                        </div>
                    </>
                )}

                {/* --- STEP 2: VEHICLE SELECTION VIEW --- */}
                {view === "select" && (
                    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500">
                        <button onClick={() => setView("search")} className="mb-4 text-blue-600 font-bold flex items-center gap-2">
                            ← Edit Route
                        </button>
                        <div className="bg-white dark:bg-slate-900 rounded-[2rem] border border-gray-100 dark:border-slate-800 overflow-hidden shadow-lg">
                            <div className="p-6 border-b border-gray-50 dark:border-slate-800 bg-gray-50/50 dark:bg-slate-800/30">
                                <p className="text-xs font-bold text-gray-400 uppercase">Your Route</p>
                                <p className="font-bold dark:text-white">{pickup} → {destination}</p>
                            </div>

                            <div className="p-6 space-y-4">
                                {vehicles.map((v) => (
                                    <div key={v.id} className="flex items-center justify-between p-4 rounded-2xl border-2 border-transparent hover:border-blue-500 bg-gray-50 dark:bg-slate-800/50 cursor-pointer transition-all">
                                        <div className="flex items-center gap-4">
                                            <span className="text-3xl">{v.icon}</span>
                                            <div>
                                                <p className="font-bold dark:text-white">{v.name}</p>
                                                <p className="text-xs text-gray-500">{v.time} away • {v.id === 'bike' ? '1 seat' : '4 seats'}</p>
                                            </div>
                                        </div>
                                        <div className="text-right">
                                            <p className="font-black text-lg dark:text-white">LKR {v.price}.00</p>
                                            <button
                                                onClick={() => setView("finding")}
                                                className="text-xs font-bold text-blue-600 uppercase hover:bg-blue-50 px-3 py-1 rounded-xl transition-all mt-1 block"
                                            >
                                                Book Now
                                            </button>
                                        </div>
                                    </div>
                                ))}
                            </div>
                        </div>
                    </div>
                )}

                {/* --- STEP 3: FINDING DRIVER VIEW --- */}
                {view === "finding" && (
                    <div className="flex flex-col items-center justify-center py-20 animate-in zoom-in duration-500">
                        <div className="relative">
                            <div className="absolute inset-0 bg-blue-500 rounded-full animate-ping opacity-20"></div>
                            <div className="relative bg-blue-600 w-24 h-24 rounded-full flex items-center justify-center text-4xl shadow-2xl">
                                🚗
                            </div>
                        </div>
                        <h2 className="text-2xl font-black mt-10 dark:text-white">Connecting you to a Driver...</h2>
                        <p className="text-gray-500 mt-2">GoGo is searching for the nearest available ride.</p>

                        <button
                            onClick={() => setView("select")}
                            className="mt-12 px-8 py-3 rounded-xl border-2 border-red-100 text-red-500 font-bold hover:bg-red-50 transition-all"
                        >
                            Cancel Request
                        </button>
                    </div>
                )}

                {/* --- STEP 4: DRIVER FOUND / TRACKING VIEW --- */}
                {view === "tracking" && driverInfo && (
                    <div className="animate-in slide-in-from-bottom-8 duration-700">
                        <div className="bg-white dark:bg-slate-900 rounded-[2.5rem] border border-gray-100 dark:border-slate-800 shadow-2xl overflow-hidden">

                            {/* --- ADVANCED FAKE MAP --- */}
                            <div className="h-72 bg-slate-100 dark:bg-slate-950 relative overflow-hidden">
                                {/* Map Grid Pattern */}
                                <div className="absolute inset-0 opacity-20 dark:opacity-10"
                                     style={{ backgroundImage: `radial-gradient(#3b82f6 1.5px, transparent 1.5px)`, backgroundSize: '24px 24px' }}>
                                </div>

                                {/* Simulated Roads */}
                                <svg className="absolute inset-0 w-full h-full opacity-20 dark:opacity-40" xmlns="http://www.w3.org/2000/svg">
                                    <path d="M-20 100 L600 100 M150 -20 L150 400 M400 -20 L400 400 M-20 250 L600 250"
                                          stroke="currentColor" strokeWidth="40" fill="none" className="text-white dark:text-slate-800" />
                                </svg>

                                {/* Pickup Point */}
                                <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 flex flex-col items-center">
                                    <div className="w-4 h-4 bg-blue-600 rounded-full border-4 border-white dark:border-slate-900 shadow-lg z-10"></div>
                                    <div className="bg-white dark:bg-slate-800 px-3 py-1 rounded-lg shadow-sm mt-2 border border-gray-100 dark:border-slate-700">
                                        <p className="text-[10px] font-black dark:text-white uppercase truncate max-w-[100px]">{pickup}</p>
                                    </div>
                                </div>

                                {/* Moving Driver Car */}
                                <div className="absolute top-[20%] left-[30%] animate-pulse">
                                    <div className="flex flex-col items-center animate-bounce duration-[2000ms]">
                                        <div className="bg-black text-white p-2 rounded-xl text-xl shadow-xl rotate-12">
                                            🚕
                                        </div>
                                        <div className="w-2 h-2 bg-black/20 rounded-full blur-sm mt-1"></div>
                                    </div>
                                </div>

                                {/* Arrival Badge */}
                                <div className="absolute bottom-6 left-6 right-6 flex justify-center">
                                    <div className="bg-white/90 dark:bg-slate-900/90 backdrop-blur-md px-6 py-3 rounded-2xl shadow-xl border border-white/20 flex items-center gap-3">
                                        <div className="w-2 h-2 bg-green-500 rounded-full animate-ping"></div>
                                        <span className="text-xs font-black dark:text-white uppercase tracking-widest">
                                            {driverInfo.name} is {driverInfo.eta} away
                                        </span>
                                    </div>
                                </div>
                            </div>

                            {/* Driver Details Card */}
                            <div className="p-8">
                                <div className="flex justify-between items-start mb-6">
                                    <div className="flex items-center gap-4">
                                        <div className="w-16 h-16 bg-gradient-to-tr from-blue-100 to-blue-50 dark:from-slate-700 dark:to-slate-800 rounded-2xl flex items-center justify-center text-2xl">
                                            👨🏻‍✈️
                                        </div>
                                        <div>
                                            <h3 className="text-xl font-black dark:text-white">{driverInfo.name}</h3>
                                            <p className="text-sm text-yellow-500 font-bold">⭐ {driverInfo.rating} Rating</p>
                                        </div>
                                    </div>
                                    <div className="text-right">
                                        <p className="text-lg font-black dark:text-white">{driverInfo.plate}</p>
                                        <p className="text-xs text-gray-500 uppercase font-bold">{driverInfo.vehicle}</p>
                                    </div>
                                </div>

                                <div className="grid grid-cols-2 gap-4">
                                    <button className="bg-gray-100 dark:bg-slate-800 dark:text-white font-bold py-4 rounded-2xl hover:bg-gray-200 transition-all">
                                        📞 Call
                                    </button>
                                    <button className="bg-blue-600 text-white font-bold py-4 rounded-2xl hover:bg-blue-700 transition-all shadow-lg shadow-blue-200 dark:shadow-none">
                                        💬 Message
                                    </button>
                                </div>

                                <button
                                    onClick={() => {
                                        setView("search");
                                        setDriverInfo(null);
                                    }}
                                    className="w-full mt-6 text-sm text-gray-400 font-medium hover:text-red-500 transition-colors"
                                >
                                    Cancel Ride
                                </button>
                            </div>
                        </div>
                    </div>
                )}

                {/* --- STEP 5: RIDE COMPLETED / RECEIPT VIEW --- */}
                {view === "completed" && (
                    <div className="max-w-md mx-auto animate-in zoom-in duration-500">
                        <div className="bg-white dark:bg-slate-900 rounded-[3rem] shadow-2xl border border-gray-100 dark:border-slate-800 overflow-hidden text-center">
                            {/* Success Header */}
                            <div className="bg-green-500 p-10 text-white">
                                <div className="w-20 h-20 bg-white/20 rounded-full flex items-center justify-center mx-auto mb-4 text-4xl">
                                    ✅
                                </div>
                                <h2 className="text-3xl font-black">Arrived Safely!</h2>
                                <p className="opacity-90 font-medium">Hope you enjoyed your GoGo ride.</p>
                            </div>

                            <div className="p-10">
                                <p className="text-sm font-bold text-gray-400 uppercase tracking-widest mb-2">Total Fare Paid</p>
                                <h3 className="text-5xl font-black dark:text-white mb-8">LKR 1,450</h3>

                                <div className="space-y-6">
                                    <p className="font-bold dark:text-gray-300">How was your trip with Arjun?</p>

                                    {/* Rating Stars */}
                                    <div className="flex justify-center gap-3 text-3xl">
                                        {[1, 2, 3, 4, 5].map((star) => (
                                            <button key={star} className="hover:scale-125 transition-transform">⭐</button>
                                        ))}
                                    </div>

                                    <textarea
                                        placeholder="Any feedback for the driver?"
                                        className="w-full p-4 rounded-2xl bg-gray-50 dark:bg-slate-800 dark:text-white outline-none border-2 border-transparent focus:border-blue-500 resize-none h-24 transition-all"
                                    />

                                    <button
                                        onClick={() => {
                                            setView("search");
                                            setPickup("");
                                            setDestination("");
                                        }}
                                        className="w-full py-5 bg-blue-600 text-white font-black rounded-2xl shadow-lg shadow-blue-200 dark:shadow-none hover:bg-blue-700 transition-all active:scale-95"
                                    >
                                        Done
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                )}
            </div>
        </AppLayout>
    );
}