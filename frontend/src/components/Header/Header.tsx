import NavMenuDropdown from "./NavMenuDropdown";
import NavMenu from "./NavMenu";
import UserAccount from "./UserAccount";

export default function Header() {
    return (
        <header className="flex items-center justify-between p-4 lg:p-4">
            <NavMenuDropdown className="block md:hidden" />
            <p className="text-xl font-bold">Pollparlor</p>

            <NavMenu className="hidden md:block" />

            <UserAccount />
        </header>
    );
}
