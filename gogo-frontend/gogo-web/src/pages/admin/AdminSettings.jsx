// src/pages/admin/AdminSettings.jsx
import AppLayout from "@/components/AppLayout";
import { useState } from "react";

export default function AdminSettings() {
    const [activeTab, setActiveTab] = useState("General");

    const tabs = ["General", "Pricing", "Security", "Notifications"];

    return (
        <AppLayout>
            <div className="max-w-6xl mx-auto">
                <div className="mb-8">
                    <h2 className="text-3xl font-black dark:text-white">System Settings</h2>
                    <p className="text-gray-500">Configure your platform's global parameters.</p>
                </div>

                <div className="flex flex-col lg:flex-row gap-8">
                    {/* Settings Navigation */}
                    <div className="w-full lg:w-64 space-y-2">
                        {tabs.map((tab) => (
                            <button
                                key={tab}
                                onClick={() => setActiveTab(tab)}
                                className={`w-full text-left px-6 py-4 rounded-2xl font-bold transition-all ${
                                    activeTab === tab
                                        ? "bg-blue-600 text-white shadow-lg shadow-blue-200 dark:shadow-none"
                                        : "bg-white dark:bg-slate-900 text-gray-500 dark:text-gray-400 hover:bg-gray-50 dark:hover:bg-slate-800"
                                }`}
                            >
                                {tab}
                            </button>
                        ))}
                    </div>

                    {/* Settings Content Area */}
                    <div className="flex-1 bg-white dark:bg-slate-900 rounded-[2.5rem] border border-gray-100 dark:border-slate-800 p-8 shadow-sm">
                        {activeTab === "General" && <GeneralSettings />}
                        {activeTab === "Pricing" && <PricingSettings />}
                        {activeTab === "Security" && <div className="text-gray-400 italic">Security modules loading...</div>}
                    </div>
                </div>
            </div>
        </AppLayout>
    );
}

function GeneralSettings() {
    return (
        <div className="space-y-6">
            <h3 className="text-xl font-black dark:text-white mb-6">General Configuration</h3>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div className="space-y-2">
                    <label className="text-sm font-bold text-gray-500 px-1">Platform Name</label>
                    <input type="text" defaultValue="GoGo Global" className="w-full p-4 rounded-2xl bg-gray-50 dark:bg-slate-800 dark:text-white outline-none border-2 border-transparent focus:border-blue-500 transition-all" />
                </div>
                <div className="space-y-2">
                    <label className="text-sm font-bold text-gray-500 px-1">Support Email</label>
                    <input type="email" defaultValue="admin@gogo.com" className="w-full p-4 rounded-2xl bg-gray-50 dark:bg-slate-800 dark:text-white outline-none border-2 border-transparent focus:border-blue-500 transition-all" />
                </div>
            </div>

            <div className="pt-6 border-t border-gray-100 dark:border-slate-800">
                <div className="flex items-center justify-between p-4 bg-blue-50 dark:bg-blue-900/20 rounded-2xl">
                    <div>
                        <p className="font-bold text-blue-900 dark:text-blue-300">Maintenance Mode</p>
                        <p className="text-xs text-blue-700 dark:text-blue-400">Disable passenger bookings globally.</p>
                    </div>
                    <div className="w-12 h-6 bg-gray-300 rounded-full relative cursor-pointer">
                        <div className="absolute left-1 top-1 w-4 h-4 bg-white rounded-full transition-all"></div>
                    </div>
                </div>
            </div>

            <button className="mt-8 bg-blue-600 text-white font-black px-10 py-4 rounded-2xl hover:bg-blue-700 transition-all active:scale-95">
                Save Changes
            </button>
        </div>
    );
}

function PricingSettings() {
    return (
        <div className="space-y-6">
            <h3 className="text-xl font-black dark:text-white mb-6">Fare Estimation Rules</h3>
            <div className="space-y-4">
                {[
                    { label: "Base Fare (LKR)", value: "100.00" },
                    { label: "Price per KM", value: "85.00" },
                    { label: "Waiting Charge (per min)", value: "5.00" },
                ].map((item, i) => (
                    <div key={i} className="flex items-center justify-between p-4 border border-gray-100 dark:border-slate-800 rounded-2xl">
                        <span className="font-bold text-gray-600 dark:text-gray-300">{item.label}</span>
                        <input type="text" defaultValue={item.value} className="w-24 text-right font-black text-blue-600 bg-transparent outline-none" />
                    </div>
                ))}
            </div>
        </div>
    );
}