import React from "react";
import { twMerge } from "tailwind-merge";

export type ButtonProps = React.ButtonHTMLAttributes<HTMLButtonElement> & {
    className?: string;
};

const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(function Button(
    { children, className, ...rest },
    ref
) {
    return (
        <button
            ref={ref}
            type={rest.type ?? "button"}
            className={twMerge(
                "border-border inline-flex cursor-pointer items-center justify-center gap-2 rounded-md border p-2 transition-colors disabled:cursor-default dark:bg-zinc-900 dark:hover:bg-zinc-800 dark:active:bg-zinc-700",
                className
            )}
            {...rest}
        >
            {children}
        </button>
    );
});

export default Button;
