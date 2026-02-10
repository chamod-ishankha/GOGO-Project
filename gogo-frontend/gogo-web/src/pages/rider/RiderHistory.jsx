// src/pages/rider/RiderHistory.jsx
import AppLayout from "@/components/AppLayout";
import { useState } from "react";
import HistoryCard from "@/components/cards/HistoryCard.jsx";

export default function RiderHistory() {
    // Mock History Data
    const [history] = useState([
        {
            id: "TRP-8821",
            date: "Feb 10, 2026",
            time: "02:30 PM",
            from: "Panadura Main Stand",
            to: "Colombo Fort",
            fare: "LKR 1,450.00",
            driver: "Arjun R.",
            status: "Completed",
            type: "GoGo Economy"
        },
        {
            id: "TRP-8819",
            date: "Feb 08, 2026",
            time: "09:15 AM",
            from: "Bambalapitiya",
            to: "Dehiwala",
            fare: "LKR 650.00",
            driver: "Nimal Silva",
            status: "Completed",
            type: "GoGo Bike"
        },
        {
            id: "TRP-8815",
            date: "Feb 05, 2026",
            time: "07:45 PM",
            from: "Galle Face",
            to: "Mount Lavinia",
            fare: "LKR 0.00",
            driver: "Hemal K.",
            status: "Cancelled",
            type: "GoGo Premium"
        },
    ]);

    return (
        <AppLayout>
            <div className="max-w-4xl mx-auto space-y-6">
                <div>
                    <h2 className="text-3xl font-black dark:text-white">Your Trips</h2>
                    <p className="text-gray-500">View and manage your past GoGo journeys.</p>
                </div>

                {/* Filter Tabs */}
                <div className="flex gap-2 overflow-x-auto pb-2">
                    {['All', 'Completed', 'Cancelled'].map((tab) => (
                        <button key={tab} className="px-6 py-2 rounded-full bg-white dark:bg-slate-900 border border-gray-100 dark:border-slate-800 text-sm font-bold dark:text-gray-300 hover:border-blue-500 transition-all shrink-0">
                            {tab}
                        </button>
                    ))}
                </div>

                {/* History List */}
                <div className="space-y-4">
                    {history.map((ride) => (
                        <HistoryCard ride={ride} />
                    ))}
                </div>
            </div>
        </AppLayout>
    );
}