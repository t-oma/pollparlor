"use client";

import { createUserStore, UserStore } from "@/stores/user-store";
import { type ReactNode, createContext, useRef, useContext } from "react";
import { useStore } from "zustand";

export type UserStoreApi = ReturnType<typeof createUserStore>;

export const UserStoreContext = createContext<UserStoreApi | undefined>(
    undefined
);

export default function UserStoreProvider({
    children,
}: Readonly<{ children: ReactNode }>) {
    const storeRef = useRef<UserStoreApi | null>(null);
    if (storeRef.current === null) {
        storeRef.current = createUserStore();
    }

    return (
        <UserStoreContext.Provider value={storeRef.current}>
            {children}
        </UserStoreContext.Provider>
    );
}

export function useUserStore<T>(selector: (store: UserStore) => T) {
    const userStore = useContext(UserStoreContext);
    if (!userStore) {
        throw new Error("useUserStore must be used within a UserStoreProvider");
    }

    return useStore(userStore, selector);
}
