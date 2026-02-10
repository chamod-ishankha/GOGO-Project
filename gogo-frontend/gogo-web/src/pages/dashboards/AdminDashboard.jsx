import AppLayout from "@/components/AppLayout";
import StatCard from "@/components/dashboard/StatCard";

export default function AdminDashboard() {
    // Mock data - eventually this comes from your API
    const stats = [
        { title: "Active Rides", value: "12", icon: "🚗", trend: 12, colors: { bg: "bg-blue-500", text: "text-blue-500" } },
        { title: "Drivers Online", value: "48", icon: "🟢", trend: 5, colors: { bg: "bg-green-500", text: "text-green-500" } },
        { title: "Today's Revenue", value: "$1,240", icon: "💰", trend: -2, colors: { bg: "bg-purple-500", text: "text-purple-500" } },
        { title: "New Users", value: "156", icon: "👥", trend: 18, colors: { bg: "bg-orange-500", text: "text-orange-500" } },
    ];

    const recentRides = [
        { id: "#G-1024", rider: "Chamod M.", driver: "Kasun P.", status: "In Progress", amount: "$15.00" },
        { id: "#G-1023", rider: "John Doe", driver: "Arjun K.", status: "Completed", amount: "$22.50" },
        { id: "#G-1022", rider: "Sarah W.", driver: "Nimal S.", status: "Cancelled", amount: "$0.00" },
    ];

    return (
        <AppLayout>
            <div className="space-y-8">
                {/* Stats Grid */}
                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
                    {stats.map((s, i) => (
                        <StatCard key={i} title={s.title} value={s.value} icon={s.icon} trend={s.trend} colorClass={s.colors} />
                    ))}
                </div>

                {/* Table Section */}
                <div className="bg-white dark:bg-slate-900 rounded-3xl shadow-sm border border-gray-100 dark:border-slate-800 overflow-hidden">
                    <div className="p-6 border-b border-gray-50 dark:border-slate-800 flex justify-between items-center">
                        <h3 className="font-bold text-lg dark:text-white">Recent Rides</h3>
                        <button className="text-blue-600 text-sm font-bold hover:underline">View All</button>
                    </div>
                    <div className="overflow-x-auto">
                        <table className="w-full text-left">
                            <thead className="bg-gray-50 dark:bg-slate-800/50 text-gray-500 dark:text-gray-400 text-xs uppercase uppercase tracking-wider">
                            <tr>
                                <th className="px-6 py-4">Ride ID</th>
                                <th className="px-6 py-4">Rider</th>
                                <th className="px-6 py-4">Driver</th>
                                <th className="px-6 py-4">Status</th>
                                <th className="px-6 py-4">Amount</th>
                            </tr>
                            </thead>
                            <tbody className="divide-y divide-gray-100 dark:divide-slate-800">
                            {recentRides.map((ride) => (
                                <tr key={ride.id} className="hover:bg-gray-50 dark:hover:bg-slate-800/30 transition-colors">
                                    <td className="px-6 py-4 font-medium dark:text-white">{ride.id}</td>
                                    <td className="px-6 py-4 dark:text-gray-300">{ride.rider}</td>
                                    <td className="px-6 py-4 dark:text-gray-300">{ride.driver}</td>
                                    <td className="px-6 py-4">
                                            <span className={`px-3 py-1 rounded-full text-[10px] font-black uppercase tracking-tighter ${
                                                ride.status === 'Completed' ? 'bg-green-100 text-green-700' :
                                                    ride.status === 'Cancelled' ? 'bg-red-100 text-red-700' : 'bg-blue-100 text-blue-700'
                                            }`}>
                                                {ride.status}
                                            </span>
                                    </td>
                                    <td className="px-6 py-4 font-bold dark:text-white">{ride.amount}</td>
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