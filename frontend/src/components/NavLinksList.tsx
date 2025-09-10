"use client";

import { useLinksStore } from "@/stores/links-store";
import Link from "next/link";

type NavLinksListProps = {
    onClose: () => void;
};

export default function NavLinksList({ onClose }: NavLinksListProps) {
    const links = useLinksStore((state) => state.links);

    return (
        <ul className="bg-background-accent border-border flex flex-col gap-1 rounded-lg border p-1 md:flex-row md:bg-transparent">
            {links.map((item) => (
                <li
                    key={item.href}
                    className=""
                >
                    <Link
                        href={item.href}
                        onClick={onClose}
                        className="text-foreground inline-flex w-full items-center gap-3 rounded-md p-2 px-4 dark:hover:bg-zinc-800"
                    >
                        {item.icon}
                        <span>{item.label}</span>
                    </Link>
                </li>
            ))}
        </ul>
    );
}
