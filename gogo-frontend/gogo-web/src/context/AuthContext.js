import { createContext, useContext } from "react";

// Just the context and the hook here
export const AuthContext = createContext({
    user: null,
    login: () => {},  // Placeholder function
    logout: () => {}  // Placeholder function
});

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error("useAuth must be used within an AuthProvider");
    }
    return context;
};