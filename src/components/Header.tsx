"use client";

import { LogIn, User } from "lucide-react";
import Button from "./Button";
import { useUserStore } from "@/providers/user-store-provider";

export default function Header() {
    const { user, login, logout, register } = useUserStore((store) => store);

    return (
        <header className="flex items-center justify-between p-4 lg:p-8">
            <h1 className="text-xl font-bold">Pollparlor</h1>

            <div>
                {user ? (
                    <Button
                        onClick={() => logout()}
                        className=""
                    >
                        <span className="sr-only font-medium md:not-sr-only">
                            Account
                        </span>
                        <User />
                    </Button>
                ) : (
                    <Button
                        onClick={() => login("test@test.com", "test")}
                        className=""
                    >
                        <span className="sr-only font-medium md:not-sr-only">
                            Login
                        </span>
                        <LogIn />
                    </Button>
                )}
            </div>
        </header>
    );
}
