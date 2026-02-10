import { createContext, useContext } from "react";

// Provide a default shape so the IDE knows what's inside
export const ThemeContext = createContext({
    isDarkMode: false,
    toggleTheme: () => {},
});

export const useTheme = () => {
    const context = useContext(ThemeContext);
    if (!context) {
        throw new Error("useTheme must be used within a ThemeProvider");
    }
    return context;
};