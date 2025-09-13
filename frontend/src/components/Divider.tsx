import { twMerge } from "tailwind-merge";

type DividerProps = {
    className?: string;
    orientation?: "horizontal" | "vertical";
};

export default function Divider({
    className,
    orientation = "horizontal",
}: Readonly<DividerProps>) {
    if (orientation === "horizontal") {
        return (
            <hr className={twMerge("border-border my-2 w-full", className)} />
        );
    }

    return <hr className={twMerge("border-border", className)} />;
}
