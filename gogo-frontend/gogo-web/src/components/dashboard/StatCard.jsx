export default function StatCard({ title, value, icon, trend, colorClass }) {
    return (
        <div className="bg-white dark:bg-slate-900 p-6 rounded-3xl shadow-sm border border-gray-100 dark:border-slate-800 transition-all hover:shadow-md">
            <div className="flex justify-between items-start">
                <div className={`p-3 rounded-2xl bg-opacity-10 ${colorClass.bg} ${colorClass.text}`}>
                    <span className="text-2xl">{icon}</span>
                </div>
                {trend && (
                    <span className={`text-xs font-bold px-2 py-1 rounded-lg ${trend > 0 ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-600'}`}>
                        {trend > 0 ? '↑' : '↓'} {Math.abs(trend)}%
                    </span>
                )}
            </div>
            <div className="mt-4">
                <p className="text-sm font-medium text-gray-500 dark:text-gray-400">{title}</p>
                <h3 className="text-3xl font-black dark:text-white mt-1">{value}</h3>
            </div>
        </div>
    );
}