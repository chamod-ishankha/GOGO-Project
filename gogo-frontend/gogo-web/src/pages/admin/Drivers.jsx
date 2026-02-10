// src/pages/admin/AdminDrivers.jsx
import AppLayout from "@/components/AppLayout";
import { useState } from "react";

export default function AdminDrivers() {
    const [drivers] = useState([
        { id: "DRV-001", name: "Nimal Silva", email: "nimal@gogo.com", phone: "+94 77 123 4567", vehicle: "Toyota Prius (WP KH-1234)", rating: 4.9, status: "Active", verified: true },
        { id: "DRV-002", name: "Sunil Kasun", email: "sunil.k@gogo.com", phone: "+94 71 987 6543", vehicle: "Suzuki WagonR (WP CAD-5566)", rating: 4.7, status: "On Trip", verified: true },
        { id: "DRV-003", name: "Arjun Raj", email: "arjun@gogo.com", phone: "+94 76 555 4433", vehicle: "Honda Vezel (WP PH-9900)", rating: 3.2, status: "Suspended", verified: true },
        { id: "DRV-004", name: "Kamal Perera", email: "kamal.p@gogo.com", phone: "+94 70 111 2222", vehicle: "Nissan Leaf (WP BEQ-1122)", rating: 0.0, status: "Pending", verified: false },
    ]);

    const getStatusStyle = (status) => {
        switch (status) {
            case 'Active': return 'bg-green-100 text-green-700 dark:bg-green-900/30 dark:text-green-400';
            case 'On Trip': return 'bg-blue-100 text-blue-700 dark:bg-blue-900/30 dark:text-blue-400';
            case 'Suspended': return 'bg-red-100 text-red-700 dark:bg-red-900/30 dark:text-red-400';
            case 'Pending': return 'bg-yellow-100 text-yellow-700 dark:bg-yellow-900/30 dark:text-yellow-400';
            default: return 'bg-gray-100 text-gray-700';
        }
    };

    return (
        <AppLayout>
            <div className="space-y-6">
                {/* Header */}
                <div className="flex flex-col md:flex-row md:items-center justify-between gap-4">
                    <div>
                        <h2 className="text-2xl font-black dark:text-white">Driver Fleet</h2>
                        <p className="text-sm text-gray-500">Manage and monitor your GoGo partners.</p>
                    </div>
                    <button className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white font-bold text-sm rounded-xl transition-all shadow-lg shadow-blue-200 dark:shadow-none active:scale-95">
                        + Add New Driver
                    </button>
                </div>

                {/* Filter Bar */}
                <div className="bg-white dark:bg-slate-900 p-4 rounded-2xl border border-gray-100 dark:border-slate-800 flex flex-wrap gap-4 items-center">
                    <div className="flex-1 min-w-[200px] relative">
                        <span className="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400">🔍</span>
                        <input
                            type="text"
                            placeholder="Search by name, vehicle, or phone..."
                            className="w-full bg-gray-50 dark:bg-slate-800 border-none rounded-xl pl-11 pr-4 py-2 text-sm outline-none focus:ring-2 focus:ring-blue-500 dark:text-white"
                        />
                    </div>
                    <select className="bg-gray-50 dark:bg-slate-800 border-none rounded-xl px-4 py-2 text-sm dark:text-gray-300 outline-none">
                        <option>All Status</option>
                        <option>Active</option>
                        <option>On Trip</option>
                        <option>Suspended</option>
                    </select>
                </div>

                {/* Drivers Table */}
                <div className="bg-white dark:bg-slate-900 rounded-3xl border border-gray-100 dark:border-slate-800 overflow-hidden shadow-sm">
                    <div className="overflow-x-auto">
                        <table className="w-full text-left">
                            <thead className="bg-gray-50/50 dark:bg-slate-800/50 border-b border-gray-100 dark:border-slate-800">
                            <tr>
                                <th className="px-6 py-4 text-xs font-bold text-gray-400 uppercase tracking-wider">Driver</th>
                                <th className="px-6 py-4 text-xs font-bold text-gray-400 uppercase tracking-wider">Vehicle Details</th>
                                <th className="px-6 py-4 text-xs font-bold text-gray-400 uppercase tracking-wider">Rating</th>
                                <th className="px-6 py-4 text-xs font-bold text-gray-400 uppercase tracking-wider">Status</th>
                                <th className="px-6 py-4 text-xs font-bold text-gray-400 uppercase tracking-wider">Actions</th>
                            </tr>
                            </thead>
                            <tbody className="divide-y divide-gray-100 dark:divide-slate-800">
                            {drivers.map((driver) => (
                                <tr key={driver.id} className="hover:bg-gray-50/80 dark:hover:bg-slate-800/30 transition-colors group">
                                    <td className="px-6 py-4">
                                        <div className="flex items-center gap-4">
                                            <div className="w-12 h-12 rounded-2xl bg-gradient-to-br from-slate-100 to-slate-200 dark:from-slate-700 dark:to-slate-800 flex items-center justify-center text-xl shadow-sm group-hover:scale-110 transition-transform">
                                                👤
                                            </div>
                                            <div>
                                                <div className="flex items-center gap-2">
                                                    <p className="font-bold dark:text-white">{driver.name}</p>
                                                    {driver.verified && (
                                                        <span className="text-blue-500 text-sm" title="Verified Driver">✔️</span>
                                                    )}
                                                </div>
                                                <p className="text-xs text-gray-400">{driver.phone}</p>
                                            </div>
                                        </div>
                                    </td>
                                    <td className="px-6 py-4">
                                        <p className="text-sm font-semibold dark:text-gray-200">{driver.vehicle}</p>
                                        <p className="text-[10px] font-bold text-blue-600 dark:text-blue-400 uppercase tracking-tighter mt-1">Plate Verified</p>
                                    </td>
                                    <td className="px-6 py-4">
                                        <div className="flex items-center gap-1">
                                            <span className="text-yellow-500 font-bold">★</span>
                                            <span className="text-sm font-black dark:text-white">{driver.rating > 0 ? driver.rating : "New"}</span>
                                        </div>
                                    </td>
                                    <td className="px-6 py-4">
                                            <span className={`px-4 py-1.5 rounded-xl text-[10px] font-black uppercase tracking-wider ${getStatusStyle(driver.status)}`}>
                                                {driver.status}
                                            </span>
                                    </td>
                                    <td className="px-6 py-4">
                                        <div className="flex gap-2">
                                            <button className="p-2.5 bg-gray-50 dark:bg-slate-800 hover:bg-blue-50 dark:hover:bg-blue-900/30 text-gray-400 hover:text-blue-600 rounded-xl transition-all">
                                                ✏️
                                            </button>
                                            <button className="p-2.5 bg-gray-50 dark:bg-slate-800 hover:bg-red-50 dark:hover:bg-red-900/30 text-gray-400 hover:text-red-600 rounded-xl transition-all">
                                                🚫
                                            </button>
                                        </div>
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