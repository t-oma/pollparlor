"use client";

import { Menu, X } from "lucide-react";
import Button from "./Button";
import { useState } from "react";
import {
    autoUpdate,
    flip,
    FloatingFocusManager,
    offset,
    shift,
    useClick,
    useDismiss,
    useFloating,
    useInteractions,
    useRole,
} from "@floating-ui/react";

export default function NavMenuDropdown({ className }: { className?: string }) {
    const [isOpen, setIsOpen] = useState(false);

    const { refs, floatingStyles, context } = useFloating({
        open: isOpen,
        onOpenChange: setIsOpen,
        middleware: [offset({ mainAxis: 10, crossAxis: 28 }), flip(), shift()],
        whileElementsMounted: autoUpdate,
    });

    const click = useClick(context);
    const dismiss = useDismiss(context);
    const role = useRole(context);

    const { getReferenceProps, getFloatingProps } = useInteractions([
        click,
        dismiss,
        role,
    ]);

    return (
        <>
            <Button
                ref={refs.setReference}
                {...getReferenceProps()}
                onClick={() => setIsOpen((prev) => !prev)}
                className={className}
            >
                <span className="sr-only">Menu</span>
                {isOpen ? <X /> : <Menu />}
            </Button>
            {isOpen && (
                <FloatingFocusManager
                    context={context}
                    modal={false}
                >
                    <nav
                        ref={refs.setFloating}
                        style={floatingStyles}
                        {...getFloatingProps()}
                        className="z-50 shadow-lg"
                    >
                        <ul className="bg-background-accent border-border flex flex-col gap-1 rounded-lg border p-1">
                            {["Home", "Polls", "Settings"].map((item) => (
                                <li
                                    key={item}
                                    className="rounded-md dark:hover:bg-zinc-800"
                                >
                                    <a
                                        href={"#/" + item}
                                        className="text-foreground block p-2 px-4"
                                    >
                                        {item}
                                    </a>
                                </li>
                            ))}
                        </ul>
                    </nav>
                </FloatingFocusManager>
            )}
        </>
    );
}
