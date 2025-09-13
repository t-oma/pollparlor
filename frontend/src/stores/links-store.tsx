import { BookCopy, House, SquarePen } from "lucide-react";
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
        { href: "/", label: "Home", icon: <House className="h-4 w-4" /> },
        {
            href: "/polls",
            label: "Polls",
            icon: <BookCopy className="h-4 w-4" />,
        },
        {
            href: "/polls/create",
            label: "Create poll",
            icon: <SquarePen className="h-4 w-4" />,
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
