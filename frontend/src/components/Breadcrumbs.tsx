"use client";

import clsx from "clsx";
import Link from "next/link";
import { usePathname } from "next/navigation";

type Breadcrumb = {
    href: string;
    label: string;
};

export default function Breadcrumbs() {
    const path = usePathname();
    const parts = path.split("/");
    const crumbs: Breadcrumb[] = parts.map((crumb, i) => {
        if (crumb === "") return { href: "/", label: "Home" };

        if (crumb.match(/\d+/)) {
            if (crumb.length >= 10) {
                return {
                    href: "/" + parts.slice(1, i + 1).join("/"),
                    label: crumb.slice(0, 10) + "...",
                };
            }

            return {
                href: "/" + parts.slice(1, i + 1).join("/"),
                label: crumb,
            };
        }

        return {
            href: "/" + parts.slice(1, i + 1).join("/"),
            label: crumb[0].toUpperCase() + crumb.slice(1),
        };
    });

    return (
        <ul className="flex items-center px-4 py-2">
            {crumbs.map((crumb, i) => {
                return (
                    <li
                        key={i}
                        className="flex items-center text-sm"
                    >
                        <Link
                            href={crumb.href}
                            className={clsx(crumb.href === path && "font-bold")}
                        >
                            {crumb.label}
                        </Link>
                        <span className="mx-1.5 text-lg font-bold">
                            {i < crumbs.length - 1 && "/"}
                        </span>
                    </li>
                );
            })}
        </ul>
    );
}
