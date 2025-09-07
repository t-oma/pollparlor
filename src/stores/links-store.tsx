import { BookCopy, House, Settings } from "lucide-react";
import React from "react";
import { create } from "zustand";

export type PageLink = {
    href: string;
    label: string;
    icon?: React.ReactNode;
};

export type LinksState = {
    links: PageLink[];
};

export type LinksActions = {
    addLink: (link: PageLink) => void;
};

export type LinksStore = LinksState & LinksActions;

const defaultInitState: LinksState = {
    links: [
        { href: "/", label: "Home", icon: <House className="h-5 w-5" /> },
        {
            href: "/polls",
            label: "Polls",
            icon: <BookCopy className="h-5 w-5" />,
        },
        {
            href: "/settings",
            label: "Settings",
            icon: <Settings className="h-5 w-5" />,
        },
    ],
};

export const useLinksStore = create<LinksStore>((set) => ({
    ...defaultInitState,
    addLink: (link: PageLink) => {
        set((state) => ({
            links: [...state.links, link],
        }));
    },
}));
