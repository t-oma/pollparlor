"use client";

import { useEffect, useRef } from "react";
import { createPortal } from "react-dom";
import Button from "./Button";
import { X } from "lucide-react";
import Divider from "./Divider";

type ModalProps = {
    open: boolean;
    onClose: () => void;
    title?: string;
    children: React.ReactNode;
    inertRootId?: string;
};

function getFocusableElements(container: HTMLElement): HTMLElement[] {
    const selector =
        'a[href], button, textarea, input, select, [tabindex]:not([tabindex="-1"])';
    return Array.from(container.querySelectorAll<HTMLElement>(selector)).filter(
        (el) => !el.hasAttribute("disabled") && el.tabIndex !== -1
    );
}

export default function Modal({
    open,
    onClose,
    title,
    children,
    inertRootId,
}: Readonly<ModalProps>) {
    const overlayRef = useRef<HTMLDivElement>(null);
    const dialogRef = useRef<HTMLDivElement>(null);
    const lastFocused = useRef<HTMLElement | null>(null);

    useEffect(() => {
        if (!open) return;

        lastFocused.current = document.activeElement as HTMLElement | null;

        const root = inertRootId ? document.getElementById(inertRootId) : null;
        if (root) root.inert = true;

        const prevOverflow = document.documentElement.style.overflow;
        document.documentElement.style.overflow = "hidden";

        const container = dialogRef.current!;
        const focusables = getFocusableElements(container);
        const initialFocus = focusables[0] ?? container;
        initialFocus.focus();

        const onKeyDown = (e: KeyboardEvent) => {
            if (e.key === "Escape") {
                e.preventDefault();
                onClose();
                return;
            }

            if (e.key === "Tab") {
                const currentFocusables = getFocusableElements(container);

                if (currentFocusables.length === 0) {
                    e.preventDefault();
                    return;
                }

                const first = currentFocusables[0];
                const last = currentFocusables[currentFocusables.length - 1];
                const { activeElement } = document;

                if (e.shiftKey && activeElement === first) {
                    e.preventDefault();
                    last.focus();
                } else if (!e.shiftKey && activeElement === last) {
                    e.preventDefault();
                    first.focus();
                }
            }
        };

        document.addEventListener("keydown", onKeyDown);

        return () => {
            document.removeEventListener("keydown", onKeyDown);
            document.documentElement.style.overflow = prevOverflow;
            if (root) root.inert = false;
            lastFocused.current?.focus();
        };
    }, [open, onClose, inertRootId]);

    if (!open) return null;

    const handleOverlayClick = (e: React.MouseEvent) => {
        if (e.target === overlayRef.current) onClose();
    };

    return createPortal(
        <div
            ref={overlayRef}
            onClick={handleOverlayClick}
            className="fixed inset-0 z-50 grid place-items-center bg-black/60"
        >
            <div
                ref={dialogRef}
                role="dialog"
                aria-modal="true"
                aria-labelledby={title ? "modal-title" : undefined}
                aria-label={title ? undefined : "Modal"}
                tabIndex={-1}
                className="bg-background flex h-full max-h-[60vh] w-full max-w-[85vw] flex-col overflow-auto rounded-lg p-4 shadow"
                onClick={(e) => e.stopPropagation()}
            >
                <div className="flex items-center justify-between gap-2 space-x-2">
                    {title && (
                        <h2
                            id="modal-title"
                            className="text-xl font-semibold"
                        >
                            {title}
                        </h2>
                    )}
                    <Button
                        onClick={onClose}
                        aria-label="Close modal"
                        className="p-1"
                    >
                        <X className="h-4 w-4" />
                    </Button>
                </div>

                <Divider className="my-4" />

                {children}
            </div>
        </div>,
        document.body
    );
}
