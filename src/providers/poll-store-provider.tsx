"use client";

import { createPollStore, PollStore } from "@/stores/poll-store";
import { type ReactNode, createContext, useRef, useContext } from "react";
import { useStore } from "zustand";

export type PollStoreApi = ReturnType<typeof createPollStore>;

export const PollStoreContext = createContext<PollStoreApi | undefined>(
    undefined
);

export default function PollStoreProvider({
    children,
}: Readonly<{ children: ReactNode }>) {
    const storeRef = useRef<PollStoreApi | null>(null);
    if (storeRef.current === null) {
        storeRef.current = createPollStore();
    }

    return (
        <PollStoreContext.Provider value={storeRef.current}>
            {children}
        </PollStoreContext.Provider>
    );
}

export function usePollStore<T>(selector: (store: PollStore) => T) {
    const pollStore = useContext(PollStoreContext);
    if (!pollStore) {
        throw new Error("usePollStore must be used within a PollStoreProvider");
    }

    return useStore(pollStore, selector);
}
