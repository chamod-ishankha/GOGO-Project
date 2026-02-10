// src/pages/admin/AdminRides.jsx
import AppLayout from "@/components/AppLayout";
import { useState } from "react";

export default function AdminRides() {
    // Mock data for the table
    const [rides] = useState([
        { id: "RID-9928", rider: "Chamod M.", driver: "Nimal Silva", pickup: "Panadura", drop: "Colombo 03", status: "In Progress", type: "GoGo Luxury", price: "$12.40" },
        { id: "RID-9927", rider: "Kasun Perera", driver: "Sunil K.", pickup: "Galle Face", drop: "Borella", status: "Completed", type: "GoGo City", price: "$8.20" },
        { id: "RID-9926", rider: "John Doe", driver: "Unassigned", pickup: "Kandy Town", drop: "Peradeniya", status: "Searching", type: "GoGo City", price: "$5.50" },
        { id: "RID-9925", rider: "Saman Kumara", driver: "Arjun R.", pickup: "Mount Lavinia", drop: "Dehiwala", status: "Cancelled", type: "GoGo Bike", price: "$0.00" },
    ]);

    const getStatusStyle = (status) => {
        switch (status) {
            case 'Completed': return 'bg-green-100 text-green-700 dark:bg-green-900/30 dark:text-green-400';
            case 'In Progress': return 'bg-blue-100 text-blue-700 dark:bg-blue-900/30 dark:text-blue-400';
            case 'Cancelled': return 'bg-red-100 text-red-700 dark:bg-red-900/30 dark:text-red-400';
            case 'Searching': return 'bg-yellow-100 text-yellow-700 dark:bg-yellow-900/30 dark:text-yellow-400 animate-pulse';
            default: return 'bg-gray-100 text-gray-700';
        }
    };

    return (
        <AppLayout>
            <div className="space-y-6">
                {/* Header Actions */}
                <div className="flex flex-col md:flex-row md:items-center justify-between gap-4">
                    <div>
                        <h2 className="text-2xl font-black dark:text-white">Ride Management</h2>
                        <p className="text-sm text-gray-500">Track and manage all ride activities on your platform.</p>
                    </div>
                    <div className="flex gap-2">
                        <button className="px-4 py-2 bg-white dark:bg-slate-800 border border-gray-200 dark:border-slate-700 rounded-xl text-sm font-bold dark:text-white hover:bg-gray-50 transition-all">
                            📥 Export CSV
                        </button>
                    </div>
                </div>

                {/* Filter Bar */}
                <div className="bg-white dark:bg-slate-900 p-4 rounded-2xl border border-gray-100 dark:border-slate-800 flex flex-wrap gap-4 items-center">
                    <div className="flex-1 min-w-[200px] relative">
                        <span className="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400">🔍</span>
                        <input
                            type="text"
                            placeholder="Search by Ride ID or Rider name..."
                            className="w-full bg-gray-50 dark:bg-slate-800 border-none rounded-xl pl-11 pr-4 py-2 text-sm outline-none focus:ring-2 focus:ring-blue-500 dark:text-white"
                        />
                    </div>
                    <select className="bg-gray-50 dark:bg-slate-800 border-none rounded-xl px-4 py-2 text-sm dark:text-gray-300 outline-none">
                        <option>All Status</option>
                        <option>Completed</option>
                        <option>In Progress</option>
                        <option>Cancelled</option>
                    </select>
                </div>

                {/* Data Table */}
                <div className="bg-white dark:bg-slate-900 rounded-3xl border border-gray-100 dark:border-slate-800 overflow-hidden shadow-sm">
                    <div className="overflow-x-auto">
                        <table className="w-full text-left">
                            <thead className="bg-gray-50 dark:bg-slate-800/50 border-b border-gray-100 dark:border-slate-800">
                            <tr>
                                <th className="px-6 py-4 text-xs font-bold text-gray-400 uppercase tracking-wider">Ride Details</th>
                                <th className="px-6 py-4 text-xs font-bold text-gray-400 uppercase tracking-wider">Participants</th>
                                <th className="px-6 py-4 text-xs font-bold text-gray-400 uppercase tracking-wider">Route</th>
                                <th className="px-6 py-4 text-xs font-bold text-gray-400 uppercase tracking-wider">Status</th>
                                <th className="px-6 py-4 text-xs font-bold text-gray-400 uppercase tracking-wider">Fare</th>
                                <th className="px-6 py-4 text-xs font-bold text-gray-400 uppercase tracking-wider">Actions</th>
                            </tr>
                            </thead>
                            <tbody className="divide-y divide-gray-100 dark:divide-slate-800">
                            {rides.map((ride) => (
                                <tr key={ride.id} className="hover:bg-gray-50/80 dark:hover:bg-slate-800/30 transition-colors group">
                                    <td className="px-6 py-4">
                                        <p className="font-bold dark:text-white">{ride.id}</p>
                                        <p className="text-xs text-blue-500 font-medium">{ride.type}</p>
                                    </td>
                                    <td className="px-6 py-4">
                                        <p className="text-sm font-bold dark:text-gray-200">{ride.rider}</p>
                                        <p className="text-xs text-gray-400">Driver: {ride.driver}</p>
                                    </td>
                                    <td className="px-6 py-4">
                                        <p className="text-xs font-medium dark:text-gray-400">From: <span className="text-gray-600 dark:text-gray-200">{ride.pickup}</span></p>
                                        <p className="text-xs font-medium dark:text-gray-400">To: <span className="text-gray-600 dark:text-gray-200">{ride.drop}</span></p>
                                    </td>
                                    <td className="px-6 py-4">
                                            <span className={`px-3 py-1 rounded-full text-[10px] font-black uppercase ${getStatusStyle(ride.status)}`}>
                                                {ride.status}
                                            </span>
                                    </td>
                                    <td className="px-6 py-4 font-black dark:text-white">
                                        {ride.price}
                                    </td>
                                    <td className="px-6 py-4">
                                        <button className="p-2 hover:bg-gray-100 dark:hover:bg-slate-700 rounded-lg transition-all text-gray-400">
                                            👁️
                                        </button>
                                    </td>
                                </tr>
                            ))}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </AppLayout>
    );
}