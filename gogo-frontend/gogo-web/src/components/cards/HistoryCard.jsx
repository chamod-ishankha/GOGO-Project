
const getStatusStyle = (status) => {
    return status === "Completed"
        ? "bg-green-100 text-green-700 dark:bg-green-900/30 dark:text-green-400"
        : "bg-red-100 text-red-700 dark:bg-red-900/30 dark:text-red-400";
};

export default function HistoryCard({ ride }) {
    return (
        <div key={ride.id} className="bg-white dark:bg-slate-900 rounded-[2.5rem] border border-gray-100 dark:border-slate-800 p-6 md:p-8 hover:shadow-xl transition-all group">
            <div className="flex flex-col md:flex-row justify-between gap-6">

                {/* Left Side: Route & Driver */}
                <div className="space-y-4 flex-1">
                    <div className="flex items-center gap-3">
                                        <span className={`px-3 py-1 rounded-lg text-[10px] font-black uppercase ${getStatusStyle(ride.status)}`}>
                                            {ride.status}
                                        </span>
                        <span className="text-xs font-bold text-gray-400">{ride.date} • {ride.time}</span>
                    </div>

                    <div className="space-y-3 relative">
                        <div className="absolute left-1.5 top-2 bottom-2 w-0.5 bg-gray-100 dark:bg-slate-800"></div>
                        <div className="flex items-center gap-4 relative">
                            <div className="w-3 h-3 rounded-full bg-blue-500 border-2 border-white dark:border-slate-900"></div>
                            <p className="text-sm font-bold dark:text-white">{ride.from}</p>
                        </div>
                        <div className="flex items-center gap-4 relative">
                            <div className="w-3 h-3 rounded-full bg-green-500 border-2 border-white dark:border-slate-900"></div>
                            <p className="text-sm font-bold dark:text-white">{ride.to}</p>
                        </div>
                    </div>

                    <div className="flex items-center gap-2 pt-2">
                        <div className="w-8 h-8 rounded-full bg-gray-100 dark:bg-slate-800 flex items-center justify-center text-xs">👤</div>
                        <p className="text-xs font-medium text-gray-500">Driver: <span className="text-gray-800 dark:text-gray-200">{ride.driver}</span></p>
                    </div>
                </div>

                {/* Right Side: Fare & Actions */}
                <div className="md:text-right flex flex-row md:flex-col justify-between items-center md:items-end gap-4 border-t md:border-t-0 pt-4 md:pt-0 border-gray-50 dark:border-slate-800">
                    <div>
                        <p className="text-2xl font-black dark:text-white">{ride.fare}</p>
                        <p className="text-xs font-bold text-blue-600 uppercase tracking-tighter">{ride.type}</p>
                    </div>
                    <div className="flex gap-2">
                        <button className="px-5 py-2.5 bg-gray-50 dark:bg-slate-800 rounded-xl text-xs font-bold dark:text-white hover:bg-gray-100 transition-all">
                            Receipt
                        </button>
                        <button className="px-5 py-2.5 bg-blue-600 text-white rounded-xl text-xs font-bold hover:bg-blue-700 shadow-lg shadow-blue-100 dark:shadow-none transition-all">
                            Rebook
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}