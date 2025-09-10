import { User } from "@/types/user";
import { createStore } from "zustand";

export type UserState = {
    user: User | null;
};

export type UserActions = {
    login: (email: string, password: string) => void;
    register: (email: string, password: string) => void;
    logout: () => void;
};

export type UserStore = UserState & UserActions;

const defaultInitState: UserState = {
    user: null,
};

export const createUserStore = (initState: UserState = defaultInitState) => {
    return createStore<UserStore>()((set, get) => ({
        ...initState,
        login: (email, password) => {
            console.log("login", { email, password });

            set({
                user: {
                    id: crypto.randomUUID(),
                    email,
                    createdAt: new Date(),
                    passwordChangedAt: new Date(),
                    verification: {
                        valid: false,
                    },
                },
            });
        },
        register: (email, password) => {
            console.log("register", { email, password });

            set({
                user: {
                    id: crypto.randomUUID(),
                    email,
                    createdAt: new Date(),
                    passwordChangedAt: new Date(),
                    verification: {
                        valid: false,
                    },
                },
            });
        },
        logout: () => {
            console.log("logout");

            set({
                user: null,
            });
        },
    }));
};
