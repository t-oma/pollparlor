"use client";

import { useUserStore } from "@/providers/user-store-provider";
import { LogIn, User } from "lucide-react";
import Button from "../Button";

export default function UserAccount() {
    const { user, login, logout, register } = useUserStore((store) => store);

    return (
        <div>
            {user ? (
                <Button
                    onClick={() => logout()}
                    className="md:px-4"
                >
                    <User className="h-5 w-5" />
                    <span className="sr-only font-medium md:not-sr-only">
                        Account
                    </span>
                </Button>
            ) : (
                <Button
                    onClick={() => login("test@test.com", "test")}
                    className="md:px-4"
                >
                    <LogIn className="h-5 w-5" />
                    <span className="sr-only font-medium md:not-sr-only">
                        Login
                    </span>
                </Button>
            )}
        </div>
    );
}
