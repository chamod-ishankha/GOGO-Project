import AppLayout from "@/components/AppLayout";
import { useState, useEffect } from "react";

export default function DriverDashboard() {
    const [isOnline, setIsOnline] = useState(false);
    const [incomingRequest, setIncomingRequest] = useState(null);
    const [currentRide, setCurrentRide] = useState(null); // Track the active ride

    useEffect(() => {
        let timer;
        // Only show requests if online and NOT currently on a ride
        if (isOnline && !currentRide) {
            timer = setTimeout(() => {
                setIncomingRequest({
                    rider: "Chamod M.",
                    pickup: "Panadura Main Stand",
                    drop: "Colombo Fort",
                    fare: "LKR 1,450.00",
                    distance: "24.5 km",
                    rating: "4.8 ⭐"
                });
            }, 5000);
        }
        return () => clearTimeout(timer);
    }, [isOnline, currentRide]);

    const handleAccept = () => {
        setCurrentRide(incomingRequest);
        setIncomingRequest(null);
    };

    const handleCompleteRide = () => {
        alert(`Ride Completed! ${currentRide.fare} added to your wallet.`);
        setCurrentRide(null);
    };

    return (
        <AppLayout>
            <div className="space-y-6">

                {/* --- CASE 1: ON A RIDE (Navigation View) --- */}
                {currentRide ? (
                    <div className="animate-in slide-in-from-right duration-500">
                        <div className="bg-white dark:bg-slate-900 rounded-[2.5rem] border border-gray-100 dark:border-slate-800 overflow-hidden shadow-xl">
                            {/* Simple Mock Map / Progress Area */}
                            <div className="h-48 bg-blue-50 dark:bg-slate-800 flex flex-col items-center justify-center p-6 text-center">
                                <div className="w-16 h-16 bg-blue-600 rounded-full flex items-center justify-center text-white text-3xl mb-3 shadow-lg animate-bounce">
                                    📍
                                </div>
                                <p className="font-black dark:text-white text-lg">Heading to {currentRide.drop}</p>
                                <p className="text-sm text-blue-600 font-bold">Estimated arrival: 12 mins</p>
                            </div>

                            {/* Passenger Summary */}
                            <div className="p-8">
                                <div className="flex items-center justify-between mb-8">
                                    <div className="flex items-center gap-4">
                                        <div className="w-14 h-14 bg-gray-100 dark:bg-slate-800 rounded-2xl flex items-center justify-center text-2xl">👤</div>
                                        <div>
                                            <p className="text-xs text-gray-400 font-bold uppercase tracking-widest">Passenger</p>
                                            <h3 className="text-xl font-black dark:text-white">{currentRide.rider}</h3>
                                        </div>
                                    </div>
                                    <div className="flex gap-2">
                                        <button className="p-4 bg-blue-50 dark:bg-blue-900/20 text-blue-600 rounded-2xl">📞</button>
                                        <button className="p-4 bg-blue-50 dark:bg-blue-900/20 text-blue-600 rounded-2xl">💬</button>
                                    </div>
                                </div>

                                <button
                                    onClick={handleCompleteRide}
                                    className="w-full py-5 bg-green-500 hover:bg-green-600 text-white font-black rounded-3xl shadow-lg shadow-green-100 dark:shadow-none transition-all active:scale-95"
                                >
                                    COMPLETE RIDE & COLLECT {currentRide.fare}
                                </button>
                            </div>
                        </div>
                    </div>
                ) : (
                    /* --- CASE 2: NORMAL DASHBOARD (Online/Offline) --- */
                    <>
                        {/* Status Toggle Header */}
                        <div className={`p-6 rounded-3xl flex items-center justify-between transition-all duration-500 ${isOnline ? 'bg-green-500 text-white' : 'bg-slate-200 dark:bg-slate-800 text-slate-600 dark:text-slate-400'}`}>
                            <div>
                                <h2 className="text-2xl font-black">{isOnline ? "You are Online" : "You are Offline"}</h2>
                                <p className="text-sm opacity-80">{isOnline ? "Looking for nearby passengers..." : "Go online to start earning"}</p>
                            </div>
                            <button
                                onClick={() => setIsOnline(!isOnline)}
                                className={`px-8 py-3 rounded-2xl font-black transition-all active:scale-95 ${isOnline ? 'bg-white text-green-600' : 'bg-blue-600 text-white'}`}
                            >
                                {isOnline ? "Go Offline" : "Go Online"}
                            </button>
                        </div>

                        {/* Driver Stats */}
                        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                            <div className="bg-white dark:bg-slate-900 p-6 rounded-3xl border border-gray-100 dark:border-slate-800">
                                <p className="text-gray-400 text-sm font-bold uppercase">Today's Earnings</p>
                                <h3 className="text-4xl font-black dark:text-white mt-1">$142.00</h3>
                                <p className="text-green-500 text-xs font-bold mt-2">↑ 15% from yesterday</p>
                            </div>
                            <div className="bg-white dark:bg-slate-900 p-6 rounded-3xl border border-gray-100 dark:border-slate-800">
                                <p className="text-gray-400 text-sm font-bold uppercase">Trips Today</p>
                                <h3 className="text-4xl font-black dark:text-white mt-1">08</h3>
                            </div>
                            <div className="bg-white dark:bg-slate-900 p-6 rounded-3xl border border-gray-100 dark:border-slate-800">
                                <p className="text-gray-400 text-sm font-bold uppercase">Rating</p>
                                <h3 className="text-4xl font-black dark:text-white mt-1">4.92 ⭐</h3>
                            </div>
                        </div>

                        {/* Live Demand Map Placeholder */}
                        <div className="bg-white dark:bg-slate-900 rounded-3xl border border-gray-100 dark:border-slate-800 overflow-hidden">
                            <div className="p-6 border-b border-gray-50 dark:border-slate-800">
                                <h3 className="font-bold dark:text-white">High Demand Zones</h3>
                            </div>
                            <div className="p-8 flex flex-col items-center justify-center text-center opacity-40">
                                <div className="text-5xl mb-4">🗺️</div>
                                <p className="font-bold dark:text-white italic">Interactive demand map coming soon...</p>
                            </div>
                        </div>
                    </>
                )}

                {/* --- INCOMING REQUEST MODAL --- */}
                {incomingRequest && (
                    <div className="fixed inset-0 z-50 flex items-center justify-center p-6 bg-slate-900/60 backdrop-blur-sm animate-in fade-in duration-300">
                        <div className="bg-white dark:bg-slate-900 w-full max-w-md rounded-[2.5rem] shadow-2xl overflow-hidden animate-in zoom-in duration-300">
                            {/* Header */}
                            <div className="bg-blue-600 p-6 text-white text-center">
                                <p className="text-xs font-bold uppercase tracking-widest opacity-80 mb-1">New Request Found</p>
                                <h3 className="text-2xl font-black">Incoming Job!</h3>
                            </div>

                            <div className="p-8">
                                <div className="flex justify-between items-center mb-8">
                                    <div className="flex items-center gap-3">
                                        <div className="w-12 h-12 bg-blue-50 dark:bg-slate-800 rounded-full flex items-center justify-center text-xl">👤</div>
                                        <div>
                                            <p className="font-black dark:text-white">{incomingRequest.rider}</p>
                                            <p className="text-xs text-blue-600 font-bold">{incomingRequest.rating}</p>
                                        </div>
                                    </div>
                                    <div className="text-right">
                                        <p className="text-2xl font-black text-green-600">{incomingRequest.fare}</p>
                                        <p className="text-xs text-gray-400 font-bold">{incomingRequest.distance}</p>
                                    </div>
                                </div>

                                {/* Route Info */}
                                <div className="space-y-6 relative mb-10">
                                    <div className="absolute left-2.5 top-3 bottom-3 w-0.5 bg-gray-200 dark:bg-slate-700 border-dashed border-l"></div>
                                    <div className="flex items-start gap-4 relative">
                                        <div className="w-5 h-5 rounded-full bg-blue-600 border-4 border-white dark:border-slate-900 z-10"></div>
                                        <div>
                                            <p className="text-[10px] font-bold text-gray-400 uppercase tracking-widest">Pickup</p>
                                            <p className="font-bold dark:text-gray-200">{incomingRequest.pickup}</p>
                                        </div>
                                    </div>
                                    <div className="flex items-start gap-4 relative">
                                        <div className="w-5 h-5 rounded-full bg-green-500 border-4 border-white dark:border-slate-900 z-10"></div>
                                        <div>
                                            <p className="text-[10px] font-bold text-gray-400 uppercase tracking-widest">Drop-off</p>
                                            <p className="font-bold dark:text-gray-200">{incomingRequest.drop}</p>
                                        </div>
                                    </div>
                                </div>

                                {/* Buttons */}
                                <div className="grid grid-cols-2 gap-4">
                                    <button
                                        onClick={() => setIncomingRequest(null)}
                                        className="py-4 rounded-2xl font-bold text-gray-400 hover:bg-gray-100 dark:hover:bg-slate-800 transition-all"
                                    >
                                        Decline
                                    </button>
                                    <button
                                        onClick={handleAccept}
                                        className="py-4 rounded-2xl font-black bg-blue-600 text-white hover:bg-blue-700 shadow-lg shadow-blue-200 dark:shadow-none transition-all active:scale-95"
                                    >
                                        ACCEPT
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