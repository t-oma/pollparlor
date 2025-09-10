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
import NavLinksList from "./NavLinksList";

export default function NavMenuDropdown({ className }: { className?: string }) {
    const [isOpen, setIsOpen] = useState(false);

    const { refs, floatingStyles, context } = useFloating({
        open: isOpen,
        onOpenChange: setIsOpen,
        middleware: [offset({ mainAxis: 10, crossAxis: 48 }), flip(), shift()],
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
                {isOpen ? (
                    <X className="h-5 w-5" />
                ) : (
                    <Menu className="h-5 w-5" />
                )}
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
                        <NavLinksList onClose={() => setIsOpen(false)} />
                    </nav>
                </FloatingFocusManager>
            )}
        </>
    );
}
