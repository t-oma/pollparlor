import { twMerge } from "tailwind-merge";
import NavLinksList from "./NavLinksList";

export default function NavMenu({ className }: { className?: string }) {
    return (
        <nav className={twMerge(className)}>
            <NavLinksList />
        </nav>
    );
}
