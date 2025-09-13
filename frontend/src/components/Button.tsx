import Link from "next/link";
import React from "react";
import { twMerge } from "tailwind-merge";

type BaseProps = {
    className?: string;
};

type ButtonAsButton = BaseProps &
    Omit<React.ButtonHTMLAttributes<HTMLButtonElement>, "href"> & {
        href?: never;
    };

type ButtonAsLink = BaseProps & React.ComponentProps<typeof Link>;

export type ButtonProps = ButtonAsButton | ButtonAsLink;

const Button = React.forwardRef<
    HTMLButtonElement | HTMLAnchorElement,
    ButtonProps
>(function Button({ children, className, ...rest }, ref) {
    const baseClasses =
        "border-border inline-flex cursor-pointer items-center justify-center gap-2 rounded-md border p-2 transition-colors disabled:cursor-default dark:bg-zinc-900 dark:hover:bg-zinc-800 dark:active:bg-zinc-700";

    if ("href" in rest && rest.href) {
        const { href, ...linkProps } = rest;
        return (
            <Link
                ref={ref as React.Ref<HTMLAnchorElement>}
                href={href}
                className={twMerge(baseClasses, className)}
                {...linkProps}
            >
                {children}
            </Link>
        );
    }

    const { type, ...buttonProps } = rest;
    return (
        <button
            ref={ref as React.Ref<HTMLButtonElement>}
            type={type ?? "button"}
            className={twMerge(baseClasses, className)}
            {...buttonProps}
        >
            {children}
        </button>
    );
});

export default Button;
